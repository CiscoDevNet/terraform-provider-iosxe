resource "iosxe_crypto_ipsec_profile" "example" {
  profile = [
    {
      name              = "vpn200"
      set_transform_set = ["test123"]
    }
  ]
}
