resource "iosxe_voice_class_tls_profile" "example" {
  tag                    = 100
  description            = "CUBE TLS Profile"
  trustpoint             = "CUBE-CERT"
  cn_san_validate_server = true
  cn_san_names = [
    {
      tag  = 1
      name = "cube.example.com"
    }
  ]
  strict_cipher = true
  sni_send      = true
}
