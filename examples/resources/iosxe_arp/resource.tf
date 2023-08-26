resource "iosxe_arp" "example" {
  incomplete_entries = 10
  proxy_disable      = true
}
