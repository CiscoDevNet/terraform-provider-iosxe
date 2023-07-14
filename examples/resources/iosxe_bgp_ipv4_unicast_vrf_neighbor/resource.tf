resource "iosxe_bgp_ipv4_unicast_vrf_neighbor" "example" {
  asn                    = "65000"
  vrf                    = "VRF1"
  ip                     = "3.3.3.3"
  remote_as              = "65000"
  description            = "BGP Neighbor Description"
  shutdown               = false
  update_source_loopback = "100"
  activate               = true
  send_community         = "both"
  route_reflector_client = false
  route_maps = [
    {
      in_out         = "in"
      route_map_name = "RM1"
    }
  ]
}
