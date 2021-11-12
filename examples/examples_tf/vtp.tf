
# Adding Password to the VTP configuration
resource "iosxe_rest" "vtp_example_post" {
  method = "POST"
  path = "/data/Cisco-IOS-XE-native:native/vtp/password"
  payload = jsonencode(
    {
      "password": "mypassword"
    }
  )
}

# Updating the available VTP configuration
resource "iosxe_rest" "vtp_example_put" {
  method = "PUT"
  path = "/data/Cisco-IOS-XE-native:native/vtp"
  payload = jsonencode(
    {
      "Cisco-IOS-XE-native:vtp": {
        "Cisco-IOS-XE-vtp:password": {
          "password": "PUT_password"
        },
        "Cisco-IOS-XE-vtp:domain": "domain_updated"
      }
    }
  )
}

# Updating the domain of VTP configuration
resource "iosxe_rest" "vtp_example_patch" {
  method = "PATCH"
  path = "/data/Cisco-IOS-XE-native:native/vtp"
  payload = jsonencode(
    {
				"Cisco-IOS-XE-native:vtp": {
					"Cisco-IOS-XE-vtp:password": {
						"password": "some_password"
					},
					"Cisco-IOS-XE-vtp:domain": "domain_updated"
				}
			}
  )
}

# Fetch the available VTP configuration
resource "iosxe_rest" "vtp_example_get" {
  method = "GET"
  path = "/data/Cisco-IOS-XE-native:native/vtp"
}

# Removing password form VTP configuration
resource "iosxe_rest" "vtp_example_delete" {
  method = "DELETE"
  path = "/data/Cisco-IOS-XE-native:native/vtp/password"
}
