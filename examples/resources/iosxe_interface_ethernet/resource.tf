resource "iosxe_interface_ethernet" "example" {
  type                           = "GigabitEthernet"
  name                           = "3"
  mtu                            = 1600
  bandwidth                      = 1000000
  description                    = "My Interface Description"
  shutdown                       = false
  ip_proxy_arp                   = false
  ip_redirects                   = false
  ip_unreachables                = false
  ipv4_address                   = "15.1.1.1"
  ipv4_address_mask              = "255.255.255.252"
  ip_dhcp_relay_source_interface = "Loopback100"
  ip_access_group_in_enable      = true
  ip_access_group_in             = "1"
  ip_access_group_out_enable     = true
  ip_access_group_out            = "1"
  helper_addresses = [
    {
      address = "10.10.10.10"
      global  = false
      vrf     = "VRF1"
    }
  ]
  source_template = [
    {
      template_name = "TEMP1"
      merge         = false
    }
  ]
  bfd_template            = "bfd_template1"
  bfd_enable              = false
  bfd_local_address       = "1.2.3.4"
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
  arp_timeout                             = 300
  spanning_tree_link_type                 = "point-to-point"
  bpduguard_enable                        = false
  bpduguard_disable                       = false
  spanning_tree_portfast                  = true
  spanning_tree_portfast_disable          = false
  spanning_tree_portfast_trunk            = true
  spanning_tree_portfast_edge             = false
  ip_dhcp_relay_information_option_vpn_id = true
  negotiation_auto                        = false
  service_policy_input                    = "POLICY1"
  service_policy_output                   = "POLICY1"
  ip_flow_monitors = [
    {
      name      = "MON1"
      direction = "input"
    }
  ]
  load_interval                    = 30
  snmp_trap_link_status            = false
  logging_event_link_status_enable = false
  cdp_enable                       = true
  cdp_tlv_app                      = false
  cdp_tlv_location                 = false
  cdp_tlv_server_location          = false
  ip_nat_inside                    = true
  evpn_ethernet_segments = [
    {
      es_value = 1
    }
  ]
}
