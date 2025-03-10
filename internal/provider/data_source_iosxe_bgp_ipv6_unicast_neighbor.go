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
	_ datasource.DataSource              = &BGPIPv6UnicastNeighborDataSource{}
	_ datasource.DataSourceWithConfigure = &BGPIPv6UnicastNeighborDataSource{}
)

func NewBGPIPv6UnicastNeighborDataSource() datasource.DataSource {
	return &BGPIPv6UnicastNeighborDataSource{}
}

type BGPIPv6UnicastNeighborDataSource struct {
	data *IosxeProviderData
}

func (d *BGPIPv6UnicastNeighborDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_bgp_ipv6_unicast_neighbor"
}

func (d *BGPIPv6UnicastNeighborDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the BGP IPv6 Unicast Neighbor configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"asn": schema.StringAttribute{
				MarkdownDescription: "",
				Required:            true,
			},
			"ip": schema.StringAttribute{
				MarkdownDescription: "",
				Required:            true,
			},
			"activate": schema.BoolAttribute{
				MarkdownDescription: "Enable the address family for this neighbor",
				Computed:            true,
			},
			"send_community": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"route_reflector_client": schema.BoolAttribute{
				MarkdownDescription: "Configure a neighbor as Route Reflector client",
				Computed:            true,
			},
			"soft_reconfiguration": schema.StringAttribute{
				MarkdownDescription: "Per neighbor soft reconfiguration",
				Computed:            true,
			},
			"default_originate": schema.BoolAttribute{
				MarkdownDescription: "Originate default route to this neighbor",
				Computed:            true,
			},
			"default_originate_route_map": schema.StringAttribute{
				MarkdownDescription: "Route-map to specify criteria to originate default",
				Computed:            true,
			},
			"route_maps": schema.ListNestedAttribute{
				MarkdownDescription: "Apply route map to neighbor",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"in_out": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"route_map_name": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *BGPIPv6UnicastNeighborDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.data = req.ProviderData.(*IosxeProviderData)
}

func (d *BGPIPv6UnicastNeighborDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config BGPIPv6UnicastNeighborData

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
		config = BGPIPv6UnicastNeighborData{Device: config.Device}
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
