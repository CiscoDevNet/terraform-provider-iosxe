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
