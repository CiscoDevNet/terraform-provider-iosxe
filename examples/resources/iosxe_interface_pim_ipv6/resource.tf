resource "iosxe_interface_pim_ipv6" "example" {
  type        = "Loopback"
  name        = "100"
  pim         = true
  bfd         = false
  bsr_border  = false
  dr_priority = 10
}
