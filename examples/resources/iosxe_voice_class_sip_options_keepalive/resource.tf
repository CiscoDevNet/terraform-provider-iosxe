resource "iosxe_voice_class_sip_options_keepalive" "example" {
  tag               = 200
  description       = "ITSP keepalive profile"
  up_interval       = 60
  down_interval     = 30
  retry             = 5
  transport_tcp     = true
  transport_tcp_tls = true
}
