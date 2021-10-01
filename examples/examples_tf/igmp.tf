
# Adding IGMP configuration
resource "iosxe_rest" "igmp_example_post" {
  method = "POST"
  path   = "/data/Cisco-IOS-XE-native:native/ip/igmp"
  payload = jsonencode(
    {
      "Cisco-IOS-XE-igmp:profile" : {
        "id" : "3",
        "action" : "permit",
        "range" : [
          {
            "low_ip" : "229.9.9.0"
          }
        ]
      }
    }
  )
}

# Updating the available IGMP configuration
resource "iosxe_rest" "igmp_example_put" {
  method = "PUT"
  path   = "/data/Cisco-IOS-XE-native:native/ip/igmp/profile=3"
  payload = jsonencode(
    {
      "Cisco-IOS-XE-igmp:profile" : {
        "id" : "3",
        "action" : "permit",
        "range" : [
          {
            "low_ip" : "229.9.9.1"
          }
        ]
      }
    }
  )
}

# Adding/Updating the IGMP configuration
resource "iosxe_rest" "igmp_example_patch" {
  method = "PATCH"
  path   = "/data/Cisco-IOS-XE-native:native/ip/igmp"
  payload = jsonencode(
    {
      "Cisco-IOS-XE-igmp:igmp" : {
        "profile" : [
          {
            "id" : "3",
            "action" : "permit",
            "range" : [
              {
                "low_ip" : "229.9.9.1"
              },
              {
                "low_ip" : "229.9.9.2"
              }
            ]
          }
        ]
      }
    }
  )
}

# Fetch the available IGMP configuration
resource "iosxe_rest" "igmp_example_get" {
  method = "GET"
  path   = "data/Cisco-IOS-XE-native:native/ip/igmp"
}

# Fetch the available IGMP configuration by id
resource "iosxe_rest" "igmp_example_get" {
  method = "GET"
  path   = "data/Cisco-IOS-XE-native:native/ip/igmp/profile=3"
}


# Remove IGMP profile configuration
resource "iosxe_rest" "igmp_example_delete" {
  method = "DELETE"
  path   = "data/Cisco-IOS-XE-native:native/ip/igmp/profile=3"
}
