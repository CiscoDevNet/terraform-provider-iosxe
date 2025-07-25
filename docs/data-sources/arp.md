---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxe_arp Data Source - terraform-provider-iosxe"
subcategory: "System"
description: |-
  This data source can read the ARP configuration.
---

# iosxe_arp (Data Source)

This data source can read the ARP configuration.

## Example Usage

```terraform
data "iosxe_arp" "example" {
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `entry_learn` (Number) Maximum learn entry limit
- `id` (String) The path of the retrieved object.
- `incomplete_entries` (Number) Specify the number of IP addresses to resolve
- `inspection_filters` (Attributes List) Specify ARP acl to be applied (see [below for nested schema](#nestedatt--inspection_filters))
- `inspection_log_buffer_entries` (Number) Number of entries for log buffer
- `inspection_log_buffer_logs_entries` (Number) Number of entries for log buffer
- `inspection_log_buffer_logs_interval` (Number) Interval for controlling logging rate
- `inspection_validate_allow_zeros` (Boolean) Allow 0.0.0.0 sender IP address
- `inspection_validate_dst_mac` (Boolean) Validate destination MAC address
- `inspection_validate_ip` (Boolean) Validate IP addresses
- `inspection_validate_src_mac` (Boolean) Validate source MAC address
- `inspection_vlan` (String) Enable/Disable ARP Inspection on vlans(Deprecated)
- `proxy_disable` (Boolean) Disable proxy ARP on all interfaces

<a id="nestedatt--inspection_filters"></a>
### Nested Schema for `inspection_filters`

Read-Only:

- `name` (String)
- `vlan` (Attributes List) Vlans to apply the filter (see [below for nested schema](#nestedatt--inspection_filters--vlan))

<a id="nestedatt--inspection_filters--vlan"></a>
### Nested Schema for `inspection_filters.vlan`

Read-Only:

- `static` (Boolean) Apply the ACL statically
- `vlan_range` (String)
