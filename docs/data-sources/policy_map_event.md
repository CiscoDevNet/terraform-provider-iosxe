---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxe_policy_map_event Data Source - terraform-provider-iosxe"
subcategory: "QoS"
description: |-
  This data source can read the Policy Map Event configuration.
---

# iosxe_policy_map_event (Data Source)

This data source can read the Policy Map Event configuration.

## Example Usage

```terraform
data "iosxe_policy_map_event" "example" {
  name       = "dot1x_policy"
  event_type = "authentication-success"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `event_type` (String) The event this control class-map triggers upon
- `name` (String) Name of the policy map

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `class_numbers` (Attributes List) class number, 1 for 1st class, 2 for 2nd... (see [below for nested schema](#nestedatt--class_numbers))
- `id` (String) The path of the retrieved object.
- `match_type` (String) Matching criteria for first or all events.

<a id="nestedatt--class_numbers"></a>
### Nested Schema for `class_numbers`

Read-Only:

- `action_numbers` (Attributes List) action number, 1 for 1st class, 2 for 2nd... (see [below for nested schema](#nestedatt--class_numbers--action_numbers))
- `class` (String) The class type this control policy-map triggers upon
- `execution_type` (String) Policy execution strategy
- `number` (Number) class number, 1 for 1st class, 2 for 2nd...

<a id="nestedatt--class_numbers--action_numbers"></a>
### Nested Schema for `class_numbers.action_numbers`

Read-Only:

- `activate_interface_template` (String) activate interface template
- `activate_policy_type_control_subscriber` (String) policy type control subscriber
- `activate_service_template_config_aaa_list` (String) Named Method List
- `activate_service_template_config_precedence` (Number) Template precedence
- `activate_service_template_config_replace_all` (Boolean) Replace all existing authorization data and services
- `activate_service_template_config_service_template` (String) activate service template
- `authenticate_using_aaa_authc_list` (String) Specify authentication method list
- `authenticate_using_aaa_authz_list` (String) Specify authorization method list
- `authenticate_using_both` (Boolean) Enabling Dot1x Authenticator & Supplicant
- `authenticate_using_method` (String) method/protocol to be used for authentication
- `authenticate_using_parameter_map` (String) Specify parameter map name
- `authenticate_using_priority` (Number) Method priority
- `authenticate_using_retries` (Number) Number of times to retry failed authentications
- `authenticate_using_retry_time` (Number) Time interval between retries
- `authentication_restart` (Number) restarts the auth sequence after the specified number of sec
- `authorize` (Boolean) authorize session
- `clear_authenticated_data_hosts_on_port` (Boolean) clears authenticated data hosts on the port
- `clear_session` (Boolean) clears an active session
- `deactivate_interface_template` (String) activate interface template
- `deactivate_policy_type_control_subscriber` (String) policy type control subscriber
- `deactivate_service_template` (String) activate service template
- `err_disable` (Boolean) temporarily disable port
- `map_attribute_to_service_table` (String) map identity-update attribute to a auto-conf templates
- `notify` (Boolean) notifies the session attributes
- `number` (Number) action number, 1 for 1st class, 2 for 2nd...
- `pause_reauthentication` (Boolean) pause reauthentication
- `protect` (Boolean) silently drop violating packets
- `replace` (Boolean) clear existing session and create session for violating host
- `restrict` (Boolean) drop violating packets and generate a syslog
- `resume_reauthentication` (Boolean) resume reauthentication
- `set_domain` (String) set domain
- `set_timer_name` (String) timer name
- `set_timer_value` (Number) Enter a value between 1 and 65535
- `terminate_config` (String) terminate auth method
- `unauthorize` (Boolean) unauthorize session
