resource "iosxe_igmp_snooping" "example" {
  querier                         = true
  querier_entry_version           = 2
  querier_entry_max_response_time = 10
  querier_entry_timer_expiry      = 120
}
