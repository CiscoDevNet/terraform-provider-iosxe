resource "iosxe_crypto_ikev2_proposal" "example" {
  name                   = "PROPOSAL1"
  encryption_aes_cbc_256 = true
  group_sixteen          = true
  integrity_sha256       = true
}
