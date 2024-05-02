resource "iosxe_crypto_ikev2_profile" "example" {
  name                            = "profile1"
  description                     = "My description"
  ivrf                            = "I-VRF of the profile"
  authentication_remote_pre_share = true
  authentication_local_pre_share  = true
  identity_local_key_id           = "key1"
  match_address_local_ip          = "1.2.3.4"
  match_fvrf_any                  = true
  match_identity_remote_ipv4_addresses = [
    {
      address = "1.2.3.4"
      mask    = "255.255.255.0"
    }
  ]
  match_identity_remote_keys = ["key1"]
  keyring_local              = "test"
  dpd_interval               = 10
  dpd_retry                  = 2
  dpd_query                  = "periodic"
  config_exchange_request    = false
}
