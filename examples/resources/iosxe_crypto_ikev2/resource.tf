resource "iosxe_crypto_ikev2" "example" {
  nat_keepalive                = 20
  dpd_container_dpd            = 10
  dpd_container_retry_interval = 5
  dpd_container_dpd_query      = "periodic"
}
