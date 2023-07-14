resource "iosxe_interface_port_channel_subinterface" "example" {
  name                        = "10.666"
  description                 = "My Interface Description"
  shutdown                    = false
  ip_proxy_arp                = false
  ip_redirects                = false
  unreachables                = false
  vrf_forwarding              = "VRF1"
  ipv4_address                = "192.0.2.2"
  ipv4_address_mask           = "255.255.255.0"
  encapsulation_dot1q_vlan_id = 666
  ip_access_group_in          = "1"
  ip_access_group_in_enable   = true
  ip_access_group_out         = "1"
  ip_access_group_out_enable  = true
  helper_addresses = [
    {
      address = "10.10.10.10"
      global  = false
    }
  ]
}
