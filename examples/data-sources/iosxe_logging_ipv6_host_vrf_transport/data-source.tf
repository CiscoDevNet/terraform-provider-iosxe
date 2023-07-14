data "iosxe_logging_ipv6_host_vrf_transport" "example" {
  ipv6_host = "2001::1"
  vrf       = "VRF1"
}
