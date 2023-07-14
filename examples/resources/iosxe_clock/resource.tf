resource "iosxe_clock" "example" {
  calendar_valid                      = true
  summer_time_zone                    = "CET"
  summer_time_recurring               = true
  summer_time_recurring_start_week    = "1"
  summer_time_recurring_start_weekday = "Mon"
  summer_time_recurring_start_month   = "Jan"
  summer_time_recurring_start_time    = "00:00"
  summer_time_recurring_end_week      = "1"
  summer_time_recurring_end_weekday   = "Mon"
  summer_time_recurring_end_month     = "Dec"
  summer_time_recurring_end_time      = "00:00"
  summer_time_recurring_offset        = 60
}
