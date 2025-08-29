resource "iosxe_aaa" "example" {
  new_model                    = true
  server_radius_dynamic_author = true
  session_id                   = "common"
  server_radius_dynamic_author_clients = [
    {
      ip              = "11.1.1.1"
      server_key_type = "0"
      server_key      = "abcd123"
    }
  ]
  group_server_radius = [
    {
      name = "T-Group"
      server_names = [
        {
          name = "TESTRADIUS"
        }
      ]
      ip_radius_source_interface_loopback = 0
    }
  ]
  group_server_tacacsplus = [
    {
      name = "tacacs-group"
      server_names = [
        {
          name = "tacacs_10.10.15.12"
        }
      ]
      ip_tacacs_source_interface_loopback = 0
      vrf                                 = "VRF1"
    }
  ]
}
