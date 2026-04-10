resource "iosxe_interface_ospfv3" "example" {
  type                             = "Loopback"
  name                             = "1"
  network_type_broadcast           = false
  network_type_non_broadcast       = false
  network_type_point_to_multipoint = false
  network_type_point_to_point      = true
  bfd                              = true
  cost                             = 1000
  dead_interval                    = 30
  hello_interval                   = 5
  mtu_ignore                       = true
  priority                         = 10
}
