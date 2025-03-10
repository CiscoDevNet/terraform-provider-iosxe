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
	_ datasource.DataSource              = &SystemDataSource{}
	_ datasource.DataSourceWithConfigure = &SystemDataSource{}
)

func NewSystemDataSource() datasource.DataSource {
	return &SystemDataSource{}
}

type SystemDataSource struct {
	data *IosxeProviderData
}

func (d *SystemDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system"
}

func (d *SystemDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the System configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"hostname": schema.StringAttribute{
				MarkdownDescription: "Set system's network name",
				Computed:            true,
			},
			"ip_bgp_community_new_format": schema.BoolAttribute{
				MarkdownDescription: "select aa:nn format for BGP community",
				Computed:            true,
			},
			"ip_routing": schema.BoolAttribute{
				MarkdownDescription: "Enable or disable IP routing",
				Computed:            true,
			},
			"ipv6_unicast_routing": schema.BoolAttribute{
				MarkdownDescription: "Enable unicast routing",
				Computed:            true,
			},
			"mtu": schema.Int64Attribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"ip_source_route": schema.BoolAttribute{
				MarkdownDescription: "Process packets with source routing header options",
				Computed:            true,
			},
			"ip_domain_lookup": schema.BoolAttribute{
				MarkdownDescription: "Enable IP Domain Name System hostname translation",
				Computed:            true,
			},
			"ip_domain_name": schema.StringAttribute{
				MarkdownDescription: "Define the default domain name",
				Computed:            true,
			},
			"login_delay": schema.Int64Attribute{
				MarkdownDescription: "Set delay between successive fail login",
				Computed:            true,
			},
			"login_on_failure": schema.BoolAttribute{
				MarkdownDescription: "Set options for failed login attempt",
				Computed:            true,
			},
			"login_on_failure_log": schema.BoolAttribute{
				MarkdownDescription: "Generate syslogs on failure logins",
				Computed:            true,
			},
			"login_on_success": schema.BoolAttribute{
				MarkdownDescription: "Set options for successful login attempt",
				Computed:            true,
			},
			"login_on_success_log": schema.BoolAttribute{
				MarkdownDescription: "Generate syslogs on successful logins",
				Computed:            true,
			},
			"ip_multicast_routing": schema.BoolAttribute{
				MarkdownDescription: "Enable IP multicast forwarding",
				Computed:            true,
			},
			"multicast_routing_switch": schema.BoolAttribute{
				MarkdownDescription: "Enable IP multicast forwarding, some XE devices use this option instead of `multicast_routing`.",
				Computed:            true,
			},
			"ip_multicast_routing_distributed": schema.BoolAttribute{
				MarkdownDescription: "Distributed multicast switching",
				Computed:            true,
			},
			"multicast_routing_vrfs": schema.ListNestedAttribute{
				MarkdownDescription: "Select VPN Routing/Forwarding instance",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"vrf": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"distributed": schema.BoolAttribute{
							MarkdownDescription: "Distributed multicast switching",
							Computed:            true,
						},
					},
				},
			},
			"ip_http_access_class": schema.Int64Attribute{
				MarkdownDescription: "Restrict http server access by access-class",
				Computed:            true,
			},
			"ip_http_authentication_aaa": schema.BoolAttribute{
				MarkdownDescription: "Use AAA access control methods",
				Computed:            true,
			},
			"ip_http_authentication_aaa_exec_authorization": schema.StringAttribute{
				MarkdownDescription: "Set method list for exec authorization",
				Computed:            true,
			},
			"ip_http_authentication_aaa_login_authentication": schema.StringAttribute{
				MarkdownDescription: "Set method list for login authentication",
				Computed:            true,
			},
			"ip_http_authentication_aaa_command_authorization": schema.ListNestedAttribute{
				MarkdownDescription: "Set method list for command authorization",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"level": schema.Int64Attribute{
							MarkdownDescription: "Enable level",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "Use an authorization list with this name",
							Computed:            true,
						},
					},
				},
			},
			"ip_http_authentication_local": schema.BoolAttribute{
				MarkdownDescription: "Use local username and passwords",
				Computed:            true,
			},
			"ip_http_server": schema.BoolAttribute{
				MarkdownDescription: "Enable http server",
				Computed:            true,
			},
			"ip_http_secure_server": schema.BoolAttribute{
				MarkdownDescription: "Enable HTTP secure server",
				Computed:            true,
			},
			"ip_http_secure_trustpoint": schema.StringAttribute{
				MarkdownDescription: "Set http secure server certificate trustpoint",
				Computed:            true,
			},
			"ip_http_tls_version": schema.StringAttribute{
				MarkdownDescription: "Set TLS version for HTTP secure server",
				Computed:            true,
			},
			"ip_http_client_secure_trustpoint": schema.StringAttribute{
				MarkdownDescription: "Set http client certificate secure trustpoint",
				Computed:            true,
			},
			"ip_http_client_source_interface": schema.StringAttribute{
				MarkdownDescription: "Specify interface for source address in all HTTP(S) client connections",
				Computed:            true,
			},
		},
	}
}

func (d *SystemDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.data = req.ProviderData.(*IosxeProviderData)
}

func (d *SystemDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config SystemData

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
		config = SystemData{Device: config.Device}
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
