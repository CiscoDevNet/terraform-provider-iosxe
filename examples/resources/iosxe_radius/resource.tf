resource "iosxe_radius" "example" {
  name                                                            = "radius_10.10.15.12"
  radius_host_address_ipv4                                        = "10.10.15.12"
  address_auth_port                                               = 1813
  address_acct_port                                               = 1812
  timeout                                                         = 4
  retransmit                                                      = 3
  key_key                                                         = "123"
  automate_tester_username                                        = "dummy"
  automate_tester_ignore_acct_port                                = true
  automate_tester_type_of_testing_probe_on_config_probe_on_config = true
  pac_key_key                                                     = "testtest1"
  pac_key_encryption                                              = "0"
}
