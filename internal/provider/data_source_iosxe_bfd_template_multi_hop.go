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
	_ datasource.DataSource              = &BFDTemplateMultiHopDataSource{}
	_ datasource.DataSourceWithConfigure = &BFDTemplateMultiHopDataSource{}
)

func NewBFDTemplateMultiHopDataSource() datasource.DataSource {
	return &BFDTemplateMultiHopDataSource{}
}

type BFDTemplateMultiHopDataSource struct {
	data *IosxeProviderData
}

func (d *BFDTemplateMultiHopDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_bfd_template_multi_hop"
}

func (d *BFDTemplateMultiHopDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the BFD Template Multi Hop configuration.",

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
				MarkdownDescription: "",
				Required:            true,
			},
			"echo": schema.BoolAttribute{
				MarkdownDescription: "Use echo adjunct as bfd detection mechanism",
				Computed:            true,
			},
			"interval_milliseconds_both": schema.Int64Attribute{
				MarkdownDescription: "Minimum transmit and receive interval capability",
				Computed:            true,
			},
			"interval_milliseconds_min_tx": schema.Int64Attribute{
				MarkdownDescription: "Minimum transmit interval capability",
				Computed:            true,
			},
			"interval_milliseconds_min_rx": schema.Int64Attribute{
				MarkdownDescription: "Minimum receive interval capability",
				Computed:            true,
			},
			"interval_milliseconds_multiplier": schema.Int64Attribute{
				MarkdownDescription: "Multiplier value used to compute holddown",
				Computed:            true,
			},
			"interval_microseconds": schema.BoolAttribute{
				MarkdownDescription: "Specify BFD timers in microseconds",
				Computed:            true,
			},
			"interval_microseconds_both": schema.Int64Attribute{
				MarkdownDescription: "Minimum transmit and receive interval capability",
				Computed:            true,
			},
			"interval_microseconds_min_tx": schema.Int64Attribute{
				MarkdownDescription: "Minimum transmit interval capability",
				Computed:            true,
			},
			"interval_microseconds_min_rx": schema.Int64Attribute{
				MarkdownDescription: "Minimum receive interval capability",
				Computed:            true,
			},
			"interval_microseconds_multiplier": schema.Int64Attribute{
				MarkdownDescription: "Multiplier value used to compute holddown",
				Computed:            true,
			},
			"authentication_md5_keychain": schema.StringAttribute{
				MarkdownDescription: "keychain name",
				Computed:            true,
			},
			"authentication_meticulous_md5_keychain": schema.StringAttribute{
				MarkdownDescription: "keychain name",
				Computed:            true,
			},
			"authentication_meticulous_sha_1keychain": schema.StringAttribute{
				MarkdownDescription: "keychain name",
				Computed:            true,
			},
			"authentication_sha_1_keychain": schema.StringAttribute{
				MarkdownDescription: "keychain name",
				Computed:            true,
			},
			"dampening_half_time": schema.Int64Attribute{
				MarkdownDescription: "Half-life time for the penalty",
				Computed:            true,
			},
			"dampening_unsuppress_time": schema.Int64Attribute{
				MarkdownDescription: "Value to unsuppress a session",
				Computed:            true,
			},
			"dampening_suppress_time": schema.Int64Attribute{
				MarkdownDescription: "Value to start suppressing a session",
				Computed:            true,
			},
			"dampening_max_suppressing_time": schema.Int64Attribute{
				MarkdownDescription: "Maximum duration to suppress a session",
				Computed:            true,
			},
			"dampening_threshold": schema.Int64Attribute{
				MarkdownDescription: "Stability threshold to enter dampening in down dampened state(seconds)",
				Computed:            true,
			},
			"dampening_down_monitoring": schema.BoolAttribute{
				MarkdownDescription: "down monitoring",
				Computed:            true,
			},
		},
	}
}

func (d *BFDTemplateMultiHopDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.data = req.ProviderData.(*IosxeProviderData)
}

func (d *BFDTemplateMultiHopDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config BFDTemplateMultiHopData

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
		config = BFDTemplateMultiHopData{Device: config.Device}
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
