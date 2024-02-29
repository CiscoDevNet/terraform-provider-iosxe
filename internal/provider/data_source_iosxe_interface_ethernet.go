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
	"github.com/netascode/go-restconf"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &InterfaceEthernetDataSource{}
	_ datasource.DataSourceWithConfigure = &InterfaceEthernetDataSource{}
)

func NewInterfaceEthernetDataSource() datasource.DataSource {
	return &InterfaceEthernetDataSource{}
}

type InterfaceEthernetDataSource struct {
	clients map[string]*restconf.Client
}

func (d *InterfaceEthernetDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_ethernet"
}

func (d *InterfaceEthernetDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Interface Ethernet configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Interface type",
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "",
				Required:            true,
			},
			"media_type": schema.StringAttribute{
				MarkdownDescription: "Media type",
				Computed:            true,
			},
			"bandwidth": schema.Int64Attribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"switchport": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Interface specific description",
				Computed:            true,
			},
			"shutdown": schema.BoolAttribute{
				MarkdownDescription: "Shutdown the selected interface",
				Computed:            true,
			},
			"ip_proxy_arp": schema.BoolAttribute{
				MarkdownDescription: "Enable proxy ARP",
				Computed:            true,
			},
			"ip_redirects": schema.BoolAttribute{
				MarkdownDescription: "Enable sending ICMP Redirect messages",
				Computed:            true,
			},
			"ip_unreachables": schema.BoolAttribute{
				MarkdownDescription: "Enable sending ICMP Unreachable messages",
				Computed:            true,
			},
			"vrf_forwarding": schema.StringAttribute{
				MarkdownDescription: "Configure forwarding table",
				Computed:            true,
			},
			"ipv4_address": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"ipv4_address_mask": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"unnumbered": schema.StringAttribute{
				MarkdownDescription: "Enable IP processing without an explicit address",
				Computed:            true,
			},
			"encapsulation_dot1q_vlan_id": schema.Int64Attribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"channel_group_number": schema.Int64Attribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"channel_group_mode": schema.StringAttribute{
				MarkdownDescription: "Etherchannel Mode of the interface",
				Computed:            true,
			},
			"ip_dhcp_relay_source_interface": schema.StringAttribute{
				MarkdownDescription: "Set source interface for relayed messages",
				Computed:            true,
			},
			"ip_access_group_in": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"ip_access_group_in_enable": schema.BoolAttribute{
				MarkdownDescription: "inbound packets",
				Computed:            true,
			},
			"ip_access_group_out": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"ip_access_group_out_enable": schema.BoolAttribute{
				MarkdownDescription: "outbound packets",
				Computed:            true,
			},
			"spanning_tree_guard": schema.StringAttribute{
				MarkdownDescription: "Change an interface's spanning tree guard mode",
				Computed:            true,
			},
			"auto_qos_classify": schema.BoolAttribute{
				MarkdownDescription: "Configure classification for untrusted devices",
				Computed:            true,
			},
			"auto_qos_classify_police": schema.BoolAttribute{
				MarkdownDescription: "Configure QoS policing for untrusted devices",
				Computed:            true,
			},
			"auto_qos_trust": schema.BoolAttribute{
				MarkdownDescription: "Trust the DSCP/CoS marking",
				Computed:            true,
			},
			"auto_qos_trust_cos": schema.BoolAttribute{
				MarkdownDescription: "Trust the CoS marking",
				Computed:            true,
			},
			"auto_qos_trust_dscp": schema.BoolAttribute{
				MarkdownDescription: "Trust the DSCP marking",
				Computed:            true,
			},
			"auto_qos_video_cts": schema.BoolAttribute{
				MarkdownDescription: "Trust the QoS marking of the Cisco Telepresence System",
				Computed:            true,
			},
			"auto_qos_video_ip_camera": schema.BoolAttribute{
				MarkdownDescription: "Trust the QoS marking of the Ip Video Surveillance camera",
				Computed:            true,
			},
			"auto_qos_video_media_player": schema.BoolAttribute{
				MarkdownDescription: "Trust the Qos marking of the Cisco Media Player",
				Computed:            true,
			},
			"auto_qos_voip": schema.BoolAttribute{
				MarkdownDescription: "Configure AutoQoS for VoIP",
				Computed:            true,
			},
			"auto_qos_voip_cisco_phone": schema.BoolAttribute{
				MarkdownDescription: "Trust the QoS marking of Cisco IP Phone",
				Computed:            true,
			},
			"auto_qos_voip_cisco_softphone": schema.BoolAttribute{
				MarkdownDescription: "Trust the QoS marking of Cisco IP SoftPhone",
				Computed:            true,
			},
			"auto_qos_voip_trust": schema.BoolAttribute{
				MarkdownDescription: "Trust the DSCP/CoS marking",
				Computed:            true,
			},
			"trust_device": schema.StringAttribute{
				MarkdownDescription: "trusted device class",
				Computed:            true,
			},
			"helper_addresses": schema.ListNestedAttribute{
				MarkdownDescription: "Specify a destination address for UDP broadcasts",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"global": schema.BoolAttribute{
							MarkdownDescription: "Helper-address is global",
							Computed:            true,
						},
						"vrf": schema.StringAttribute{
							MarkdownDescription: "VRF name for helper-address (if different from interface VRF)",
							Computed:            true,
						},
					},
				},
			},
			"source_template": schema.ListNestedAttribute{
				MarkdownDescription: "",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"template_name": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"merge": schema.BoolAttribute{
							MarkdownDescription: "merge option of binding",
							Computed:            true,
						},
					},
				},
			},
			"bfd_template": schema.StringAttribute{
				MarkdownDescription: "BFD template",
				Computed:            true,
			},
			"bfd_enable": schema.BoolAttribute{
				MarkdownDescription: "Enable BFD under the interface",
				Computed:            true,
			},
			"bfd_local_address": schema.StringAttribute{
				MarkdownDescription: "The Source IP address to be used for BFD sessions over this interface.",
				Computed:            true,
			},
			"bfd_interval": schema.Int64Attribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"bfd_interval_min_rx": schema.Int64Attribute{
				MarkdownDescription: "Minimum receive interval capability",
				Computed:            true,
			},
			"bfd_interval_multiplier": schema.Int64Attribute{
				MarkdownDescription: "Multiplier value used to compute holddown",
				Computed:            true,
			},
			"bfd_echo": schema.BoolAttribute{
				MarkdownDescription: "Use echo adjunct as bfd detection mechanism",
				Computed:            true,
			},
			"ipv6_enable": schema.BoolAttribute{
				MarkdownDescription: "Enable IPv6 on interface",
				Computed:            true,
			},
			"ipv6_mtu": schema.Int64Attribute{
				MarkdownDescription: "Set IPv6 Maximum Transmission Unit",
				Computed:            true,
			},
			"ipv6_nd_ra_suppress_all": schema.BoolAttribute{
				MarkdownDescription: "Suppress all IPv6 RA",
				Computed:            true,
			},
			"ipv6_address_autoconfig_default": schema.BoolAttribute{
				MarkdownDescription: "Insert default route",
				Computed:            true,
			},
			"ipv6_address_dhcp": schema.BoolAttribute{
				MarkdownDescription: "Obtain IPv6 address from DHCP server",
				Computed:            true,
			},
			"ipv6_link_local_addresses": schema.ListNestedAttribute{
				MarkdownDescription: "",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"link_local": schema.BoolAttribute{
							MarkdownDescription: "Use link-local address",
							Computed:            true,
						},
					},
				},
			},
			"ipv6_addresses": schema.ListNestedAttribute{
				MarkdownDescription: "",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"prefix": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"eui_64": schema.BoolAttribute{
							MarkdownDescription: "Use eui-64 interface identifier",
							Computed:            true,
						},
					},
				},
			},
			"arp_timeout": schema.Int64Attribute{
				MarkdownDescription: "Set ARP cache timeout",
				Computed:            true,
			},
			"spanning_tree_link_type": schema.StringAttribute{
				MarkdownDescription: "Specify a link type for spanning tree tree protocol use",
				Computed:            true,
			},
			"spanning_tree_portfast_trunk": schema.BoolAttribute{
				MarkdownDescription: "Enable portfast on the interface even in trunk mode",
				Computed:            true,
			},
			"ip_arp_inspection_trust": schema.BoolAttribute{
				MarkdownDescription: "Configure Trust state",
				Computed:            true,
			},
			"ip_arp_inspection_limit_rate": schema.Int64Attribute{
				MarkdownDescription: "Rate Limit",
				Computed:            true,
			},
			"ip_dhcp_snooping_trust": schema.BoolAttribute{
				MarkdownDescription: "DHCP Snooping trust config",
				Computed:            true,
			},
			"speed_100": schema.BoolAttribute{
				MarkdownDescription: "100 Mbps operation",
				Computed:            true,
			},
			"speed_1000": schema.BoolAttribute{
				MarkdownDescription: "1000 Mbps operation",
				Computed:            true,
			},
			"speed_2500": schema.BoolAttribute{
				MarkdownDescription: "2500 Mbps operation",
				Computed:            true,
			},
			"speed_5000": schema.BoolAttribute{
				MarkdownDescription: "5000 Mbps operation",
				Computed:            true,
			},
			"speed_10000": schema.BoolAttribute{
				MarkdownDescription: "10000 Mbps operation",
				Computed:            true,
			},
			"speed_25000": schema.BoolAttribute{
				MarkdownDescription: "25000 Mbps operation",
				Computed:            true,
			},
			"speed_40000": schema.BoolAttribute{
				MarkdownDescription: "40000 Mbps operation",
				Computed:            true,
			},
			"speed_100000": schema.BoolAttribute{
				MarkdownDescription: "100000 Mbps operation",
				Computed:            true,
			},
			"negotiation_auto": schema.BoolAttribute{
				MarkdownDescription: "Enable link autonegotiation",
				Computed:            true,
			},
			"speed_nonegotiate": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"authentication_host_mode": schema.StringAttribute{
				MarkdownDescription: "Set the Host mode for authentication on this interface",
				Computed:            true,
			},
			"authentication_order_dot1x": schema.BoolAttribute{
				MarkdownDescription: "Authentication method dot1x allowed",
				Computed:            true,
			},
			"authentication_order_dot1x_mab": schema.BoolAttribute{
				MarkdownDescription: "Authentication method mab allowed",
				Computed:            true,
			},
			"authentication_order_dot1x_webauth": schema.BoolAttribute{
				MarkdownDescription: "Authentication method webauth allowed",
				Computed:            true,
			},
			"authentication_order_mab": schema.BoolAttribute{
				MarkdownDescription: "Authentication method mab allowed",
				Computed:            true,
			},
			"authentication_order_mab_dot1x": schema.BoolAttribute{
				MarkdownDescription: "Authentication method dot1x allowed",
				Computed:            true,
			},
			"authentication_order_mab_webauth": schema.BoolAttribute{
				MarkdownDescription: "Authentication method webauth allowed",
				Computed:            true,
			},
			"authentication_order_webauth": schema.BoolAttribute{
				MarkdownDescription: "Authentication method webauth allowed",
				Computed:            true,
			},
			"authentication_priority_dot1x": schema.BoolAttribute{
				MarkdownDescription: "Authentication method dot1x allowed",
				Computed:            true,
			},
			"authentication_priority_dot1x_mab": schema.BoolAttribute{
				MarkdownDescription: "Authentication method mab allowed",
				Computed:            true,
			},
			"authentication_priority_dot1x_webauth": schema.BoolAttribute{
				MarkdownDescription: "Authentication method webauth allowed",
				Computed:            true,
			},
			"authentication_priority_mab": schema.BoolAttribute{
				MarkdownDescription: "Authentication method mab allowed",
				Computed:            true,
			},
			"authentication_priority_mab_dot1x": schema.BoolAttribute{
				MarkdownDescription: "Authentication method dot1x allowed",
				Computed:            true,
			},
			"authentication_priority_mab_webauth": schema.BoolAttribute{
				MarkdownDescription: "Authentication method webauth allowed",
				Computed:            true,
			},
			"authentication_priority_webauth": schema.BoolAttribute{
				MarkdownDescription: "Authentication method webauth allowed",
				Computed:            true,
			},
			"authentication_port_control": schema.StringAttribute{
				MarkdownDescription: "set the port-control value",
				Computed:            true,
			},
			"authentication_periodic": schema.BoolAttribute{
				MarkdownDescription: "Enable or Disable Reauthentication for this port",
				Computed:            true,
			},
			"authentication_timer_reauthenticate": schema.Int64Attribute{
				MarkdownDescription: "Enter a value between 1 and 1073741823",
				Computed:            true,
			},
			"authentication_timer_reauthenticate_server": schema.BoolAttribute{
				MarkdownDescription: "Obtain re-authentication timeout value from the server",
				Computed:            true,
			},
			"mab": schema.BoolAttribute{
				MarkdownDescription: "MAC Authentication Bypass Interface Config Commands",
				Computed:            true,
			},
			"mab_eap": schema.BoolAttribute{
				MarkdownDescription: "Use EAP authentication for MAC Auth Bypass",
				Computed:            true,
			},
			"dot1x_pae": schema.StringAttribute{
				MarkdownDescription: "Set 802.1x interface pae type",
				Computed:            true,
			},
			"dot1x_timeout_auth_period": schema.Int64Attribute{
				MarkdownDescription: "Timeout for authenticator reply",
				Computed:            true,
			},
			"dot1x_timeout_held_period": schema.Int64Attribute{
				MarkdownDescription: "Timeout for authentication retries",
				Computed:            true,
			},
			"dot1x_timeout_quiet_period": schema.Int64Attribute{
				MarkdownDescription: "QuietPeriod in Seconds",
				Computed:            true,
			},
			"dot1x_timeout_ratelimit_period": schema.Int64Attribute{
				MarkdownDescription: "Ratelimit Period in seconds",
				Computed:            true,
			},
			"dot1x_timeout_server_timeout": schema.Int64Attribute{
				MarkdownDescription: "Timeout for Radius Retries",
				Computed:            true,
			},
			"dot1x_timeout_start_period": schema.Int64Attribute{
				MarkdownDescription: "Timeout for EAPOL-start retries",
				Computed:            true,
			},
			"dot1x_timeout_supp_timeout": schema.Int64Attribute{
				MarkdownDescription: "Timeout for supplicant reply",
				Computed:            true,
			},
			"dot1x_timeout_tx_period": schema.Int64Attribute{
				MarkdownDescription: "Timeout for supplicant retries",
				Computed:            true,
			},
			"dot1x_max_req": schema.Int64Attribute{
				MarkdownDescription: "Max No. of Retries",
				Computed:            true,
			},
			"dot1x_max_reauth_req": schema.Int64Attribute{
				MarkdownDescription: "Max No. of Reauthentication Attempts",
				Computed:            true,
			},
			"service_policy_input": schema.StringAttribute{
				MarkdownDescription: "Assign policy-map to the input of an interface",
				Computed:            true,
			},
			"service_policy_output": schema.StringAttribute{
				MarkdownDescription: "Assign policy-map to the output of an interface",
				Computed:            true,
			},
		},
	}
}

func (d *InterfaceEthernetDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.clients = req.ProviderData.(map[string]*restconf.Client)
}

func (d *InterfaceEthernetDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config InterfaceEthernetData

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if _, ok := d.clients[config.Device.ValueString()]; !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", config.Device.ValueString()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getPath()))

	res, err := d.clients[config.Device.ValueString()].GetData(config.getPath())
	if res.StatusCode == 404 {
		config = InterfaceEthernetData{Device: config.Device}
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
