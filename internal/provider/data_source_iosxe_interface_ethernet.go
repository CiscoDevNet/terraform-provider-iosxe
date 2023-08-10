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
			"unreachables": schema.BoolAttribute{
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
			"template": schema.StringAttribute{
				MarkdownDescription: "BFD template",
				Computed:            true,
			},
			"enable": schema.BoolAttribute{
				MarkdownDescription: "Enable BFD under the interface",
				Computed:            true,
			},
			"local_address": schema.StringAttribute{
				MarkdownDescription: "The Source IP address to be used for BFD sessions over this interface.",
				Computed:            true,
			},
			"interval_interface_msecs": schema.Int64Attribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"interval_interface_min_rx": schema.Int64Attribute{
				MarkdownDescription: "Minimum receive interval capability",
				Computed:            true,
			},
			"interval_interface_multiplier": schema.Int64Attribute{
				MarkdownDescription: "Multiplier value used to compute holddown",
				Computed:            true,
			},
			"echo": schema.BoolAttribute{
				MarkdownDescription: "Use echo adjunct as bfd detection mechanism",
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
