resource "iosxe_interface_loopback" "example" {
  name                       = 201
  description                = "My Interface Description"
  shutdown                   = false
  ip_proxy_arp               = false
  ip_redirects               = false
  ip_unreachables            = false
  vrf_forwarding             = "VRF1"
  ipv4_address               = "200.1.1.1"
  ipv4_address_mask          = "255.255.255.255"
  ip_access_group_in_enable  = true
  ip_access_group_in         = "1"
  ip_access_group_out_enable = true
  ip_access_group_out        = "1"
  ipv6_enable                = true
  ipv6_mtu                   = 1300
  ipv6_address_dhcp          = true
  ipv6_link_local_addresses = [
    {
      address    = "fe80::9656:d028:8652:66b7"
      link_local = true
    }
  ]
  ipv6_addresses = [
    {
      prefix = "2002:DB8::/32"
      eui_64 = true
    }
  ]
  arp_timeout     = 2147
  ip_igmp_version = 3
}
