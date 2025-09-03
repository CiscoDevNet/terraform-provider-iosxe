resource "iosxe_bgp_address_family_ipv4_vrf" "example" {
  asn     = "65000"
  af_name = "unicast"
  vrfs = [
    {
      name                                = "VRF1"
      ipv4_unicast_advertise_l2vpn_evpn   = true
      ipv4_unicast_redistribute_connected = true
      ipv4_unicast_router_id_loopback     = 101
      ipv4_unicast_aggregate_addresses = [
        {
          ipv4_address = "50.0.0.0"
          ipv4_mask    = "255.255.0.0"
        }
      ]
      ipv4_unicast_redistribute_static = true
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
      ipv4_unicast_admin_distances = [
        {
          distance  = 200
          source_ip = "10.1.1.1"
          wildcard  = "0.0.0.0"
        }
      ]
      ipv4_unicast_distance_bgp_external = 20
      ipv4_unicast_distance_bgp_internal = 200
      ipv4_unicast_distance_bgp_local    = 200
    }
  ]
}
