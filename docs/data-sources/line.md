---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxe_line Data Source - terraform-provider-iosxe"
subcategory: "System"
description: |-
  This data source can read the Line configuration.
---

# iosxe_line (Data Source)

This data source can read the Line configuration.

## Example Usage

```terraform
data "iosxe_line" "example" {
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `console` (Attributes List) Primary terminal line (see [below for nested schema](#nestedatt--console))
- `id` (String) The path of the retrieved object.
- `vty` (Attributes List) Virtual terminal (see [below for nested schema](#nestedatt--vty))

<a id="nestedatt--console"></a>
### Nested Schema for `console`

Read-Only:

- `exec_timeout_minutes` (Number) <0-35791>;;Timeout in minutes
- `exec_timeout_seconds` (Number) <0-2147483>;;Timeout in seconds
- `first` (String) Console line number
- `login_authentication` (String)
- `login_local` (Boolean)
- `password` (String)
- `password_level` (Number) Set exec level password
- `password_type` (String)
- `privilege_level` (Number)
- `stopbits` (String) Set async line stop bits


<a id="nestedatt--vty"></a>
### Nested Schema for `vty`

Read-Only:

- `access_classes` (Attributes List) Choose direction of the access list (see [below for nested schema](#nestedatt--vty--access_classes))
- `authorization_exec` (String) Use an authorization list with this name
- `authorization_exec_default` (Boolean) Use the default authorization list
- `escape_character` (String)
- `exec_timeout_minutes` (Number) <0-35791>;;Timeout in minutes
- `exec_timeout_seconds` (Number) <0-2147483>;;Timeout in seconds
- `first` (Number) Vty first line number
- `last` (Number) Vty last line number
- `login_authentication` (String) Authentication list
- `password` (String)
- `password_level` (Number) Set exec level password
- `password_type` (String)
- `transport_input` (String) Define which protocols to use when connecting to the terminal server
- `transport_input_all` (Boolean) All protocols
- `transport_input_none` (Boolean) Define no transport protocols for line
- `transport_preferred_protocol` (String)

<a id="nestedatt--vty--access_classes"></a>
### Nested Schema for `vty.access_classes`

Read-Only:

- `access_list` (String)
- `direction` (String) Filter connections based on the incoming/outgoing direction
- `vrf_also` (Boolean) Same access list is applied for all VRFs
