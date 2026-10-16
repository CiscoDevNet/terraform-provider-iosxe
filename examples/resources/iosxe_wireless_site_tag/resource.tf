resource "iosxe_wireless_site_tag" "example" {
  site_tag_name   = "S1-SITE-TAG"
  description     = "My Site Tag"
  ap_join_profile = "S1-AP-JOIN"
  flex_profile    = "S1-FLEX"
  local_site      = false
}
