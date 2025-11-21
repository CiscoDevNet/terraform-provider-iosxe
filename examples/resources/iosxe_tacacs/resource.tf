resource "iosxe_tacacs" "example" {
  name         = "tacacs_10.10.15.13"
  address_ipv4 = "10.10.15.13"
  timeout      = 4
  port         = 490
  encryption   = "0"
  key          = "123"
}
