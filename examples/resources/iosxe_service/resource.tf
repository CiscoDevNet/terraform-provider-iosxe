resource "iosxe_service" "example" {
  pad                                     = true
  password_encryption                     = true
  password_recovery                       = true
  timestamps                              = true
  timestamps_debug                        = true
  timestamps_debug_datetime               = true
  timestamps_debug_datetime_msec          = true
  timestamps_debug_datetime_localtime     = true
  timestamps_debug_datetime_show_timezone = true
  timestamps_debug_datetime_year          = true
  timestamps_debug_uptime                 = true
  timestamps_log                          = true
  timestamps_log_datetime                 = true
  timestamps_log_datetime_msec            = true
  timestamps_log_datetime_localtime       = true
  timestamps_log_datetime_show_timezone   = true
  timestamps_log_datetime_year            = true
  timestamps_log_uptime                   = true
  dhcp                                    = true
  tcp_keepalives_in                       = true
  tcp_keepalives_out                      = true
}
