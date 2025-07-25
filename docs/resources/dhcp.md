---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxe_dhcp Resource - terraform-provider-iosxe"
subcategory: "System"
description: |-
  This resource can manage the DHCP configuration.
---

# iosxe_dhcp (Resource)

This resource can manage the DHCP configuration.

## Example Usage

```terraform
resource "iosxe_dhcp" "example" {
  relay_information_trust_all      = false
  relay_information_option_default = false
  relay_information_option_vpn     = true
  snooping                         = true
  snooping_vlans = [
    {
      vlan_id = "3-4"
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `compatibility_suboption_link_selection` (String) - Choices: `cisco`, `standard`
- `compatibility_suboption_server_override` (String) - Choices: `cisco`, `standard`
- `delete_mode` (String) Configure behavior when deleting/destroying the resource. Either delete the entire object (YANG container) being managed, or only delete the individual resource attributes configured explicitly and leave everything else as-is. Default value is `all`.
  - Choices: `all`, `attributes`
- `device` (String) A device name from the provider configuration.
- `relay_information_option_default` (Boolean) Default option, no vpn
- `relay_information_option_vpn` (Boolean) Insert VPN sub-options and change the giaddr to the outgoing interface
- `relay_information_trust_all` (Boolean) Received DHCP packets may contain relay info option with zero giaddr
- `snooping` (Boolean) DHCP Snooping
- `snooping_information_option_format_remote_id_hostname` (Boolean) Use configured hostname for remote id
- `snooping_vlans` (Attributes List) DHCP Snooping vlan (OBSOLETE) (see [below for nested schema](#nestedatt--snooping_vlans))

### Read-Only

- `id` (String) The path of the object.

<a id="nestedatt--snooping_vlans"></a>
### Nested Schema for `snooping_vlans`

Required:

- `vlan_id` (String) DHCP Snooping vlan first number or vlan range,example: 1,3-5,7,9-11

## Import

Import is supported using the following syntax:

```shell
terraform import iosxe_dhcp.example "Cisco-IOS-XE-native:native/ip/dhcp"
```
