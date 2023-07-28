resource "iosxe_crypto_ikev2_profile" "example" {
  name = "aws_vpg_1"
  match_identity_remote_address_ipv4 = [
    {
      ipv4_address = "1.2.3.4"
      ipv4_mask    = "255.255.255.0"
    }
  ]
  authentication_remote_pre_share = true
  authentication_local_pre_share  = true
  keyring_local_name              = "test"
  dpd_interval                    = 10
  dpd_retry                       = 2
  dpd_query                       = "periodic"
  config_exchange_request_1       = false
}
