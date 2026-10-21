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

	"github.com/hashicorp/terraform-plugin-framework-validators/resourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// IOS-XE refuses a VRRPv3 group without an address family, and the provider treats each family's
// primary address as the anchor of that family: removing it removes the whole family, so settings of
// a family whose primary address is not configured would be written once and then removed again.
var _ resource.ResourceWithConfigValidators = &InterfaceVRRPResource{}

var interfaceVRRPFamilyAttributes = map[string][]string{
	"ipv4": {"secondary_addresses", "priority", "preempt_delay_minimum", "timers_advertise", "description", "tracks", "shutdown"},
	"ipv6": {"prefixes", "priority", "preempt_delay_minimum", "timers_advertise", "description", "tracks", "shutdown"},
}

func (r *InterfaceVRRPResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{
		resourcevalidator.AtLeastOneOf(path.MatchRoot("ipv4_primary_address"), path.MatchRoot("ipv6_primary_address")),
		interfaceVRRPFamilyValidator{},
	}
}

type interfaceVRRPFamilyValidator struct{}

func (v interfaceVRRPFamilyValidator) Description(_ context.Context) string {
	return "Address family settings require the primary address of that family."
}

func (v interfaceVRRPFamilyValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v interfaceVRRPFamilyValidator) ValidateResource(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	for family, attrs := range interfaceVRRPFamilyAttributes {
		var primary attr.Value
		resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root(family+"_primary_address"), &primary)...)
		if primary == nil || !primary.IsNull() {
			continue
		}
		for _, name := range attrs {
			var value attr.Value
			resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root(family+"_"+name), &value)...)
			if value != nil && !value.IsNull() {
				resp.Diagnostics.AddAttributeError(path.Root(family+"_"+name), "Missing primary address",
					fmt.Sprintf("%s_%s configures the %s address family, which also requires %s_primary_address.", family, name, family, family))
			}
		}
	}
}
