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
	_ datasource.DataSource              = &InterfaceOSPFDataSource{}
	_ datasource.DataSourceWithConfigure = &InterfaceOSPFDataSource{}
)

func NewInterfaceOSPFDataSource() datasource.DataSource {
	return &InterfaceOSPFDataSource{}
}

type InterfaceOSPFDataSource struct {
	data *IosxeProviderData
}

func (d *InterfaceOSPFDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_ospf"
}

func (d *InterfaceOSPFDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Interface OSPF configuration.",

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
			"cost": schema.Int64Attribute{
				MarkdownDescription: "Route cost of this interface",
				Computed:            true,
			},
			"dead_interval": schema.Int64Attribute{
				MarkdownDescription: "Interval after which a neighbor is declared dead",
				Computed:            true,
			},
			"hello_interval": schema.Int64Attribute{
				MarkdownDescription: "Time between HELLO packets",
				Computed:            true,
			},
			"mtu_ignore": schema.BoolAttribute{
				MarkdownDescription: "Ignores the MTU in DBD packets",
				Computed:            true,
			},
			"network_type_broadcast": schema.BoolAttribute{
				MarkdownDescription: "Specify OSPF broadcast multi-access network",
				Computed:            true,
			},
			"network_type_non_broadcast": schema.BoolAttribute{
				MarkdownDescription: "Specify OSPF NBMA network",
				Computed:            true,
			},
			"network_type_point_to_multipoint": schema.BoolAttribute{
				MarkdownDescription: "Specify OSPF point-to-multipoint network",
				Computed:            true,
			},
			"network_type_point_to_point": schema.BoolAttribute{
				MarkdownDescription: "Specify OSPF point-to-point network",
				Computed:            true,
			},
			"priority": schema.Int64Attribute{
				MarkdownDescription: "Router priority",
				Computed:            true,
			},
			"ttl_security_hops": schema.Int64Attribute{
				MarkdownDescription: "IP hops",
				Computed:            true,
			},
			"process_ids": schema.ListNestedAttribute{
				MarkdownDescription: "",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							MarkdownDescription: "Process ID",
							Computed:            true,
						},
						"areas": schema.ListNestedAttribute{
							MarkdownDescription: "",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"area_id": schema.StringAttribute{
										MarkdownDescription: "Set the OSPF area ID",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
			"message_digest_keys": schema.ListNestedAttribute{
				MarkdownDescription: "Message digest authentication password (key)",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							MarkdownDescription: "Key ID",
							Computed:            true,
						},
						"md5_auth_key": schema.StringAttribute{
							MarkdownDescription: "The OSPF password (key) (only the first 16 characters are used)",
							Computed:            true,
						},
						"md5_auth_type": schema.Int64Attribute{
							MarkdownDescription: "Encryption type (0 for not yet encrypted, 7 for proprietary)",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *InterfaceOSPFDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.data = req.ProviderData.(*IosxeProviderData)
}

func (d *InterfaceOSPFDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config InterfaceOSPFData

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
		config = InterfaceOSPFData{Device: config.Device}
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
