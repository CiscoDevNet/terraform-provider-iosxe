resource "iosxe_radius" "example" {
  attributes = [
    {
      number = "31"
      attri31 = [
        {
          calling_station_id = "mac"
          attri31_format     = "ietf"
          attri31_lu_case    = "lower-case"
        }
      ]
    }
  ]
}
