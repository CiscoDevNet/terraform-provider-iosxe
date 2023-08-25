resource "iosxe_bfd_template_single_hop" "example" {
  name                                       = "singelHop"
  authentication_md5_keychain                = "KEYC1"
  interval_singlehop_v2_mill_unit_min_tx     = 200
  interval_singlehop_v2_mill_unit_min_rx     = 200
  interval_singlehop_v2_mill_unit_multiplier = 4
  echo                                       = true
  dampening_half_time                        = 30
  dampening_unsuppress_time                  = 30
  dampening_suppress_time                    = 100
  dampening_max_suppressing_time             = 60
}
