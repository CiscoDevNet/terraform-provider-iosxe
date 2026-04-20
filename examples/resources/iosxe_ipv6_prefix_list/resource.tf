resource "iosxe_ipv6_prefix_list" "example" {
  prefixes = [
    {
      name   = "PREFIX_LIST_1"
      seq    = 10
      action = "permit"
      ip     = "2001:db8::/32"
      ge     = 48
      le     = 128
    }
  ]
  prefix_list_description = [
    {
      name        = "PREFIX_LIST_1"
      description = "My IPv6 prefix list"
    }
  ]
}
