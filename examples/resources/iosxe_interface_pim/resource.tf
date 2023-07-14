resource "iosxe_interface_pim" "example" {
  type              = "Loopback"
  name              = "100"
  passive           = false
  dense_mode        = false
  sparse_mode       = true
  sparse_dense_mode = false
  bfd               = false
  border            = false
  bsr_border        = false
  dr_priority       = 10
}
