---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxe_logging Data Source - terraform-provider-iosxe"
subcategory: "Management"
description: |-
  This data source can read the Logging configuration.
---

# iosxe_logging (Data Source)

This data source can read the Logging configuration.

## Example Usage

```terraform
data "iosxe_logging" "example" {
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `buffered_severity` (String) DEPRECATED. Logging severity level
- `buffered_size` (Number) DEPRECATED. Logging buffer size
- `console_severity` (String)
- `facility` (String) Facility parameter for syslog messages
- `file_max_size` (Number)
- `file_min_size` (Number)
- `file_name` (String)
- `file_severity` (String)
- `history_severity` (String)
- `history_size` (Number) Set history table size
- `id` (String) The path of the retrieved object.
- `ipv4_hosts` (Attributes List) (see [below for nested schema](#nestedatt--ipv4_hosts))
- `ipv4_vrf_hosts` (Attributes List) (see [below for nested schema](#nestedatt--ipv4_vrf_hosts))
- `ipv6_hosts` (Attributes List) (see [below for nested schema](#nestedatt--ipv6_hosts))
- `ipv6_vrf_hosts` (Attributes List) (see [below for nested schema](#nestedatt--ipv6_vrf_hosts))
- `monitor_severity` (String)
- `origin_id_name` (String) Define a unique text string as ID
- `origin_id_type` (String) Use origin hostname/ip/ipv6 as ID
- `source_interface` (String)
- `source_interfaces_vrf` (Attributes List) Specify interface and vrf for source address in logging transactions (see [below for nested schema](#nestedatt--source_interfaces_vrf))
- `trap` (Boolean) Set trap server logging level
- `trap_severity` (String)

<a id="nestedatt--ipv4_hosts"></a>
### Nested Schema for `ipv4_hosts`

Read-Only:

- `ipv4_host` (String)


<a id="nestedatt--ipv4_vrf_hosts"></a>
### Nested Schema for `ipv4_vrf_hosts`

Read-Only:

- `ipv4_host` (String)
- `vrf` (String) Set VRF option


<a id="nestedatt--ipv6_hosts"></a>
### Nested Schema for `ipv6_hosts`

Read-Only:

- `ipv6_host` (String)


<a id="nestedatt--ipv6_vrf_hosts"></a>
### Nested Schema for `ipv6_vrf_hosts`

Read-Only:

- `ipv6_host` (String)
- `vrf` (String) Set VRF option


<a id="nestedatt--source_interfaces_vrf"></a>
### Nested Schema for `source_interfaces_vrf`

Read-Only:

- `interface_name` (String)
- `vrf` (String) Specify the vrf of source interface for logging transactions
