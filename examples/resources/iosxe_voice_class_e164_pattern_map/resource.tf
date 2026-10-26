resource "iosxe_voice_class_e164_pattern_map" "example" {
  tag = 200
  patterns = [
    {
      pattern = "+1408.......T"
    }
  ]
}
