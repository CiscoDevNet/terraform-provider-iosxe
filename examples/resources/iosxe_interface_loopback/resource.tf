resource "iosxe_interface_loopback" "example" {
  name                       = 100
  description                = "My Interface Description"
  shutdown                   = false
  ip_proxy_arp               = false
  ip_redirects               = false
  unreachables               = false
  vrf_forwarding             = "VRF1"
  ipv4_address               = "200.1.1.1"
  ipv4_address_mask          = "255.255.255.255"
  ip_access_group_in         = "1"
  ip_access_group_in_enable  = true
  ip_access_group_out        = "1"
  ip_access_group_out_enable = true
}
