resource "iosxe_interface_bdi" "example" {
  name        = "100"
  mac_address = "0000.11AA.22BB"
  ip_mtu      = 1400
}
