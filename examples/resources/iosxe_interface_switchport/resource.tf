resource "iosxe_interface_switchport" "example" {
  type                          = "GigabitEthernet"
  name                          = "1/0/3"
  mode_access                   = false
  mode_dot1q_tunnel             = false
  mode_private_vlan_trunk       = false
  mode_private_vlan_host        = false
  mode_private_vlan_promiscuous = false
  mode_trunk                    = true
  nonegotiate                   = false
  access_vlan                   = "100"
  trunk_allowed_vlans           = "100,101"
  trunk_native_vlan             = 100
  host                          = false
}
