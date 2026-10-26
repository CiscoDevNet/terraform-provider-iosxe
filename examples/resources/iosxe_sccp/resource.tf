resource "iosxe_sccp" "example" {
  local_interface = "GigabitEthernet1"
  ccm_servers = [
    {
      host       = "192.0.2.50"
      identifier = 1
    }
  ]
  ccm_groups = [
    {
      group_id = 1
      associate_ccm = [
        {
          ccm_id   = 1
          priority = 1
        }
      ]
    }
  ]
}
