resource "iosxe_crypto_ikev2_policy" "example" {
  name           = "aws_vpg_1"
  match_fvrf_any = true
  proposals = [
    {
      proposals = "proposal123"
    }
  ]
}
