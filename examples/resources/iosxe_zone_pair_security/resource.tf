resource "iosxe_zone_pair_security" "example" {
  name                        = "ZP_IN_OUT"
  source                      = "INSIDE"
  destination                 = "OUTSIDE"
  description                 = "Inside to outside traffic policy"
  service_policy_type_inspect = "PM_IN_TO_OUT"
}
