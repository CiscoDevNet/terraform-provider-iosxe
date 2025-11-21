resource "iosxe_ipv6_pim" "example" {
  rp_address             = "2001:db8::100"
  rp_address_access_list = "ipv6_acl_1"
  rp_address_bidir       = false
  vrfs = [
    {
      vrf                    = "VRF1"
      rp_address             = "2001:db8::100"
      rp_address_access_list = "ipv6_acl_2"
      rp_address_bidir       = false
    }
  ]
}
