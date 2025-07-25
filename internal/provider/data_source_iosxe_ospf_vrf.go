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
	_ datasource.DataSource              = &OSPFVRFDataSource{}
	_ datasource.DataSourceWithConfigure = &OSPFVRFDataSource{}
)

func NewOSPFVRFDataSource() datasource.DataSource {
	return &OSPFVRFDataSource{}
}

type OSPFVRFDataSource struct {
	data *IosxeProviderData
}

func (d *OSPFVRFDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ospf_vrf"
}

func (d *OSPFVRFDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the OSPF VRF configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"process_id": schema.Int64Attribute{
				MarkdownDescription: "Process ID",
				Required:            true,
			},
			"vrf": schema.StringAttribute{
				MarkdownDescription: "VPN Routing/Forwarding Instance",
				Required:            true,
			},
			"bfd_all_interfaces": schema.BoolAttribute{
				MarkdownDescription: "Enable BFD on all interfaces",
				Computed:            true,
			},
			"default_information_originate": schema.BoolAttribute{
				MarkdownDescription: "Distribute a default route",
				Computed:            true,
			},
			"default_information_originate_always": schema.BoolAttribute{
				MarkdownDescription: "Always advertise default route",
				Computed:            true,
			},
			"default_metric": schema.Int64Attribute{
				MarkdownDescription: "Set metric of redistributed routes",
				Computed:            true,
			},
			"distance": schema.Int64Attribute{
				MarkdownDescription: "Administrative distance",
				Computed:            true,
			},
			"domain_tag": schema.Int64Attribute{
				MarkdownDescription: "OSPF domain-tag",
				Computed:            true,
			},
			"mpls_ldp_autoconfig": schema.BoolAttribute{
				MarkdownDescription: "Configure LDP automatic configuration",
				Computed:            true,
			},
			"mpls_ldp_sync": schema.BoolAttribute{
				MarkdownDescription: "Configure LDP-IGP Synchronization",
				Computed:            true,
			},
			"neighbor": schema.ListNestedAttribute{
				MarkdownDescription: "Specify a neighbor router",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip": schema.StringAttribute{
							MarkdownDescription: "Neighbor address",
							Computed:            true,
						},
						"priority": schema.Int64Attribute{
							MarkdownDescription: "OSPF priority of non-broadcast neighbor",
							Computed:            true,
						},
						"cost": schema.Int64Attribute{
							MarkdownDescription: "OSPF cost for point-to-multipoint neighbor",
							Computed:            true,
						},
					},
				},
			},
			"network": schema.ListNestedAttribute{
				MarkdownDescription: "Enable routing on an IP network",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip": schema.StringAttribute{
							MarkdownDescription: "Network number",
							Computed:            true,
						},
						"wildcard": schema.StringAttribute{
							MarkdownDescription: "OSPF wild card bits",
							Computed:            true,
						},
						"area": schema.StringAttribute{
							MarkdownDescription: "Set the OSPF area ID",
							Computed:            true,
						},
					},
				},
			},
			"priority": schema.Int64Attribute{
				MarkdownDescription: "OSPF topology priority",
				Computed:            true,
			},
			"router_id": schema.StringAttribute{
				MarkdownDescription: "Configure router identifier. New router-id will take effect immediately (peers will reset)",
				Computed:            true,
			},
			"shutdown": schema.BoolAttribute{
				MarkdownDescription: "Shutdown the OSPF protocol under the current instance",
				Computed:            true,
			},
			"summary_address": schema.ListNestedAttribute{
				MarkdownDescription: "Configure IP address summaries",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip": schema.StringAttribute{
							MarkdownDescription: "IP summary address",
							Computed:            true,
						},
						"mask": schema.StringAttribute{
							MarkdownDescription: "Summary mask",
							Computed:            true,
						},
					},
				},
			},
			"areas": schema.ListNestedAttribute{
				MarkdownDescription: "OSPF area parameters",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"area_id": schema.StringAttribute{
							MarkdownDescription: "OSPF area ID",
							Computed:            true,
						},
						"authentication_message_digest": schema.BoolAttribute{
							MarkdownDescription: "Use message-digest authentication",
							Computed:            true,
						},
						"nssa": schema.BoolAttribute{
							MarkdownDescription: "Specify a NSSA area",
							Computed:            true,
						},
						"nssa_default_information_originate": schema.BoolAttribute{
							MarkdownDescription: "Originate Type 7 default into NSSA area",
							Computed:            true,
						},
						"nssa_default_information_originate_metric": schema.Int64Attribute{
							MarkdownDescription: "OSPF default metric",
							Computed:            true,
						},
						"nssa_default_information_originate_metric_type": schema.Int64Attribute{
							MarkdownDescription: "OSPF metric type for default routes",
							Computed:            true,
						},
						"nssa_no_summary": schema.BoolAttribute{
							MarkdownDescription: "Do not send summary LSA into NSSA",
							Computed:            true,
						},
						"nssa_no_redistribution": schema.BoolAttribute{
							MarkdownDescription: "No redistribution into this NSSA area",
							Computed:            true,
						},
					},
				},
			},
			"passive_interface_default": schema.BoolAttribute{
				MarkdownDescription: "Suppress routing updates on all interfaces",
				Computed:            true,
			},
		},
	}
}

func (d *OSPFVRFDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.data = req.ProviderData.(*IosxeProviderData)
}

func (d *OSPFVRFDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config OSPFVRFData

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
		config = OSPFVRFData{Device: config.Device}
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
