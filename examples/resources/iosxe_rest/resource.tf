provider "iosxe" {
  host = "https://localhost:8443"
  insecure = true
  device_password = "Cisco123"
  device_username = "admin"
}

resource "iosxe_rest" "rest1" {
  method = "GET"
  path = "/data/Cisco-IOS-XE-native:native"
}