resource "iosxe_bgp_address_family_ipv6" "example" {
  asn                                           = "65000"
  af_name                                       = "unicast"
  ipv6_unicast_redistribute_connected           = true
  ipv6_unicast_redistribute_connected_route_map = "RM_BGP6_CONNECTED"
  ipv6_unicast_redistribute_connected_metric    = 100
  ipv6_unicast_redistribute_static              = true
  ipv6_unicast_redistribute_static_route_map    = "RM_BGP6_STATIC"
  ipv6_unicast_redistribute_static_metric       = 200
  ipv6_unicast_networks = [
    {
      network   = "2001:1234::/64"
      route_map = "RM1"
      backdoor  = true
    }
  ]
}
