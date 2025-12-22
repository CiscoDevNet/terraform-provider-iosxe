---
subcategory: "Guides"
page_title: "NETCONF"
description: |-
    Howto use NETCONF with the IOSXE provider.
---

# NETCONF

## Overview

The IOSXE provider uses **NETCONF** (SSH-based) as the default protocol for device communication. NETCONF support was introduced in version 0.10.0 and became the default in version 0.15.0.

### Key NETCONF Features

- **Transactional Commits**: Stage multiple configuration changes in the candidate datastore and commit them atomically
- **Connection Reuse**: Keep SSH connections open between operations for improved performance
- **Standardized Protocol**: Industry-standard NETCONF protocol (RFC 6241) over SSH
- **Better Error Handling**: More detailed error messages and validation

## Device Configuration

### Enabling NETCONF on IOS-XE

To use the NETCONF protocol, enable it on your IOS-XE device:

```
configure terminal
netconf-yang
```

### Optional: Candidate Datastore for Transactional Commits

For transactional commits using the candidate datastore, also enable the candidate datastore feature:

```
netconf-yang feature candidate-datastore
```

The candidate datastore allows you to stage multiple configuration changes and commit them atomically as a single transaction. Without this feature, the provider will edit the running configuration directly.

### Verifying NETCONF Capabilities

You can verify that NETCONF is enabled and check supported capabilities:

```
show platform software yang-management process
show netconf-yang sessions
```

## Provider Configuration

### Basic NETCONF Configuration

NETCONF is the default protocol, so you can use the provider without explicitly specifying `protocol`:

```terraform
provider "iosxe" {
  username = "admin"
  password = "password"
  host     = "10.1.1.1"
  # protocol = "netconf"  # Optional - this is the default
}
```

### Provider Attributes

The following provider attributes control NETCONF behavior:

| Attribute | Description | Default | Environment Variable |
|-----------|-------------|---------|----------------------|
| `host` | Device hostname/IP address. Optionally add `:port` (default: 830) | Required | `IOSXE_HOST` |
| `username` | Device username | Required | `IOSXE_USERNAME` |
| `password` | Device password | Required | `IOSXE_PASSWORD` |
| `insecure` | Skip SSH host key verification | `true` | `IOSXE_INSECURE` |
| `retries` | Number of retries for operations | `3` | `IOSXE_RETRIES` |
| `lock_release_timeout` | Seconds to wait for datastore lock release | `120` | `IOSXE_LOCK_RELEASE_TIMEOUT` |
| `reuse_connection` | Keep SSH connections open between operations | `true` | `IOSXE_REUSE_CONNECTION` |
| `auto_commit` | Automatically commit changes after each operation | `true` | `IOSXE_AUTO_COMMIT` |

### Port Configuration

The default NETCONF port is **830** (SSH). You can specify a custom port in the `host` attribute:

```terraform
provider "iosxe" {
  host = "10.1.1.1:2830"  # Custom NETCONF port
}
```

## Candidate Datastore Support

### How Candidate Datastore Works

When your device supports the candidate datastore capability (enabled with `netconf-yang feature candidate-datastore`), the provider uses a transactional workflow:

1. **Lock** both the running and candidate datastores
2. **Edit** the candidate configuration
3. **Validate** the candidate configuration (implicit)
4. **Commit** candidate to running configuration (when `auto_commit=true` or using `iosxe_commit`)
5. **Unlock** both datastores

This provides atomic commits where all changes succeed or fail together, preventing partial configuration states.

### Without Candidate Datastore

If the device does not support the candidate datastore capability, the provider falls back to editing the running datastore directly:

1. **Lock** the running datastore
2. **Edit** the running configuration directly
3. **Unlock** the running datastore

## Configuration Dependencies

### Manual Commit Requires Connection Reuse

When using manual commit mode (`auto_commit = false`), you **must** also enable connection reuse (`reuse_connection = true`, which is the default).

**Why?** Manual commit mode stages configuration changes in the candidate datastore across multiple resource operations. Without connection reuse, each resource would open a new connection, losing access to previously staged changes.

```terraform
# ✅ Valid Configuration
provider "iosxe" {
  protocol         = "netconf"
  auto_commit      = false
  reuse_connection = true  # Required (or omit for default true)
}

# ❌ Invalid Configuration - Provider will return an error
provider "iosxe" {
  protocol         = "netconf"
  auto_commit      = false
  reuse_connection = false  # Error: manual commit requires reuse
}
```

The provider will validate this configuration and return an error if you attempt to use manual commit mode without connection reuse.

## Configuration Commit Modes

### Auto-Commit Mode (Default)

With `auto_commit = true` (default), each resource commits its changes immediately after applying them to the candidate datastore:

```terraform
provider "iosxe" {
  host        = "10.1.1.1"
  protocol    = "netconf"
  auto_commit = true  # Default behavior
}

resource "iosxe_interface_loopback" "example" {
  name = "100"
  description = "Managed by Terraform"
}
# Changes are committed automatically after this resource is created/updated
```

**Use auto-commit when:**
- You want immediate configuration changes
- Each resource represents an independent configuration change
- You prefer simplicity over transactional grouping

### Manual Commit Mode

With `auto_commit = false`, changes are staged in the candidate datastore and must be explicitly committed using the `iosxe_commit` resource:

```terraform
provider "iosxe" {
  host        = "10.1.1.1"
  protocol    = "netconf"
  auto_commit = false  # Disable automatic commits
}

resource "iosxe_interface_loopback" "loopback100" {
  name = "100"
  description = "Loopback 100"
}

resource "iosxe_interface_loopback" "loopback101" {
  name = "101"
  description = "Loopback 101"
}

# Explicitly commit all staged changes as a single transaction
resource "iosxe_commit" "commit_changes" {
  save_config = true  # Optional: also save to startup-config
  depends_on = [
    iosxe_interface_loopback.loopback100,
    iosxe_interface_loopback.loopback101,
  ]
}
```

**Use manual commit when:**
- You want to group multiple configuration changes into a single atomic transaction
- You need to ensure all related changes are committed together or not at all
- You want to stage configuration changes before committing them
- You're implementing complex configuration workflows

## The iosxe_commit Resource

### Purpose

The `iosxe_commit` resource explicitly commits the candidate configuration to the running configuration. It is only meaningful when:
- Using the NETCONF protocol (`protocol = "netconf"`)
- Auto-commit is disabled (`auto_commit = false`)
- The device supports the candidate datastore capability

### Usage

```terraform
resource "iosxe_commit" "example" {
  # Optional: specify device name if using multi-device configuration
  # device = "device1"

  # Optional: save to startup-config after commit
  # save_config = true

  # The commit attribute is managed internally
  depends_on = [
    # List all resources that should be staged before committing
  ]
}
```

### Attributes

- `device` (Optional): Device name from the provider's `devices` list (for multi-device scenarios)
- `commit` (Computed): Internal attribute used to track commit state
- `save_config` (Optional): Save running configuration to startup configuration after commit. Equivalent to 'copy running-config startup-config'. Defaults to `false`

### Saving Configuration to Startup

The `iosxe_commit` resource can optionally save the running configuration to startup configuration after committing, eliminating the need for a separate `iosxe_save_config` resource:

```terraform
resource "iosxe_interface_loopback" "loopback100" {
  name        = "100"
  description = "Loopback 100"
}

resource "iosxe_interface_loopback" "loopback101" {
  name        = "101"
  description = "Loopback 101"
}

# Commit and save in a single atomic operation
resource "iosxe_commit" "commit_and_save" {
  save_config = true  # Saves to startup-config after commit
  depends_on = [
    iosxe_interface_loopback.loopback100,
    iosxe_interface_loopback.loopback101,
  ]
}
```

This combines commit and save operations in a single transaction, ensuring your committed changes persist across device reboots.

### Important Notes

1. **No effect with auto_commit=true**: If `auto_commit` is enabled, the commit resource has no practical effect since changes are already committed automatically
2. **Requires candidate datastore**: The commit resource only works when the device supports the candidate datastore capability
3. **Use depends_on**: Always use `depends_on` to ensure resources are created/updated before committing
4. **Destroy Behavior**: When the `iosxe_commit` resource is destroyed (e.g., during `terraform destroy`), it automatically enables auto-commit mode for subsequent resource deletions. This ensures that deletion operations commit their changes and prevents uncommitted configuration from remaining in the candidate datastore
5. **Saving Configuration**: Set `save_config = true` to persist changes to startup-config in the same transaction, eliminating the need for a separate `iosxe_save_config` resource

## Migration to NETCONF (from versions < 0.15.0)

If you're upgrading from a version before 0.15.0 where RESTCONF was the default:

### Option 1: Adopt NETCONF (Recommended)

1. **Enable NETCONF on your devices**:
   ```
   configure terminal
   netconf-yang
   ```

   Optionally, enable the candidate datastore for transactional commits:
   ```
   netconf-yang feature candidate-datastore
   ```

2. **Remove explicit `protocol = "restconf"` from your configuration** (if present)

3. **Test the migration**:
   ```bash
   terraform plan
   ```

4. **Verify no resources need replacement** - the provider protocol change should not require recreating resources

### Option 2: Continue Using RESTCONF

If you need to continue using RESTCONF:

1. **Explicitly set `protocol = "restconf"` in your provider configuration**:
   ```terraform
   provider "iosxe" {
     username = "admin"
     password = "password"
     host     = "10.1.1.1"
     protocol = "restconf"  # Maintain RESTCONF behavior
   }
   ```

2. **No other changes required** - your existing configurations will continue to work

## Additional Resources

- [RFC 6241 - Network Configuration Protocol (NETCONF)](https://tools.ietf.org/html/rfc6241)
- [go-netconf Library Documentation](https://github.com/netascode/go-netconf)
- [IOS-XE Programmability Guide](https://www.cisco.com/c/en/us/td/docs/ios-xml/ios/prog/configuration/1715/b_1715_programmability_cg.html)
