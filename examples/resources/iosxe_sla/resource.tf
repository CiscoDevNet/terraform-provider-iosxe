resource "iosxe_sla" "example" {
  entries = [
    {
      number                = 20
      icmp_echo_destination = "192.168.10.10"
      icmp_echo_source_ip   = "192.168.1.1"
    }
  ]
  schedules = [
    {
      entry_number   = 20
      life           = 4000
      start_time_now = true
    }
  ]
}
