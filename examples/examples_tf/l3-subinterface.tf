# Adding L3-SUBINTERFACE configuration
resource "iosxe_rest" "l3_subinterface_example_post" {
  method = "POST"
  path   = "/data/Cisco-IOS-XE-native:native/interface/Port-channel-subinterface"
  payload = jsonencode(
    {
        "Port-channel": {
            "name": "2.10",
            "encapsulation": {
                "dot1Q": {
                    "vlan-id": 10
                }
            },
            "ip": {
                "address": {
                    "primary": {
                        "address": "10.10.10.10",
                        "mask": "255.255.255.0"
                    }
                }
            }
        }
    }
  )
}

# Updating the available L3-SUBINTERFACE configuration 
resource "iosxe_rest" "l3_subinterface_example_put" {
  method = "PUT"
  path   = "/data/Cisco-IOS-XE-native:native/interface/Port-channel-subinterface/Port-channel=2.10"
  payload = jsonencode(
    {
        "Port-channel": {
            "name": "2.10",
            "encapsulation": {
                "dot1Q": {
                    "vlan-id": 10
                }
            },
            "ip": {
                "address": {
                    "primary": {
                        "address": "10.10.10.11",
                        "mask": "255.255.255.0"
                    }
                }
            }
        }
    }
  )
}

# Adding/Updating the available L3-SUBINTERFACE configuration 
resource "iosxe_rest" "l3_subinterface_example_patch" {
  method = "PATCH"
  path   = "/data/Cisco-IOS-XE-native:native/interface/Port-channel-subinterface/Port-channel"
  payload = jsonencode(
    {
        "Port-channel": {
            "name": "2.10",
            "encapsulation": {
                "dot1Q": {
                    "vlan-id": 10
                }
            },
            "ip": {
                "address": {
                    "primary": {
                        "address": "10.10.10.10",
                        "mask": "255.255.255.0"
                    }
                }
            }
        }
    }
    )
}

# Fetch the available L3-SUBINTERFACE configuration 
resource "iosxe_rest" "l3_subinterface_example_get" {
  method = "GET"
  path   = "/data/Cisco-IOS-XE-native:native/interface/Port-channel-subinterface/Port-channel"
}

# Fetch the available L3-SUBINTERFACE configuration by Port Channel id
resource "iosxe_rest" "l3_subinterface_example_get_id" {
  method = "GET"
  path   = "/data/Cisco-IOS-XE-native:native/interface/Port-channel-subinterface/Port-channel=2.10"
}

# Remove the available L3-SUBINTERFACE configuration 
resource "iosxe_rest" "l3_subinterface_example_delete" {
  method = "DELETE"
  path = "/data/Cisco-IOS-XE-native:native/interface/Port-channel-subinterface/Port-channel=2.10"
}
