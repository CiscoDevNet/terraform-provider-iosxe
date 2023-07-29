resource "iosxe_interface_mpls" "example" {
  type = "GigabitEthernet"
  name = "1"
  ip   = true
  mtu  = "1200"
}
