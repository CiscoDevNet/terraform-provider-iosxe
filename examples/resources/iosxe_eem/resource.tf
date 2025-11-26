resource "iosxe_eem" "example" {
  environment_variables = [
    {
      name  = "IOSXE_TEST_VAR"
      value = "test_pass"
    }
  ]
  session_cli_username                  = "test_user"
  session_cli_username_privilege        = 15
  history_size_events                   = 25
  history_size_traps                    = 25
  directory_user_policy                 = "test/test_path"
  scheduler_applet_thread_class_default = true
  scheduler_applet_thread_class_number  = 1
  detector_routing_bootup_delay         = 2
  applets = [
    {
      name          = "test_applet_10"
      authorization = "bypass"
      class         = "A"
      description   = "test describing applet"
      actions = [
        {
          name        = "10"
          cli_command = "enable"
        }
      ]
      event_syslog_pattern   = "%EXAMPLE-5-TEST: Example syslog for testing purposes"
      event_syslog_occurs    = 1
      event_syslog_maxrun    = 30
      event_syslog_ratelimit = 10
      event_syslog_period    = 60
    }
  ]
}
