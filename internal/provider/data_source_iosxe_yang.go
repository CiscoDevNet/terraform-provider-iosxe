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
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &YangDataSource{}
	_ datasource.DataSourceWithConfigure = &YangDataSource{}
)

func NewYangDataSource() datasource.DataSource {
	return &YangDataSource{}
}

type YangDataSource struct {
	data *IosxeProviderData
}

func (d *YangDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_yang"
}

func (d *YangDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can retrieve one or more attributes via YANG paths.",

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
				MarkdownDescription: "A YANG path, e.g. `openconfig-interfaces:interfaces`.",
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

func (d *YangDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.data = req.ProviderData.(*IosxeProviderData)
}

func (d *YangDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config, state YangDataSourceModel

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Path.ValueString()))

	device, ok := d.data.Devices[config.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", config.Device.ValueString()))
		return
	}

	if device.Protocol == "restconf" {
		res, err := device.RestconfClient.GetData(config.getPath())
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

			for attr, value := range res.Res.Get(helpers.LastElement(config.getPath())).Map() {
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
	} else {
		// Serialize NETCONF operations (all ops when reuse disabled, reads concurrent when reuse enabled)
		locked := helpers.AcquireNetconfLock(&device.NetconfOpMutex, device.ReuseConnection, false)
		if locked {
			defer device.NetconfOpMutex.Unlock()
		}
		defer helpers.CloseNetconfConnection(ctx, device.NetconfClient, device.ReuseConnection)

		res, err := device.NetconfClient.GetConfig(ctx, "running", helpers.GetXpathFilter(config.Path.ValueString()))
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		}

		state.Path = config.Path
		state.Id = config.Path

		// Get the result at the specified path
		result := helpers.GetFromXPath(res.Res, config.Path.ValueString())

		if !result.Exists() {
			state.Attributes = types.MapValueMust(types.StringType, map[string]attr.Value{})
		} else {
			attributes := make(map[string]attr.Value)

			// If the result is an array, use the first element
			if result.IsArray() {
				resultArray := result.Array()
				if len(resultArray) > 0 {
					result = resultArray[0]
				}
			}

			// Parse the raw XML to get child elements
			// The result.Raw contains the XML content, we need to parse it to get attributes
			// For simplicity, we'll just return the raw content as a single attribute if needed
			// In a real implementation, you would parse child elements here

			// For now, store the entire content as a single value if it's a simple type
			if result.String() != "" {
				attributes["value"] = types.StringValue(result.String())
			}

			state.Attributes = types.MapValueMust(types.StringType, attributes)
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Path.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}
