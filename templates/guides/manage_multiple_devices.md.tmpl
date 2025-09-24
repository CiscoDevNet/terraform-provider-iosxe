---
subcategory: "Guides"
page_title: "Manage Multiple Devices"
description: |-
    Howto manage multiple devices.
---

# Manage Multiple Devices

When it comes to managing multiple IOS-XE devices, one can create multiple provider configurations and distinguish them by `alias` ([documentation](https://www.terraform.io/language/providers/configuration#alias-multiple-provider-configurations)).

```terraform
provider "iosxe" {
  alias    = "ROUTER-1"
  url      = "https://10.1.1.1"
}

provider "iosxe" {
  alias    = "ROUTER-2"
  url      = "https://10.1.1.2"
}

resource "iosxe_system" "ROUTER-1" {
  provider = iosxe.ROUTER-1
  hostname = "ROUTER-1"
}

resource "iosxe_system" "ROUTER-2" {
  provider = iosxe.ROUTER-2
  hostname = "ROUTER-2"
}
```

The disadvantages here is that the `provider` attribute of resources cannot be dynamic and therefore cannot be used in combination with `for_each` as an example. The issue is being tracked [here](https://github.com/hashicorp/terraform/issues/24476).

This provider offers an alternative approach where mutliple devices can be managed by a single provider configuration and the optional `device` attribute, which is available in every resource and data source, can then be used to select the respective device. This assumes that every device uses the same credentials.

```terraform
locals {
  routers = [
    {
      name = "ROUTER-1"
      url  = "https://10.1.1.1"
    },
    {
      name = "ROUTER-2"
      url  = "https://10.1.1.2"
    },
  ]
}

provider "iosxe" {
  devices  = local.routers
}

resource "iosxe_system" "hostname" {
  for_each = toset([for router in local.routers : router.name])
  device   = each.key
  hostname = each.key
}
```

## Device-Level Management Control

### The "managed" Flag

Each device in the `devices` list supports an optional `managed` attribute that controls whether the device is actively managed by Terraform:

- **`managed = true`** (default): Device is actively managed - configurations are applied
- **`managed = false`**: Device is "frozen" - state preserved but no changes applied
- **Not specified**: Defaults to `true`

### Basic Managed Flag Usage

```hcl
locals {
  devices = [
    {
      name    = "production-sw01"
      url     = "https://10.1.1.10"
      managed = true   # Actively managed
    },
    {
      name    = "maintenance-sw02"
      url     = "https://10.1.1.20"
      managed = false  # Temporarily frozen for maintenance
    },
    {
      name    = "active-sw03"
      url     = "https://10.1.1.30"
      # managed defaults to true when not specified
    }
  ]
}

provider "iosxe" {
  devices  = local.devices
}

resource "iosxe_banner" "login_banner" {
  for_each = toset([for device in local.devices : device.name])
  device   = each.key
  login    = "Authorized Access Only - ${each.key}"
}
```

**Result**:
- `production-sw01` and `active-sw03`: Banner configuration applied
- `maintenance-sw02`: Configuration skipped, existing state preserved

### Relationship with `selected_devices`

**Important**: The [`selected_devices` provider attribute](selective_deploy.md) takes precedence over individual `managed` flags:

- **When `selected_devices` is specified**: Individual `managed` flags are ignored
- **When `selected_devices` is not specified**: Individual `managed` flags are respected

#### Example: selected_devices Override
```hcl
provider "iosxe" {
  selected_devices = ["switch-01", "switch-03"]  # Only these devices managed
  devices = [
    { name = "switch-01", url = "https://10.1.1.10", managed = false },  # Overridden to managed=true
    { name = "switch-02", url = "https://10.1.1.20", managed = true },   # Overridden to managed=false
    { name = "switch-03", url = "https://10.1.1.30", managed = true }    # Remains managed=true
  ]
}
```

**Result**: Only `switch-01` and `switch-03` are managed, regardless of their individual `managed` flags.
