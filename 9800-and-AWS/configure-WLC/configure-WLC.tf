# initialize terraform using 'terraform init'
# apply this terraform file using 'terraform apply -auto-approve -var-file="9800.tfvars"'

terraform {
  required_providers {
    iosxe = {
      source  = "CiscoDevNet/iosxe" 
    }
  }
}

provider "iosxe" {
  # variables initialized in variables.tf and values stored in 9800.tfvars
  host            = var.host_url
  insecure        = var.insecure
  device_username = var.device_username
  device_password = var.device_password
}


# Configure the a new WLAN
resource "iosxe_rest" "wlan_example_set" {
  method = "PATCH"
  path = "/data/Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data"
  payload = jsonencode(
    {
        "Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data": {
        "wlan-cfg-entries": {
            "wlan-cfg-entry": [
            {
                "profile-name": "WLAN_Configured_by_Terraform",
                "wlan-id": 200,
                "auth-key-mgmt-psk": true,
                "auth-key-mgmt-dot1x": false,
                "psk": "Cisco123!",
                "apf-vap-id-data": {
                "ssid": "WLAN_Configured_by_Terraform",
                "wlan-status": true
                }
            }
            ]
        }
        }
    }
  )
}

# Fetch the available WLAN configuration
# resource "iosxe_rest" "wlan_example_get" {
#   method = "GET"
#   path = "/data/Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data"
# }