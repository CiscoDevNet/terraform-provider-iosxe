resource "iosxe_wireless_ap_tags" "example" {
  items = {
    "00:53:00:03:00:01" = {
      policy_tag = "S3-PT-TEST"
      site_tag   = "S3-ST-TEST"
      rf_tag     = "S3-RT-TEST"
    }
  }
}
