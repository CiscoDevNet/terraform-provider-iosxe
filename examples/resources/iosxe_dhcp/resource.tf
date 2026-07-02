resource "iosxe_dhcp" "example" {
  relay_information_trust_all                 = false
  relay_information_option_default            = false
  relay_information_option_vpn                = true
  relay_bootp_ignore                          = true
  snooping                                    = true
  snooping_information_option                 = true
  snooping_information_option_allow_untrusted = true
  pools = [
    {
      name            = "POOL1"
      domain_name     = "example.com"
      bootfile        = "boot.cfg"
      network_number  = "10.1.1.0"
      network_mask    = "255.255.255.0"
      default_routers = ["10.1.1.1"]
      dns_servers     = ["10.1.1.1"]
      lease_days      = 1
      lease_hours     = 12
      lease_minutes   = 30
      options = [
        {
          option_code = 150
          ascii       = "10.1.1.1"
        }
      ]
    }
  ]
}
