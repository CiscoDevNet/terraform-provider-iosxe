
# Adding PIM configuration on the interface Gi1/0/15
#
# # POST method is not available for PIM on the interface Gi1/0/15
# resource "iosxe_rest" "pim_example_post" {
#   method = "POST"
#   path = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1%2f0%2f15/ip/pim"
#   payload = jsonencode()
# }

# Updating the available PIM configuration on the interface Gi1/0/15
resource "iosxe_rest" "pim_example_put" {
  method = "PUT"
  path = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1%2f0%2f15/ip/pim"
  payload = jsonencode(
    {
    "Cisco-IOS-XE-native:pim": {
        "Cisco-IOS-XE-multicast:pim-mode-choice-cfg": {
            "passive": [
                null
            ]
        }
    }
}
  )
}

# Adding/Updating the PIM configuration on the interface Gi1/0/15
resource "iosxe_rest" "pim_example_patch" {
  method = "PATCH"
  path = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1%2f0%2f15/ip/pim"
  payload = jsonencode(
    {
    "Cisco-IOS-XE-native:pim": {
        "Cisco-IOS-XE-multicast:pim-mode-choice-cfg": {
            "passive": [
                null
            ]
        }
    }
}
  )
}

# Fetch the available PIM configuration for interface Gi1/0/15
resource "iosxe_rest" "pim_example_get" {
  method = "GET"
  path = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1%2f0%2f15/ip/pim"
}

# Remove PIM configurations from the interface Gi1/0/15
resource "iosxe_rest" "pim_example_delete" {
  method = "DELETE"
  path = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1%2f0%2f15/ip/pim"
}
