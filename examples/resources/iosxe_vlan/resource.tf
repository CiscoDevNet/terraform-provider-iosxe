resource "iosxe_vlan" "example" {
  vlan_id  = 123
  name     = "Vlan123"
  shutdown = false
}
