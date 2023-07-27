resource "iosxe_interface_ospfv3" "example" {
  type                   = "GigabitEthernet"
  name                   = "1"
  network_point_to_point = true
  cost_config_value      = 1000
}
