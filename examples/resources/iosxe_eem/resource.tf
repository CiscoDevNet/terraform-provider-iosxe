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
  detector_rpc_max_sessions             = 8
  detector_routing_bootup_delay         = 2
  applets = [
    {
      name              = "test_applet_10"
      authorization     = "bypass"
      class             = "A"
      description       = "test describing applet"
      event_cli_pattern = "shutdown"
      event_cli_sync    = "no"
      event_cli_skip    = "no"
      actions = [
        {
          name        = "10"
          cli_command = "enable"
        }
      ]
      event_timer_watchdog_time      = 1800
      event_timer_watchdog_name      = "test_time"
      event_timer_watchdog_maxrun    = 10
      event_timer_watchdog_ratelimit = 10
      event_timer_cron_entry         = "0 12 * * 1-5"
      event_timer_cron_name          = "test_time"
      event_timer_cron_maxrun        = 10
      event_timer_cron_ratelimit     = 10
    }
  ]
}
