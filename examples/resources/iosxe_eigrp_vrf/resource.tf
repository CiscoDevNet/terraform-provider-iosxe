resource "iosxe_eigrp_vrf" "example" {
  name              = "TOPGEN"
  vrf               = "VRF1"
  autonomous_system = 100
  router_id         = "10.255.1.1"
  networks = [
    {
      ip       = "10.20.0.0"
      wildcard = "0.0.255.255"
    }
  ]
  auto_summary = false
  shutdown     = false
}
