resource "iosxe_interface_ospf_process" "example" {
  type       = "Loopback"
  name       = "1"
  process_id = 1
  area = [
    {
      area_id = "1"
    }
  ]
}
