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
	_ datasource.DataSource              = &DHCPDataSource{}
	_ datasource.DataSourceWithConfigure = &DHCPDataSource{}
)

func NewDHCPDataSource() datasource.DataSource {
	return &DHCPDataSource{}
}

type DHCPDataSource struct {
	data *IosxeProviderData
}

func (d *DHCPDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dhcp"
}

func (d *DHCPDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the DHCP configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"compatibility_suboption_link_selection": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"compatibility_suboption_server_override": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"relay_information_trust_all": schema.BoolAttribute{
				MarkdownDescription: "Received DHCP packets may contain relay info option with zero giaddr",
				Computed:            true,
			},
			"relay_information_option_default": schema.BoolAttribute{
				MarkdownDescription: "Default option, no vpn",
				Computed:            true,
			},
			"relay_information_option_vpn": schema.BoolAttribute{
				MarkdownDescription: "Insert VPN sub-options and change the giaddr to the outgoing interface",
				Computed:            true,
			},
			"snooping": schema.BoolAttribute{
				MarkdownDescription: "DHCP Snooping",
				Computed:            true,
			},
			"snooping_information_option_format_remote_id_hostname": schema.BoolAttribute{
				MarkdownDescription: "Use configured hostname for remote id",
				Computed:            true,
			},
			"snooping_vlans": schema.ListNestedAttribute{
				MarkdownDescription: "DHCP Snooping vlan (OBSOLETE)",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"vlan_id": schema.StringAttribute{
							MarkdownDescription: "DHCP Snooping vlan first number or vlan range,example: 1,3-5,7,9-11",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *DHCPDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.data = req.ProviderData.(*IosxeProviderData)
}

func (d *DHCPDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DHCPData

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
		config = DHCPData{Device: config.Device}
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
