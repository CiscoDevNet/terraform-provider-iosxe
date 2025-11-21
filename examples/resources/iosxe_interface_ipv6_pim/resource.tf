resource "iosxe_interface_ipv6_pim" "example" {
  type        = "Loopback"
  name        = "100"
  pim         = true
  bfd         = false
  bsr_border  = false
  dr_priority = 10
}
