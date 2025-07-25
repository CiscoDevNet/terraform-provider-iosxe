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
	_ datasource.DataSource              = &NTPDataSource{}
	_ datasource.DataSourceWithConfigure = &NTPDataSource{}
)

func NewNTPDataSource() datasource.DataSource {
	return &NTPDataSource{}
}

type NTPDataSource struct {
	data *IosxeProviderData
}

func (d *NTPDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ntp"
}

func (d *NTPDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the NTP configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"authenticate": schema.BoolAttribute{
				MarkdownDescription: "Authenticate time sources",
				Computed:            true,
			},
			"logging": schema.BoolAttribute{
				MarkdownDescription: "Enable NTP message logging",
				Computed:            true,
			},
			"access_group_peer_acl": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"access_group_query_only_acl": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"access_group_serve_acl": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"access_group_serve_only_acl": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"authentication_keys": schema.ListNestedAttribute{
				MarkdownDescription: "Authentication key for trusted time sources",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"number": schema.Int64Attribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"md5": schema.StringAttribute{
							MarkdownDescription: "MD5 authentication",
							Computed:            true,
						},
						"cmac_aes_128": schema.StringAttribute{
							MarkdownDescription: "CMAC-AES-128 (digest length = 128 bits,  key length = [16 or 32] bytes)",
							Computed:            true,
						},
						"hmac_sha1": schema.StringAttribute{
							MarkdownDescription: "HMAC-SHA1 (digest length = 160 bits,  key length = [1-32] bytes)",
							Computed:            true,
						},
						"hmac_sha2_256": schema.StringAttribute{
							MarkdownDescription: "HMAC-SHA2-256 (digest length = 256 bits,  key length = [1-32] bytes)",
							Computed:            true,
						},
						"sha1": schema.StringAttribute{
							MarkdownDescription: "SHA1 (digest length = 160 bits,  key length = [1-32] bytes)",
							Computed:            true,
						},
						"sha2": schema.StringAttribute{
							MarkdownDescription: "SHA-256 (digest length = 256 bits,  key length = [1-32] bytes)",
							Computed:            true,
						},
						"encryption_type": schema.Int64Attribute{
							MarkdownDescription: "Authentication key encryption type",
							Computed:            true,
						},
					},
				},
			},
			"clock_period": schema.Int64Attribute{
				MarkdownDescription: "Length of hardware clock tick",
				Computed:            true,
			},
			"master": schema.BoolAttribute{
				MarkdownDescription: "Act as NTP master clock",
				Computed:            true,
			},
			"master_stratum": schema.Int64Attribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"passive": schema.BoolAttribute{
				MarkdownDescription: "NTP passive mode",
				Computed:            true,
			},
			"update_calendar": schema.BoolAttribute{
				MarkdownDescription: "Periodically update calendar with NTP time",
				Computed:            true,
			},
			"trap_source_gigabit_ethernet": schema.StringAttribute{
				MarkdownDescription: "GigabitEthernet IEEE 802.3z",
				Computed:            true,
			},
			"trap_source_ten_gigabit_ethernet": schema.StringAttribute{
				MarkdownDescription: "Ten Gigabit Ethernet",
				Computed:            true,
			},
			"trap_source_forty_gigabit_ethernet": schema.StringAttribute{
				MarkdownDescription: "Forty GigabitEthernet ",
				Computed:            true,
			},
			"trap_source_hundred_gig_e": schema.StringAttribute{
				MarkdownDescription: "Hundred GigabitEthernet",
				Computed:            true,
			},
			"trap_source_loopback": schema.Int64Attribute{
				MarkdownDescription: "Loopback interface",
				Computed:            true,
			},
			"trap_source_port_channel": schema.Int64Attribute{
				MarkdownDescription: "Ethernet Channel of interfaces",
				Computed:            true,
			},
			"trap_source_port_channel_subinterface": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"trap_source_vlan": schema.Int64Attribute{
				MarkdownDescription: "Iosxr Vlans",
				Computed:            true,
			},
			"servers": schema.ListNestedAttribute{
				MarkdownDescription: "",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip_address": schema.StringAttribute{
							MarkdownDescription: "Configure ip/ipv6 address/hostname of peer",
							Computed:            true,
						},
						"source": schema.StringAttribute{
							MarkdownDescription: "Interface for source address",
							Computed:            true,
						},
						"key": schema.Int64Attribute{
							MarkdownDescription: "Configure peer authentication key",
							Computed:            true,
						},
						"prefer": schema.BoolAttribute{
							MarkdownDescription: "Prefer this peer when possible",
							Computed:            true,
						},
						"version": schema.Int64Attribute{
							MarkdownDescription: "    Configure NTP version",
							Computed:            true,
						},
					},
				},
			},
			"server_vrfs": schema.ListNestedAttribute{
				MarkdownDescription: "VPN Routing/Forwarding Information",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"servers": schema.ListNestedAttribute{
							MarkdownDescription: "",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"ip_address": schema.StringAttribute{
										MarkdownDescription: "Configure ip/ipv6 address/hostname of peer",
										Computed:            true,
									},
									"key": schema.Int64Attribute{
										MarkdownDescription: "Configure peer authentication key",
										Computed:            true,
									},
									"prefer": schema.BoolAttribute{
										MarkdownDescription: "Prefer this peer when possible",
										Computed:            true,
									},
									"version": schema.Int64Attribute{
										MarkdownDescription: "    Configure NTP version",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
			"peers": schema.ListNestedAttribute{
				MarkdownDescription: "",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip_address": schema.StringAttribute{
							MarkdownDescription: "Configure ip/ipv6 address/hostname of peer",
							Computed:            true,
						},
						"source": schema.StringAttribute{
							MarkdownDescription: "Interface for source address",
							Computed:            true,
						},
						"key": schema.Int64Attribute{
							MarkdownDescription: "Configure peer authentication key",
							Computed:            true,
						},
						"prefer": schema.BoolAttribute{
							MarkdownDescription: "Prefer this peer when possible",
							Computed:            true,
						},
						"version": schema.Int64Attribute{
							MarkdownDescription: "    Configure NTP version",
							Computed:            true,
						},
					},
				},
			},
			"peer_vrfs": schema.ListNestedAttribute{
				MarkdownDescription: "VPN Routing/Forwarding Information",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"peers": schema.ListNestedAttribute{
							MarkdownDescription: "",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"ip_address": schema.StringAttribute{
										MarkdownDescription: "Configure ip/ipv6 address/hostname of peer",
										Computed:            true,
									},
									"key": schema.Int64Attribute{
										MarkdownDescription: "Configure peer authentication key",
										Computed:            true,
									},
									"prefer": schema.BoolAttribute{
										MarkdownDescription: "Prefer this peer when possible",
										Computed:            true,
									},
									"version": schema.Int64Attribute{
										MarkdownDescription: "    Configure NTP version",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
			"trusted_keys": schema.ListNestedAttribute{
				MarkdownDescription: "Key numbers for trusted time sources",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"number": schema.Int64Attribute{
							MarkdownDescription: "",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NTPDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.data = req.ProviderData.(*IosxeProviderData)
}

func (d *NTPDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config NTPData

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
		config = NTPData{Device: config.Device}
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
