resource "iosxe_license" "example" {
  boot_level_network_advantage       = true
  boot_level_network_advantage_addon = "dna-advantage"
  smart_transport_type               = "Off"
}
