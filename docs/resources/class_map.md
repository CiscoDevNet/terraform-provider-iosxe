---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxe_class_map Resource - terraform-provider-iosxe"
subcategory: "QoS"
description: |-
  This resource can manage the Class Map configuration.
---

# iosxe_class_map (Resource)

This resource can manage the Class Map configuration.

## Example Usage

```terraform
resource "iosxe_class_map" "example" {
  name                                  = "CM1"
  type                                  = "control"
  subscriber                            = true
  prematch                              = "match-all"
  match_authorization_status_authorized = true
  match_result_type_aaa_timeout         = true
  match_activated_service_templates = [
    {
      service_name = "CRITICAL_AUTH_ACCESS"
    }
  ]
  match_authorizing_method_priority_greater_than = [20]
  match_method_dot1x                             = true
  match_result_type_method_dot1x_authoritative   = true
  match_result_type_method_dot1x_agent_not_found = true
  match_result_type_method_dot1x_method_timeout  = true
  match_method_mab                               = true
  match_result_type_method_mab_authoritative     = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) name of the class map
- `prematch` (String) Logical-AND/Logical-OR of all matching statements under this class map
  - Choices: `match-all`, `match-any`, `match-none`

### Optional

- `description` (String) Class-Map description
- `device` (String) A device name from the provider configuration.
- `match_activated_service_templates` (Attributes List) match name of service template activated on session (see [below for nested schema](#nestedatt--match_activated_service_templates))
- `match_authorization_status_authorized` (Boolean) authorized
- `match_authorization_status_unauthorized` (Boolean) unauthorized
- `match_authorizing_method_priority_greater_than` (List of Number) greater than
- `match_method_dot1x` (Boolean) dot1x
- `match_method_mab` (Boolean) mab
- `match_result_type_aaa_timeout` (Boolean) aaa timeout type
- `match_result_type_method_dot1x_agent_not_found` (Boolean) agent not found type
- `match_result_type_method_dot1x_authoritative` (Boolean) failure type
- `match_result_type_method_dot1x_method_timeout` (Boolean) method_timeout type
- `match_result_type_method_mab_authoritative` (Boolean) failure type
- `subscriber` (Boolean) Domain name of the class map
- `type` (String) type of the class-map
  - Choices: `access-control`, `appnav`, `control`, `inspect`, `multicast-flows`, `site-manager`, `stack`, `traffic`

### Read-Only

- `id` (String) The path of the object.

<a id="nestedatt--match_activated_service_templates"></a>
### Nested Schema for `match_activated_service_templates`

Required:

- `service_name` (String) Enter service name

## Import

Import is supported using the following syntax:

```shell
terraform import iosxe_class_map.example "Cisco-IOS-XE-native:native/policy/Cisco-IOS-XE-policy:class-map=CM1"
```