resource "iosxe_logging_ipv4_host_transport" "example" {
  ipv4_host = "2.2.2.2"
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
