resource "iosxe_flow_monitor" "example" {
  flow_record = [
    {
      name                             = "FNF-input"
      description                      = "IPv4-NetFlow"
      match_ipv4_source_address        = true
      match_ipv4_destination_address   = true
      match_ipv4_protocol              = true
      match_ipv4_tos                   = true
      match_transport_source_port      = true
      match_transport_destination_port = true
      match_interface_input            = true
      match_flow_direction             = true
      collect_interface_output         = true
      collect_counter_bytes_long       = true
      collect_counter_packets_long     = true
      collect_transport_tcp_flags      = true
      collect_timestamp_absolute_first = true
      collect_timestamp_absolute_last  = true
    }
  ]
  flow_exporter = [
    {
      name                                            = "Scrutinizer"
      description                                     = "Export to Scrutinizer"
      destination_destination_choice_ipdest_ipdest_ip = "1.1.1.1"
      source_loopback                                 = 1474
      transport_udp                                   = 655
      template_data_timeout                           = 60
    }
  ]
  flow_monitor = [
    {
      name        = "Scrut_mon_input"
      description = "IPv4 FNF ingress exports"
      flow_monitor_exporter = [
        {
          name = "Scrutinizer"
        }
      ]
      cache_timeout_active = 60
      record_type          = "FNF_Input"
    }
  ]
}
