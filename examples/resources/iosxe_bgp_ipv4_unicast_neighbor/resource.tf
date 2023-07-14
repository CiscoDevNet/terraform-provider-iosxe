resource "iosxe_bgp_ipv4_unicast_neighbor" "example" {
  asn                    = "65000"
  ip                     = "3.3.3.3"
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
