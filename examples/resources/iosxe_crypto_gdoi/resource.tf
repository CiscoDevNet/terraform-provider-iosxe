resource "iosxe_crypto_gdoi" "example" {
  name                                 = "GETVPN-GROUP"
  identity_number                      = 100
  server_local                         = true
  server_local_address_ipv4            = "198.51.100.1"
  server_local_rekey_transport_unicast = true
  server_local_sa_ipsec = [
    {
      sequence                   = 100
      replay_counter             = true
      replay_counter_window_size = "512"
    }
  ]
}
