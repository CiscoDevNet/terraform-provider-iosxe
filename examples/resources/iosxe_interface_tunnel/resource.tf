resource "iosxe_interface_tunnel" "example" {
  name              = 90
  description       = "My Interface Description"
  shutdown          = false
  ip_proxy_arp      = false
  ip_redirects      = false
  unreachables      = false
  vrf_forwarding    = "VRF1"
  ipv6_enable       = true
  ipv6_mtu          = 1300
  ra_suppress_all   = true
  ipv6_address_dhcp = true
  ipv6_link_local_addresses = [
    {
      address    = "fe80::9656:d028:8652:66b6"
      link_local = true
    }
  ]
  ipv6_address_prefix_lists = [
    {
      prefix = "2001:DB8::/32"
      eui_64 = true
    }
  ]
  tunnel_destination_ipv4        = "2.2.2.2"
  crypto_ipsec_df_bit            = "clear"
  ipv4_address                   = "10.1.1.1"
  ipv4_address_mask              = "255.255.255.0"
  ip_dhcp_relay_source_interface = "Loopback100"
  ip_access_group_in             = "1"
  ip_access_group_in_enable      = true
  ip_access_group_out            = "1"
  ip_access_group_out_enable     = true
  helper_addresses = [
    {
      address = "10.10.10.10"
      global  = false
      vrf     = "VRF1"
    }
  ]
  template                      = "Tunnel_template1"
  enable                        = true
  local_address                 = "1.2.3.4"
  interval_interface_msecs      = 50
  interval_interface_min_rx     = 50
  interval_interface_multiplier = 3
  echo                          = true
}
