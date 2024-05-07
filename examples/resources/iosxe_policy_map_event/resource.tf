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
          number                                            = 10
          activate_service_template_config_service_template = "DEFAULT_LINK_POLICY"
          activate_service_template_config_aaa_list         = "methodlist1"
          activate_service_template_config_precedence       = 1
          activate_interface_template                       = "templ1"
          activate_policy_type_control_subscriber           = "subscriber1"
          authenticate_using_method                         = "dot1x"
          authenticate_using_retries                        = 2
          authenticate_using_retry_time                     = 0
          authenticate_using_priority                       = 10
          authenticate_using_aaa_authc_list                 = "listname1"
          authenticate_using_aaa_authz_list                 = "listname2"
          authenticate_using_both                           = true
          replace                                           = true
          restrict                                          = true
          clear_session                                     = true
          clear_authenticated_data_hosts_on_port            = true
          protect                                           = true
          err_disable                                       = true
          resume_reauthentication                           = true
          authentication_restart                            = 2
          set_domain                                        = "data"
          unauthorize                                       = true
          notify                                            = true
          set_timer_name                                    = "timer1"
          set_timer_value                                   = 3600
        }
      ]
    }
  ]
}
