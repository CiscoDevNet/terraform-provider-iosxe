resource "iosxe_vlan_access_map" "example" {
  name               = "accessmap1"
  value              = 1000
  match_ipv6_address = ["ipv6_address1"]
  match_ip_address   = ["ip_address1"]
  action             = "forward"
}
