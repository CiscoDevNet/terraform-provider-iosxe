resource "iosxe_device_sensor" "example" {
  filter_spec_dhcp_includes = [
    {
      name = "dhcp1"
    }
  ]
  filter_spec_lldp_includes = [
    {
      name = "lldp1"
    }
  ]
  filter_lists_lldp = [
    {
      name                        = "lldp1"
      tlv_name_port_id            = true
      tlv_name_port_description   = true
      tlv_name_system_name        = true
      tlv_name_system_description = true
    }
  ]
  filter_lists_dhcp = [
    {
      name                               = "dhcp1"
      option_name_host_name              = true
      option_name_default_ip_ttl         = true
      option_name_requested_address      = true
      option_name_parameter_request_list = true
      option_name_class_identifier       = true
      option_name_client_identifier      = true
      option_name_client_fqdn            = true
    }
  ]
  notify_all_changes = true
}
