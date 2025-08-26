resource "iosxe_crypto_pki" "example" {
  trustpoints = [
    {
      id                       = "trustpoint1"
      enrollment_pkcs12        = true
      enrollment_pkcs12_legacy = true
      revocation_check         = ["none"]
    }
  ]
}
