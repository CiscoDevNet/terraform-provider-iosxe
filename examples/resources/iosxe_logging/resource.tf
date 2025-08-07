resource "iosxe_logging" "example" {
  monitor_severity  = "informational"
  buffered_size     = 16000
  buffered_severity = "informational"
  console_severity  = "informational"
  facility          = "local0"
  history_size      = 100
  history_severity  = "informational"
  trap              = true
  trap_severity     = "informational"
  origin_id_type    = "hostname"
  source_interface  = "Loopback0"
  source_interfaces_vrf = [
    {
      vrf            = "VRF1"
      interface_name = "Loopback100"
    }
  ]
  ipv4_hosts = [
    {
      ipv4_host = "1.1.1.1"
    }
  ]
  ipv4_hosts_transport = [
    {
      ipv4_host = "1.1.1.1"
      transport_udp_ports = [
        {
          port_number = 10000
        }
      ]
      transport_tcp_ports = [
        {
          port_number = 10001
        }
      ]
      transport_tls_ports = [
        {
          port_number = 10002
        }
      ]
    }
  ]
  ipv4_vrf_hosts = [
    {
      ipv4_host = "1.1.1.1"
      vrf       = "VRF1"
    }
  ]
  ipv4_vrf_hosts_transport = [
    {
      ipv4_host = "1.1.1.1"
      vrf       = "VRF1"
      transport_udp_ports = [
        {
          port_number = 10000
        }
      ]
      transport_tcp_ports = [
        {
          port_number = 10001
        }
      ]
      transport_tls_ports = [
        {
          port_number = 10002
        }
      ]
    }
  ]
  ipv6_hosts = [
    {
      ipv6_host = "2001::1"
    }
  ]
  ipv6_hosts_transport = [
    {
      ipv6_host = "2001::1"
      transport_udp_ports = [
        {
          port_number = 10000
        }
      ]
      transport_tcp_ports = [
        {
          port_number = 10001
        }
      ]
      transport_tls_ports = [
        {
          port_number = 10002
        }
      ]
    }
  ]
  ipv6_vrf_hosts = [
    {
      ipv6_host = "2001::1"
      vrf       = "VRF1"
    }
  ]
  ipv6_vrf_hosts_transport = [
    {
      ipv6_host = "2001::1"
      vrf       = "VRF1"
      transport_udp_ports = [
        {
          port_number = 10000
        }
      ]
      transport_tcp_ports = [
        {
          port_number = 10001
        }
      ]
      transport_tls_ports = [
        {
          port_number = 10002
        }
      ]
    }
  ]
}
