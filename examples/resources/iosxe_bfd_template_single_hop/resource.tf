resource "iosxe_bfd_template_single_hop" "example" {
  name                           = "SH-TEMPLATE-1"
  authentication_md5_keychain    = "KEYC1"
  echo                           = true
  dampening_half_time            = 30
  dampening_unsuppress_time      = 30
  dampening_suppress_time        = 100
  dampening_max_suppressing_time = 60
}
