resource "iosxe_sip_ua" "example" {
  transport_udp               = true
  transport_tcp               = true
  transport_tcp_tls           = true
  sip_server                  = "ipv4:10.0.0.1"
  redirection                 = true
  handle_replaces             = true
  reason_header_override      = true
  host_registrar              = true
  max_forwards                = 70
  disable_early_media         = "180"
  retry_invite                = 6
  retry_register              = 6
  retry_rel1xx                = 6
  retry_response              = 6
  retry_keepalive             = 6
  timers_expires              = 180000
  timers_trying               = 500
  timers_connection_aging_tls = 720
  timers_options              = 500
  timers_keepalive_active     = 120
  timers_keepalive_down       = 30
  connection_reuse            = true
  connection_reuse_via_port   = true
}
