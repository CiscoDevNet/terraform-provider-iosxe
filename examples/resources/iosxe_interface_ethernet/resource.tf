resource "iosxe_interface_ethernet" "example" {
  type                           = "GigabitEthernet"
  name                           = "3"
  bandwidth                      = 1000000
  description                    = "My Interface Description"
  shutdown                       = false
  ip_proxy_arp                   = false
  ip_redirects                   = false
  ip_unreachables                = false
  ipv4_address                   = "15.1.1.1"
  ipv4_address_mask              = "255.255.255.252"
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
  ipv6_enable             = true
  ipv6_mtu                = 1300
  ipv6_nd_ra_suppress_all = true
  ipv6_address_dhcp       = true
  ipv6_link_local_addresses = [
    {
      address    = "fe80::9656:d028:8652:66b6"
      link_local = true
    }
  ]
  ipv6_addresses = [
    {
      prefix = "2001:DB8::/32"
      eui_64 = true
    }
  ]
  arp_timeout             = 300
  spanning_tree_link_type = "point-to-point"
  negotiation_auto        = false
  service_policy_input    = "POLICY1"
  service_policy_output   = "POLICY1"
  ip_flow_monitors = [
    {
      name      = "MON1"
      direction = "input"
    }
  ]
  load_interval                    = 30
  snmp_trap_link_status            = true
  logging_event_link_status_enable = true
}
