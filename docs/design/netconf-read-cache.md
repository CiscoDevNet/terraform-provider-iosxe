# NETCONF Read Cache Design Document

> **Status**: Implemented (Experimental - Disabled by Default)
> **Implementation Date**: 2025-12-23
> **Last Updated**: 2025-12-23

## Performance Note

The cache is currently **disabled by default** due to performance characteristics observed during testing:

- **Parsing overhead**: Each XPath query on the cached 100KB+ config takes ~300-950ms
- **Break-even point**: Estimated at 1000+ resources where network latency savings exceed parsing overhead
- **Recommendation**: Only enable for very large deployments where NETCONF round-trip time dominates

For typical deployments (<500 resources), individual filtered NETCONF queries are faster because:
1. Each filtered response is small (~1-5KB) and fast to parse
2. The full config (100KB+) requires expensive path traversal for each resource read

## Overview

This document describes the design and requirements for implementing a read cache for NETCONF read operations in the IOS-XE Terraform Provider. The cache significantly improves performance by reducing the number of NETCONF operations to the device during Terraform plan and apply phases.

## Problem Statement

Currently, every resource read operation performs an individual `GetConfig` NETCONF RPC to the device:

```
Terraform Plan (100 resources) → 100 GetConfig operations to device
Terraform Apply (50 changes)  → 50+ reads + writes
```

This results in:
- **High latency**: Each NETCONF RPC has overhead (connection, lock acquisition, XML parsing)
- **Device load**: Multiple concurrent sessions or sequential requests burden the device
- **Slow plans**: `terraform plan` can take minutes for large configurations

## Proposed Solution

Implement a read cache that:
1. Retrieves the **full device configuration** using a single `GetConfig` operation with `NoFilter()`
2. Caches this configuration in memory
3. Serves all subsequent resource reads from the cached configuration
4. Invalidates the cache on any write operation (Create/Update/Delete)

```
Terraform Plan (100 resources) → 1 GetConfig (full config) + 100 local XPath queries
Terraform Apply (50 changes)  → Cache invalidation after each write, refresh as needed
```

## Architecture

### Cache Structure

```go
// NetconfReadCache holds the cached device configuration
// Located in: internal/provider/cache/cache.go
type NetconfReadCache struct {
    mu          sync.RWMutex
    enabled     bool          // Whether caching is enabled for this device
    populated   bool          // Whether cache contains data
    configData  xmldot.Result // Cached full configuration (rpc-reply structure)
    populatedAt time.Time     // When cache was last populated
}
```

**Key Methods:**
- `New(enabled bool) *NetconfReadCache` - Create new cache instance
- `IsEnabled() bool` - Check if caching is enabled
- `Get(ctx) (xmldot.Result, bool)` - Get cached data (returns cache hit indicator)
- `Populate(ctx, client) (xmldot.Result, error)` - Fetch and cache full config, returns cached data
- `Invalidate(ctx)` - Clear the cache

### Provider Data Enhancement

```go
type IosxeProviderDataDevice struct {
    RestconfClient  *restconf.Client
    NetconfClient   *netconf.Client
    Protocol        string
    ReuseConnection bool
    AutoCommit      bool
    Managed         bool
    NetconfOpMutex  sync.Mutex

    // NEW: Read cache for NETCONF operations
    ReadCache       *NetconfReadCache
    ReadCacheEnabled bool  // Derived from provider config
}
```

### Provider Configuration

Add new provider attribute:

```hcl
provider "iosxe" {
  username = "admin"
  password = "password"
  host     = "192.168.1.1"
  protocol = "netconf"

  # NEW: Enable read cache (default: true)
  read_cache = true
}
```

Schema definition:

```go
"read_cache": schema.BoolAttribute{
    MarkdownDescription: "Cache the full device configuration on first read to improve performance. " +
        "Subsequent read operations query the cached configuration instead of the device. " +
        "Cache is invalidated after any write operation. " +
        "Only applies to NETCONF protocol. " +
        "This can also be set as the IOSXE_READ_CACHE environment variable. " +
        "Defaults to `true`.",
    Optional: true,
},
```

## Implementation Details

### Cache Population (Lazy Loading)

The cache is populated on the **first read operation** rather than at provider initialization:

```go
// GetCachedConfig returns cached config or fetches and caches the full config
func (device *IosxeProviderDataDevice) GetCachedConfig(ctx context.Context) (xmldot.Result, error) {
    // Fast path: return cached config if valid
    device.ReadCache.mu.RLock()
    if device.ReadCache.Valid && device.ReadCache.Config.Exists() {
        config := device.ReadCache.Config
        device.ReadCache.mu.RUnlock()
        return config, nil
    }
    device.ReadCache.mu.RUnlock()

    // Slow path: fetch and cache full config
    device.ReadCache.mu.Lock()
    defer device.ReadCache.mu.Unlock()

    // Double-check after acquiring write lock
    if device.ReadCache.Valid && device.ReadCache.Config.Exists() {
        return device.ReadCache.Config, nil
    }

    // Fetch full configuration from device
    res, err := device.NetconfClient.GetConfig(ctx, "running", netconf.NoFilter())
    if err != nil {
        return xmldot.Result{}, fmt.Errorf("failed to fetch full config for cache: %w", err)
    }

    // Cache the configuration
    device.ReadCache.Config = res.Res
    device.ReadCache.PopulatedAt = time.Now()
    device.ReadCache.Valid = true

    tflog.Debug(ctx, "NETCONF read cache populated",
        "configSize", len(res.Res.Raw),
        "populatedAt", device.ReadCache.PopulatedAt)

    return device.ReadCache.Config, nil
}
```

### Cache Invalidation

Cache is invalidated after any write operation completes:

```go
// InvalidateCache marks the cache as invalid
func (device *IosxeProviderDataDevice) InvalidateCache() {
    device.ReadCache.mu.Lock()
    defer device.ReadCache.mu.Unlock()

    device.ReadCache.Valid = false
    device.ReadCache.Config = xmldot.Result{}
}
```

Write operations (Create/Update/Delete) call `InvalidateCache()` after successful completion:

```go
// In resource Create/Update/Delete
if err := helpers.EditConfig(ctx, device.NetconfClient, body, device.AutoCommit); err != nil {
    resp.Diagnostics.AddError("Client Error", err.Error())
    return
}

// Invalidate cache after successful write
if device.ReadCacheEnabled {
    device.InvalidateCache()
}
```

### Helper Function for Cached Reads

New helper function in `helpers/netconf.go`:

```go
// GetConfigFromCache retrieves config for an XPath from cache or device
//
// When cache is enabled and valid:
//   - Returns data from cached full config using local XPath query
// When cache is disabled or invalid:
//   - Performs direct GetConfig to device with XPath filter
func GetConfigFromCache(
    ctx context.Context,
    device *IosxeProviderDataDevice,
    xPath string,
) (xmldot.Result, error) {
    // If cache disabled, do direct query
    if !device.ReadCacheEnabled {
        filter := GetXpathFilter(xPath)
        res, err := device.NetconfClient.GetConfig(ctx, "running", filter)
        if err != nil {
            return xmldot.Result{}, err
        }
        return res.Res, nil
    }

    // Get full cached config
    fullConfig, err := device.GetCachedConfig(ctx)
    if err != nil {
        return xmldot.Result{}, err
    }

    // Query the cached config with XPath
    return GetFromXPath(fullConfig, xPath), nil
}
```

### Template Changes

Update `gen/templates/resource.go` for the Read operation:

```go
// Current implementation
filter := helpers.GetXpathFilter(state.getXPath())
res, err := device.NetconfClient.GetConfig(ctx, "running", filter)

// New implementation with cache support
configResult, err := helpers.GetConfigFromCache(ctx, device, state.getXPath())
if err != nil {
    resp.Diagnostics.AddError("Client Error",
        fmt.Sprintf("Failed to retrieve object (%s), got error: %s", state.getPath(), err))
    return
}

// Check if config exists at XPath
if !configResult.Exists() && helpers.IsListPath(state.getXPath()) {
    tflog.Debug(ctx, fmt.Sprintf("%s: Resource does not exist", state.Id.ValueString()))
    resp.State.RemoveResource(ctx)
    return
}
```

## Behavioral Considerations

### Thread Safety

The cache uses `sync.RWMutex` to allow:
- Multiple concurrent reads (RLock)
- Exclusive writes for cache population and invalidation (Lock)

This aligns with Terraform's concurrent resource processing model.

### Memory Usage

The full device configuration is stored in memory. Typical IOS-XE configurations:
- Small deployment: 50KB - 500KB
- Medium deployment: 500KB - 5MB
- Large deployment: 5MB - 50MB

The cache stores an `xmldot.Result` which contains both the parsed XML structure and the raw XML string (accessible via `.Raw`).

For very large configurations, users can disable caching via `read_cache = false`.

### Cache Coherency

The cache reflects a point-in-time snapshot of device configuration:

1. **Within a Terraform run**: Cache is populated once, invalidated on writes, repopulated as needed
2. **Between Terraform runs**: Cache starts empty, repopulated on first read
3. **External changes**: Not reflected until cache is repopulated (invalidation or new run)

This behavior is acceptable because:
- Terraform expects to be the sole configuration manager
- Each plan/apply is a discrete operation
- Write operations invalidate cache, ensuring consistency

### Write Operation Impact

After any write operation, the cache is invalidated. If subsequent reads are needed:
- Cache is lazily repopulated on next read
- This ensures reads after writes see the updated configuration

Consider a scenario:
```
Read A → Cache populated
Read B → From cache
Write C → Cache invalidated
Read D → Cache repopulated (new GetConfig)
Read E → From cache
```

### Multi-Device Support

Each device in the provider configuration gets its own cache instance:

```go
for _, device := range config.Devices {
    // ... existing device setup ...

    data.Devices[device.Name.ValueString()] = &IosxeProviderDataDevice{
        // ... existing fields ...
        ReadCache: &NetconfReadCache{Valid: false},
        ReadCacheEnabled: readCacheEnabled,
    }
}
```

### Data Source Consideration

Data sources should also benefit from the read cache. The same `GetConfigFromCache` helper works for both resources and data sources.

### RESTCONF Protocol

The read cache only applies to NETCONF protocol. RESTCONF operations continue to work as before (individual GET requests). A future enhancement could add RESTCONF caching, but it's out of scope for this implementation.

## Configuration Options

### Provider Attribute

| Attribute | Type | Default | Description |
|-----------|------|---------|-------------|
| `read_cache` | bool | `false` | Enable/disable read cache for NETCONF operations (experimental) |

### Environment Variable

| Variable | Default | Description |
|----------|---------|-------------|
| `IOSXE_READ_CACHE` | `false` | Enable/disable read cache (provider attribute takes precedence) |

## Testing Strategy

### Unit Tests

1. **Cache population**: Verify cache is populated on first read
2. **Cache hit**: Verify subsequent reads use cache
3. **Cache invalidation**: Verify writes invalidate cache
4. **Cache repopulation**: Verify reads after invalidation repopulate cache
5. **Thread safety**: Concurrent read/write operations

### Integration Tests

1. **Performance comparison**: Measure time for plan with/without cache
2. **Correctness**: Verify cached reads match direct reads
3. **Write-then-read**: Verify configuration changes are reflected
4. **Multi-device**: Verify independent caches per device

### Acceptance Tests

Existing acceptance tests should pass without modification. The cache is transparent to resource logic.

## Performance Expectations

Estimated improvements (based on typical deployments):

| Scenario | Without Cache | With Cache | Improvement |
|----------|--------------|------------|-------------|
| 10 resources plan | ~10s | ~2s | 5x |
| 50 resources plan | ~50s | ~3s | 16x |
| 100 resources plan | ~100s | ~5s | 20x |

Actual improvements depend on:
- Network latency to device
- Device processing time
- Configuration complexity

## Rollout Plan

### Phase 1: Implementation
- Add cache structure to provider data
- Add provider configuration attribute
- Implement cache population and invalidation
- Update helper functions

### Phase 2: Template Updates
- Modify resource.go template for cached reads
- Modify data_source.go template for cached reads
- Regenerate all resources and data sources

### Phase 3: Testing
- Unit tests for cache logic
- Integration tests with real devices
- Performance benchmarking

### Phase 4: Documentation
- Update provider documentation
- Add migration guide (if needed)
- Document configuration options

## Open Questions

1. **Cache TTL**: Should we add a TTL for cache expiration? Currently, cache is valid until explicitly invalidated.

2. **Selective caching**: Should some resources bypass the cache? (e.g., frequently changing operational data)

3. **Cache warming**: Should we support eager cache population at provider initialization?

4. **Memory limits**: Should we add configuration for maximum cache size?

## Alternatives Considered

### 1. Per-Resource Caching
Cache individual resource configurations instead of full config.
- **Pro**: Lower memory usage
- **Con**: Still requires one GetConfig per unique resource type

### 2. RESTCONF Caching
Implement similar caching for RESTCONF protocol.
- **Pro**: Benefits RESTCONF users
- **Con**: Different implementation, HTTP caching semantics differ

### 3. Provider-Level Caching
Cache at the provider level across all devices.
- **Pro**: Simpler implementation
- **Con**: Would require device-specific keys, more complex invalidation

## Conclusion

The proposed read cache provides significant performance improvements for NETCONF-based Terraform operations by reducing device queries from N (one per resource) to 1 (one for all resources). The lazy loading and write-invalidation strategy ensures cache coherency while maintaining simplicity.

The implementation is:
- **Backward compatible**: Default behavior is caching enabled, can be disabled
- **Thread safe**: Uses appropriate locking for concurrent operations
- **Transparent**: No changes required to existing resource/data source definitions
- **Configurable**: Can be enabled/disabled via provider configuration or environment variable
