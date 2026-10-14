resource "iosxe_crypto_map_gdoi" "example" {
  name            = "GETVPN-MAP"
  sequence_number = 10
  gdoi            = true
  set_group       = "GETVPN-GROUP"
}
