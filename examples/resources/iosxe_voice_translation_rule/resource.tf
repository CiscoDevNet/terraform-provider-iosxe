resource "iosxe_voice_translation_rule" "example" {
  tag = 200
  rules = [
    {
      rule_id                      = 1
      matching_replacement_pattern = "/408/ /555/"
    }
  ]
}
