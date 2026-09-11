---
subcategory: "Guides"
page_title: "Bulk Resources"
description: |-
    Bulk Resources
---

# Bulk Resources

Bulk resources manage every instance of a configuration object on a device from a single resource block. Instead of one Terraform resource per interface, VLAN or neighbor, a bulk resource takes a map of items and manages them collectively.

The main benefit is the number of NETCONF operations. A regular resource issues one `get-config` per object during a refresh and one `edit-config` per object during an apply. A bulk resource issues exactly one of each, no matter how many items it manages. On a 48-port switch this turns 48 round trips into one.

->Bulk resources are additional to, not a replacement for, the regular per-object resources. `iosxe_interface_ethernet` and `iosxe_interface_ethernets` both exist and you can pick whichever fits your configuration. Do not manage the same object with both.

## How Bulk Resources Work

All items of a bulk resource live below a common parent path in the YANG model. `iosxe_interface_ethernets`, for example, manages entries below `/Cisco-IOS-XE-native:native/interface`.

- **Read**: a single `get-config` with an XPath filter on the parent path returns every item at once.
- **Create and update**: every item is serialized into one `edit-config` payload as a sibling element below the parent, together with the delete operations for any item or attribute that was removed.
- **Delete**: one `edit-config` removes all managed items. For objects that cannot be deleted from the device, such as physical interfaces, only the attributes that were explicitly configured are removed, exactly as the corresponding regular resource does.

Because a bulk resource is a single Terraform resource, all of its items also share one `device` attribute. To manage the same object type on several devices, combine a bulk resource with `for_each` over your devices, as described in [Manage Multiple Devices](manage_multiple_devices).

## Example Usage

```terraform
resource "iosxe_interface_ethernets" "example" {
  items = {
    "GigabitEthernet;1" = {
      description = "Uplink"
      mtu         = 9000
      shutdown    = false
    }
    "GigabitEthernet;2" = {
      description       = "Workstation 1"
      ipv4_address      = "10.1.1.1"
      ipv4_address_mask = "255.255.255.0"
    }
    "GigabitEthernet;3" = {
      description = "Workstation 2"
    }
    "TenGigabitEthernet;1/0/1" = {
      description = "Core"
    }
  }
}
```

Across multiple devices:

```terraform
locals {
  switches = [
    { name = "SWITCH-1", host = "10.1.1.1" },
    { name = "SWITCH-2", host = "10.1.1.2" },
  ]
}

provider "iosxe" {
  devices = local.switches
}

resource "iosxe_interface_ethernets" "access_ports" {
  for_each = toset([for switch in local.switches : switch.name])
  device   = each.key

  items = {
    for id in range(1, 25) : "GigabitEthernet;1/0/${id}" => {
      description = "Access port ${id}"
      shutdown    = false
    }
  }
}
```

## The Item Map Key

The `items` collection is a map. Its key identifies an item and is built from the attributes that identify the object on the device, joined by a semicolon (`;`). Those attributes are therefore not part of the item body itself, they are only present in the key.

For `iosxe_interface_ethernets` the key is composed of the interface `type` and `name`, so `"GigabitEthernet;3"` addresses `GigabitEthernet3`. The documentation of every bulk resource lists the attributes that make up its key, and the valid choices where applicable.

Using a map rather than a list keeps plans readable and stable: items are matched by key, so adding, removing or reordering items only shows the items that actually changed.

```
Terraform will perform the following actions:

  # iosxe_interface_ethernets.example will be updated in-place
  ~ resource "iosxe_interface_ethernets" "example" {
        id    = "Cisco-IOS-XE-native:native/interface"
      ~ items = {
          ~ "GigabitEthernet;2" = {
              ~ description = "Workstation 1" -> "Workstation 5"
                # (2 unchanged attributes hidden)
            },
            # (3 unchanged elements hidden)
        }
    }

Plan: 0 to add, 1 to change, 0 to destroy.
```

Removing an item from the map removes that item's configuration from the device on the next apply. For objects that cannot be deleted, such as physical interfaces, only the attributes that were managed by the resource are removed and the object itself stays in place.

## Importing Bulk Resources

Importing a bulk resource reads every existing item below the parent path into state, not just the ones present in your configuration. On a device with 48 interfaces, importing `iosxe_interface_ethernets` yields 48 items.

```shell
terraform import iosxe_interface_ethernets.example ""
```

When managing multiple devices, append the device name:

```shell
terraform import iosxe_interface_ethernets.example "SWITCH-1"
```

->After importing, run `terraform plan` and reconcile your configuration with the imported state before applying. Any item present in state but missing from your configuration will be removed from the device on the next apply.
