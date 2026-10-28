resource "iosxe_voice_port" "example" {
  port        = "0/1/0:23"
  description = "PSTN PRI trunk"
}
