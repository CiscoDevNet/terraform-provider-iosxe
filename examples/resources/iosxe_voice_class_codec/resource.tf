resource "iosxe_voice_class_codec" "example" {
  tag = 1
  preferences = [
    {
      preference = 1
      codec      = "g711ulaw"
    }
  ]
}
