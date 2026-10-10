resource "iosxe_dial_peer_voice" "example" {
  tag                                       = 100
  type                                      = "voip"
  description                               = "Outbound to CUCM"
  preference                                = 1
  huntstop                                  = false
  shutdown                                  = false
  vad                                       = true
  session_protocol                          = "sipv2"
  session_server_group                      = "10"
  voice_class_sip_options_keepalive_profile = "100"
  destination_dpg                           = "100"
  voice_class_sip_tenant                    = "100"
  voice_class_sip_profiles                  = "100"
}
