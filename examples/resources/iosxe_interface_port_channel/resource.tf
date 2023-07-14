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
}
