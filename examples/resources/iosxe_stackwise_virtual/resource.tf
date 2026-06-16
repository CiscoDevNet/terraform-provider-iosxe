resource "iosxe_stackwise_virtual" "example" {
  domain                                         = 10
  dual_active_detection_pagp                     = true
  dual_active_detection_pagp_trust_channel_group = 1
}
