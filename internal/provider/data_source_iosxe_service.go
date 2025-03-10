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
	_ datasource.DataSource              = &ServiceDataSource{}
	_ datasource.DataSourceWithConfigure = &ServiceDataSource{}
)

func NewServiceDataSource() datasource.DataSource {
	return &ServiceDataSource{}
}

type ServiceDataSource struct {
	data *IosxeProviderData
}

func (d *ServiceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service"
}

func (d *ServiceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Service configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"pad": schema.BoolAttribute{
				MarkdownDescription: "Enable PAD commands",
				Computed:            true,
			},
			"password_encryption": schema.BoolAttribute{
				MarkdownDescription: "Encrypt system passwords",
				Computed:            true,
			},
			"password_recovery": schema.BoolAttribute{
				MarkdownDescription: "Enable password recovery",
				Computed:            true,
			},
			"timestamps": schema.BoolAttribute{
				MarkdownDescription: "Timestamp debug/log messages",
				Computed:            true,
			},
			"timestamps_debug": schema.BoolAttribute{
				MarkdownDescription: "Timestamp debug messages",
				Computed:            true,
			},
			"timestamps_debug_datetime": schema.BoolAttribute{
				MarkdownDescription: "Timestamp with date and time",
				Computed:            true,
			},
			"timestamps_debug_datetime_msec": schema.BoolAttribute{
				MarkdownDescription: "Include milliseconds in timestamp",
				Computed:            true,
			},
			"timestamps_debug_datetime_localtime": schema.BoolAttribute{
				MarkdownDescription: "Use local time zone for timestamps",
				Computed:            true,
			},
			"timestamps_debug_datetime_show_timezone": schema.BoolAttribute{
				MarkdownDescription: "Add time zone information to timestamp",
				Computed:            true,
			},
			"timestamps_debug_datetime_year": schema.BoolAttribute{
				MarkdownDescription: "Include year in timestamp",
				Computed:            true,
			},
			"timestamps_debug_uptime": schema.BoolAttribute{
				MarkdownDescription: "Timestamp with system uptime",
				Computed:            true,
			},
			"timestamps_log": schema.BoolAttribute{
				MarkdownDescription: "Timestamp log messages",
				Computed:            true,
			},
			"timestamps_log_datetime": schema.BoolAttribute{
				MarkdownDescription: "Timestamp with date and time",
				Computed:            true,
			},
			"timestamps_log_datetime_msec": schema.BoolAttribute{
				MarkdownDescription: "Include milliseconds in timestamp",
				Computed:            true,
			},
			"timestamps_log_datetime_localtime": schema.BoolAttribute{
				MarkdownDescription: "Use local time zone for timestamps",
				Computed:            true,
			},
			"timestamps_log_datetime_show_timezone": schema.BoolAttribute{
				MarkdownDescription: "Add time zone information to timestamp",
				Computed:            true,
			},
			"timestamps_log_datetime_year": schema.BoolAttribute{
				MarkdownDescription: "Include year in timestamp",
				Computed:            true,
			},
			"timestamps_log_uptime": schema.BoolAttribute{
				MarkdownDescription: "Timestamp with system uptime",
				Computed:            true,
			},
			"dhcp": schema.BoolAttribute{
				MarkdownDescription: "Enable DHCP server and relay agent",
				Computed:            true,
			},
			"tcp_keepalives_in": schema.BoolAttribute{
				MarkdownDescription: "Generate keepalives on idle incoming network connections",
				Computed:            true,
			},
			"tcp_keepalives_out": schema.BoolAttribute{
				MarkdownDescription: "Generate keepalives on idle outgoing network connections",
				Computed:            true,
			},
			"compress_config": schema.BoolAttribute{
				MarkdownDescription: "Compress the configuration file",
				Computed:            true,
			},
			"sequence_numbers": schema.BoolAttribute{
				MarkdownDescription: "Stamp logger messages with a sequence number",
				Computed:            true,
			},
			"call_home": schema.BoolAttribute{
				MarkdownDescription: "Enable call-home service",
				Computed:            true,
			},
		},
	}
}

func (d *ServiceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.data = req.ProviderData.(*IosxeProviderData)
}

func (d *ServiceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ServiceData

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
		config = ServiceData{Device: config.Device}
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
