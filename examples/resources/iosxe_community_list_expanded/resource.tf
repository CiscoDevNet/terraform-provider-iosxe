resource "iosxe_community_list_expanded" "example" {
  name = "CLE1"
  entries = [
    {
      action = "permit"
      regex  = "65000:500"
    }
  ]
}
