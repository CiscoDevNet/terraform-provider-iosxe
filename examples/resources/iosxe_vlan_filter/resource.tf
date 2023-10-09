resource "iosxe_vlan_filter" "example" {
  word       = "f1"
  vlan_lists = [1]
}
