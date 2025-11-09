resource "iosxe_evpn_ethernet_segment" "example" {
  es_value                 = 1
  df_election_wait_time    = 3
  redundancy_all_active    = false
  redundancy_single_active = true
  identifier_types = [
    {
      type       = 3
      system_mac = "0011.2233.4455"
    }
  ]
}
