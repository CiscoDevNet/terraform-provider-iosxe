resource "iosxe_dhcp" "example" {
  relay_information_trust_all      = false
  relay_information_option_default = false
  relay_information_option_vpn     = true
  snooping                         = true
  snooping_vlans = [
    {
      vlan_id = "3-4"
    }
  ]
}
