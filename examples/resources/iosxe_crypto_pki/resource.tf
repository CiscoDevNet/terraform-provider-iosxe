resource "iosxe_crypto_pki" "example" {
  trustpoints = [
    {
      id                = "trustpoint1"
      enrollment_pkcs12 = true
      revocation_check  = ["none"]
      hash              = "sha256"
    }
  ]
}
