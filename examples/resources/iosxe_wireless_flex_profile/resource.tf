resource "iosxe_wireless_flex_profile" "example" {
  name                = "S2-BRANCH-FLEX"
  description         = "Branch FlexConnect profile"
  native_vlan_id      = 10
  arp_caching         = false
  fallback_radio_shut = true
  predownload         = false
  vlan_names = [
    {
      name    = "data"
      vlan_id = 20
    }
  ]
  acl_policies = [
    {
      name = "S2-FLEX-ACL"
    }
  ]
}
