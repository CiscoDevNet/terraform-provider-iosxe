// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Package cache provides caching for NETCONF read operations.
//
// The cache stores the complete device configuration retrieved via NETCONF GetConfig
// and allows subsequent read operations to query the cached data instead of making
// additional NETCONF requests. This significantly improves performance for Terraform
// plans and applies with many resources.
//
// Cache Lifecycle:
//   - Cache is created per-device during provider configuration
//   - Cache is populated on first read operation via GetConfig with NoFilter()
//   - Subsequent reads query the cached XML using XPath
//   - Cache is invalidated after any write operation (Create/Update/Delete)
//   - Next read after invalidation repopulates the cache
//
// Thread Safety:
//   - All operations are protected by sync.RWMutex
//   - Multiple concurrent reads are allowed
//   - Writes (populate, invalidate) are exclusive
package cache

import (
	"context"
	"sync"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-netconf"
	"github.com/netascode/xmldot"
)

// NetconfReadCache provides caching for NETCONF GetConfig operations.
// It stores the complete device configuration and allows filtering via XPath.
//
// Thread-safe: All operations protected by RWMutex.
// Lifecycle: Cache is populated on first read, invalidated after writes.
type NetconfReadCache struct {
	mu          sync.RWMutex
	enabled     bool
	populated   bool
	configData  xmldot.Result // Cached full configuration (the <data> element content)
	populatedAt time.Time     // When cache was last populated
}

// New creates a new NetconfReadCache instance.
//
// Parameters:
//   - enabled: Whether caching is enabled for this device
//
// Returns:
//   - *NetconfReadCache: New cache instance (never nil)
func New(enabled bool) *NetconfReadCache {
	return &NetconfReadCache{
		enabled: enabled,
	}
}

// IsEnabled returns whether caching is enabled for this device.
func (c *NetconfReadCache) IsEnabled() bool {
	return c.enabled
}

// IsPopulated returns whether the cache contains data.
// Thread-safe: Uses read lock.
func (c *NetconfReadCache) IsPopulated() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.populated
}

// PopulatedAt returns when the cache was last populated.
// Returns zero time if cache is not populated.
// Thread-safe: Uses read lock.
func (c *NetconfReadCache) PopulatedAt() time.Time {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.populatedAt
}

// Populate fetches the full device configuration and stores it in the cache.
// This is called on the first read operation when cache is enabled but not populated.
//
// Parameters:
//   - ctx: Context for logging and cancellation
//   - client: NETCONF client to fetch configuration from
//
// Returns:
//   - xmldot.Result: The cached configuration data (can be used directly without calling Get)
//   - error: Any error from NETCONF GetConfig operation
//
// Thread-safe: Uses write lock.
// Idempotent: Safe to call multiple times (subsequent calls return cached data).
func (c *NetconfReadCache) Populate(ctx context.Context, client *netconf.Client) (xmldot.Result, error) {
	if !c.enabled {
		return xmldot.Result{}, nil
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	// Double-check under lock to avoid race condition
	if c.populated {
		tflog.Debug(ctx, "NETCONF read cache already populated, returning cached data")
		return c.configData, nil
	}

	tflog.Debug(ctx, "Populating NETCONF read cache with full device configuration")

	// Fetch full configuration without filter
	res, err := client.GetConfig(ctx, "running", netconf.NoFilter())
	if err != nil {
		tflog.Error(ctx, "Failed to populate NETCONF read cache", map[string]any{
			"error": err.Error(),
		})
		return xmldot.Result{}, err
	}

	// Store the response - we keep the full rpc-reply structure
	// so that existing fromBodyXML/updateFromBodyXML methods work unchanged
	c.configData = res.Res
	c.populated = true
	c.populatedAt = time.Now()

	// Log cache size for diagnostics
	configSize := len(res.Res.Raw)
	tflog.Debug(ctx, "NETCONF read cache populated successfully", map[string]any{
		"size_bytes": configSize,
	})

	return c.configData, nil
}

// Get retrieves the cached configuration data.
// The returned xmldot.Result contains the full rpc-reply structure,
// allowing callers to use GetFromXPath with "data" prefix as usual.
//
// Parameters:
//   - ctx: Context for logging
//
// Returns:
//   - xmldot.Result: Cached configuration data (full rpc-reply)
//   - bool: Whether data was retrieved from cache (true) or cache miss (false)
//
// Thread-safe: Uses read lock.
func (c *NetconfReadCache) Get(ctx context.Context) (xmldot.Result, bool) {
	if !c.enabled {
		return xmldot.Result{}, false
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	if !c.populated {
		tflog.Debug(ctx, "NETCONF read cache miss (not populated)")
		return xmldot.Result{}, false
	}

	tflog.Debug(ctx, "NETCONF read cache hit")
	return c.configData, true
}

// Invalidate clears the cache.
// Called after any write operation (Create/Update/Delete) to ensure fresh reads.
//
// Parameters:
//   - ctx: Context for logging
//
// Thread-safe: Uses write lock.
// Idempotent: Safe to call when cache is already empty.
func (c *NetconfReadCache) Invalidate(ctx context.Context) {
	if !c.enabled {
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.populated {
		return // Already empty, nothing to do
	}

	tflog.Debug(ctx, "Invalidating NETCONF read cache after write operation")
	c.configData = xmldot.Result{}
	c.populated = false
	c.populatedAt = time.Time{}
}
