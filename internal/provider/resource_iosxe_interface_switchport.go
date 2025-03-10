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

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"regexp"

	"github.com/CiscoDevNet/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-restconf"
)

func NewInterfaceSwitchportResource() resource.Resource {
	return &InterfaceSwitchportResource{}
}

type InterfaceSwitchportResource struct {
	data *IosxeProviderData
}

func (r *InterfaceSwitchportResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_switchport"
}

func (r *InterfaceSwitchportResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the Interface Switchport configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the object.",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"delete_mode": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure behavior when deleting/destroying the resource. Either delete the entire object (YANG container) being managed, or only delete the individual resource attributes configured explicitly and leave everything else as-is. Default value is `all`.").AddStringEnumDescription("all", "attributes").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("all", "attributes"),
				},
			},
			"type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interface type").AddStringEnumDescription("GigabitEthernet", "TwoGigabitEthernet", "FiveGigabitEthernet", "TenGigabitEthernet", "TwentyFiveGigE", "FortyGigabitEthernet", "HundredGigE", "TwoHundredGigE", "FourHundredGigE", "Port-channel").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("GigabitEthernet", "TwoGigabitEthernet", "FiveGigabitEthernet", "TenGigabitEthernet", "TwentyFiveGigE", "FortyGigabitEthernet", "HundredGigE", "TwoHundredGigE", "FourHundredGigE", "Port-channel"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`(0|[1-9][0-9]*)(/(0|[1-9][0-9]*))*(\.[0-9]*)?`), ""),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"mode_access": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set trunking mode to ACCESS unconditionally").String,
				Optional:            true,
			},
			"mode_dot1q_tunnel": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("set trunking mode to TUNNEL unconditionally").String,
				Optional:            true,
			},
			"mode_private_vlan_trunk": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the mode to private-vlan trunk").String,
				Optional:            true,
			},
			"mode_private_vlan_host": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the mode to private-vlan host").String,
				Optional:            true,
			},
			"mode_private_vlan_promiscuous": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the mode to private-vlan promiscuous").String,
				Optional:            true,
			},
			"mode_trunk": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set trunking mode to TRUNK unconditionally").String,
				Optional:            true,
			},
			"nonegotiate": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Device will not engage in negotiation protocol on this interface").String,
				Optional:            true,
			},
			"access_vlan": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
			},
			"trunk_allowed_vlans": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
			},
			"trunk_allowed_vlans_none": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("no VLANs").String,
				Optional:            true,
			},
			"trunk_native_vlan_tag": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
			},
			"trunk_native_vlan": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddIntegerRangeDescription(1, 4094).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 4094),
				},
			},
			"host": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set port host").String,
				Optional:            true,
			},
		},
	}
}

func (r *InterfaceSwitchportResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.data = req.ProviderData.(*IosxeProviderData)
}

func (r *InterfaceSwitchportResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceSwitchport

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.getPath()))

	device, ok := r.data.Devices[plan.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", plan.Device.ValueString()))
		return
	}

	if device.Managed {
		// Create object
		body := plan.toBody(ctx)

		emptyLeafsDelete := plan.getEmptyLeafsDelete(ctx)
		tflog.Debug(ctx, fmt.Sprintf("List of empty leafs to delete: %+v", emptyLeafsDelete))

		if YangPatch {
			edits := []restconf.YangPatchEdit{restconf.NewYangPatchEdit("merge", plan.getPath(), restconf.Body{Str: body})}
			for _, i := range emptyLeafsDelete {
				edits = append(edits, restconf.NewYangPatchEdit("remove", i, restconf.Body{}))
			}
			_, err := device.Client.YangPatchData("", "1", "", edits)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object, got error: %s", err))
				return
			}
		} else {
			res, err := device.Client.PatchData(plan.getPathShort(), body)
			if len(res.Errors.Error) > 0 && res.Errors.Error[0].ErrorMessage == "patch to a nonexistent resource" {
				_, err = device.Client.PutData(plan.getPath(), body)
			}
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PATCH), got error: %s", err))
				return
			}
			for _, i := range emptyLeafsDelete {
				res, err := device.Client.DeleteData(i)
				if err != nil && res.StatusCode != 404 {
					resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
					return
				}
			}
		}
	}

	plan.Id = types.StringValue(plan.getPath())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.getPath()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *InterfaceSwitchportResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceSwitchport

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.ValueString()))

	device, ok := r.data.Devices[state.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", state.Device.ValueString()))
		return
	}

	if device.Managed {
		res, err := device.Client.GetData(state.Id.ValueString())
		if res.StatusCode == 404 {
			state = InterfaceSwitchport{Device: state.Device, Id: state.Id}
		} else {
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
				return
			}

			imp, diags := helpers.IsFlagImporting(ctx, req)
			if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
				return
			}

			// After `terraform import` we switch to a full read.
			if imp {
				state.fromBody(ctx, res.Res)
			} else {
				state.updateFromBody(ctx, res.Res)
			}
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *InterfaceSwitchportResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceSwitchport

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read state
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	device, ok := r.data.Devices[plan.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", plan.Device.ValueString()))
		return
	}

	if device.Managed {
		body := plan.toBody(ctx)

		deletedItems := plan.getDeletedItems(ctx, state)
		tflog.Debug(ctx, fmt.Sprintf("Removed items to delete: %+v", deletedItems))

		emptyLeafsDelete := plan.getEmptyLeafsDelete(ctx)
		tflog.Debug(ctx, fmt.Sprintf("List of empty leafs to delete: %+v", emptyLeafsDelete))

		if YangPatch {
			edits := []restconf.YangPatchEdit{restconf.NewYangPatchEdit("merge", plan.getPath(), restconf.Body{Str: body})}
			for _, i := range deletedItems {
				edits = append(edits, restconf.NewYangPatchEdit("remove", i, restconf.Body{}))
			}
			for _, i := range emptyLeafsDelete {
				edits = append(edits, restconf.NewYangPatchEdit("remove", i, restconf.Body{}))
			}
			_, err := device.Client.YangPatchData("", "1", "", edits)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object, got error: %s", err))
				return
			}
		} else {
			res, err := device.Client.PatchData(plan.getPathShort(), body)
			if len(res.Errors.Error) > 0 && res.Errors.Error[0].ErrorMessage == "patch to a nonexistent resource" {
				_, err = device.Client.PutData(plan.getPath(), body)
			}
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PATCH), got error: %s", err))
				return
			}
			for _, i := range deletedItems {
				res, err := device.Client.DeleteData(i)
				if err != nil && res.StatusCode != 404 {
					resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
					return
				}
			}
			for _, i := range emptyLeafsDelete {
				res, err := device.Client.DeleteData(i)
				if err != nil && res.StatusCode != 404 {
					resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
					return
				}
			}
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *InterfaceSwitchportResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceSwitchport

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	device, ok := r.data.Devices[state.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", state.Device.ValueString()))
		return
	}

	if device.Managed {
		deleteMode := "all"
		if state.DeleteMode.ValueString() == "all" {
			deleteMode = "all"
		} else if state.DeleteMode.ValueString() == "attributes" {
			deleteMode = "attributes"
		}

		if deleteMode == "all" {
			res, err := device.Client.DeleteData(state.Id.ValueString())
			if err != nil && res.StatusCode != 404 && res.StatusCode != 400 {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
				return
			}
		} else {
			deletePaths := state.getDeletePaths(ctx)
			tflog.Debug(ctx, fmt.Sprintf("Paths to delete: %+v", deletePaths))

			if YangPatch {
				edits := []restconf.YangPatchEdit{}
				for _, i := range deletePaths {
					edits = append(edits, restconf.NewYangPatchEdit("remove", i, restconf.Body{}))
				}
				_, err := device.Client.YangPatchData("", "1", "", edits)
				if err != nil {
					resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
					return
				}
			} else {
				for _, i := range deletePaths {
					res, err := device.Client.DeleteData(i)
					if err != nil && res.StatusCode != 404 {
						resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
						return
					}
				}
			}
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *InterfaceSwitchportResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}
