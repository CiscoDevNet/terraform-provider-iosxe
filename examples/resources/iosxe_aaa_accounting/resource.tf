resource "iosxe_aaa_accounting" "example" {
  update_newinfo_periodic           = 2880
  identity_default_start_stop_group = "RADIUS-GROUP"
  execs = [
    {
      name              = "default"
      start_stop_group1 = "T-Group"
    }
  ]
  networks = [
    {
      id                = "network1"
      start_stop_group1 = "radius"
      start_stop_group2 = "tacacs+"
    }
  ]
  system_guarantee_first = false
}
