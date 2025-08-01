resource "iosxe_lldp" "example" {
  run      = true
  holdtime = 60
  timer    = 60
}
