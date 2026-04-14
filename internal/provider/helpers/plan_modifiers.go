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
	// When attribute is not configured, explicitly set to null so that
	// Computed attributes don't remain "unknown" after apply.
	if req.ConfigValue.IsNull() {
		resp.PlanValue = types.StringNull()
		return
	}
	if req.ConfigValue.IsUnknown() {
		return
	}

	// Do nothing on CREATE (no prior state). Terraform Core requires plan == config
	// for Optional+Computed attributes when no prior state exists. The first apply
	// will store the raw config value; subsequent reads will normalize from the device,
	// and the drift-prevention logic below will handle idempotency.
	if req.StateValue.IsNull() || req.StateValue.IsUnknown() {
		return
	}

	// Normalize both values by trimming leading/trailing whitespace from each line.
	// IOS-XE strips this whitespace on the round-trip regardless of protocol.
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
