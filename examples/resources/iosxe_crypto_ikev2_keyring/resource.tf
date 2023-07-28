resource "iosxe_crypto_ikev2_keyring" "example" {
  name = "aws_vpg_1"
  peer = [
    {
      name                                                                   = "aws_vpg_1"
      ipv4_address                                                           = "1.2.3.4"
      ipv4_mask                                                              = "255.255.255.248"
      identity_identity_key_id_key_id_number                                 = "vpn200"
      pre_shared_key_local_option_encryption_hex_encryption_case_encryption  = "6"
      pre_shared_key_local_option_encryption_hex_encryption_case_key         = "cisco123"
      pre_shared_key_remote_option_encryption_hex_encryption_case_encryption = "6"
      pre_shared_key_remote_option_encryption_hex_encryption_case_key        = "cisco123"
    }
  ]
}
