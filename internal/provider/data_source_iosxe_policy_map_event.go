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
	_ datasource.DataSource              = &PolicyMapEventDataSource{}
	_ datasource.DataSourceWithConfigure = &PolicyMapEventDataSource{}
)

func NewPolicyMapEventDataSource() datasource.DataSource {
	return &PolicyMapEventDataSource{}
}

type PolicyMapEventDataSource struct {
	data *IosxeProviderData
}

func (d *PolicyMapEventDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policy_map_event"
}

func (d *PolicyMapEventDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Policy Map Event configuration.",

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
			"event_type": schema.StringAttribute{
				MarkdownDescription: "",
				Required:            true,
			},
			"match_type": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"class_numbers": schema.ListNestedAttribute{
				MarkdownDescription: "class number, 1 for 1st class, 2 for 2nd...",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"number": schema.Int64Attribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"class": schema.StringAttribute{
							MarkdownDescription: "The class type this control policy-map triggers upon",
							Computed:            true,
						},
						"execution_type": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"action_numbers": schema.ListNestedAttribute{
							MarkdownDescription: "action number, 1 for 1st class, 2 for 2nd...",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"number": schema.Int64Attribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"pause_reauthentication": schema.BoolAttribute{
										MarkdownDescription: "pause reauthentication",
										Computed:            true,
									},
									"authorize": schema.BoolAttribute{
										MarkdownDescription: "authorize session",
										Computed:            true,
									},
									"terminate_config": schema.StringAttribute{
										MarkdownDescription: "terminate auth method",
										Computed:            true,
									},
									"activate_service_template_config_service_template": schema.StringAttribute{
										MarkdownDescription: "activate service template",
										Computed:            true,
									},
									"activate_service_template_config_aaa_list": schema.StringAttribute{
										MarkdownDescription: "Named Method List",
										Computed:            true,
									},
									"activate_service_template_config_precedence": schema.Int64Attribute{
										MarkdownDescription: "Template precedence",
										Computed:            true,
									},
									"activate_service_template_config_replace_all": schema.BoolAttribute{
										MarkdownDescription: "Replace all existing authorization data and services",
										Computed:            true,
									},
									"activate_interface_template": schema.StringAttribute{
										MarkdownDescription: "activate interface template",
										Computed:            true,
									},
									"activate_policy_type_control_subscriber": schema.StringAttribute{
										MarkdownDescription: "policy type control subscriber",
										Computed:            true,
									},
									"deactivate_interface_template": schema.StringAttribute{
										MarkdownDescription: "activate interface template",
										Computed:            true,
									},
									"deactivate_service_template": schema.StringAttribute{
										MarkdownDescription: "activate service template",
										Computed:            true,
									},
									"deactivate_policy_type_control_subscriber": schema.StringAttribute{
										MarkdownDescription: "policy type control subscriber",
										Computed:            true,
									},
									"authenticate_using_method": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"authenticate_using_retries": schema.Int64Attribute{
										MarkdownDescription: "Number of times to retry failed authentications",
										Computed:            true,
									},
									"authenticate_using_retry_time": schema.Int64Attribute{
										MarkdownDescription: "Time interval between retries",
										Computed:            true,
									},
									"authenticate_using_priority": schema.Int64Attribute{
										MarkdownDescription: "Method priority",
										Computed:            true,
									},
									"authenticate_using_aaa_authc_list": schema.StringAttribute{
										MarkdownDescription: "Specify authentication method list",
										Computed:            true,
									},
									"authenticate_using_aaa_authz_list": schema.StringAttribute{
										MarkdownDescription: "Specify authorization method list",
										Computed:            true,
									},
									"authenticate_using_both": schema.BoolAttribute{
										MarkdownDescription: "Enabling Dot1x Authenticator & Supplicant",
										Computed:            true,
									},
									"authenticate_using_parameter_map": schema.StringAttribute{
										MarkdownDescription: "Specify parameter map name",
										Computed:            true,
									},
									"replace": schema.BoolAttribute{
										MarkdownDescription: "clear existing session and create session for violating host",
										Computed:            true,
									},
									"restrict": schema.BoolAttribute{
										MarkdownDescription: "drop violating packets and generate a syslog",
										Computed:            true,
									},
									"clear_session": schema.BoolAttribute{
										MarkdownDescription: "clears an active session",
										Computed:            true,
									},
									"clear_authenticated_data_hosts_on_port": schema.BoolAttribute{
										MarkdownDescription: "clears authenticated data hosts on the port",
										Computed:            true,
									},
									"protect": schema.BoolAttribute{
										MarkdownDescription: "silently drop violating packets",
										Computed:            true,
									},
									"err_disable": schema.BoolAttribute{
										MarkdownDescription: "temporarily disable port",
										Computed:            true,
									},
									"resume_reauthentication": schema.BoolAttribute{
										MarkdownDescription: "resume reauthentication",
										Computed:            true,
									},
									"authentication_restart": schema.Int64Attribute{
										MarkdownDescription: "restarts the auth sequence after the specified number of sec",
										Computed:            true,
									},
									"set_domain": schema.StringAttribute{
										MarkdownDescription: "set domain",
										Computed:            true,
									},
									"unauthorize": schema.BoolAttribute{
										MarkdownDescription: "unauthorize session",
										Computed:            true,
									},
									"notify": schema.BoolAttribute{
										MarkdownDescription: "notifies the session attributes",
										Computed:            true,
									},
									"set_timer_name": schema.StringAttribute{
										MarkdownDescription: "timer name",
										Computed:            true,
									},
									"set_timer_value": schema.Int64Attribute{
										MarkdownDescription: "Enter a value between 1 and 65535",
										Computed:            true,
									},
									"map_attribute_to_service_table": schema.StringAttribute{
										MarkdownDescription: "map identity-update attribute to a auto-conf templates",
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

func (d *PolicyMapEventDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.data = req.ProviderData.(*IosxeProviderData)
}

func (d *PolicyMapEventDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config PolicyMapEventData

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
		config = PolicyMapEventData{Device: config.Device}
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
