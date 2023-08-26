resource "iosxe_bgp_address_family_ipv4" "example" {
  asn     = "65000"
  af_name = "unicast"
  ipv4_unicast_networks_mask = [
    {
      network   = "12.0.0.0"
      mask      = "255.255.0.0"
      route_map = "RM1"
      backdoor  = true
    }
  ]
  ipv4_unicast_networks = [
    {
      network   = "13.0.0.0"
      route_map = "RM1"
      backdoor  = true
    }
  ]
}
