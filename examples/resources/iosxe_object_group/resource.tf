resource "iosxe_object_group" "example" {
  fqdn = [
    {
      name        = "FQDN_GROUP_1"
      description = "My FQDN object group"
      patterns = [
        {
          fqdn_pattern = ".*\\.cisco\\.com"
        }
      ]
    }
  ]
  network = [
    {
      name        = "NETWORK_GROUP_1"
      description = "My network object group"
      hosts = [
        {
          ipv4_host = "10.1.1.1"
        }
      ]
      network_addresses = [
        {
          ipv4_address = "10.1.2.0"
          ipv4_mask    = "255.255.255.0"
        }
      ]
      address_ranges = [
        {
          start = "10.1.3.1"
          end   = "10.1.3.10"
        }
      ]
    }
  ]
  service = [
    {
      name                      = "SERVICE_GROUP_1"
      description               = "My service object group"
      protocol_numbers          = []
      ahp                       = true
      eigrp                     = false
      esp                       = false
      gre                       = false
      icmp                      = false
      igmp                      = false
      ip                        = false
      ipinip                    = false
      nos                       = false
      ospf                      = false
      pcp                       = false
      pim                       = false
      tcp                       = false
      udp                       = false
      icmp_port_number          = 8
      icmp_alternate_address    = false
      icmp_conversion_error     = false
      icmp_echo                 = false
      icmp_echo_reply           = false
      icmp_information_reply    = false
      icmp_information_request  = false
      icmp_mask_reply           = false
      icmp_mask_request         = false
      icmp_mobile_redirect      = false
      icmp_parameter_problem    = false
      icmp_redirect             = false
      icmp_router_advertisement = false
      icmp_router_solicitation  = false
      icmp_source_quench        = false
      icmp_time_exceeded        = false
      icmp_timestamp_reply      = false
      icmp_timestamp_request    = false
      icmp_traceroute           = false
      icmp_unreachable          = false
      tcp_dst_port_list_op = [
        {
          operator = "eq"
          port     = "www"
        }
      ]
      tcp_dst_port_list = [
        {
          port = "www"
        }
      ]
      tcp_dst_port_ranges = [
        {
          min_port = "1024"
          max_port = "2048"
        }
      ]
      tcp_src_port_list_op = [
        {
          operator = "eq"
          port     = "8080"
        }
      ]
      tcp_src_port_list = [
        {
          port = "8080"
        }
      ]
      tcp_src_port_ranges = [
        {
          min_port = "4000"
          max_port = "5000"
        }
      ]
      tcp_src_dst_port_list_op = [
        {
          src_operator = "eq"
          src_port     = "8080"
          dst_operator = "eq"
          dst_port     = "www"
        }
      ]
      tcp_src_dst_port_list = [
        {
          src_port = "8080"
          dst_port = "www"
        }
      ]
      tcp_src_dst_port_list_src_op = [
        {
          src_operator = "eq"
          src_port     = "8080"
          dst_port     = "www"
        }
      ]
      tcp_src_dst_port_list_dst_op = [
        {
          src_port     = "8080"
          dst_operator = "eq"
          dst_port     = "www"
        }
      ]
      tcp_src_range_dst_port_list_op = [
        {
          src_min_port = "1024"
          src_max_port = "2048"
          operator     = "eq"
          dst_port     = "www"
        }
      ]
      tcp_src_range_dst_port_list = [
        {
          src_min_port = "1024"
          src_max_port = "2048"
          dst_port     = "www"
        }
      ]
      tcp_src_dst_range_port_list_op = [
        {
          operator     = "eq"
          src_port     = "8080"
          dst_min_port = "3000"
          dst_max_port = "4000"
        }
      ]
      tcp_src_dst_range_port_list = [
        {
          src_port     = "8080"
          dst_min_port = "3000"
          dst_max_port = "4000"
        }
      ]
      tcp_src_range_dst_range_port_list = [
        {
          src_min_port = "1024"
          src_max_port = "2048"
          dst_min_port = "3000"
          dst_max_port = "4000"
        }
      ]
      udp_dst_port_list_op = [
        {
          operator = "eq"
          port     = "syslog"
        }
      ]
      udp_dst_port_list = [
        {
          port = "syslog"
        }
      ]
      udp_dst_port_ranges = [
        {
          min_port = "1024"
          max_port = "2048"
        }
      ]
      udp_src_port_list_op = [
        {
          operator = "eq"
          port     = "5000"
        }
      ]
      udp_src_port_list = [
        {
          port = "5000"
        }
      ]
      udp_src_port_ranges = [
        {
          min_port = "4000"
          max_port = "5000"
        }
      ]
      udp_src_dst_port_list_op = [
        {
          src_operator = "eq"
          src_port     = "5000"
          dst_operator = "eq"
          dst_port     = "syslog"
        }
      ]
      udp_src_dst_port_list = [
        {
          src_port = "5000"
          dst_port = "syslog"
        }
      ]
      udp_src_dst_port_list_src_op = [
        {
          src_operator = "eq"
          src_port     = "5000"
          dst_port     = "syslog"
        }
      ]
      udp_src_dst_port_list_dst_op = [
        {
          src_port     = "5000"
          dst_operator = "eq"
          dst_port     = "syslog"
        }
      ]
      udp_src_range_dst_port_list_op = [
        {
          src_min_port = "1024"
          src_max_port = "2048"
          operator     = "eq"
          dst_port     = "syslog"
        }
      ]
      udp_src_range_dst_port_list = [
        {
          src_min_port = "1024"
          src_max_port = "2048"
          dst_port     = "syslog"
        }
      ]
      udp_src_dst_range_port_list_op = [
        {
          operator     = "eq"
          src_port     = "5000"
          dst_min_port = "3000"
          dst_max_port = "4000"
        }
      ]
      udp_src_dst_range_port_list = [
        {
          src_port     = "5000"
          dst_min_port = "3000"
          dst_max_port = "4000"
        }
      ]
      udp_src_range_dst_range_port_list = [
        {
          src_min_port = "1024"
          src_max_port = "2048"
          dst_min_port = "3000"
          dst_max_port = "4000"
        }
      ]
      tcp_udp_dst_port_list_op = [
        {
          operator = "eq"
          port     = "3389"
        }
      ]
      tcp_udp_dst_port_list = [
        {
          port = "3389"
        }
      ]
      tcp_udp_dst_port_ranges = [
        {
          min_port = "1024"
          max_port = "2048"
        }
      ]
      tcp_udp_src_port_list_op = [
        {
          operator = "eq"
          port     = "5000"
        }
      ]
      tcp_udp_src_port_list = [
        {
          port = "5000"
        }
      ]
      tcp_udp_src_port_ranges = [
        {
          min_port = "4000"
          max_port = "5000"
        }
      ]
      tcp_udp_src_dst_port_list_op = [
        {
          src_operator = "eq"
          src_port     = "5000"
          dst_operator = "eq"
          dst_port     = "3389"
        }
      ]
      tcp_udp_src_dst_port_list = [
        {
          src_port = "5000"
          dst_port = "3389"
        }
      ]
      tcp_udp_src_dst_port_list_src_op = [
        {
          src_operator = "eq"
          src_port     = "5000"
          dst_port     = "3389"
        }
      ]
      tcp_udp_src_dst_port_list_dst_op = [
        {
          src_port     = "5000"
          dst_operator = "eq"
          dst_port     = "3389"
        }
      ]
      tcp_udp_src_range_dst_port_list_op = [
        {
          src_min_port = "1024"
          src_max_port = "2048"
          operator     = "eq"
          dst_port     = "3389"
        }
      ]
      tcp_udp_src_range_dst_port_list = [
        {
          src_min_port = "1024"
          src_max_port = "2048"
          dst_port     = "3389"
        }
      ]
      tcp_udp_src_dst_range_port_list_op = [
        {
          operator     = "eq"
          src_port     = "5000"
          dst_min_port = "3000"
          dst_max_port = "4000"
        }
      ]
      tcp_udp_src_dst_range_port_list = [
        {
          src_port     = "5000"
          dst_min_port = "3000"
          dst_max_port = "4000"
        }
      ]
      tcp_udp_src_range_dst_range_port_list = [
        {
          src_min_port = "1024"
          src_max_port = "2048"
          dst_min_port = "3000"
          dst_max_port = "4000"
        }
      ]
    }
  ]
}
