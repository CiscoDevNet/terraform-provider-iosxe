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
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &SNMPServerDataSource{}
	_ datasource.DataSourceWithConfigure = &SNMPServerDataSource{}
)

func NewSNMPServerDataSource() datasource.DataSource {
	return &SNMPServerDataSource{}
}

type SNMPServerDataSource struct {
	data *IosxeProviderData
}

func (d *SNMPServerDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snmp_server"
}

func (d *SNMPServerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the SNMP Server configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"chassis_id": schema.StringAttribute{
				MarkdownDescription: "String to uniquely identify this chassis",
				Computed:            true,
			},
			"contact": schema.StringAttribute{
				MarkdownDescription: "Text for mib object sysContact",
				Computed:            true,
			},
			"ifindex_persist": schema.BoolAttribute{
				MarkdownDescription: "Persist interface indices",
				Computed:            true,
			},
			"location": schema.StringAttribute{
				MarkdownDescription: "Text for mib object sysLocation",
				Computed:            true,
			},
			"packetsize": schema.Int64Attribute{
				MarkdownDescription: "Largest SNMP packet size",
				Computed:            true,
			},
			"queue_length": schema.Int64Attribute{
				MarkdownDescription: "Message queue length for each TRAP host",
				Computed:            true,
			},
			"enable_logging_getop": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP GET Operation logging",
				Computed:            true,
			},
			"enable_logging_setop": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP SET Operation logging",
				Computed:            true,
			},
			"enable_informs": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP Informs",
				Computed:            true,
			},
			"enable_traps": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP Traps",
				Computed:            true,
			},
			"enable_traps_snmp_authentication": schema.BoolAttribute{
				MarkdownDescription: "Enable authentication trap",
				Computed:            true,
			},
			"enable_traps_snmp_coldstart": schema.BoolAttribute{
				MarkdownDescription: "Enable coldStart trap",
				Computed:            true,
			},
			"enable_traps_snmp_linkdown": schema.BoolAttribute{
				MarkdownDescription: "Enable linkDown trap",
				Computed:            true,
			},
			"enable_traps_snmp_linkup": schema.BoolAttribute{
				MarkdownDescription: "Enable linkUp trap",
				Computed:            true,
			},
			"enable_traps_snmp_warmstart": schema.BoolAttribute{
				MarkdownDescription: "Enable warmStart trap",
				Computed:            true,
			},
			"hosts": schema.ListNestedAttribute{
				MarkdownDescription: "Specify hosts keyed by (ip-address, community-or-user)",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip_address": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"community_or_user": schema.StringAttribute{
							MarkdownDescription: "SNMPv1/v2c community string or SNMPv3 user name",
							Computed:            true,
						},
						"version": schema.StringAttribute{
							MarkdownDescription: "SNMP version to use for notification messages",
							Computed:            true,
						},
						"encryption": schema.StringAttribute{
							MarkdownDescription: "Specifies an encryption type for community string",
							Computed:            true,
						},
					},
				},
			},
			"system_shutdown": schema.BoolAttribute{
				MarkdownDescription: "Enable use of the SNMP reload command",
				Computed:            true,
			},
			"enable_traps_flowmon": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP flowmon notifications",
				Computed:            true,
			},
			"enable_traps_entity_perf_throughput_notif": schema.BoolAttribute{
				MarkdownDescription: "Enable ENTITY PERFORMANCE MIB throughput traps",
				Computed:            true,
			},
			"enable_traps_call_home_message_send_fail": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"enable_traps_call_home_server_fail": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"enable_traps_tty": schema.BoolAttribute{
				MarkdownDescription: "Enable TCP connection traps",
				Computed:            true,
			},
			"enable_traps_ospfv3_config_state_change": schema.BoolAttribute{
				MarkdownDescription: "Enable all traps of state-change",
				Computed:            true,
			},
			"enable_traps_ospfv3_config_errors": schema.BoolAttribute{
				MarkdownDescription: "Enable all traps of errors",
				Computed:            true,
			},
			"enable_traps_ospf_config_retransmit": schema.BoolAttribute{
				MarkdownDescription: "Enable all traps of retransmit",
				Computed:            true,
			},
			"enable_traps_ospf_config_lsa": schema.BoolAttribute{
				MarkdownDescription: "Enable all traps of lsa",
				Computed:            true,
			},
			"enable_traps_ospf_nssa_trans_change": schema.BoolAttribute{
				MarkdownDescription: "Nssa translator state changes",
				Computed:            true,
			},
			"enable_traps_ospf_shamlink_interface": schema.BoolAttribute{
				MarkdownDescription: "Sham link interface state changes",
				Computed:            true,
			},
			"enable_traps_ospf_shamlink_neighbor": schema.BoolAttribute{
				MarkdownDescription: "Sham link neighbor state changes",
				Computed:            true,
			},
			"enable_traps_ospf_errors_enable": schema.BoolAttribute{
				MarkdownDescription: "Enable all traps of errors",
				Computed:            true,
			},
			"enable_traps_ospf_retransmit_enable": schema.BoolAttribute{
				MarkdownDescription: "Enable all traps of retransmit",
				Computed:            true,
			},
			"enable_traps_ospf_lsa_enable": schema.BoolAttribute{
				MarkdownDescription: "Enable all traps of lsa",
				Computed:            true,
			},
			"enable_traps_eigrp": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP EIGRP traps",
				Computed:            true,
			},
			"enable_traps_auth_framework_sec_violation": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"enable_traps_rep": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"enable_traps_vtp": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"enable_traps_vlancreate": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"enable_traps_vlandelete": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"enable_traps_port_security": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP port security traps",
				Computed:            true,
			},
			"enable_traps_license": schema.BoolAttribute{
				MarkdownDescription: "Enable license traps",
				Computed:            true,
			},
			"enable_traps_smart_license": schema.BoolAttribute{
				MarkdownDescription: "Enable smart license traps",
				Computed:            true,
			},
			"enable_traps_cpu_threshold": schema.BoolAttribute{
				MarkdownDescription: "Allow CPU utilization threshold violation traps",
				Computed:            true,
			},
			"enable_traps_memory_bufferpeak": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP Memory Bufferpeak traps",
				Computed:            true,
			},
			"enable_traps_stackwise": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP stackwise traps",
				Computed:            true,
			},
			"enable_traps_udld_link_fail_rpt": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP cudldpFastHelloLinkFailRptNotification traps",
				Computed:            true,
			},
			"enable_traps_udld_status_change": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP cudldpFastHelloStatusChangeNotification traps",
				Computed:            true,
			},
			"enable_traps_fru_ctrl": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP entity FRU control traps",
				Computed:            true,
			},
			"enable_traps_flash_insertion": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP Flash Insertion notifications",
				Computed:            true,
			},
			"enable_traps_flash_removal": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP Flash Removal notifications",
				Computed:            true,
			},
			"enable_traps_flash_lowspace": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP Flash Lowspace notifications",
				Computed:            true,
			},
			"enable_traps_energywise": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"enable_traps_power_ethernet_group": schema.StringAttribute{
				MarkdownDescription: "Enable SNMP inline power group based traps",
				Computed:            true,
			},
			"enable_traps_power_ethernet_police": schema.BoolAttribute{
				MarkdownDescription: "Enable Policing Trap",
				Computed:            true,
			},
			"enable_traps_entity": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP entity traps",
				Computed:            true,
			},
			"enable_traps_pw_vc": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"enable_traps_envmon": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP environmental monitor traps",
				Computed:            true,
			},
			"enable_traps_cef_resource_failure": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"enable_traps_cef_peer_state_change": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"enable_traps_cef_peer_fib_state_change": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"enable_traps_cef_inconsistency": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"enable_traps_isis": schema.BoolAttribute{
				MarkdownDescription: "Enable ISIS traps traps",
				Computed:            true,
			},
			"enable_traps_ipsla": schema.BoolAttribute{
				MarkdownDescription: "Enable IPSLA traps traps",
				Computed:            true,
			},
			"enable_traps_entity_diag_boot_up_fail": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP ceDiagBootUpFailedNotif traps",
				Computed:            true,
			},
			"enable_traps_entity_diag_hm_test_recover": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP ceDiagHMTestRecoverNotif traps",
				Computed:            true,
			},
			"enable_traps_entity_diag_hm_thresh_reached": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP ceDiagHMThresholdReachedNotif traps",
				Computed:            true,
			},
			"enable_traps_entity_diag_scheduled_test_fail": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP ceDiagScheduledTestFailedNotif traps",
				Computed:            true,
			},
			"enable_traps_bfd": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"enable_traps_ike_policy_add": schema.BoolAttribute{
				MarkdownDescription: "Enable IKE Policy add trap",
				Computed:            true,
			},
			"enable_traps_ike_policy_delete": schema.BoolAttribute{
				MarkdownDescription: "Enable IKE Policy delete trap",
				Computed:            true,
			},
			"enable_traps_ike_tunnel_start": schema.BoolAttribute{
				MarkdownDescription: "Enable IKE Tunnel start trap",
				Computed:            true,
			},
			"enable_traps_ike_tunnel_stop": schema.BoolAttribute{
				MarkdownDescription: "Enable IKE Tunnel stop trap",
				Computed:            true,
			},
			"enable_traps_ipsec_cryptomap_add": schema.BoolAttribute{
				MarkdownDescription: "Enable IPsec Cryptomap add trap",
				Computed:            true,
			},
			"enable_traps_ipsec_cryptomap_attach": schema.BoolAttribute{
				MarkdownDescription: "Enable IPsec Cryptomap Attach trap",
				Computed:            true,
			},
			"enable_traps_ipsec_cryptomap_delete": schema.BoolAttribute{
				MarkdownDescription: "Enable IPsec Cryptomap delete trap",
				Computed:            true,
			},
			"enable_traps_ipsec_cryptomap_detach": schema.BoolAttribute{
				MarkdownDescription: "Enable IPsec Cryptomap Detach trap",
				Computed:            true,
			},
			"enable_traps_ipsec_tunnel_start": schema.BoolAttribute{
				MarkdownDescription: "Enable IPsec Tunnel Start trap",
				Computed:            true,
			},
			"enable_traps_ipsec_tunnel_stop": schema.BoolAttribute{
				MarkdownDescription: "Enable IPsec Tunnel Stop trap",
				Computed:            true,
			},
			"enable_traps_ipsec_too_many_sas": schema.BoolAttribute{
				MarkdownDescription: "Enable IPsec Tunnel Start trap",
				Computed:            true,
			},
			"enable_traps_config_copy": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP config-copy traps",
				Computed:            true,
			},
			"enable_traps_config": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP config traps",
				Computed:            true,
			},
			"enable_traps_config_ctid": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"enable_traps_dhcp": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP dhcp traps",
				Computed:            true,
			},
			"enable_traps_event_manager": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP Embedded Event Manager traps",
				Computed:            true,
			},
			"enable_traps_hsrp": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP HSRP traps",
				Computed:            true,
			},
			"enable_traps_ipmulticast": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP ipmulticast traps",
				Computed:            true,
			},
			"enable_traps_msdp": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP MSDP traps",
				Computed:            true,
			},
			"enable_traps_ospf_config_state_change": schema.BoolAttribute{
				MarkdownDescription: "Enable all traps of state-change",
				Computed:            true,
			},
			"enable_traps_ospf_config_errors": schema.BoolAttribute{
				MarkdownDescription: "Enable all traps of errors",
				Computed:            true,
			},
			"enable_traps_pim_invalid_pim_message": schema.BoolAttribute{
				MarkdownDescription: "Enable invalid pim message trap",
				Computed:            true,
			},
			"enable_traps_pim_neighbor_change": schema.BoolAttribute{
				MarkdownDescription: "Enable neighbor change trap",
				Computed:            true,
			},
			"enable_traps_pim_rp_mapping_change": schema.BoolAttribute{
				MarkdownDescription: "Enable rp mapping change trap",
				Computed:            true,
			},
			"enable_traps_bridge_newroot": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"enable_traps_bridge_topologychange": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"enable_traps_stpx_inconsistency": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"enable_traps_stpx_root_inconsistency": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"enable_traps_stpx_loop_inconsistency": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"enable_traps_syslog": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP syslog traps",
				Computed:            true,
			},
			"enable_traps_bgp_cbgp2": schema.BoolAttribute{
				MarkdownDescription: "Enable BGP MIBv2 traps",
				Computed:            true,
			},
			"enable_traps_nhrp_nhs": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"enable_traps_nhrp_nhc": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"enable_traps_nhrp_nhp": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"enable_traps_nhrp_quota_exceeded": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"enable_traps_mpls_traffic_eng": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"enable_traps_mpls_vpn": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"enable_traps_mpls_rfc_ldp": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"enable_traps_mpls_ldp": schema.BoolAttribute{
				MarkdownDescription: "SNMP MPLS label distribution protocol traps",
				Computed:            true,
			},
			"enable_traps_fast_reroute_protected": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"enable_traps_local_auth": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP local auth traps",
				Computed:            true,
			},
			"enable_traps_vlan_membership": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"enable_traps_errdisable": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP errdisable notifications",
				Computed:            true,
			},
			"enable_traps_rf": schema.BoolAttribute{
				MarkdownDescription: "Enable all SNMP traps defined in CISCO-RF-MIB",
				Computed:            true,
			},
			"enable_traps_transceiver_all": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"enable_traps_bulkstat_collection": schema.BoolAttribute{
				MarkdownDescription: "Enable Data-Collection-MIB Collection notifications",
				Computed:            true,
			},
			"enable_traps_bulkstat_transfer": schema.BoolAttribute{
				MarkdownDescription: "Enable Data-Collection-MIB Transfer notifications",
				Computed:            true,
			},
			"enable_traps_mac_notification_change": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP Change traps",
				Computed:            true,
			},
			"enable_traps_mac_notification_move": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP Move traps",
				Computed:            true,
			},
			"enable_traps_mac_notification_threshold": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP Threshold traps",
				Computed:            true,
			},
			"enable_traps_vrfmib_vrf_up": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"enable_traps_vrfmib_vrf_down": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"enable_traps_vrfmib_vnet_trunk_up": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"enable_traps_vrfmib_vnet_trunk_down": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"source_interface_informs_gigabit_ethernet": schema.StringAttribute{
				MarkdownDescription: "GigabitEthernet IEEE 802.3z",
				Computed:            true,
			},
			"source_interface_informs_ten_gigabit_ethernet": schema.StringAttribute{
				MarkdownDescription: "Ten Gigabit Ethernet",
				Computed:            true,
			},
			"source_interface_informs_forty_gigabit_ethernet": schema.StringAttribute{
				MarkdownDescription: "Forty GigabitEthernet ",
				Computed:            true,
			},
			"source_interface_informs_hundred_gig_e": schema.StringAttribute{
				MarkdownDescription: "Hundred GigabitEthernet",
				Computed:            true,
			},
			"source_interface_informs_loopback": schema.Int64Attribute{
				MarkdownDescription: "Loopback interface",
				Computed:            true,
			},
			"source_interface_informs_port_channel": schema.Int64Attribute{
				MarkdownDescription: "Ethernet Channel of interfaces",
				Computed:            true,
			},
			"source_interface_informs_port_channel_subinterface": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"source_interface_informs_vlan": schema.Int64Attribute{
				MarkdownDescription: "Iosxr Vlans",
				Computed:            true,
			},
			"source_interface_traps_gigabit_ethernet": schema.StringAttribute{
				MarkdownDescription: "GigabitEthernet IEEE 802.3z",
				Computed:            true,
			},
			"source_interface_traps_ten_gigabit_ethernet": schema.StringAttribute{
				MarkdownDescription: "Ten Gigabit Ethernet",
				Computed:            true,
			},
			"source_interface_traps_forty_gigabit_ethernet": schema.StringAttribute{
				MarkdownDescription: "Forty GigabitEthernet ",
				Computed:            true,
			},
			"source_interface_traps_hundred_gig_e": schema.StringAttribute{
				MarkdownDescription: "Hundred GigabitEthernet",
				Computed:            true,
			},
			"source_interface_traps_loopback": schema.Int64Attribute{
				MarkdownDescription: "Loopback interface",
				Computed:            true,
			},
			"source_interface_traps_port_channel": schema.Int64Attribute{
				MarkdownDescription: "Ethernet Channel of interfaces",
				Computed:            true,
			},
			"source_interface_traps_port_channel_subinterface": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"source_interface_traps_vlan": schema.Int64Attribute{
				MarkdownDescription: "Iosxr Vlans",
				Computed:            true,
			},
			"trap_source_gigabit_ethernet": schema.StringAttribute{
				MarkdownDescription: "GigabitEthernet IEEE 802.3z",
				Computed:            true,
			},
			"trap_source_ten_gigabit_ethernet": schema.StringAttribute{
				MarkdownDescription: "Ten Gigabit Ethernet",
				Computed:            true,
			},
			"trap_source_forty_gigabit_ethernet": schema.StringAttribute{
				MarkdownDescription: "Forty GigabitEthernet ",
				Computed:            true,
			},
			"trap_source_hundred_gig_e": schema.StringAttribute{
				MarkdownDescription: "Hundred GigabitEthernet",
				Computed:            true,
			},
			"trap_source_loopback": schema.Int64Attribute{
				MarkdownDescription: "Loopback interface",
				Computed:            true,
			},
			"trap_source_port_channel": schema.Int64Attribute{
				MarkdownDescription: "Ethernet Channel of interfaces",
				Computed:            true,
			},
			"trap_source_port_channel_subinterface": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"trap_source_vlan": schema.Int64Attribute{
				MarkdownDescription: "Iosxr Vlans",
				Computed:            true,
			},
			"snmp_communities": schema.ListNestedAttribute{
				MarkdownDescription: "Enable SNMP; set community string and access privs",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"view": schema.StringAttribute{
							MarkdownDescription: "Restrict this community to a named MIB view",
							Computed:            true,
						},
						"permission": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"ipv6": schema.StringAttribute{
							MarkdownDescription: "Specify IPv6 Named Access-List",
							Computed:            true,
						},
						"access_list_name": schema.StringAttribute{
							MarkdownDescription: "Access-list name",
							Computed:            true,
						},
					},
				},
			},
			"contexts": schema.ListNestedAttribute{
				MarkdownDescription: "Create/Delete a context apart from default",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
					},
				},
			},
			"views": schema.ListNestedAttribute{
				MarkdownDescription: "Define an SNMPv2 MIB view",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"mib": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"inc_exl": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *SNMPServerDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.data = req.ProviderData.(*IosxeProviderData)
}

func (d *SNMPServerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config SNMPServerData

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getPath()))

	device, ok := d.data.Devices[config.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", config.Device.ValueString()))
		return
	}

	res, err := device.Client.GetData(config.getPath())
	if res.StatusCode == 404 {
		config = SNMPServerData{Device: config.Device}
	} else {
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		}

		config.fromBody(ctx, res.Res)
	}

	config.Id = types.StringValue(config.getPath())

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.getPath()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
