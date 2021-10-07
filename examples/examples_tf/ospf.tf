# Adding OSPF configuration
resource "iosxe_rest" "ospf_example_post" {
  method = "POST"
  path = "/data/Cisco-IOS-XE-native:native/router/router-ospf/ospf"
  payload = jsonencode(
    {
    "process-id": [
        {
            "id": 22,
            "network": [
                {
                    "ip": "10.1.1.0",
                    "wildcard": "255.240.0.0",
                    "area": 20
                }
            ],
            "nsf": {
                "nsf-cisco": {
                    "enforce": {
                        "global": [
                            null
                        ]
                    }
                }
            }
        }
    ]
}
  )
}

# Updating the available OSPF configuration
resource "iosxe_rest" "ospf_example_put" {
  method = "PUT"
  path = "/data/Cisco-IOS-XE-native:native/router/router-ospf/ospf/process-id=15"
  payload = jsonencode(
    {
    "Cisco-IOS-XE-ospf:process-id": {
        "id": 15,
        "network": [
            {
                "ip": "10.1.1.0",
                "wildcard": "255.240.0.0",
                "area": 20
            }
        ],
        "nsf": {
            "nsf-cisco": {
                "enforce": {
                    "global": [
                        null
                    ]
                }
            }
        }
    }
}
  )
}

# Adding/Updating the available OSPF configuration
resource "iosxe_rest" "ospf_example_patch" {
  method = "PATCH"
  path = "/data/Cisco-IOS-XE-native:native/router/router-ospf/ospf"
  payload = jsonencode(
    {
    "ospf": {
        "process-id": [
            {
                "id": 15,
                "network": [
                    {
                        "ip": "10.1.1.1",
                        "wildcard": "255.240.0.0",
                        "area": 20
                    }
                ],
                "nsf": {
                    "nsf-cisco": {
                        "enforce": {
                            "global": [
                                null
                            ]
                        }
                    }
                }
            }
        ]
    }
}
  )
}

# Fetch the available OSPF configuration
resource "iosxe_rest" "ospf_example_get" {
  method = "GET"
  path = "/data/Cisco-IOS-XE-native:native/router/router-ospf/ospf"
}

# Fetch the available OSPF configuration by id
resource "iosxe_rest" "ospf_example_get_id" {
  method = "GET"
  path = "/data/Cisco-IOS-XE-native:native/router/router-ospf/ospf/process-id=15"
}

# Remove the available OSPF configuration by id
resource "iosxe_rest" "ospf_example_delete" {
  method = "DELETE"
  path = "/data/Cisco-IOS-XE-native:native/router/router-ospf/ospf/process-id=15"
}
