data "iosxe_bgp_ipv4_unicast_vrf_neighbor" "example" {
  asn = "65000"
  vrf = "VRF1"
  ip  = "3.3.3.3"
}
