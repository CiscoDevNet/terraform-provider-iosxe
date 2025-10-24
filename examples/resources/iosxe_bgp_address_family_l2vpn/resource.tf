resource "iosxe_bgp_address_family_l2vpn" "example" {
  asn                    = "65000"
  af_name                = "evpn"
  rewrite_evpn_rt_asn    = true
  nexthop_trigger_enable = true
  nexthop_trigger_delay  = 10
}
