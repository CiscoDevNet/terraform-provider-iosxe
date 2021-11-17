# Adding BGP configuration(s)
resource "iosxe_rest" "bgp_example_post" {
  method = "POST"
  path = "/data/Cisco-IOS-XE-native:native/router"
  payload = jsonencode(
    {
    "Cisco-IOS-XE-bgp:bgp": [
        {
            "id": 46000,
            "bgp": {
                "log-neighbor-changes": true
            },
            "neighbor": [
                {
                    "id": "10.108.1.2",
                    "remote-as": 65200
                }
            ]
        }
    ]
}
  )
}

# Updating the available BGP configuration
resource "iosxe_rest" "bgp_example_put" {
  method = "PUT"
  path = "/data/Cisco-IOS-XE-native:native/router/bgp=46000"
  payload = jsonencode(
    {
    "Cisco-IOS-XE-bgp:bgp": [
        {
            "id": 46000,
            "neighbor": [
                {
                    "id": "10.106.1.2",
                    "remote-as": 65000
                },
                {
                    "id": "10.107.1.2",
                    "remote-as": 65200
                }
            ]
        }
      ]
  }
  )
}

# Adding/Updating the available BGP configuration
resource "iosxe_rest" "bgp_example_patch" {
		method = "PATCH"
		path = "/data/Cisco-IOS-XE-native:native/router/bgp"
		payload = jsonencode(
        {
            "Cisco-IOS-XE-bgp:bgp": [
            {
                "id": 46000,
                "neighbor": [
                    {
                        "id": "10.109.1.4",
                        "remote-as": 66000
                    }
                ]
            }
            ]
        }
		)
	}

# Fetch the available BGP configuration
resource "iosxe_rest" "bgp_example_get" {
  method = "GET"
  path = "/data/Cisco-IOS-XE-native:native/router/bgp"
}

# Fetch the available BGP configuration by id
resource "iosxe_rest" "bgp_example_get_id" {
  method = "GET"
  path = "/data/Cisco-IOS-XE-native:native/router/bgp=46000"
}

# Remove the available BGP configuration by id
resource "iosxe_rest" "bgp_example_delete" {
  method = "DELETE"
  path = "/data/Cisco-IOS-XE-native:native/router/bgp=46000"
}
