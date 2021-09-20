provider "iosxe" {
  // Required but Optional if env variable are set
  host = "https://10.1.1.5"
  device_username = "Cisco123"
  device_password = "somePassword"

  // Optional Parameters
  insecure = true
  request_timeout = 30
  ca_file = ""
  proxy_url = ""
  proxy_creds = ""
}