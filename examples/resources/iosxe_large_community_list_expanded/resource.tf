resource "iosxe_large_community_list_expanded" "example" {
  name = "LCLE1"
  entries = [
    {
      action = "permit"
      regex  = "65000:500"
    }
  ]
}
