resource "iosxe_vlan_group" "example" {
  name       = "GROUP1"
  vlan_lists = [1]
}
