resource "iosxe_crypto_ipsec_transform_set" "example" {
  tag                                      = "TEST"
  esp                                      = "esp-aes"
  esp_hmac                                 = "esp-sha-hmac"
  mode_mode_type_tunnel_case_tunnel_choice = true
}
