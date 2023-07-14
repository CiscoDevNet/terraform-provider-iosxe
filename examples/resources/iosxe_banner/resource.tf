resource "iosxe_banner" "example" {
  exec_banner           = "My Exec Banner"
  login_banner          = "My Login Banner"
  prompt_timeout_banner = "My Prompt-Timeout Banner"
  motd_banner           = "My MOTD Banner"
}
