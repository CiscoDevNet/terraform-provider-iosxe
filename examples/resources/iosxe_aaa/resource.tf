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
    }
  ]
  group_tacacsplus = [
    {
      name = "tacacs-group"
      servers = [
        {
          name = "tacacs_10.10.15.12"
        }
      ]
    }
  ]
}
