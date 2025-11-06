resource "iosxe_bgp_l2vpn_evpn_neighbor" "example" {
  asn                    = "65000"
  ip                     = "3.3.3.3"
  activate               = true
  send_community         = "both"
  route_reflector_client = false
  soft_reconfiguration   = "inbound"
  route_map = [
    {
      inout          = "in"
      route_map_name = "RM1"
    }
  ]
}
