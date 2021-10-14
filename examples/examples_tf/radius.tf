# Adding Password to the RADIUS configuration
resource "iosxe_rest" "radius_example_post" {
  method = "POST"
  path = "/data/Cisco-IOS-XE-native:native/radius"
  payload = jsonencode(
    {
        "Cisco-IOS-XE-aaa:server": [
            {
                "id": "rsim",
                "address": {
                    "ipv4": "124.2.2.12",
                    "auth-port": 1612,
                    "acct-port": 1813
                },
                "timeout": 60,
                "key": {
                    "key": "rad123"
                },
                "retransmit": 10
            }
        ]
    }
  )
}

# Updating the available RADIUS configuration
resource "iosxe_rest" "radius_example_put" {
  method = "PUT"
  path = "/data/Cisco-IOS-XE-native:native/radius/server=rsim"
  payload = jsonencode(
    {
    "Cisco-IOS-XE-aaa:server": {
        "id": "rsim",
        "address": {
            "ipv4": "124.2.2.12",
            "auth-port": 1612,
            "acct-port": 1813
        },
        "timeout": 60,
        "key": {
            "key": "rad123"
        },
        "retransmit": 10
    }
}
  )
}

# Updating the domain of RADIUS configuration
resource "iosxe_rest" "radius_example_patch" {
  method = "PATCH"
  path = "/data/Cisco-IOS-XE-native:native/radius"
  payload = jsonencode(
   {
    "Cisco-IOS-XE-native:radius": {
        "Cisco-IOS-XE-aaa:server": [
            {
                "id": "rsim",
                "address": {
                    "ipv4": "124.2.2.12",
                    "auth-port": 1612,
                    "acct-port": 1813
                },
                "timeout": 60,
                "key": {
                    "key": "rad123"
                },
                "retransmit": 10
            }
        ]
    }
}
  )
}

# Fetch the available RADIUS configuration
resource "iosxe_rest" "radius_example_get" {
  method = "GET"
  path = "/data/Cisco-IOS-XE-native:native/radius"
}

# Fetch the available RADIUS configuration by ID
resource "iosxe_rest" "radius_example_get" {
  method = "GET"
  path = "/data/Cisco-IOS-XE-native:native/radius/server=rsim"
}

# Removing password form RADIUS configuration
resource "iosxe_rest" "radius_example_delete" {
  method = "DELETE"
  path = "/data/Cisco-IOS-XE-native:native/radius/server=rsim"
}
