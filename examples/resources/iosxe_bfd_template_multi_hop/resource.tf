resource "iosxe_bfd_template_multi_hop" "example" {
  name                           = "T11"
  echo                           = true
  authentication_md5_keychain    = "KEYNAME"
  dampening_half_time            = 21
  dampening_unsuppress_time      = 1800
  dampening_suppress_time        = 1900
  dampening_max_suppressing_time = 70
}
