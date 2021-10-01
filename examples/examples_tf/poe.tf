
# Adding POE configuration on the interface Gi1/0/12
resource "iosxe_rest" "poe_example_post" {
  method = "POST"
  path = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1%2f0%2f12/Cisco-IOS-XE-power:power/inline"
  payload = jsonencode(
    {
      "static-choice": {
        "max": 5000
      }
    }
  )
}

# Updating the available POE configuration on the interface Gi1/0/12
resource "iosxe_rest" "poe_example_put" {
  method = "PUT"
  path = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1%2f0%2f12/Cisco-IOS-XE-power:power/inline"
  payload = jsonencode(
    {
      "Cisco-IOS-XE-power:inline": {
        "auto-choice": {
          "max": 4500
        }
      }
    }
  )
}

# Adding/Updating the POE configuration on the interface Gi1/0/12
resource "iosxe_rest" "poe_example_patch" {
  method = "PATCH"
  path = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1%2f0%2f12/Cisco-IOS-XE-power:power/inline"
  payload = jsonencode(
    {
      "Cisco-IOS-XE-power:inline": {
        "auto-choice": {
          "max": 5000
        }
      }
    }
  )
}

# Fetch the available POE configuration for interface Gi1/0/12
resource "iosxe_rest" "poe_example_get" {
  method = "GET"
  path = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1%2f0%2f12/Cisco-IOS-XE-power:power"
}

# Remove POE configurations from the interface Gi1/0/12
resource "iosxe_rest" "poe_example_delete" {
  method = "DELETE"
  path = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1%2f0%2f12/Cisco-IOS-XE-power:power/inline"
}
