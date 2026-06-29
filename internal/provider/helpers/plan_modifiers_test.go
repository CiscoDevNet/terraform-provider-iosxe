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
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestNetconfTrailingWhitespaceTrimModifier_CreateNoOp(t *testing.T) {
	// On CREATE (no state), the modifier must be a no-op.
	// Terraform Core requires plan == config when no prior state exists.
	modifier := UseNetconfTrailingWhitespaceTrimming()
	configVal := "  Authorized access only  "
	req := planmodifier.StringRequest{
		ConfigValue: types.StringValue(configVal),
		StateValue:  types.StringNull(),
		PlanValue:   types.StringValue(configVal),
	}
	resp := &planmodifier.StringResponse{
		PlanValue: req.PlanValue,
	}

	modifier.PlanModifyString(context.Background(), req, resp)

	if resp.PlanValue.ValueString() != configVal {
		t.Errorf("CREATE should be no-op: got %q, want %q", resp.PlanValue.ValueString(), configVal)
	}
}

func TestNetconfTrailingWhitespaceTrimModifier_UpdateNormalizedMatch(t *testing.T) {
	// On UPDATE, when normalized config matches normalized state, use state to prevent drift.
	modifier := UseNetconfTrailingWhitespaceTrimming()
	req := planmodifier.StringRequest{
		ConfigValue: types.StringValue("  Banner text  "),
		StateValue:  types.StringValue("Banner text"),
		PlanValue:   types.StringValue("  Banner text  "),
	}
	resp := &planmodifier.StringResponse{
		PlanValue: req.PlanValue,
	}

	modifier.PlanModifyString(context.Background(), req, resp)

	if resp.PlanValue.ValueString() != "Banner text" {
		t.Errorf("UPDATE normalized match: got %q, want %q", resp.PlanValue.ValueString(), "Banner text")
	}
}

func TestNetconfTrailingWhitespaceTrimModifier_UpdateMultilineMatch(t *testing.T) {
	modifier := UseNetconfTrailingWhitespaceTrimming()
	req := planmodifier.StringRequest{
		ConfigValue: types.StringValue("  Welcome to Router1 \n  Network Device  \n"),
		StateValue:  types.StringValue("Welcome to Router1\nNetwork Device"),
		PlanValue:   types.StringValue("  Welcome to Router1 \n  Network Device  \n"),
	}
	resp := &planmodifier.StringResponse{
		PlanValue: req.PlanValue,
	}

	modifier.PlanModifyString(context.Background(), req, resp)

	expected := "Welcome to Router1\nNetwork Device"
	if resp.PlanValue.ValueString() != expected {
		t.Errorf("UPDATE multiline match: got %q, want %q", resp.PlanValue.ValueString(), expected)
	}
}

func TestNetconfTrailingWhitespaceTrimModifier_UpdateValuesDiffer(t *testing.T) {
	// When config and state are genuinely different (drift or intentional change),
	// the plan modifier must return the raw config value to satisfy Terraform's
	// plan validation (planned value must == config value when changing).
	modifier := UseNetconfTrailingWhitespaceTrimming()
	req := planmodifier.StringRequest{
		ConfigValue: types.StringValue("New banner text  "),
		StateValue:  types.StringValue("Old banner text"),
		PlanValue:   types.StringValue("New banner text  "),
	}
	resp := &planmodifier.StringResponse{
		PlanValue: req.PlanValue,
	}

	modifier.PlanModifyString(context.Background(), req, resp)

	// Must return raw config value, not normalized, to pass Terraform plan validation
	expected := "New banner text  "
	if resp.PlanValue.ValueString() != expected {
		t.Errorf("UPDATE values differ: got %q, want %q", resp.PlanValue.ValueString(), expected)
	}
}

func TestNetconfTrailingWhitespaceTrimModifier_ConfigNull(t *testing.T) {
	modifier := UseNetconfTrailingWhitespaceTrimming()
	req := planmodifier.StringRequest{
		ConfigValue: types.StringNull(),
		StateValue:  types.StringValue("Banner text"),
		PlanValue:   types.StringNull(),
	}
	resp := &planmodifier.StringResponse{
		PlanValue: req.PlanValue,
	}

	modifier.PlanModifyString(context.Background(), req, resp)

	if !resp.PlanValue.IsNull() {
		t.Errorf("Config null: expected plan to remain null, got %q", resp.PlanValue.ValueString())
	}
}

func TestNetconfTrailingWhitespaceTrimModifier_ConfigUnknown(t *testing.T) {
	modifier := UseNetconfTrailingWhitespaceTrimming()
	req := planmodifier.StringRequest{
		ConfigValue: types.StringUnknown(),
		StateValue:  types.StringValue("Banner text"),
		PlanValue:   types.StringUnknown(),
	}
	resp := &planmodifier.StringResponse{
		PlanValue: req.PlanValue,
	}

	modifier.PlanModifyString(context.Background(), req, resp)

	if !resp.PlanValue.IsUnknown() {
		t.Errorf("Config unknown: expected plan to remain unknown, got %q", resp.PlanValue.ValueString())
	}
}

func TestIPv6NormalizationModifier_SameAddressDifferentCase(t *testing.T) {
	ctx := context.Background()
	modifier := UseIPv6Normalization()

	req := planmodifier.StringRequest{
		ConfigValue: types.StringValue("2001:DB8::/32"),
		StateValue:  types.StringValue("2001:db8::/32"),
	}
	resp := &planmodifier.StringResponse{
		PlanValue: req.ConfigValue,
	}

	modifier.PlanModifyString(ctx, req, resp)

	if resp.PlanValue.ValueString() != "2001:db8::/32" {
		t.Errorf("expected state value %q, got %q", "2001:db8::/32", resp.PlanValue.ValueString())
	}
}

func TestIPv6NormalizationModifier_GenuineChange(t *testing.T) {
	ctx := context.Background()
	modifier := UseIPv6Normalization()

	req := planmodifier.StringRequest{
		ConfigValue: types.StringValue("2001:db8:cafe::/48"),
		StateValue:  types.StringValue("2001:db8::/32"),
	}
	resp := &planmodifier.StringResponse{
		PlanValue: req.ConfigValue,
	}

	modifier.PlanModifyString(ctx, req, resp)

	if resp.PlanValue.ValueString() != "2001:db8:cafe::/48" {
		t.Errorf("expected config value %q, got %q", "2001:db8:cafe::/48", resp.PlanValue.ValueString())
	}
}

func TestIPv6NormalizationModifier_NullConfig(t *testing.T) {
	ctx := context.Background()
	modifier := UseIPv6Normalization()

	req := planmodifier.StringRequest{
		ConfigValue: types.StringNull(),
		StateValue:  types.StringValue("2001:db8::/32"),
	}
	resp := &planmodifier.StringResponse{
		PlanValue: req.ConfigValue,
	}

	modifier.PlanModifyString(ctx, req, resp)

	if !resp.PlanValue.IsNull() {
		t.Errorf("expected null plan value, got %q", resp.PlanValue.ValueString())
	}
}

func TestIPv6NormalizationModifier_Create(t *testing.T) {
	ctx := context.Background()
	modifier := UseIPv6Normalization()

	req := planmodifier.StringRequest{
		ConfigValue: types.StringValue("2001:DB8::/32"),
		StateValue:  types.StringNull(),
	}
	resp := &planmodifier.StringResponse{
		PlanValue: req.ConfigValue,
	}

	modifier.PlanModifyString(ctx, req, resp)

	if resp.PlanValue.ValueString() != "2001:DB8::/32" {
		t.Errorf("expected config value %q on create, got %q", "2001:DB8::/32", resp.PlanValue.ValueString())
	}
}

func TestIPv6NormalizationModifier_ExpandedVsCompressed(t *testing.T) {
	ctx := context.Background()
	modifier := UseIPv6Normalization()

	req := planmodifier.StringRequest{
		ConfigValue: types.StringValue("2001:0db8:0000:0000:0000:0000:0000:0001"),
		StateValue:  types.StringValue("2001:db8::1"),
	}
	resp := &planmodifier.StringResponse{
		PlanValue: req.ConfigValue,
	}

	modifier.PlanModifyString(ctx, req, resp)

	if resp.PlanValue.ValueString() != "2001:db8::1" {
		t.Errorf("expected state value %q, got %q", "2001:db8::1", resp.PlanValue.ValueString())
	}
}
