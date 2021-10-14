
# Adding etherChannel configuration on the interface Gi1/0/19
# 
# POST method is not suitable for etherChannel, try using PATCH instead 
# resource "iosxe_rest" "etherChannel_example_post" {
#   method = "POST"
#   path = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1%2f0%2f19"
#   payload = jsonencode(
#     {       
#     }
#   )
# }

# Updating the available etherChannel configuration on the interface Gi1/0/19
resource "iosxe_rest" "etherChannel_example_put" {
  method = "PUT"
  path   = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1%2f0%2f19"
  payload = jsonencode(
    {
      "Cisco-IOS-XE-native:GigabitEthernet" : {
        "name" : "1/0/19",
        "Cisco-IOS-XE-ethernet:channel-group" : {
          "number" : 5,
          "mode" : "auto"
        },
        "Cisco-IOS-XE-ethernet:lacp" : {
          "port-priority" : 30000,
          "rate" : "normal"
        }
      }
    }
  )
}

# Adding/Updating the etherChannel configuration on the interface Gi1/0/19
resource "iosxe_rest" "etherChannel_example_patch" {
  method = "PATCH"
  path   = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1%2f0%2f19"
  payload = jsonencode(
    {
      "Cisco-IOS-XE-native:GigabitEthernet" : {
        "name" : "1/0/19",
        "Cisco-IOS-XE-ethernet:channel-group" : {
          "number" : 5,
          "mode" : "desirable"
        },
        "Cisco-IOS-XE-ethernet:lacp" : {
          "port-priority" : 34000,
          "rate" : "fast"
        }
      }
    }
  )
}

# Fetch the available etherChannel configuration for interface Gi1/0/19
resource "iosxe_rest" "etherChannel_example_get" {
  method = "GET"
  path   = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1%2f0%2f19"
}

# Remove `channel-group` from configurations for the interface Gi1/0/19
resource "iosxe_rest" "etherChannel_example_delete" {
  method = "DELETE"
  path   = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1%2f0%2f19/Cisco-IOS-XE-ethernet:channel-group"
}
