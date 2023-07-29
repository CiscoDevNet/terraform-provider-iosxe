resource "iosxe_crypto_ipsec_transform_set" "example" {
  tag         = "TEST"
  esp         = "esp-aes"
  esp_hmac    = "esp-sha-hmac"
  mode_tunnel = true
}
