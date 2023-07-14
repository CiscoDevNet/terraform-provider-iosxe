resource "iosxe_system" "example" {
  hostname             = "ROUTER-1"
  ipv6_unicast_routing = true
  ip_source_route      = false
  ip_domain_lookup     = false
  ip_domain_name       = "test.com"
  login_delay          = 10
  login_on_failure     = true
  login_on_failure_log = true
  login_on_success     = true
  login_on_success_log = true
  multicast_routing_vrfs = [
    {
      vrf = "VRF1"
    }
  ]
}
