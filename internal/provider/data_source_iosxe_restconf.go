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

package provider

import (
	"context"
	"fmt"

	"github.com/CiscoDevNet/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-restconf"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &RestconfDataSource{}
	_ datasource.DataSourceWithConfigure = &RestconfDataSource{}
)

func NewRestconfDataSource() datasource.DataSource {
	return &RestconfDataSource{}
}

type RestconfDataSource struct {
	clients map[string]*restconf.Client
}

func (d *RestconfDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_restconf"
}

func (d *RestconfDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can retrieve one or more attributes via RESTCONF.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"path": schema.StringAttribute{
				MarkdownDescription: "A RESTCONF path, e.g. `openconfig-interfaces:interfaces`.",
				Required:            true,
			},
			"attributes": schema.MapAttribute{
				MarkdownDescription: "Map of key-value pairs which represents the attributes and its values.",
				Computed:            true,
				ElementType:         types.StringType,
			},
		},
	}
}

func (d *RestconfDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.clients = req.ProviderData.(map[string]*restconf.Client)
}

func (d *RestconfDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config, state RestconfDataSourceModel

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

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.ValueString()))

	res, err := d.clients[config.Device.ValueString()].GetData(config.Path.ValueString())
	if res.StatusCode == 404 {
		state.Attributes = types.MapValueMust(types.StringType, map[string]attr.Value{})
	} else {
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		}

		state.Path = config.Path
		state.Id = config.Path

		attributes := make(map[string]attr.Value)

		for attr, value := range res.Res.Get(helpers.LastElement(config.Path.ValueString())).Map() {
			// handle empty maps
			if value.IsObject() && len(value.Map()) == 0 {
				attributes[attr] = types.StringValue("")
			} else if value.Raw == "[null]" {
				attributes[attr] = types.StringValue("")
			} else {
				attributes[attr] = types.StringValue(value.String())
			}
		}
		state.Attributes = types.MapValueMust(types.StringType, attributes)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}
