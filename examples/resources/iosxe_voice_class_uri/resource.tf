resource "iosxe_voice_class_uri" "example" {
  tag     = "200"
  type    = "sip"
  pattern = "cube.example.com"
}
