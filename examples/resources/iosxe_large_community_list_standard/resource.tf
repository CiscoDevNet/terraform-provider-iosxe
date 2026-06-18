resource "iosxe_large_community_list_standard" "example" {
  name           = "LCLS1"
  deny_entries   = ["65000:500:1"]
  permit_entries = ["65000:501:1"]
}
