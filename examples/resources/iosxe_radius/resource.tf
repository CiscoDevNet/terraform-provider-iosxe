resource "iosxe_radius" "example" {
  name                     = "radius_10.10.15.12"
  radius_host_address_ipv4 = "10.10.15.12"
  address_auth_port        = 1813
  timeout                  = 4
  retransmit               = 3
  key_key                  = "123"
}
