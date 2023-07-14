data "iosxe_static_route" "example" {
  prefix = "5.5.5.5"
  mask   = "255.255.255.255"
}
