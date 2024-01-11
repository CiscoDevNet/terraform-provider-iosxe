resource "iosxe_policy_map" "example" {
  name        = "POLICY1"
  description = "My first policy-map"
  classes = [
    {
      name = "CLASS1"
      actions = [
        {
          type              = "bandwidth"
          bandwidth_percent = 10
        }
      ]
    }
  ]
}
