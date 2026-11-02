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

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Hand-written, not generated: the generator has no way to express that two
// attributes of the same list item are mutually exclusive.

var _ resource.ResourceWithValidateConfig = &DHCPResource{}

// ValidateConfig rejects pools that set both client_identifier and
// hardware_address. They are cases of the same YANG choice, so the device
// silently keeps only one of them and every later plan shows a diff.
func (r *DHCPResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var pools types.List
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("pools"), &pools)...)
	if resp.Diagnostics.HasError() || pools.IsNull() || pools.IsUnknown() {
		return
	}

	var items []DHCPPools
	resp.Diagnostics.Append(pools.ElementsAs(ctx, &items, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	for i, p := range items {
		if dhcpIsSet(p.ClientIdentifier) && dhcpIsSet(p.HardwareAddress) {
			resp.Diagnostics.AddAttributeError(
				path.Root("pools").AtListIndex(i).AtName("client_identifier"),
				"Conflicting DHCP pool attributes",
				"Set either client_identifier or hardware_address on a pool, not both. "+
					"IOS XE only keeps one of them.",
			)
		}
	}
}

// dhcpIsSet reports whether a value is present in the config. Unknown values are
// treated as set, since they will be at apply time.
func dhcpIsSet(v types.String) bool {
	return !v.IsNull()
}
