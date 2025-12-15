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
	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/action/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ action.Action = (*CommitAction)(nil)

func NewCommitAction() action.Action {
	return &CommitAction{}
}

type CommitAction struct {
	data *IosxeProviderData
}

func (r *CommitAction) Metadata(ctx context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_commit"
}

func (r *CommitAction) Schema(ctx context.Context, req action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This action is used to commit the candidate config to the running config (NETCONF only) and optionally save the running config to startup config (both NETCONF and RESTCONF).",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"save_config": schema.BoolAttribute{
				MarkdownDescription: "Save running configuration to startup configuration. Equivalent to 'copy running-config startup-config'. For NETCONF devices, this saves after commit. For RESTCONF devices, this saves the current running configuration (RESTCONF is stateless, no commit needed). Defaults to `false`.",
				Optional:            true,
			},
		},
	}
}

func (r *CommitAction) Configure(_ context.Context, req action.ConfigureRequest, _ *action.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.data = req.ProviderData.(*IosxeProviderData)
}

func (r *CommitAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var device types.String
	var saveConfig types.Bool

	diags := req.Config.GetAttribute(ctx, path.Root("device"), &device)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	diags = req.Config.GetAttribute(ctx, path.Root("save_config"), &saveConfig)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Beginning to commit config")

	d, ok := r.data.Devices[device.ValueString()]
	if !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", device.ValueString()))
		return
	}

	if d.Managed && d.Protocol == "netconf" {
		// Serialize NETCONF operations when reuse disabled (concurrent reads allowed when reuse enabled)
		locked := helpers.AcquireNetconfLock(&d.NetconfOpMutex, d.ReuseConnection, true)
		if locked {
			defer d.NetconfOpMutex.Unlock()
		}
		defer helpers.CloseNetconfConnection(ctx, d.NetconfClient, d.ReuseConnection)

		// Commit candidate to running
		if err := helpers.Commit(ctx, d.NetconfClient); err != nil {
			resp.Diagnostics.AddError("Client Error", err.Error())
			return
		}

		// Optionally save running to startup
		if saveConfig.ValueBool() {
			if err := helpers.SaveConfig(ctx, d.NetconfClient); err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Commit succeeded but save failed: %s", err.Error()))
				return
			}
		}
	} else if d.Managed && d.Protocol == "restconf" {
		// RESTCONF is stateless (no commit needed), but save if requested
		if saveConfig.ValueBool() {
			if err := helpers.SaveConfigRestconf(d.RestconfClient); err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to save config: %s", err))
				return
			}
		}
	}

	tflog.Debug(ctx, "Commit config finished successfully")
}
