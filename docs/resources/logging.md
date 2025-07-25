---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxe_logging Resource - terraform-provider-iosxe"
subcategory: "Management"
description: |-
  This resource can manage the Logging configuration.
---

# iosxe_logging (Resource)

This resource can manage the Logging configuration.

## Example Usage

```terraform
resource "iosxe_logging" "example" {
  monitor_severity  = "informational"
  buffered_size     = 16000
  buffered_severity = "informational"
  console_severity  = "informational"
  facility          = "local0"
  history_size      = 100
  history_severity  = "informational"
  trap              = true
  trap_severity     = "informational"
  origin_id_type    = "hostname"
  source_interface  = "Loopback0"
  source_interfaces_vrf = [
    {
      vrf            = "VRF1"
      interface_name = "Loopback100"
    }
  ]
  ipv4_hosts = [
    {
      ipv4_host = "1.1.1.1"
    }
  ]
  ipv4_vrf_hosts = [
    {
      ipv4_host = "1.1.1.1"
      vrf       = "VRF1"
    }
  ]
  ipv6_hosts = [
    {
      ipv6_host = "2001::1"
    }
  ]
  ipv6_vrf_hosts = [
    {
      ipv6_host = "2001::1"
      vrf       = "VRF1"
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `buffered_severity` (String) DEPRECATED. Logging severity level
- `buffered_size` (Number) DEPRECATED. Logging buffer size
  - Range: `4096`-`2147483647`
- `console_severity` (String)
- `device` (String) A device name from the provider configuration.
- `facility` (String) Facility parameter for syslog messages
  - Choices: `auth`, `cron`, `daemon`, `kern`, `local0`, `local1`, `local2`, `local3`, `local4`, `local5`, `local6`, `local7`, `lpr`, `mail`, `news`, `sys10`, `sys11`, `sys12`, `sys13`, `sys14`, `sys9`, `syslog`, `user`, `uucp`
- `file_max_size` (Number) - Range: `0`-`4294967295`
- `file_min_size` (Number) - Range: `0`-`4294967295`
- `file_name` (String)
- `file_severity` (String)
- `history_severity` (String)
- `history_size` (Number) Set history table size
  - Range: `0`-`65535`
- `ipv4_hosts` (Attributes List) (see [below for nested schema](#nestedatt--ipv4_hosts))
- `ipv4_vrf_hosts` (Attributes List) (see [below for nested schema](#nestedatt--ipv4_vrf_hosts))
- `ipv6_hosts` (Attributes List) (see [below for nested schema](#nestedatt--ipv6_hosts))
- `ipv6_vrf_hosts` (Attributes List) (see [below for nested schema](#nestedatt--ipv6_vrf_hosts))
- `monitor_severity` (String)
- `origin_id_name` (String) Define a unique text string as ID
- `origin_id_type` (String) Use origin hostname/ip/ipv6 as ID
  - Choices: `hostname`, `ip`, `ipv6`
- `source_interface` (String)
- `source_interfaces_vrf` (Attributes List) Specify interface and vrf for source address in logging transactions (see [below for nested schema](#nestedatt--source_interfaces_vrf))
- `trap` (Boolean) Set trap server logging level
- `trap_severity` (String)

### Read-Only

- `id` (String) The path of the object.

<a id="nestedatt--ipv4_hosts"></a>
### Nested Schema for `ipv4_hosts`

Required:

- `ipv4_host` (String)


<a id="nestedatt--ipv4_vrf_hosts"></a>
### Nested Schema for `ipv4_vrf_hosts`

Required:

- `ipv4_host` (String)
- `vrf` (String) Set VRF option


<a id="nestedatt--ipv6_hosts"></a>
### Nested Schema for `ipv6_hosts`

Required:

- `ipv6_host` (String)


<a id="nestedatt--ipv6_vrf_hosts"></a>
### Nested Schema for `ipv6_vrf_hosts`

Required:

- `ipv6_host` (String)
- `vrf` (String) Set VRF option


<a id="nestedatt--source_interfaces_vrf"></a>
### Nested Schema for `source_interfaces_vrf`

Required:

- `interface_name` (String)
- `vrf` (String) Specify the vrf of source interface for logging transactions

## Import

Import is supported using the following syntax:

```shell
terraform import iosxe_logging.example "Cisco-IOS-XE-native:native/logging"
```
