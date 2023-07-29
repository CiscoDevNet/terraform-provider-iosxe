resource "iosxe_interface_ospfv3" "example" {
  type                   = "Loopback"
  name                   = "1"
  network_point_to_point = true
  cost                   = 1000
}
