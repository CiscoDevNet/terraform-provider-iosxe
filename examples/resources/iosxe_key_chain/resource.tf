resource "iosxe_key_chain" "example" {
  name = "OSPF-AUTH"
  keys = [
    {
      id                          = "1"
      key_string_encryption       = "0"
      key_string_key              = "mySecretKey123"
      accept_lifetime_start_time  = "00:00:00"
      accept_lifetime_start_month = "Jan"
      accept_lifetime_start_day   = 1
      accept_lifetime_start_year  = 2025
      accept_lifetime_infinite    = true
      send_lifetime_start_time    = "00:00:00"
      send_lifetime_start_month   = "Jan"
      send_lifetime_start_day     = 1
      send_lifetime_start_year    = 2025
      send_lifetime_infinite      = true
    }
  ]
}
