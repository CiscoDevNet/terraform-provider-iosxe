resource "iosxe_vlan_configuration" "example" {
  vlan_id           = 123
  evpn_instance     = 123
  evpn_instance_vni = 10123
}
