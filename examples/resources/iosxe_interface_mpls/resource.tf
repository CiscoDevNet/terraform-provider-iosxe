resource "iosxe_interface_mpls" "example" {
  type = "Loopback"
  name = "1"
  ip   = true
  mtu  = "1200"
}
