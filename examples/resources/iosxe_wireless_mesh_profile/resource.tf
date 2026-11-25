resource "iosxe_wireless_mesh_profile" "example" {
  profile_name       = "S1-MESH-OUTDOOR"
  description        = "Outdoor mesh profile"
  convergence_method = "mesh-convergence-fast"
  multicast_mode     = "mesh-multicast-mode-in"
  range              = 20000
  security_mode      = "mesh-security-mode-psk"
}
