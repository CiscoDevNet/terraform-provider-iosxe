resource "iosxe_cli" "example" {
  cli = <<-EOT
  interface Loopback123
  description configured-via-cli
  EOT
}
