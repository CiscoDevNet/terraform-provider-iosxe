resource "iosxe_voice_translation_profile" "example" {
  tag               = "TP_INBOUND"
  translate_called  = 200
  translate_calling = 200
}
