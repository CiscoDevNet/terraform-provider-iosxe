resource "iosxe_wireless_wlan_profile" "example" {
  profile_name                    = "S1-CORP-PSK"
  wlan_id                         = 10
  ssid                            = "S1-CORP-PSK"
  wlan_status                     = true
  broadcast_ssid                  = false
  ccx_aironet_ie                  = true
  peer_blocking                   = "p2p-blocking-action-drop"
  security_wpa_akm_dot1x          = false
  security_wpa_akm_psk            = true
  security_wpa_akm_ft_psk         = true
  security_wpa_psk_key_format     = "key-hex"
  security_wpa_psk_key_encryption = "clear"
  security_wpa_psk_key            = "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"
  security_web_auth               = true
  security_ft                     = "dot11r-enabled"
  security_ft_over_the_ds         = true
  load_balance                    = true
  wmm                             = "apf-vap-wme-required"
}
