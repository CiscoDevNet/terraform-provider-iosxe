resource "iosxe_wireless_policy_profile" "example" {
  policy_profile_name   = "S2-PP-CORP"
  description           = "S2 Corporate users"
  status                = true
  vlan                  = "100"
  aaa_override          = true
  nac                   = true
  nac_type              = "nac-support-radius"
  session_timeout       = 3600
  idle_timeout          = 600
  exclusionlist_timeout = 180
  static_ip_mobility    = true
  dhcp_tlv_caching      = true
  http_tlv_caching      = true
  ipv4_flow_monitors_input = [
    {
      name = "S2-AVC-MON-IN"
    }
  ]
  ipv4_flow_monitors_output = [
    {
      name = "S2-AVC-MON-OUT"
    }
  ]
}
