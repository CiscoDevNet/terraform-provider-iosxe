resource "iosxe_vtp" "example" {
  file                     = "TEST"
  version                  = 3
  interface                = "Gi1/0/1"
  password                 = "test123"
  password_hidden          = true
  domain                   = "TESTDOMAIN"
  mode_client_mst          = true
  mode_client_unknown      = true
  mode_client_vlan         = true
  mode_off_mst             = true
  mode_off_vlan            = true
  mode_server_mst          = true
  mode_server_unknown      = true
  mode_server_vlan         = true
  mode_transparent_mst     = true
  mode_transparent_unknown = true
  mode_transparent_vlan    = true
}
