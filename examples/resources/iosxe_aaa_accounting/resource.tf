resource "iosxe_aaa_accounting" "example" {
  update_newinfo_periodic   = 2880
  identity_start_stop_group = "RADIUS-GROUP"
  exec = [
    {
      name   = "default"
      group1 = "T-Group"
    }
  ]
  network = [
    {
      id     = "network1"
      group1 = "radius"
      group2 = "tacacs+"
    }
  ]
  system_guarantee_first = false
}
