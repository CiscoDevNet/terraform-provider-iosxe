resource "iosxe_interface_vlan" "example" {
  name                           = 10
  autostate                      = false
  description                    = "My Interface Description"
  shutdown                       = false
  ip_proxy_arp                   = false
  ip_redirects                   = false
  ip_unreachables                = false
  vrf_forwarding                 = "VRF1"
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
  bfd_template            = "bfd_template1"
  bfd_enable              = true
  bfd_local_address       = "1.2.3.4"
  ipv6_enable             = true
  ipv6_mtu                = 1300
  ipv6_nd_ra_suppress_all = true
  ipv6_address_dhcp       = true
  ipv6_link_local_addresses = [
    {
      address    = "fe80::9656:d028:8652:66bb"
      link_local = true
    }
  ]
  ipv6_addresses = [
    {
      prefix = "2006:DB8::/32"
      eui_64 = true
    }
  ]
  load_interval = 30
}
