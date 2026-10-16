resource "iosxe_wireless_rf_profile" "example" {
  name                           = "S2-RF-5GHz-Test"
  band                           = "dot11-5-ghz-band"
  status                         = true
  description                    = "5GHz RF Profile"
  data_rate_1m                   = "apf-tx-rate-basic"
  data_rate_2m                   = "apf-tx-rate-basic"
  data_rate_5_5m                 = "apf-tx-rate-basic"
  data_rate_11m                  = "apf-tx-rate-basic"
  data_rate_6m                   = "apf-tx-rate-basic"
  data_rate_9m                   = "apf-tx-rate-supported"
  data_rate_12m                  = "apf-tx-rate-basic"
  data_rate_18m                  = "apf-tx-rate-supported"
  data_rate_24m                  = "apf-tx-rate-basic"
  data_rate_36m                  = "apf-tx-rate-supported"
  data_rate_48m                  = "apf-tx-rate-supported"
  data_rate_54m                  = "apf-tx-rate-supported"
  band_select_client_rssi        = -70
  band_select_cycle_count        = 3
  band_select_cycle_threshold    = 300
  band_select_expire_dual_band   = 80
  band_select_expire_suppression = 30
  dca_allowed_channels = [
    {
      channel = 36
    }
  ]
  coverage_data_rssi_threshold  = -75
  coverage_voice_rssi_threshold = -75
  coverage_level_global         = 5
  load_balancing_window         = 7
  load_balancing_denial         = 5
  max_clients                   = 100
  tx_power_v1_threshold         = -65
  tx_power_max                  = 30
  tx_power_min                  = 7
  trap_threshold_clients        = 20
  trap_threshold_interference   = 15
  trap_threshold_noise          = -60
}
