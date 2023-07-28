resource "iosxe_crypto_ikev2_proposal" "example" {
  ikev2_proposals = [
    {
      name                   = "aws_vpg_2"
      encryption_aes_cbc_256 = true
      group_fourteen         = true
      group_nineteen         = true
      group_twenty           = true
      integrity_sha1         = true
    }
  ]
}
