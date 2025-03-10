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
	_ datasource.DataSource              = &PIMDataSource{}
	_ datasource.DataSourceWithConfigure = &PIMDataSource{}
)

func NewPIMDataSource() datasource.DataSource {
	return &PIMDataSource{}
}

type PIMDataSource struct {
	data *IosxeProviderData
}

func (d *PIMDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_pim"
}

func (d *PIMDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the PIM configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"autorp": schema.BoolAttribute{
				MarkdownDescription: "Configure AutoRP global operations",
				Computed:            true,
			},
			"autorp_listener": schema.BoolAttribute{
				MarkdownDescription: "Allow AutoRP packets across sparse mode interface",
				Computed:            true,
			},
			"bsr_candidate_loopback": schema.Int64Attribute{
				MarkdownDescription: "Loopback interface",
				Computed:            true,
			},
			"bsr_candidate_mask": schema.Int64Attribute{
				MarkdownDescription: "Hash Mask length for RP selection",
				Computed:            true,
			},
			"bsr_candidate_priority": schema.Int64Attribute{
				MarkdownDescription: "Priority value for candidate bootstrap router",
				Computed:            true,
			},
			"bsr_candidate_accept_rp_candidate": schema.StringAttribute{
				MarkdownDescription: "BSR RP candidate filter",
				Computed:            true,
			},
			"ssm_range": schema.StringAttribute{
				MarkdownDescription: "ACL for group range to be used for SSM",
				Computed:            true,
			},
			"ssm_default": schema.BoolAttribute{
				MarkdownDescription: "Use 232/8 group range for SSM",
				Computed:            true,
			},
			"rp_address": schema.StringAttribute{
				MarkdownDescription: "IP address of Rendezvous-point for group",
				Computed:            true,
			},
			"rp_address_override": schema.BoolAttribute{
				MarkdownDescription: "Overrides dynamically learnt RP mappings",
				Computed:            true,
			},
			"rp_address_bidir": schema.BoolAttribute{
				MarkdownDescription: "Group range treated in bidirectional shared-tree mode",
				Computed:            true,
			},
			"rp_addresses": schema.ListNestedAttribute{
				MarkdownDescription: "PIM RP-address (Rendezvous Point)",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"access_list": schema.StringAttribute{
							MarkdownDescription: "IP Access-list",
							Computed:            true,
						},
						"rp_address": schema.StringAttribute{
							MarkdownDescription: "IP address of Rendezvous-point for group",
							Computed:            true,
						},
						"override": schema.BoolAttribute{
							MarkdownDescription: "Overrides dynamically learnt RP mappings",
							Computed:            true,
						},
						"bidir": schema.BoolAttribute{
							MarkdownDescription: "Group range treated in bidirectional shared-tree mode",
							Computed:            true,
						},
					},
				},
			},
			"rp_candidates": schema.ListNestedAttribute{
				MarkdownDescription: "To be a PIM version 2 RP candidate",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface": schema.StringAttribute{
							MarkdownDescription: "Autonomic-Networking virtual interface",
							Computed:            true,
						},
						"group_list": schema.StringAttribute{
							MarkdownDescription: "IP Access list",
							Computed:            true,
						},
						"interval": schema.Int64Attribute{
							MarkdownDescription: "RP candidate advertisement interval",
							Computed:            true,
						},
						"priority": schema.Int64Attribute{
							MarkdownDescription: "RP candidate priority",
							Computed:            true,
						},
						"bidir": schema.BoolAttribute{
							MarkdownDescription: "Group range treated in bidirectional shared-tree mode",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *PIMDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.data = req.ProviderData.(*IosxeProviderData)
}

func (d *PIMDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config PIMData

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
		config = PIMData{Device: config.Device}
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
