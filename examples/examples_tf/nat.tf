# Adding NAT configuration
#
# POST method is not available for NAT
# resource "iosxe_rest" "nat_example_post" {
#   method = "POST"
#   path   = "/data/Cisco-IOS-XE-native:native/ip/nat"
#   payload = jsonencode(
#   )
# }

# Updating the available NAT configuration 
resource "iosxe_rest" "nat_example_put" {
  method = "PUT"
  path   = "/data/Cisco-IOS-XE-native:native/ip/nat"
  payload = jsonencode(
    {
        "Cisco-IOS-XE-nat:nat": {
            "pool": [
                {
                    "id": "net-208",
                    "prefix-length": 27
                }
            ],
            "inside": {
                "source": {
                    "list": [
                        {
                            "id": 1,
                            "pool": "net-208"
                        }
                    ]
                }
            }
        }
    }
  )
}

# Adding/Updating the available NAT configuration 
resource "iosxe_rest" "nat_example_patch" {
  method = "PATCH"
  path   = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet"
  payload = jsonencode(
   {
        "Cisco-IOS-XE-nat:nat": {
            "pool": [
                {
                    "id": "net-208",
                    "prefix-length": 27
                }
            ],
            "inside": {
                "source": {
                    "list": [
                        {
                            "id": 1,
                            "pool": "net-208"
                        }
                    ]
                }
            }
        }
    }
    )
}

# Fetch the available NAT configuration 
resource "iosxe_rest" "nat_example_get" {
  method = "GET"
  path   = "/data/Cisco-IOS-XE-native:native/ip/nat"
}

# Remove the available NAT configuration 
resource "iosxe_rest" "nat_example_delete" {
  method = "DELETE"
  path = "/data/Cisco-IOS-XE-native:native/ip/nat"
}
