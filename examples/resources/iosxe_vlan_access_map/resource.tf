resource "iosxe_vlan_access_map" "example" {
  name               = "VAM1"
  sequence           = 10
  match_ipv6_address = ["ACL2"]
  match_ip_address   = ["ACL1"]
  action             = "forward"
}
