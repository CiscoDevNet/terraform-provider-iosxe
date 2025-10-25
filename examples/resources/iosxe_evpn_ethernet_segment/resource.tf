resource "iosxe_evpn_ethernet_segment" "example" {
  es_value                 = 1
  df_election_wait_time    = 3
  redundancy_all_active    = false
  redundancy_single_active = true
  identifier_types = [
    {
      type       = 0
      hex_string = "0001.0000.0000.0000.000A"
      system_mac = "00:11:22:33:44:55"
    }
  ]
}
