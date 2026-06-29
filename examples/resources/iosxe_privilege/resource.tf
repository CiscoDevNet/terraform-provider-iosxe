resource "iosxe_privilege" "example" {
  name = "exec"
  levels = [
    {
      level = 7
      commands = [
        {
          command = "configure"
        }
      ]
    }
  ]
}
