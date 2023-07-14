resource "iosxe_prefix_list" "example" {
  prefixes = [
    {
      name   = "PREFIX_LIST_1"
      seq    = 10
      action = "permit"
      ip     = "10.0.0.0/8"
      ge     = 24
      le     = 32
    }
  ]
}
