---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxe_tacacs_server Data Source - terraform-provider-iosxe"
subcategory: "AAA"
description: |-
  This data source can read the TACACS Server configuration.
---

# iosxe_tacacs_server (Data Source)

This data source can read the TACACS Server configuration.

## Example Usage

```terraform
data "iosxe_tacacs_server" "example" {
  name = "tacacs_10.10.15.13"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name for the tacacs server configuration

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `address_ipv4` (String) IPv4 address or Hostname for tacacs server
- `encryption` (String) 0 - Specifies an UNENCRYPTED key will follow 6 - Specifies an ENCRYPTED key will follow 7 - Specifies HIDDEN key will follow
- `id` (String) The path of the retrieved object.
- `key` (String) The UNENCRYPTED (cleartext) server key
- `timeout` (Number) Time to wait for this TACACS server to reply (overrides default)
