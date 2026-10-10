resource "iosxe_voice_register_pool" "example" {
  tag                             = 1
  id_network_address              = "192.0.2.0"
  id_network_mask                 = "255.255.255.0"
  call_forward_b2bua_all          = "1001"
  call_forward_b2bua_busy         = "1001"
  call_forward_b2bua_mailbox      = "1001"
  call_forward_b2bua_noan         = "1001"
  call_forward_b2bua_noan_timeout = 20
  dtmf_relay_rtp_nte              = true
  dtmf_relay_sip_kpml             = true
  dtmf_relay_sip_notify           = true
}
