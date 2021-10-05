
# Adding DHCP configuration
resource "iosxe_rest" "dhcp_example_post" {
  method = "POST"
  path   = "/data/Cisco-IOS-XE-native:native/ipv6/dhcp"
  payload = jsonencode(
    {
      "Cisco-IOS-XE-dhcp:pool" : [
        {
          "name" : "7",
          "link-address" : [
            {
              "address" : "2001:1002::/64"
            }
          ]
        }
      ]
    }
  )
}

# Updating the available DHCP configuration
resource "iosxe_rest" "dhcp_example_put" {
  method = "PUT"
  path   = "/data/Cisco-IOS-XE-native:native/ipv6/dhcp"
  payload = jsonencode(
    {
      "Cisco-IOS-XE-native:dhcp" : {
        "Cisco-IOS-XE-dhcp:pool" : [
          {
            "name" : "7",
            "link-address" : [
              {
                "address" : "2001:1002::/64"
              }
            ]
          },
          {
            "name" : "8",
            "link-address" : [
              {
                "address" : "2001:1004::/64"
              }
            ]
          }
        ]
      }
    }
  )
}

# Adding/Updating the NTP configuration
resource "iosxe_rest" "dhcp_example_patch" {
  method = "PATCH"
  path   = "/data/Cisco-IOS-XE-native:native/ipv6/dhcp"
  payload = jsonencode(
    {
      "Cisco-IOS-XE-native:dhcp" : {
        "Cisco-IOS-XE-dhcp:pool" : [
          {
            "name" : "7",
            "link-address" : [
              {
                "address" : "2001:1002::/64"
              }
            ]
          },
          {
            "name" : "8",
            "link-address" : [
              {
                "address" : "2001:1003::/64"
              }
            ]
          }
        ]
      }
    }
  )
}

# Fetch the available NTP configuration
resource "iosxe_rest" "dhcp_example_get" {
  method = "GET"
  path   = "/data/Cisco-IOS-XE-native:native/ipv6/dhcp"
}

# Removes all the NTP configuration
resource "iosxe_rest" "dhcp_example_delete" {
  method = "DELETE"
  path   = "/data/Cisco-IOS-XE-native:native/ipv6/dhcp/pool=7"
}
