// Copyright © 2025 Cisco Systems, Inc. and its affiliates.
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

package helpers

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// netconfTrailingWhitespaceTrimModifier implements the plan modifier for trimming trailing whitespace
type netconfTrailingWhitespaceTrimModifier struct{}

// UseNetconfTrailingWhitespaceTrimming returns a plan modifier that prevents drift caused by
// trailing whitespace differences in NETCONF responses.
//
// IOS-XE NETCONF responses often include trailing whitespace on each line of multi-line
// strings (like banners). This whitespace differs from the original configuration,
// causing Terraform to detect false drift on subsequent plan/apply cycles.
//
// This plan modifier normalizes both the config value and state value by trimming
// trailing whitespace from each line before comparison. If the normalized values
// match, the state value is preserved to prevent unnecessary drift.
//
// Usage in schema:
//
//	"banner": schema.StringAttribute{
//	    PlanModifiers: []planmodifier.String{
//	        helpers.UseNetconfTrailingWhitespaceTrimming(),
//	    },
//	},
func UseNetconfTrailingWhitespaceTrimming() planmodifier.String {
	return netconfTrailingWhitespaceTrimModifier{}
}

// Description returns a human-readable description of the plan modifier.
func (m netconfTrailingWhitespaceTrimModifier) Description(_ context.Context) string {
	return "Normalizes trailing whitespace differences in NETCONF responses to prevent drift."
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (m netconfTrailingWhitespaceTrimModifier) MarkdownDescription(_ context.Context) string {
	return "Normalizes trailing whitespace differences in NETCONF responses to prevent drift."
}

// PlanModifyString implements the plan modification logic.
func (m netconfTrailingWhitespaceTrimModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	// Do nothing if there is no state value (resource is being created)
	if req.StateValue.IsNull() || req.StateValue.IsUnknown() {
		return
	}

	// Do nothing if there is no config value (attribute not configured)
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	// Normalize both values by trimming trailing whitespace from each line
	configNormalized := TrimNetconfTrailingWhitespace(req.ConfigValue.ValueString())
	stateNormalized := TrimNetconfTrailingWhitespace(req.StateValue.ValueString())

	// If normalized values match, use the state value to prevent drift
	if configNormalized == stateNormalized {
		resp.PlanValue = req.StateValue
		return
	}

	// Values are different - use the normalized config value
	resp.PlanValue = types.StringValue(configNormalized)
}
