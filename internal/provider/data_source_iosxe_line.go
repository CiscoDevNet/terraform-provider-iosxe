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
	_ datasource.DataSource              = &LineDataSource{}
	_ datasource.DataSourceWithConfigure = &LineDataSource{}
)

func NewLineDataSource() datasource.DataSource {
	return &LineDataSource{}
}

type LineDataSource struct {
	clients map[string]*restconf.Client
}

func (d *LineDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_line"
}

func (d *LineDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Line configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"console": schema.ListNestedAttribute{
				MarkdownDescription: "Primary terminal line",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"first": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"exec_timeout_minutes": schema.Int64Attribute{
							MarkdownDescription: "<0-35791>;;Timeout in minutes",
							Computed:            true,
						},
						"exec_timeout_seconds": schema.Int64Attribute{
							MarkdownDescription: "<0-2147483>;;Timeout in seconds",
							Computed:            true,
						},
						"login_local": schema.BoolAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"login_authentication": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"privilege_level": schema.Int64Attribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"stopbits": schema.StringAttribute{
							MarkdownDescription: "Set async line stop bits",
							Computed:            true,
						},
						"password_level": schema.Int64Attribute{
							MarkdownDescription: "Set exec level password",
							Computed:            true,
						},
						"password_type": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"password": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
					},
				},
			},
			"vty": schema.ListNestedAttribute{
				MarkdownDescription: "Virtual terminal",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"first": schema.Int64Attribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"last": schema.Int64Attribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"access_classes": schema.ListNestedAttribute{
							MarkdownDescription: "",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"direction": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"access_list": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"vrf_also": schema.BoolAttribute{
										MarkdownDescription: "Same access list is applied for all VRFs",
										Computed:            true,
									},
								},
							},
						},
						"exec_timeout_minutes": schema.Int64Attribute{
							MarkdownDescription: "<0-35791>;;Timeout in minutes",
							Computed:            true,
						},
						"exec_timeout_seconds": schema.Int64Attribute{
							MarkdownDescription: "<0-2147483>;;Timeout in seconds",
							Computed:            true,
						},
						"password_level": schema.Int64Attribute{
							MarkdownDescription: "Set exec level password",
							Computed:            true,
						},
						"password_type": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"password": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"login_authentication": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"transport_preferred_protocol": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"escape_character": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *LineDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.clients = req.ProviderData.(map[string]*restconf.Client)
}

func (d *LineDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LineData

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
		config = LineData{Device: config.Device}
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
