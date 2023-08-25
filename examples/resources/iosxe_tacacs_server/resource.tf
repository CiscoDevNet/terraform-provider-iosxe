resource "iosxe_tacacs_server" "example" {
  name                = "tacacs_10.10.15.13"
  tacacs_address_ipv4 = "10.10.15.13"
  timeout             = 4
  key_key             = "123"
}
