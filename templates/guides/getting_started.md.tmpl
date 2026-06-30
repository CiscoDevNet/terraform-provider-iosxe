---
subcategory: "Guides"
page_title: "Getting Started"
description: |-
    Get started with the IOSXE provider by configuring a device end to end.
---

# Getting Started

This guide walks through a minimal end-to-end example: configuring the system
hostname, creating a loopback interface, and adding a VRF on a single Cisco
IOS-XE device. Each resource shown here is documented in full detail in the
sidebar.

## Prerequisite: Enable NETCONF

The IOSXE provider communicates with devices over **NETCONF** (SSH). Before
Terraform can manage a device, enable `netconf-yang` on it:

```
configure terminal
netconf-yang
```

Optionally enable the candidate datastore for transactional commits. See the
[NETCONF guide](netconf) for details.

## Provider Configuration

First, declare the provider and point it at your device:

```terraform
terraform {
  required_providers {
    iosxe = {
      source = "CiscoDevNet/iosxe"
    }
  }
}

provider "iosxe" {
  username = "admin"
  password = "password"
  host     = "10.1.1.1"
}
```

## System Hostname

Next, configure the device hostname with the `iosxe_system` resource:

```terraform
resource "iosxe_system" "system" {
  hostname = "ROUTER-1"
}
```

## Loopback Interface

Now add a loopback interface with the `iosxe_interface_loopback` resource:

```terraform
resource "iosxe_interface_loopback" "loopback100" {
  name              = 100
  description       = "Managed by Terraform"
  ipv4_address      = "10.0.0.1"
  ipv4_address_mask = "255.255.255.255"
}
```

## VRF

And finally, create a VRF with the `iosxe_vrf` resource:

```terraform
resource "iosxe_vrf" "vrf1" {
  name                = "VRF1"
  description         = "Managed by Terraform"
  address_family_ipv4 = true
}
```

## Deploy the Configuration

Initialize the working directory to download the provider:

```shell
terraform init
```

Review the planned changes:

```shell
terraform plan
```

Apply the configuration to the device. Terraform prompts for confirmation;
type `yes` to proceed:

```shell
terraform apply
```

## Next Steps

- [NETCONF](netconf) — transactional commits and the candidate datastore
- [Manage Multiple Devices](manage_multiple_devices) — drive many devices from a single configuration
- [Importing Resources](importing_resources) — bring existing device configuration under Terraform management
