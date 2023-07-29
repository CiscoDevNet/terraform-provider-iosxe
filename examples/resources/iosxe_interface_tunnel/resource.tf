resource "iosxe_interface_tunnel" "example" {
  name                 = 90
  ipv6_enable          = true
  ipv6_mtu             = 1300
  ipv6_nd_suppress_all = true
  ipv6_address_dhcp    = true
  ipv6_link_local_address = [
    {
      address    = "fe80::9656:d028:8652:66b6"
      link_local = true
    }
  ]
  ipv6_prefix_list_address = [
    {
      prefix = "2001:DB8::/32"
      eui_64 = true
    }
  ]
  crypto_ipsec_df_bit     = "clear"
  ip_primary_address      = "170.254.10.2"
  ip_primary_address_mask = "255.255.255.252"
}
