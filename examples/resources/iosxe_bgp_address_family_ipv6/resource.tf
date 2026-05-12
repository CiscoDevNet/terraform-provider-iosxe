resource "iosxe_bgp_address_family_ipv6" "example" {
  asn                                           = "65000"
  af_name                                       = "unicast"
  ipv6_unicast_redistribute_connected           = true
  ipv6_unicast_redistribute_connected_route_map = "RM_BGP6_CONNECTED"
  ipv6_unicast_redistribute_connected_metric    = 100
  ipv6_unicast_redistribute_static              = true
  ipv6_unicast_redistribute_static_route_map    = "RM_BGP6_STATIC"
  ipv6_unicast_redistribute_static_metric       = 200
  ipv6_unicast_aggregate_addresses = [
    {
      ipv6_address = "2001:DB8::/32"
    }
  ]
  ipv6_unicast_networks = [
    {
      network   = "2001:1234::/64"
      route_map = "RM1"
      backdoor  = true
    }
  ]
  ipv6_unicast_admin_distances = [
    {
      distance            = 200
      source_ipv6_address = "2001:DB8::/48"
    }
  ]
  ipv6_unicast_distance_bgp_external = 20
  ipv6_unicast_distance_bgp_internal = 200
  ipv6_unicast_distance_bgp_local    = 200
  ipv6_unicast_maximum_paths_ebgp    = 2
  ipv6_unicast_maximum_paths_ibgp    = 2
}
