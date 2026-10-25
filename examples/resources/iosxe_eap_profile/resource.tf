resource "iosxe_eap_profile" "example" {
  name            = "EAP1"
  description     = "My EAP Profile"
  method_fast     = true
  method_gtc      = true
  method_leap     = true
  method_md5      = true
  method_mschapv2 = true
  method_peap     = true
  method_tls      = true
  pki_trustpoint  = "TP1"
}
