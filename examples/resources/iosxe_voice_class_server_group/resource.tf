resource "iosxe_voice_class_server_group" "example" {
  tag = 200
  ipv4 = [
    {
      address = "192.0.2.10"
    }
  ]
  ipv4_with_port = [
    {
      address = "192.0.2.20"
      port    = 5061
    }
  ]
}
