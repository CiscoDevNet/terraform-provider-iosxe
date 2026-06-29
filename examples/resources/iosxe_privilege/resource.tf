resource "iosxe_privilege" "example" {
  name  = "exec"
  level = 7
  commands = [
    {
      command = "configure"
    }
  ]
}
