resource "iosxe_bgp_address_family_ipv4_vrf" "example" {
  asn     = "65000"
  af_name = "unicast"
  vrfs = [
    {
      name                   = "VRF1"
      advertise_l2vpn_evpn   = true
      redistribute_connected = true
      redistribute_static    = true
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
  ]
}
