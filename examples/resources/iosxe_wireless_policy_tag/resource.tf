resource "iosxe_wireless_policy_tag" "example" {
  tag_name    = "S2-PT-CORP"
  description = "S2 Corporate campus policy tag"
  wlan_policies = [
    {
      wlan_profile_name   = "S2-WLAN-CORP"
      policy_profile_name = "S2-PP-CORP"
    }
  ]
}
