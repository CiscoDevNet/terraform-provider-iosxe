resource "iosxe_vlan_group" "example" {
  name       = "group"
  vlan_lists = [1]
}
