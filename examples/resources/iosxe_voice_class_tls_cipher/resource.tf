resource "iosxe_voice_class_tls_cipher" "example" {
  tag = 1
  ciphers = [
    {
      preference   = 1
      cipher_suite = "ECDHE_RSA_AES256_GCM_SHA384"
    }
  ]
}
