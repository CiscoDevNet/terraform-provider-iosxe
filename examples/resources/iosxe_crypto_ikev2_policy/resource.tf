resource "iosxe_crypto_ikev2_policy" "example" {
  name                   = "policy1"
  match_address_local_ip = ["1.2.3.4"]
  match_fvrf_any         = true
  proposals = [
    {
      proposals = "proposal123"
    }
  ]
}
