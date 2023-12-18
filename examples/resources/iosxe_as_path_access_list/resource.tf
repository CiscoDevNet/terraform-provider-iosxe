resource "iosxe_as_path_access_list" "example" {
  name = 10
  entries = [
    {
      action = "permit"
      regex  = "^100$"
    }
  ]
}
