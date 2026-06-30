resource "iosxe_eigrp" "example" {
  name              = "TOPGEN"
  autonomous_system = 100
  router_id         = "10.255.0.1"
  networks = [
    {
      ip       = "10.10.0.0"
      wildcard = "0.0.255.255"
    }
  ]
  auto_summary = false
  shutdown     = false
}
