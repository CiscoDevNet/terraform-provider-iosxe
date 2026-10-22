resource "iosxe_wireless_ap_filter_priority" "example" {
  priorities = [
    {
      priority    = 1000
      filter_name = "S3-AF-TEST"
    }
  ]
}
