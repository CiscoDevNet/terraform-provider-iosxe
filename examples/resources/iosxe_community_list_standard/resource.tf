resource "iosxe_community_list_standard" "example" {
  name           = "CLS1"
  deny_entries   = ["65000:500"]
  permit_entries = ["65000:501"]
}
