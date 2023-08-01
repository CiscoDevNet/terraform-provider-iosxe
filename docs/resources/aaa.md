---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxe_aaa Resource - terraform-provider-iosxe"
subcategory: "AAA"
description: |-
  This resource can manage the AAA configuration.
---

# iosxe_aaa (Resource)

This resource can manage the AAA configuration.

## Example Usage

```terraform
resource "iosxe_aaa" "example" {
  new_model                    = true
  server_radius_dynamic_author = true
  session_id                   = "common"
  server_radius_dynamic_author_clients = [
    {
      ip              = "11.1.1.1"
      server_key_type = "0"
      server_key      = "abcd123"
    }
  ]
  group_server_radius = [
    {
      name = "T-Group"
      server_names = [
        {
          name = "TESTRADIUS"
        }
      ]
    }
  ]
  group_tacacsplus = [
    {
      name = "tacacs-group"
      servers = [
        {
          name = "tacacs_10.10.15.12"
        }
      ]
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `device` (String) A device name from the provider configuration.
- `group_server_radius` (Attributes List) Radius server-group definition (see [below for nested schema](#nestedatt--group_server_radius))
- `group_tacacsplus` (Attributes List) Tacacs+ server-group definition (see [below for nested schema](#nestedatt--group_tacacsplus))
- `new_model` (Boolean) Enable NEW access control commands and functions.(Disables OLD commands.)
- `server_radius_dynamic_author` (Boolean) Local server profile for RFC 3576 support
- `server_radius_dynamic_author_clients` (Attributes List) Specify a RADIUS client (see [below for nested schema](#nestedatt--server_radius_dynamic_author_clients))
- `session_id` (String) AAA Session ID
  - Choices: `common`, `unique`

### Read-Only

- `id` (String) The path of the object.

<a id="nestedatt--group_server_radius"></a>
### Nested Schema for `group_server_radius`

Required:

- `name` (String) Radius Server-group name with max string length 32

Optional:

- `server_names` (Attributes List) Name of radius server (see [below for nested schema](#nestedatt--group_server_radius--server_names))

<a id="nestedatt--group_server_radius--server_names"></a>
### Nested Schema for `group_server_radius.server_names`

Required:

- `name` (String) Radius server name



<a id="nestedatt--group_tacacsplus"></a>
### Nested Schema for `group_tacacsplus`

Required:

- `name` (String) Server-group name with max string length 32

Optional:

- `servers` (Attributes List) Name of tacacs server (see [below for nested schema](#nestedatt--group_tacacsplus--servers))

<a id="nestedatt--group_tacacsplus--servers"></a>
### Nested Schema for `group_tacacsplus.servers`

Required:

- `name` (String)



<a id="nestedatt--server_radius_dynamic_author_clients"></a>
### Nested Schema for `server_radius_dynamic_author_clients`

Required:

- `ip` (String)

Optional:

- `server_key` (String)
- `server_key_type` (String) - Choices: `0`, `6`, `7`

## Import

Import is supported using the following syntax:

```shell
terraform import iosxe_aaa.example "Cisco-IOS-XE-native:native/aaa"
```