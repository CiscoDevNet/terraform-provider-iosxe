resource "iosxe_bfd" "example" {
  ipv4_both_vrfs = [
    {
      dst_vrf       = "dest_vrf1"
      dest_ip       = "1.2.3.4/4"
      src_vrf       = "src_vrf1"
      src_ip        = "11.22.33.44/12"
      template_name = "template1"
    }
  ]
  ipv4_without_vrfs = [
    {
      dest_ip       = "1.2.3.4/4"
      src_ip        = "11.22.33.44/12"
      template_name = "template1"
    }
  ]
  ipv4_with_src_vrfs = [
    {
      dest_ip       = "1.2.3.4/4"
      src_vrf       = "src_vrf1"
      src_ip        = "11.22.33.44/12"
      template_name = "template1"
    }
  ]
  ipv4_with_dst_vrfs = [
    {
      dst_vrf       = "dest_vrf1"
      dest_ip       = "1.2.3.4/4"
      src_ip        = "11.22.33.44/12"
      template_name = "template1"
    }
  ]
  ipv6_with_both_vrfs = [
    {
      dst_vrf       = "dst_vrf1"
      dest_ipv6     = "2001:DB8:0:1::/64"
      src_vrf       = "src_vrf1"
      src_ipv6      = "2001:DB8:0:2::/64"
      template_name = "template1"
    }
  ]
  ipv6_without_vrfs = [
    {
      dest_ipv6     = "2001:DB8:0:1::/64"
      src_ipv6      = "2001:DB8:0:2::/64"
      template_name = "template1"
    }
  ]
  ipv6_with_src_vrfs = [
    {
      dest_ipv6     = "2001:DB8:0:1::/64"
      src_vrf       = "src_vrf1"
      src_ipv6      = "2001:DB8:0:2::/64"
      template_name = "template1"
    }
  ]
  ipv6_with_dst_vrfs = [
    {
      dst_vrf       = "dst_vrf1"
      dest_ipv6     = "2001:DB8:0:1::/64"
      src_ipv6      = "2001:DB8:0:2::/64"
      template_name = "template1"
    }
  ]
  slow_timers = 1000
}
