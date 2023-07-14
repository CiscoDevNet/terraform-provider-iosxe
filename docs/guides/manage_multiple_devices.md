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
  username = "admin"
  password = "Cisco123"
  url      = "https://10.1.1.1"
}

provider "iosxe" {
  alias    = "ROUTER-2"
  username = "admin"
  password = "Cisco123"
  url      = "https://10.1.1.2"
}

resource "iosxe_restconf" "ROUTER-1" {
  provider   = iosxe.ROUTER-1
  path       = "openconfig-system:system/config"
  attributes = {
    hostname = "ROUTER-1"
  }
}

resource "iosxe_restconf" "ROUTER-2" {
  provider   = iosxe.ROUTER-2
  path       = "openconfig-system:system/config"
  attributes = {
    hostname = "ROUTER-2"
  }
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
  username = "admin"
  password = "Cisco123"
  devices  = local.routers
}

resource "iosxe_restconf" "hostname" {
  for_each   = toset([for router in local.routers : router.name])
  device     = each.key
  path       = "openconfig-system:system/config"
  attributes = {
    hostname = each.key
  }
}
```
