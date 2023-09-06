// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxeSNMPServer(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "chassis_id", "R1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "contact", "Contact1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "ifindex_persist", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "location", "Location1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "packetsize", "2000"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "queue_length", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_logging_getop", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_logging_setop", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_snmp_authentication", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_snmp_coldstart", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_snmp_linkdown", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_snmp_linkup", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_snmp_warmstart", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "system_shutdown", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_snmp_flowmon", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_entity_perf_throughput_notif", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_call_home", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_call_home_server_fail", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_tty", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_ospfv3_config_state_change", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_ospfv3_errors_change", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_ospf_config_retransmit", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_ospf_config_lsa", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_ospf_nssa_trans_change", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_ospf_shamlink_interface", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_ospf_shamlink_neighbor", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_ospf_errors_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_ospf_retransmit_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_ospf_lsa_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_eigrp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_auth_framework_sec_violation", "true"))
	if os.Getenv("C9000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_rep", "true"))
	}
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_vtp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_vlancreate", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_vlandelete", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_port_security", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_license", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_smart_license", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_cpu_threshold", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_memory_bufferpeak", "true"))
	if os.Getenv("C9000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_stackwise", "true"))
	}
	if os.Getenv("C9000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_link_fail_rpt", "true"))
	}
	if os.Getenv("C9000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_status_change", "true"))
	}
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_fru_ctrl", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_flash_insertion", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_flash_removal", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_flash_lowspace", "true"))
	if os.Getenv("C9000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_energywise", "true"))
	}
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_entity", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_pw_vc", "true"))
	if os.Getenv("C9000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_envmon", "true"))
	}
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_ipsla", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_bfd", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_ike_policy_add", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_ike_policy_delete", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_ike_tunnel_start", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_ike_tunnel_stop", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_ipsec_cryptomap_add", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_ipsec_cryptomap_attach", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_ipsec_cryptomap_delete", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_ipsec_cryptomap_detach", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_ipsec_tunnel_start", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_ipsec_tunnel_stop", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_ipsec_too_many_sas", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_config_copy", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_config", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_config_ctid", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_dhcp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_event_manager", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_ipmulticast", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_msdp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_ospf_config_state", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_ospf_config_errors", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_pim_invalid_pim_message", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_pim_invalid_neighbor_change", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_pim_rp_mapping_change", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_bridge_newroot", "true"))
	if os.Getenv("C9000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_bridge_topologychange", "true"))
	}
	if os.Getenv("C9000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_stpx_inconsistency", "true"))
	}
	if os.Getenv("C9000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_stpx_root_inconsistency", "true"))
	}
	if os.Getenv("C9000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_stpx_loop_inconsistency", "true"))
	}
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_syslog", "true"))
	if os.Getenv("C9000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_vlan_membership", "true"))
	}
	if os.Getenv("C9000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_errdisable", "true"))
	}
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_rf", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_transceiver_all", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_bulkstat_collection", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_bulkstat_transfer", "true"))
	if os.Getenv("C9000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_mac_notification_change", "true"))
	}
	if os.Getenv("C9000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_mac_notification_move", "true"))
	}
	if os.Getenv("C9000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_mac_notification_threshold", "true"))
	}
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_vrfmib_vrf_up", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_vrfmib_vrf_down", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_vnet_trunk_up", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "enable_traps_vnet_trunk_down", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "source_interface_informs_loopback", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "source_interface_traps_loopback", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "trap_source_loopback", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "snmp_communities.0.name", "COM1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "snmp_communities.0.view", "VIEW1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "snmp_communities.0.permission", "ro"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "snmp_communities.0.ipv6", "ACL1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "snmp_communities.0.access_list_name", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "contexts.0.name", "CON1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "views.0.name", "VIEW1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "views.0.mib", "interfaces"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_snmp_server.test", "views.0.inc_exl", "included"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeSNMPServerPrerequisitesConfig + testAccIosxeSNMPServerConfig_minimum(),
			},
			{
				Config: testAccIosxeSNMPServerPrerequisitesConfig + testAccIosxeSNMPServerConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxe_snmp_server.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/snmp-server",
			},
		},
	})
}

const testAccIosxeSNMPServerPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
	path = "Cisco-IOS-XE-native:native/interface/Loopback=1"
	attributes = {
		"name" = "1"
	}
}

`

func testAccIosxeSNMPServerConfig_minimum() string {
	config := `resource "iosxe_snmp_server" "test" {` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, ]` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxeSNMPServerConfig_all() string {
	config := `resource "iosxe_snmp_server" "test" {` + "\n"
	config += `	chassis_id = "R1"` + "\n"
	config += `	contact = "Contact1"` + "\n"
	config += `	ifindex_persist = true` + "\n"
	config += `	location = "Location1"` + "\n"
	config += `	packetsize = 2000` + "\n"
	config += `	queue_length = 100` + "\n"
	config += `	enable_logging_getop = true` + "\n"
	config += `	enable_logging_setop = true` + "\n"
	config += `	enable_traps = true` + "\n"
	config += `	enable_traps_snmp_authentication = true` + "\n"
	config += `	enable_traps_snmp_coldstart = true` + "\n"
	config += `	enable_traps_snmp_linkdown = true` + "\n"
	config += `	enable_traps_snmp_linkup = true` + "\n"
	config += `	enable_traps_snmp_warmstart = true` + "\n"
	config += `	system_shutdown = true` + "\n"
	config += `	enable_traps_snmp_flowmon = true` + "\n"
	config += `	enable_traps_entity_perf_throughput_notif = true` + "\n"
	config += `	enable_traps_call_home = true` + "\n"
	config += `	enable_traps_call_home_server_fail = true` + "\n"
	config += `	enable_traps_tty = true` + "\n"
	config += `	enable_traps_ospfv3_config_state_change = true` + "\n"
	config += `	enable_traps_ospfv3_errors_change = true` + "\n"
	config += `	enable_traps_ospf_config_retransmit = true` + "\n"
	config += `	enable_traps_ospf_config_lsa = true` + "\n"
	config += `	enable_traps_ospf_nssa_trans_change = true` + "\n"
	config += `	enable_traps_ospf_shamlink_interface = true` + "\n"
	config += `	enable_traps_ospf_shamlink_neighbor = true` + "\n"
	config += `	enable_traps_ospf_errors_enable = true` + "\n"
	config += `	enable_traps_ospf_retransmit_enable = true` + "\n"
	config += `	enable_traps_ospf_lsa_enable = true` + "\n"
	config += `	enable_traps_eigrp = true` + "\n"
	config += `	enable_traps_auth_framework_sec_violation = true` + "\n"
	if os.Getenv("C9000V") != "" {
		config += `	enable_traps_rep = true` + "\n"
	}
	config += `	enable_traps_vtp = true` + "\n"
	config += `	enable_traps_vlancreate = true` + "\n"
	config += `	enable_traps_vlandelete = true` + "\n"
	config += `	enable_traps_port_security = true` + "\n"
	config += `	enable_traps_license = true` + "\n"
	config += `	enable_traps_smart_license = true` + "\n"
	config += `	enable_traps_cpu_threshold = true` + "\n"
	config += `	enable_traps_memory_bufferpeak = true` + "\n"
	if os.Getenv("C9000V") != "" {
		config += `	enable_traps_stackwise = true` + "\n"
	}
	if os.Getenv("C9000V") != "" {
		config += `	enable_traps_link_fail_rpt = true` + "\n"
	}
	if os.Getenv("C9000V") != "" {
		config += `	enable_traps_status_change = true` + "\n"
	}
	config += `	enable_traps_fru_ctrl = true` + "\n"
	config += `	enable_traps_flash_insertion = true` + "\n"
	config += `	enable_traps_flash_removal = true` + "\n"
	config += `	enable_traps_flash_lowspace = true` + "\n"
	if os.Getenv("C9000V") != "" {
		config += `	enable_traps_energywise = true` + "\n"
	}
	config += `	enable_traps_entity = true` + "\n"
	config += `	enable_traps_pw_vc = true` + "\n"
	if os.Getenv("C9000V") != "" {
		config += `	enable_traps_envmon = true` + "\n"
	}
	config += `	enable_traps_ipsla = true` + "\n"
	config += `	enable_traps_bfd = true` + "\n"
	config += `	enable_traps_ike_policy_add = true` + "\n"
	config += `	enable_traps_ike_policy_delete = true` + "\n"
	config += `	enable_traps_ike_tunnel_start = true` + "\n"
	config += `	enable_traps_ike_tunnel_stop = true` + "\n"
	config += `	enable_traps_ipsec_cryptomap_add = true` + "\n"
	config += `	enable_traps_ipsec_cryptomap_attach = true` + "\n"
	config += `	enable_traps_ipsec_cryptomap_delete = true` + "\n"
	config += `	enable_traps_ipsec_cryptomap_detach = true` + "\n"
	config += `	enable_traps_ipsec_tunnel_start = true` + "\n"
	config += `	enable_traps_ipsec_tunnel_stop = true` + "\n"
	config += `	enable_traps_ipsec_too_many_sas = true` + "\n"
	config += `	enable_traps_config_copy = true` + "\n"
	config += `	enable_traps_config = true` + "\n"
	config += `	enable_traps_config_ctid = true` + "\n"
	config += `	enable_traps_dhcp = true` + "\n"
	config += `	enable_traps_event_manager = true` + "\n"
	config += `	enable_traps_ipmulticast = true` + "\n"
	config += `	enable_traps_msdp = true` + "\n"
	config += `	enable_traps_ospf_config_state = true` + "\n"
	config += `	enable_traps_ospf_config_errors = true` + "\n"
	config += `	enable_traps_pim_invalid_pim_message = true` + "\n"
	config += `	enable_traps_pim_invalid_neighbor_change = true` + "\n"
	config += `	enable_traps_pim_rp_mapping_change = true` + "\n"
	config += `	enable_traps_bridge_newroot = true` + "\n"
	if os.Getenv("C9000V") != "" {
		config += `	enable_traps_bridge_topologychange = true` + "\n"
	}
	if os.Getenv("C9000V") != "" {
		config += `	enable_traps_stpx_inconsistency = true` + "\n"
	}
	if os.Getenv("C9000V") != "" {
		config += `	enable_traps_stpx_root_inconsistency = true` + "\n"
	}
	if os.Getenv("C9000V") != "" {
		config += `	enable_traps_stpx_loop_inconsistency = true` + "\n"
	}
	config += `	enable_traps_syslog = true` + "\n"
	if os.Getenv("C9000V") != "" {
		config += `	enable_traps_vlan_membership = true` + "\n"
	}
	if os.Getenv("C9000V") != "" {
		config += `	enable_traps_errdisable = true` + "\n"
	}
	config += `	enable_traps_rf = true` + "\n"
	config += `	enable_traps_transceiver_all = true` + "\n"
	config += `	enable_traps_bulkstat_collection = true` + "\n"
	config += `	enable_traps_bulkstat_transfer = true` + "\n"
	if os.Getenv("C9000V") != "" {
		config += `	enable_traps_mac_notification_change = true` + "\n"
	}
	if os.Getenv("C9000V") != "" {
		config += `	enable_traps_mac_notification_move = true` + "\n"
	}
	if os.Getenv("C9000V") != "" {
		config += `	enable_traps_mac_notification_threshold = true` + "\n"
	}
	config += `	enable_traps_vrfmib_vrf_up = true` + "\n"
	config += `	enable_traps_vrfmib_vrf_down = true` + "\n"
	config += `	enable_traps_vnet_trunk_up = true` + "\n"
	config += `	enable_traps_vnet_trunk_down = true` + "\n"
	config += `	source_interface_informs_loopback = 1` + "\n"
	config += `	source_interface_traps_loopback = 1` + "\n"
	config += `	trap_source_loopback = 1` + "\n"
	config += `	snmp_communities = [{` + "\n"
	config += `		name = "COM1"` + "\n"
	config += `		view = "VIEW1"` + "\n"
	config += `		permission = "ro"` + "\n"
	config += `		ipv6 = "ACL1"` + "\n"
	config += `		access_list_name = "1"` + "\n"
	config += `	}]` + "\n"
	config += `	contexts = [{` + "\n"
	config += `		name = "CON1"` + "\n"
	config += `	}]` + "\n"
	config += `	views = [{` + "\n"
	config += `		name = "VIEW1"` + "\n"
	config += `		mib = "interfaces"` + "\n"
	config += `		inc_exl = "included"` + "\n"
	config += `	}]` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, ]` + "\n"
	config += `}` + "\n"
	return config
}
