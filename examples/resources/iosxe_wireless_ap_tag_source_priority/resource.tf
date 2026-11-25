resource "iosxe_wireless_ap_tag_source_priority" "example" {
  priorities = [
    {
      priority = 2
      source   = "tag-source-ap-pnp"
    }
  ]
}
