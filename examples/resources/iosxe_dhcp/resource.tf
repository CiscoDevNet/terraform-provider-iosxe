resource "iosxe_dhcp" "example" {
  relay_information_trust_all                 = false
  relay_information_option_default            = false
  relay_information_option_vpn                = true
  snooping                                    = true
  snooping_information_option                 = true
  snooping_information_option_allow_untrusted = true
}
