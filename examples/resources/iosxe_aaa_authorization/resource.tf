resource "iosxe_aaa_authorization" "example" {
  execs = [
    {
      name     = "EXEC1"
      a1_group = "GROUP1"
      a2_group = "GROUP2"
      a3_group = "GROUP3"
      a4_local = true
    }
  ]
  networks = [
    {
      id       = "NET1"
      a1_group = "RGROUP1"
      a2_group = "RGROUP2"
      a3_group = "RGROUP3"
      a4_local = true
    }
  ]
}
