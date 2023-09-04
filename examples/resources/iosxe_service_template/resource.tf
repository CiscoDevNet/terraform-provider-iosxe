resource "iosxe_service_template" "example" {
  word = "DEFAULT_LINKSEC_POLICY_MUST_SECURE"
  access_group = [
    {
      name = "ag1"
    }
  ]
  inactivity_timer_value = 25
  inactivity_timer_probe = false
  vlan                   = 27
  voice_vlan             = false
  linksec_policy         = "must-secure"
  sgt                    = 57
  absolute_timer         = 45
  description            = "service_template_desc"
  interface_template = [
    {
      name = "template1"
    }
  ]
  tunnel_capwap_name          = "tunnel_name"
  vnid                        = "vnid_1"
  redirect_append_client_mac  = "00:01:00:01:00:01"
  redirect_append_switch_mac  = "00:01:00:01:00:02"
  redirect_url_url_name       = "valid_url"
  redirect_url_match_acl_name = "acl_name"
  redirect_url_match_action   = "redirect-on-no-match"
  dns_acl_preauth             = "dns_acl_name"
  service_policy_qos_input    = "input_qos"
  service_policy_qos_output   = "output_qos"
  tag_config = [
    {
      name = "tag_name"
    }
  ]
}
