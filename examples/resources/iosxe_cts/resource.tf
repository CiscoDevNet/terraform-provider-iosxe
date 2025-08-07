resource "iosxe_cts" "example" {
  authorization_list          = "Tacacs-GROUP"
  sgt                         = 200
  sxp_enable                  = true
  sxp_default_password_type   = "0"
  sxp_default_password_secret = "MySecretPassword"
  sxp_retry_period            = 60
  sxp_connection_peer_ipv4_no_vrf = [
    {
      ip              = "2.2.2.2"
      source_ip       = "3.3.3.3"
      password        = "default"
      connection_mode = "local"
      option          = "listener"
      hold_time       = 60
      max_time        = 300
    }
  ]
  sxp_connection_peer_ipv4_with_vrf = [
    {
      ip              = "4.4.4.4"
      vrf             = "VRF1"
      source_ip       = "5.5.5.5"
      password        = "default"
      connection_mode = "local"
      option          = "listener"
      hold_time       = 60
      max_time        = 300
    }
  ]
  sxp_speaker_hold_time                   = 300
  sxp_listener_hold_min_time              = 60
  sxp_listener_hold_max_time              = 300
  role_based_enforcement_logging_interval = 300
}
