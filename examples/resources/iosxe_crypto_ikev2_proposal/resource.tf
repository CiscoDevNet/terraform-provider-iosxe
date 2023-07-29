resource "iosxe_crypto_ikev2_proposal" "example" {
  name                   = "PROPOSAL1"
  encryption_aes_gcm_256 = true
  group_one              = true
  group_two              = true
  group_fourteen         = true
  group_fifteen          = true
  group_sixteen          = true
  group_nineteen         = true
  group_twenty           = true
  group_twenty_one       = true
  group_twenty_four      = true
  integrity_sha1         = true
  prf_sha1               = true
}
