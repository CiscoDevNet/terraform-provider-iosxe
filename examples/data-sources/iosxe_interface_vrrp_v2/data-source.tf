data "iosxe_interface_vrrp_v2" "example" {
  type     = "GigabitEthernet"
  name     = "1"
  group_id = 1
}
