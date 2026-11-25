resource "iosxe_privilege" "example" {
  mode = "exec"
  levels = [
    {
      level = 7
      commands = [
        {
          command = "ping"
        }
      ]
    }
  ]
}
