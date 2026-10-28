resource "iosxe_wireless_multi_bssid_profile" "example" {
  name                    = "MBSSID1"
  description             = "6 GHz Multi BSSID Profile"
  dot11ax_downlink_ofdma  = false
  dot11ax_uplink_ofdma    = false
  dot11ax_downlink_mumimo = false
  dot11ax_uplink_mumimo   = false
  dot11ax_target_waketime = true
  dot11ax_twt_broadcast   = true
}
