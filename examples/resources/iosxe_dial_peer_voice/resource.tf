resource "iosxe_dial_peer_voice" "example" {
  tag                                       = 100
  type                                      = "voip"
  description                               = "Outbound to CUCM"
  destination_pattern                       = "+1408......."
  preference                                = 1
  huntstop                                  = false
  shutdown                                  = false
  vad                                       = true
  srtp                                      = true
  session_protocol                          = "sipv2"
  session_server_group                      = 10
  voice_class_sip_options_keepalive_profile = 100
  incoming_called_number                    = "+1.*"
  destination_dpg                           = 100
  voice_class_sip_tenant                    = 100
  voice_class_sip_profiles                  = 100
  voice_class_sip_early_offer_forced        = true
  dtmf_relay_rtp_nte                        = true
  ip_qos_dscp_signaling                     = "cs3"
  ip_qos_dscp_signaling_enable              = true
  max_conn                                  = 100
  max_conn_exempt_local_media               = true
  answer_address                            = "+15551234567"
}
