resource "iosxe_vlan_configuration" "example" {
  vlan_id                  = "123"
  evpn_instance_legacy     = 123
  evpn_instance_vni_legacy = 10123
}
