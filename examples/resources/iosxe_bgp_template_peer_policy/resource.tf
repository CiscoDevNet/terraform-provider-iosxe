resource "iosxe_bgp_template_peer_policy" "example" {
  asn                    = "65000"
  name                   = "PEERPOLICY_1"
  route_reflector_client = true
  send_community         = "both"
  route_maps = [
    {
      in_out         = "in"
      route_map_name = "ROUTEMAP_1"
    }
  ]
  allowas_in_as_number      = 2
  as_override_split_horizon = true
}
