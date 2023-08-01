resource "iosxe_crypto_ikev2_keyring" "example" {
  name = "keyring1"
  peers = [
    {
      name                             = "peer1"
      description                      = "My description"
      ipv4_address                     = "1.2.3.4"
      ipv4_mask                        = "255.255.255.248"
      identity_key_id                  = "key1"
      pre_shared_key_local_encryption  = "6"
      pre_shared_key_local             = "cisco123"
      pre_shared_key_remote_encryption = "6"
      pre_shared_key_remote            = "cisco123"
    }
  ]
}
