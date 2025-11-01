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
	"strings"

	"github.com/CiscoDevNet/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-netconf"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &SaveConfigResource{}

func NewSaveConfigResource() resource.Resource {
	return &SaveConfigResource{}
}

type SaveConfigResource struct {
	data *IosxeProviderData
}

func (r *SaveConfigResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_save_config"
}

func (r *SaveConfigResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resources is used to save the config which is the equivalent of doing a `copy running-config startup-config` in the device CLI.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"save": schema.BoolAttribute{
				MarkdownDescription: "This attribute is only used internally.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
		},
	}
}

func (r *SaveConfigResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.data = req.ProviderData.(*IosxeProviderData)
}

func (r *SaveConfigResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var device types.String
	var save types.Bool

	diags := req.Plan.GetAttribute(ctx, path.Root("device"), &device)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	diags = req.Plan.GetAttribute(ctx, path.Root("save"), &save)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Beginning to save config")

	d, ok := r.data.Devices[device.ValueString()]
	if !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", device.ValueString()))
		return
	}

	if d.Managed {
		if d.Protocol == "restconf" {
			request := d.RestconfClient.NewReq("POST", "/operations/cisco-ia:save-config/", strings.NewReader(""))
			_, err := d.RestconfClient.Do(request)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to save config, got error: %s", err))
				return
			}
		} else {
			body := netconf.Body{}
			body = helpers.SetFromXPath(body, "/cisco-ia:save-config", "")
			body = body.SetAttr("save-config", "xmlns", "http://cisco.com/yang/cisco-ia")

			if _, err := d.NetconfClient.RPC(ctx, body.Res()); err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to save config, got error: %s", err))
				return
			}
		}
	}

	tflog.Debug(ctx, "Save config finished successfully")

	diags = resp.State.SetAttribute(ctx, path.Root("device"), device)
	resp.Diagnostics.Append(diags...)
	diags = resp.State.SetAttribute(ctx, path.Root("save"), save)
	resp.Diagnostics.Append(diags...)
}

func (r *SaveConfigResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	resp.State.SetAttribute(ctx, path.Root("save"), false)
}

func (r *SaveConfigResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var device types.String
	var save types.Bool

	diags := req.Plan.GetAttribute(ctx, path.Root("device"), &device)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	diags = req.Plan.GetAttribute(ctx, path.Root("save"), &save)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Beginning to save config")

	d, ok := r.data.Devices[device.ValueString()]
	if !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", device.ValueString()))
		return
	}

	if d.Managed {
		if d.Protocol == "restconf" {
			request := d.RestconfClient.NewReq("POST", "/operations/cisco-ia:save-config/", strings.NewReader(""))
			_, err := d.RestconfClient.Do(request)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to save config, got error: %s", err))
				return
			}
		} else {
			body := netconf.Body{}
			body = helpers.SetFromXPath(body, "/save-config", "")
			body = body.SetAttr("save-config", "xmlns", "http://cisco.com/yang/cisco-ia")

			if _, err := d.NetconfClient.RPC(ctx, body.Res()); err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to save config, got error: %s", err))
				return
			}
		}
	}

	tflog.Debug(ctx, "Save config finished successfully")

	diags = resp.State.SetAttribute(ctx, path.Root("device"), device)
	resp.Diagnostics.Append(diags...)
	diags = resp.State.SetAttribute(ctx, path.Root("save"), save)
	resp.Diagnostics.Append(diags...)
}

func (r *SaveConfigResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
