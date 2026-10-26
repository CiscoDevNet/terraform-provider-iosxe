resource "iosxe_voice_class_sip_event_list" "example" {
  tag    = 100
  events = ["refer"]
}
