resource "iosxe_flow_record" "example" {
  name                                         = "FNF1"
  description                                  = "My flow record"
  match_ipv4_source_address                    = true
  match_ipv4_destination_address               = true
  match_ipv4_protocol                          = true
  match_ipv4_tos                               = true
  match_transport_source_port                  = true
  match_transport_destination_port             = true
  match_interface_input                        = true
  match_flow_direction                         = true
  collect_interface_output                     = true
  collect_counter_bytes_long                   = true
  collect_counter_packets_long                 = true
  collect_transport_tcp_flags                  = true
  collect_timestamp_absolute_first             = true
  collect_timestamp_absolute_last              = true
  match_datalink_mac_source_address_input      = true
  match_datalink_mac_destination_address_input = true
  match_ipv4_ttl                               = true
}
