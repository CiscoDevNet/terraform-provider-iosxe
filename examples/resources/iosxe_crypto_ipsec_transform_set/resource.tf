resource "iosxe_crypto_ipsec_transform_set" "example" {
  name        = "TEST"
  esp         = "esp-aes"
  esp_hmac    = "esp-sha-hmac"
  mode_tunnel = true
}
