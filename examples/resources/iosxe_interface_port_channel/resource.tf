resource "iosxe_interface_port_channel" "example" {
  name                           = 10
  description                    = "My Interface Description"
  shutdown                       = false
  vrf_forwarding                 = "VRF1"
  ipv4_address                   = "192.0.2.1"
  ipv4_address_mask              = "255.255.255.0"
  ip_access_group_in             = "1"
  ip_access_group_in_enable      = true
  ip_access_group_out            = "1"
  ip_access_group_out_enable     = true
  ip_dhcp_relay_source_interface = "Loopback100"
  helper_addresses = [
    {
      address = "10.10.10.10"
      global  = false
    }
  ]
  bfd_template            = "bfd_template1"
  bfd_enable              = true
  bfd_local_address       = "1.2.3.4"
  ipv6_enable             = true
  ipv6_mtu                = 1300
  ipv6_nd_ra_suppress_all = true
  ipv6_address_dhcp       = true
  ipv6_link_local_addresses = [
    {
      address    = "fe80::64"
      link_local = true
    }
  ]
  ipv6_addresses = [
    {
      prefix = "2224:DB8::/32"
      eui_64 = true
    }
  ]
  arp_timeout                      = 2147
  load_interval                    = 30
  snmp_trap_link_status            = true
  logging_event_link_status_enable = false
}
