---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxe_username Resource - terraform-provider-iosxe"
subcategory: "System"
description: |-
  This resource can manage the Username configuration.
---

# iosxe_username (Resource)

This resource can manage the Username configuration.

## Example Usage

```terraform
resource "iosxe_username" "example" {
  name                = "user1"
  privilege           = 15
  description         = "User1 description"
  password_encryption = "0"
  password            = "MyPassword"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)

### Optional

- `description` (String) description string with max 128 characters
- `device` (String) A device name from the provider configuration.
- `password` (String)
- `password_encryption` (String) - Choices: `0`, `6`, `7`
- `privilege` (Number) Set user privilege level
  - Range: `0`-`15`
- `secret` (String)
- `secret_encryption` (String) - Choices: `0`, `5`, `8`, `9`

### Read-Only

- `id` (String) The path of the object.

## Import

Import is supported using the following syntax:

```shell
terraform import iosxe_username.example "Cisco-IOS-XE-native:native/username=user1"
```