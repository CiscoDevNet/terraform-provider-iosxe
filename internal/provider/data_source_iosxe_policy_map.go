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
	_ datasource.DataSource              = &PolicyMapDataSource{}
	_ datasource.DataSourceWithConfigure = &PolicyMapDataSource{}
)

func NewPolicyMapDataSource() datasource.DataSource {
	return &PolicyMapDataSource{}
}

type PolicyMapDataSource struct {
	data *IosxeProviderData
}

func (d *PolicyMapDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policy_map"
}

func (d *PolicyMapDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Policy Map configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the policy map",
				Required:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "type of the policy-map",
				Computed:            true,
			},
			"subscriber": schema.BoolAttribute{
				MarkdownDescription: "Domain name of the policy map",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Policy-Map description",
				Computed:            true,
			},
			"classes": schema.ListNestedAttribute{
				MarkdownDescription: "policy criteria",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"actions": schema.ListNestedAttribute{
							MarkdownDescription: "",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"bandwidth_bits": schema.Int64Attribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"bandwidth_percent": schema.Int64Attribute{
										MarkdownDescription: "% of total Bandwidth",
										Computed:            true,
									},
									"bandwidth_remaining_option": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"bandwidth_remaining_percent": schema.Int64Attribute{
										MarkdownDescription: "% of the remaining bandwidth",
										Computed:            true,
									},
									"bandwidth_remaining_ratio": schema.Int64Attribute{
										MarkdownDescription: "ratio for sharing excess bandwidth",
										Computed:            true,
									},
									"priority_level": schema.Int64Attribute{
										MarkdownDescription: "Multi-Level Priority Queue",
										Computed:            true,
									},
									"priority_burst": schema.Int64Attribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"queue_limit": schema.Int64Attribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"queue_limit_type": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"shape_average_bit_rate": schema.Int64Attribute{
										MarkdownDescription: "Target Bit Rate (bits/sec)",
										Computed:            true,
									},
									"shape_average_bits_per_interval_sustained": schema.Int64Attribute{
										MarkdownDescription: "bits per interval, sustained. Recommend not to configure, algo finds the best value",
										Computed:            true,
									},
									"shape_average_bits_per_interval_excess": schema.Int64Attribute{
										MarkdownDescription: "bits per interval, excess.",
										Computed:            true,
									},
									"shape_average_percent": schema.Int64Attribute{
										MarkdownDescription: "% of interface bandwidth for Committed information rate",
										Computed:            true,
									},
									"shape_average_burst_size_sustained": schema.Int64Attribute{
										MarkdownDescription: "sustained burst in milliseconds. Recommend not to configure it, the algorithm will find out the best value",
										Computed:            true,
									},
									"shape_average_ms": schema.BoolAttribute{
										MarkdownDescription: "milliseconds",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *PolicyMapDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.data = req.ProviderData.(*IosxeProviderData)
}

func (d *PolicyMapDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config PolicyMapData

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
		config = PolicyMapData{Device: config.Device}
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
