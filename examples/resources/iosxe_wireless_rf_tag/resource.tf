resource "iosxe_wireless_rf_tag" "example" {
  name            = "S1-RF-Tag-Test"
  description     = "Test RF Tag"
  rf_policy_24ghz = "Typical_Client_Density_rf_24gh"
  rf_policy_5ghz  = "Typical_Client_Density_rf_5gh"
  rf_policy_6ghz  = "S1-RF-6GHZ"
}
