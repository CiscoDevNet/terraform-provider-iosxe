resource "iosxe_vlan_filter" "example" {
  word       = "VAM1"
  vlan_lists = [1]
}
