resource "iosxe_msdp_vrf" "example" {
  vrf           = "VRF1"
  originator_id = "Loopback100"
  peers = [
    {
      addr                    = "10.1.1.1"
      remote_as               = 65000
      connect_source_loopback = 100
    }
  ]
  passwords = [
    {
      addr       = "10.1.1.1"
      encryption = 0
      password   = "Cisco123"
    }
  ]
}
