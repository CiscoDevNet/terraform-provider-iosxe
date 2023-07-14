resource "iosxe_interface_vlan" "example" {
  name                           = 10
  autostate                      = false
  description                    = "My Interface Description"
  shutdown                       = false
  ip_proxy_arp                   = false
  ip_redirects                   = false
  unreachables                   = false
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
}
