resource "iosxe_radius" "example" {
  name                = "radius_10.10.15.12"
  ipv4_address        = "10.10.15.12"
  authentication_port = 1813
  timeout             = 4
  retransmit          = 3
  key                 = "123"
}
