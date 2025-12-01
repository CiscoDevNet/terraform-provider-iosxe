resource "iosxe_radius" "example" {
  name                             = "radius_10.10.15.12"
  ipv4_address                     = "10.10.15.12"
  authentication_port              = 1813
  accounting_port                  = 1812
  timeout                          = 4
  retransmit                       = 3
  key                              = "123"
  key_encryption                   = "0"
  automate_tester_username         = "dummy"
  automate_tester_ignore_acct_port = true
  automate_tester_ignore_auth_port = true
  automate_tester_probe_on_config  = true
}
