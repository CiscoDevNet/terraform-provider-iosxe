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
var _ action.Action = (*DiscardAction)(nil)

func NewDiscardAction() action.Action {
	return &DiscardAction{}
}

type DiscardAction struct {
	data *IosxeProviderData
}

func (r *DiscardAction) Metadata(ctx context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_discard"
}

func (r *DiscardAction) Schema(ctx context.Context, req action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This action is used to discard the candidate config, reverting it to the running config. The candidate datastore is shared and persists on the device, so changes left staged by a failed commit are re-attempted by every subsequent commit until they are discarded.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
		},
	}
}

func (r *DiscardAction) Configure(_ context.Context, req action.ConfigureRequest, _ *action.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.data = req.ProviderData.(*IosxeProviderData)
}

func (r *DiscardAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var device types.String

	diags := req.Config.GetAttribute(ctx, path.Root("device"), &device)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Beginning to discard candidate config")

	d, ok := r.data.Devices[device.ValueString()]
	if !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", device.ValueString()))
		return
	}

	if d.Managed {
		// Serialize NETCONF operations when reuse disabled (concurrent reads allowed when reuse enabled)
		locked := helpers.AcquireNetconfLock(&d.NetconfOpMutex, d.ReuseConnection, true)
		if locked {
			defer d.NetconfOpMutex.Unlock()
		}
		defer helpers.CloseNetconfConnection(ctx, d.NetconfClient, d.ReuseConnection)

		// Revert candidate to running
		if err := helpers.Discard(ctx, d.NetconfClient); err != nil {
			resp.Diagnostics.AddError("Client Error", err.Error())
			return
		}
	}

	tflog.Debug(ctx, "Discard candidate config finished successfully")
}
