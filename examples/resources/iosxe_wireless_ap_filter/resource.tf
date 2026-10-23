resource "iosxe_wireless_ap_filter" "example" {
  name          = "S3-AF-TEST"
  ap_name_regex = "^S3-AP-TEST-.*"
  tag_policy    = "S3-PT-TEST"
  tag_site      = "S3-ST-TEST"
  tag_rf        = "S3-RT-TEST"
  filter_type   = "ap-filter-type-tag"
}
