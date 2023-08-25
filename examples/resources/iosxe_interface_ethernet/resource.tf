resource "iosxe_interface_ethernet" "example" {
  type                           = "GigabitEthernet"
  name                           = "3"
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
  source_template = [
    {
      template_name = "TEMP1"
      merge         = false
    }
  ]
  bfd_enable              = true
  bfd_local_address       = "1.2.3.4"
  bfd_interval            = 50
  bfd_interval_min_rx     = 50
  bfd_interval_multiplier = 3
}
