resource "iosxe_aaa_accounting" "example" {
  update_newinfo_periodic            = 2880
  identity_default_start_stop_group1 = "RADIUS-GROUP"
  identity_default_start_stop_group2 = "RADIUS-GROUP2"
  identity_default_start_stop_group3 = "RADIUS-GROUP3"
  identity_default_start_stop_group4 = "RADIUS-GROUP4"
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
