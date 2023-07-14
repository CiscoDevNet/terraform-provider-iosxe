resource "iosxe_access_list_standard" "example" {
  name = "SACL1"
  entries = [
    {
      sequence         = 10
      remark           = "Description"
      deny_prefix      = "10.0.0.0"
      deny_prefix_mask = "0.0.0.255"
    }
  ]
}
