resource "iosxe_crypto_ipsec_profile" "example" {
  name                                                              = "vpn200"
  set_transform_set                                                 = ["TS1"]
  set_isakmp_profile_ikev2_profile_ikev2_profile_case_ikev2_profile = "vpn300"
}
