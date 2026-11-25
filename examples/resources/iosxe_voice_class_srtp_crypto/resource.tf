resource "iosxe_voice_class_srtp_crypto" "example" {
  tag = 200
  cryptos = [
    {
      preference   = 1
      cipher_suite = "AEAD_AES_256_GCM"
    }
  ]
}
