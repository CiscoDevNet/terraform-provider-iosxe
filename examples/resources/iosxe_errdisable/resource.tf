resource "iosxe_errdisable" "example" {
  detect_cause_dhcp_rate_limit           = true
  detect_cause_dtp_flap                  = true
  detect_cause_l2ptguard                 = true
  detect_cause_link_flap                 = true
  detect_cause_pppoe_ia_rate_limit       = true
  detect_cause_loopdetect                = true
  flap_setting_cause_dtp_flap_max_flaps  = 80
  flap_setting_cause_dtp_flap_time       = 90
  flap_setting_cause_link_flap_max_flaps = 80
  flap_setting_cause_link_flap_time      = 90
  flap_setting_cause_pagp_flap_max_flaps = 80
  flap_setting_cause_pagp_flap_time      = 90
  recovery_interval                      = 855
  recovery_cause_arp_inspection          = true
  recovery_cause_bpduguard               = true
  recovery_cause_dhcp_rate_limit         = true
  recovery_cause_dtp_flap                = true
  recovery_cause_l2ptguard               = true
  recovery_cause_link_flap               = true
  recovery_cause_port_mode_failure       = true
  recovery_cause_pppoe_ia_rate_limit     = true
  recovery_cause_psp                     = true
  recovery_cause_psecure_violation       = true
  recovery_cause_security_violation      = true
  recovery_cause_udld                    = true
  recovery_cause_loopdetect              = true
}
