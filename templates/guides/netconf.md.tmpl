---
subcategory: "Guides"
page_title: "NETCONF"
description: |-
    Howto use NETCONF with the IOSXE provider.
---

# NETCONF

## Overview

The IOSXE provider supports both **RESTCONF** (HTTPS-based) and **NETCONF** (SSH-based) protocols for device communication. NETCONF support was introduced in version 0.10.0 and provides additional capabilities, particularly for transactional configuration management using the candidate datastore.

### Key Benefits of NETCONF

- **Transactional Commits**: Stage multiple configuration changes in the candidate datastore and commit them atomically
- **Connection Reuse**: Keep SSH connections open between operations for improved performance
- **Standardized Protocol**: Industry-standard NETCONF protocol (RFC 6241) over SSH

## Device Configuration

### Enabling NETCONF on IOS-XE

To use the NETCONF protocol, you must enable it on your IOS-XE device:

```
configure terminal
netconf-yang
netconf-yang feature candidate-datastore
```

The `candidate-datastore` feature enables the candidate configuration datastore, which allows you to stage configuration changes before committing them to the running configuration.

### Verifying NETCONF Capabilities

You can verify that NETCONF is enabled and check supported capabilities:

```
show platform software yang-management process
show netconf-yang sessions
```

## Provider Configuration

### Basic NETCONF Configuration

Configure the provider to use NETCONF protocol:

```terraform
provider "iosxe" {
  username = "admin"
  password = "password"
  host     = "10.1.1.1"
  protocol = "netconf"
}
```

### Provider Attributes

The following provider attributes control NETCONF behavior:

| Attribute | Description | Default | Environment Variable |
|-----------|-------------|---------|----------------------|
| `protocol` | Protocol to use: `restconf` or `netconf` | `restconf` | `IOSXE_PROTOCOL` |
| `host` | Device hostname/IP address. Optionally add `:port` | Required | `IOSXE_HOST` |
| `username` | Device username | Required | `IOSXE_USERNAME` |
| `password` | Device password | Required | `IOSXE_PASSWORD` |
| `insecure` | Skip SSH host key verification | `true` | `IOSXE_INSECURE` |
| `retries` | Number of retries for operations | `3` (NETCONF) | `IOSXE_RETRIES` |
| `lock_release_timeout` | Seconds to wait for datastore lock release | `120` | `IOSXE_LOCK_RELEASE_TIMEOUT` |
| `reuse_connection` | Keep SSH connections open between operations | `true` | `IOSXE_REUSE_CONNECTION` |
| `auto_commit` | Automatically commit changes after each operation | `true` | `IOSXE_AUTO_COMMIT` |

### Port Configuration

- **Default NETCONF Port**: 830 (SSH)
- **Default RESTCONF Port**: 443 (HTTPS)

You can specify a custom port in the `host` attribute:

```terraform
provider "iosxe" {
  host     = "10.1.1.1:2830"  # Custom NETCONF port
  protocol = "netconf"
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

  # The commit attribute is managed internally
  depends_on = [
    # List all resources that should be staged before committing
  ]
}
```

### Attributes

- `device` (Optional): Device name from the provider's `devices` list (for multi-device scenarios)
- `commit` (Computed): Internal attribute used to track commit state

### Important Notes

1. **No effect with auto_commit=true**: If `auto_commit` is enabled, the commit resource has no practical effect since changes are already committed automatically
2. **No effect with RESTCONF**: The commit resource only works with NETCONF protocol
3. **Use depends_on**: Always use `depends_on` to ensure resources are created/updated before committing
4. **Lifecycle Management**: The commit resource's create and update operations trigger a commit; destroy is a no-op

## Migration from RESTCONF to NETCONF

### Step 1: Enable NETCONF on Device

```
configure terminal
netconf-yang
netconf-yang feature candidate-datastore
```

### Step 2: Update Provider Configuration

```terraform
provider "iosxe" {
  # username = "admin"  # Unchanged
  # password = "password"  # Unchanged

  # Change from deprecated 'url' to 'host'
  # url      = "https://10.1.1.1"  # Old
  host     = "10.1.1.1"            # New

  # Add protocol configuration
  protocol = "netconf"
}
```

### Step 3: Test Configuration

Run `terraform plan` to verify the migration doesn't require resource replacement:

```bash
terraform plan
```

### Step 4: Optional - Enable Manual Commits

If you want transactional commits:

```terraform
provider "iosxe" {
  host        = "10.1.1.1"
  protocol    = "netconf"
  auto_commit = false
}

# Add iosxe_commit resources where needed
```

## Comparison: NETCONF vs RESTCONF

| Feature | NETCONF | RESTCONF |
|---------|---------|----------|
| **Protocol** | SSH (Port 830) | HTTPS (Port 443) |
| **Data Format** | XML | JSON/XML |
| **Transaction Support** | Yes (candidate datastore) | No |
| **Connection Reuse** | Yes (configurable) | N/A (stateless HTTP) |
| **Atomic Commits** | Yes | No |
| **Default Retries** | 3 | 10 |
| **Use Case** | Transactional changes, complex workflows | Simple operations, REST-friendly environments |

## Additional Resources

- [RFC 6241 - Network Configuration Protocol (NETCONF)](https://tools.ietf.org/html/rfc6241)
- [go-netconf Library Documentation](https://github.com/netascode/go-netconf)
- [IOS-XE Programmability Guide](https://www.cisco.com/c/en/us/td/docs/ios-xml/ios/prog/configuration/1715/b_1715_programmability_cg.html)
