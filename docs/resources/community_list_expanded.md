---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxe_community_list_expanded Resource - terraform-provider-iosxe"
subcategory: "BGP"
description: |-
  This resource can manage the Community List Expanded configuration.
---

# iosxe_community_list_expanded (Resource)

This resource can manage the Community List Expanded configuration.

## Example Usage

```terraform
resource "iosxe_community_list_expanded" "example" {
  name = "CLE1"
  entries = [
    {
      action = "permit"
      regex  = "65000:500"
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)

### Optional

- `device` (String) A device name from the provider configuration.
- `entries` (Attributes List) Specify community list to accept or deny (see [below for nested schema](#nestedatt--entries))

### Read-Only

- `id` (String) The path of the object.

<a id="nestedatt--entries"></a>
### Nested Schema for `entries`

Required:

- `action` (String) - Choices: `deny`, `permit`
- `regex` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import iosxe_community_list_expanded.example "Cisco-IOS-XE-native:native/ip/Cisco-IOS-XE-bgp:community-list/expanded=CLE1"
```
