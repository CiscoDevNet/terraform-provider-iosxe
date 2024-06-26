---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxe_dot1x Data Source - terraform-provider-iosxe"
subcategory: "System"
description: |-
  This data source can read the Dot1x configuration.
---

# iosxe_dot1x (Data Source)

This data source can read the Dot1x configuration.

## Example Usage

```terraform
data "iosxe_dot1x" "example" {
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `auth_fail_eapol` (Boolean) Send EAPOL-Success on successful auth-fail Authorization
- `credentials` (Attributes List) Configure 802.1X credentials profiles (see [below for nested schema](#nestedatt--credentials))
- `critical_eapol_config_block` (Boolean) Block all EAPoL transaction on Critical Authentication
- `critical_recovery_delay` (Number) Set 802.1x Critical Authentication Recovery Delay period
- `id` (String) The path of the retrieved object.
- `logging_verbose` (Boolean) Show verbose messages in system logs
- `supplicant_controlled_transient` (Boolean) Controlled access is only applied during authentication
- `supplicant_force_multicast` (Boolean) Force 802.1X supplicant to send multicast packets
- `system_auth_control` (Boolean) Enable or Disable SysAuthControl
- `test_timeout` (Number) Timeout for device EAPOL capabilities test in seconds

<a id="nestedatt--credentials"></a>
### Nested Schema for `credentials`

Read-Only:

- `anonymous_id` (String) Set the anonymous userid
- `description` (String) Provide a description for the credentials profile
- `password` (String)
- `password_type` (String)
- `pki_trustpoint` (String) Set the default pki trustpoint
- `profile_name` (String) Specify a profile name
- `username` (String) Set the authentication userid
