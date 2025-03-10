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
	_ datasource.DataSource              = &BFDDataSource{}
	_ datasource.DataSourceWithConfigure = &BFDDataSource{}
)

func NewBFDDataSource() datasource.DataSource {
	return &BFDDataSource{}
}

type BFDDataSource struct {
	data *IosxeProviderData
}

func (d *BFDDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_bfd"
}

func (d *BFDDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the BFD configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"ipv4_both_vrfs": schema.ListNestedAttribute{
				MarkdownDescription: "IPv4 Address Family with vrf",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"dst_vrf": schema.StringAttribute{
							MarkdownDescription: "Destination VRF instance name",
							Computed:            true,
						},
						"dest_ip": schema.StringAttribute{
							MarkdownDescription: "Destination IP prefix/len",
							Computed:            true,
						},
						"src_vrf": schema.StringAttribute{
							MarkdownDescription: "source VRF instance name",
							Computed:            true,
						},
						"src_ip": schema.StringAttribute{
							MarkdownDescription: "Source IP prefix/len",
							Computed:            true,
						},
						"template_name": schema.StringAttribute{
							MarkdownDescription: "BFD template name",
							Computed:            true,
						},
					},
				},
			},
			"ipv4_without_vrfs": schema.ListNestedAttribute{
				MarkdownDescription: "IPv4 Address Family with vrf",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"dest_ip": schema.StringAttribute{
							MarkdownDescription: "Destination IP prefix/len",
							Computed:            true,
						},
						"src_ip": schema.StringAttribute{
							MarkdownDescription: "Source IP prefix/len",
							Computed:            true,
						},
						"template_name": schema.StringAttribute{
							MarkdownDescription: "BFD template name",
							Computed:            true,
						},
					},
				},
			},
			"ipv4_with_src_vrfs": schema.ListNestedAttribute{
				MarkdownDescription: "IPv4 Address Family with vrf",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"dest_ip": schema.StringAttribute{
							MarkdownDescription: "Destination IP prefix/len",
							Computed:            true,
						},
						"src_vrf": schema.StringAttribute{
							MarkdownDescription: "source VRF instance name",
							Computed:            true,
						},
						"src_ip": schema.StringAttribute{
							MarkdownDescription: "Source IP prefix/len",
							Computed:            true,
						},
						"template_name": schema.StringAttribute{
							MarkdownDescription: "BFD template name",
							Computed:            true,
						},
					},
				},
			},
			"ipv4_with_dst_vrfs": schema.ListNestedAttribute{
				MarkdownDescription: "IPv4 Address Family with vrf",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"dst_vrf": schema.StringAttribute{
							MarkdownDescription: "Destination VRF instance name",
							Computed:            true,
						},
						"dest_ip": schema.StringAttribute{
							MarkdownDescription: "Destination IP prefix/len",
							Computed:            true,
						},
						"src_ip": schema.StringAttribute{
							MarkdownDescription: "Source IP prefix/len",
							Computed:            true,
						},
						"template_name": schema.StringAttribute{
							MarkdownDescription: "BFD template name",
							Computed:            true,
						},
					},
				},
			},
			"ipv6_with_both_vrfs": schema.ListNestedAttribute{
				MarkdownDescription: "IPv6 Address Family with vrf",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"dst_vrf": schema.StringAttribute{
							MarkdownDescription: "Destination VRF instance name",
							Computed:            true,
						},
						"dest_ipv6": schema.StringAttribute{
							MarkdownDescription: "Destination IPv6 prefix/len",
							Computed:            true,
						},
						"src_vrf": schema.StringAttribute{
							MarkdownDescription: "source VRF instance name",
							Computed:            true,
						},
						"src_ipv6": schema.StringAttribute{
							MarkdownDescription: "Source IPv6 prefix/len",
							Computed:            true,
						},
						"template_name": schema.StringAttribute{
							MarkdownDescription: "BFD template name",
							Computed:            true,
						},
					},
				},
			},
			"ipv6_without_vrfs": schema.ListNestedAttribute{
				MarkdownDescription: "IPv6 Address Family with vrf",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"dest_ipv6": schema.StringAttribute{
							MarkdownDescription: "Destination IPv6 prefix/len",
							Computed:            true,
						},
						"src_ipv6": schema.StringAttribute{
							MarkdownDescription: "Source IPv6 prefix/len",
							Computed:            true,
						},
						"template_name": schema.StringAttribute{
							MarkdownDescription: "BFD template name",
							Computed:            true,
						},
					},
				},
			},
			"ipv6_with_src_vrfs": schema.ListNestedAttribute{
				MarkdownDescription: "IPv6 Address Family with vrf",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"dest_ipv6": schema.StringAttribute{
							MarkdownDescription: "Destination IPv6 prefix/len",
							Computed:            true,
						},
						"src_vrf": schema.StringAttribute{
							MarkdownDescription: "source VRF instance name",
							Computed:            true,
						},
						"src_ipv6": schema.StringAttribute{
							MarkdownDescription: "Source IPv6 prefix/len",
							Computed:            true,
						},
						"template_name": schema.StringAttribute{
							MarkdownDescription: "BFD template name",
							Computed:            true,
						},
					},
				},
			},
			"ipv6_with_dst_vrfs": schema.ListNestedAttribute{
				MarkdownDescription: "IPv6 Address Family with vrf",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"dst_vrf": schema.StringAttribute{
							MarkdownDescription: "Destination VRF instance name",
							Computed:            true,
						},
						"dest_ipv6": schema.StringAttribute{
							MarkdownDescription: "Destination IPv6 prefix/len",
							Computed:            true,
						},
						"src_ipv6": schema.StringAttribute{
							MarkdownDescription: "Source IPv6 prefix/len",
							Computed:            true,
						},
						"template_name": schema.StringAttribute{
							MarkdownDescription: "BFD template name",
							Computed:            true,
						},
					},
				},
			},
			"slow_timers": schema.Int64Attribute{
				MarkdownDescription: "Value in ms to use for slow timers",
				Computed:            true,
			},
		},
	}
}

func (d *BFDDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.data = req.ProviderData.(*IosxeProviderData)
}

func (d *BFDDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config BFDData

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
		config = BFDData{Device: config.Device}
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
