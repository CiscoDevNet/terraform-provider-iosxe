resource "iosxe_static_route_vrf" "example" {
  vrf = "VRF1"
  routes = [
    {
      prefix = "6.6.6.6"
      mask   = "255.255.255.255"
      next_hops = [
        {
          next_hop  = "7.7.7.7"
          metric    = 10
          global    = false
          name      = "Route1"
          permanent = true
          tag       = 100
        }
      ]
    }
  ]
}
