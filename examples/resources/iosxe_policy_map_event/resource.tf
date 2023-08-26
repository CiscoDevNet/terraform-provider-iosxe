resource "iosxe_policy_map_event" "example" {
  name       = "dot1x_policy"
  event_type = "authentication-success"
  match_type = "match-all"
  class_numbers = [
    {
      number         = 10
      class          = "MY_CLASS"
      execution_type = "do-until-failure"
      action_numbers = [
        {
          number                        = 10
          authenticate_using_method     = "dot1x"
          authenticate_using_retries    = 2
          authenticate_using_retry_time = 0
          authenticate_using_priority   = 10
        }
      ]
    }
  ]
}
