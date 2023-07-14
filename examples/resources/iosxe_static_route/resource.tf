resource "iosxe_static_route" "example" {
  prefix = "5.5.5.5"
  mask   = "255.255.255.255"
  next_hops = [
    {
      next_hop  = "6.6.6.6"
      metric    = 10
      global    = false
      name      = "Route1"
      permanent = true
      tag       = 100
    }
  ]
}
