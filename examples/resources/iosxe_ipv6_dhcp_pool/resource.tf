resource "iosxe_ipv6_dhcp_pool" "example" {
  name                        = "DHCPv6-PD"
  prefix_delegation_pool_name = "DHCPv6-PD"
  dns_servers                 = ["2001:4860:4860::8888"]
  domain_names                = ["example.com"]
  vendor_specifics = [
    {
      enterprise_id = 9
      suboptions = [
        {
          number  = 1
          address = "2001:db8::1"
        }
      ]
    }
  ]
}
