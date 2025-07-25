---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxe_vlan_group Data Source - terraform-provider-iosxe"
subcategory: "Switching"
description: |-
  This data source can read the VLAN Group configuration.
---

# iosxe_vlan_group (Data Source)

This data source can read the VLAN Group configuration.

## Example Usage

```terraform
data "iosxe_vlan_group" "example" {
  name = "GROUP1"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Group name starts with alphabet

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `id` (String) The path of the retrieved object.
- `vlan_lists` (List of Number) VLANs in the vlan group
