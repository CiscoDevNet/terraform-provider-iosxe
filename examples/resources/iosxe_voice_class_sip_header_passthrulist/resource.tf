resource "iosxe_voice_class_sip_header_passthrulist" "example" {
  tag                 = 100
  passthru_headers    = ["Call-Info"]
  passthru_hdr_unsupp = true
}
