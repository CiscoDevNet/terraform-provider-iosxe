resource "iosxe_wireless_site_tag" "example" {
  site_tag_name   = "S1-SITE-TAG"
  description     = "My Site Tag"
  ap_join_profile = "S1-AP-JOIN"
  local_site      = false
  flex_profile    = "S1-FLEX"
}
