resource "iosxe_username" "example" {
  name                = "user1"
  privilege           = 15
  description         = "User1 description"
  password_encryption = "0"
  password            = "MyPassword"
}
