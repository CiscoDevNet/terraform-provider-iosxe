# Adding SPAN-RSPAN configuration
resource "iosxe_rest" "span-rspan_example_post" {
  method = "POST"
  path   = "/data/Cisco-IOS-XE-native:native/monitor"
  payload = jsonencode(
    {
        "session": [
            {
                "id": 1,
                "destination": {
                    "interface": [
                        {
                            "name": "AppGigabitEthernet1/0/1"
                        }
                    ]
                },
                "source": {
                    "interface": [
                        {
                            "name": "GigabitEthernet1/0/1"
                        }
                    ]
                }
            }
        ]
    }
  )
}

# Updating the available SPAN-RSPAN configuration
resource "iosxe_rest" "span-rspan_example_put" {
  method = "PUT"
  path   = "/data/Cisco-IOS-XE-native:native/monitor"
  payload = jsonencode(
    {
        "Cisco-IOS-XE-native:monitor": {
            "session": [
                {
                    "id": 1,
                    "destination": {
                        "interface": [
                            {
                                "name": "AppGigabitEthernet1/0/1"
                            }
                        ]
                    },
                    "source": {
                        "interface": [
                            {
                                "name": "GigabitEthernet1/0/2"
                            }
                        ]
                    }
                }
            ]
        }
    }
  )
}

# Adding/Updating the available SPAN-RSPAN configuration
resource "iosxe_rest" "span-rspan_example_patch" {
  method = "PATCH"
  path   = "/data/Cisco-IOS-XE-native:native/monitor"
  payload = jsonencode(
    {
        "Cisco-IOS-XE-native:monitor": {
            "session": [
                {
                    "id": 1,
                    "destination": {
                        "interface": [
                            {
                                "name": "AppGigabitEthernet1/0/1"
                            }
                        ]
                    },
                    "source": {
                        "interface": [
                            {
                                "name": "GigabitEthernet1/0/1"
                            }
                        ]
                    }
                },
                {
                    "id": 2,
                    "destination": {
                        "interface": [
                            {
                                "name": "AppGigabitEthernet1/0/1"
                            }
                        ]
                    },
                    "source": {
                        "interface": [
                            {
                                "name": "GigabitEthernet1/0/2"
                            }
                        ]
                    }
                }
            ]
        }
    }
    )
}

# Fetch the available SPAN-RSPAN configuration 
resource "iosxe_rest" "span-rspan_example_get" {
  method = "GET"
  path   = "/data/Cisco-IOS-XE-native:native/monitor"
}

# Fetch the available SPAN-RSPAN configuration  by session id
resource "iosxe_rest" "span-rspan_example_get" {
  method = "GET"
  path   = "/data/Cisco-IOS-XE-native:native/monitor/session=1"
}

# Remove the available SPAN-RSPAN configuration 
resource "iosxe_rest" "span-rspan_example_delete" {
  method = "DELETE"
  path = "/data/Cisco-IOS-XE-native:native/monitor/session=1"
}
