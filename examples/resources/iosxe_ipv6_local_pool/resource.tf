resource "iosxe_ipv6_local_pool" "example" {
  name          = "DHCPv6-PD"
  start_address = "2001:db8::/48"
  prefix_length = 64
}
