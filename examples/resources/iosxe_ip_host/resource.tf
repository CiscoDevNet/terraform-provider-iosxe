resource "iosxe_ip_host" "example" {
  host_lists = [
    {
      name    = "test.router.com"
      ip_list = ["3.3.3.3"]
    }
  ]
  vrf = [
    {
      vrf = "VRF1"
      host_name = [
        {
          host_name = "test.router.com"
          ip_list   = ["2.2.2.2"]
        }
      ]
    }
  ]
}
