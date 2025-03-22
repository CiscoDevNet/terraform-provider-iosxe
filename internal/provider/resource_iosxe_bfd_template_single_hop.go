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

	"github.com/CiscoDevNet/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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

func NewBFDTemplateSingleHopResource() resource.Resource {
	return &BFDTemplateSingleHopResource{}
}

type BFDTemplateSingleHopResource struct {
	data *IosxeProviderData
}

func (r *BFDTemplateSingleHopResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_bfd_template_single_hop"
}

func (r *BFDTemplateSingleHopResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the BFD Template Single Hop configuration.",

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
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"authentication_md5_keychain": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("keychain name").String,
				Optional:            true,
			},
			"authentication_meticulous_md5_keychain": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("keychain name").String,
				Optional:            true,
			},
			"authentication_meticulous_sha_1_keychain": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("keychain name").String,
				Optional:            true,
			},
			"authentication_sha_1_keychain": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("keychain name").String,
				Optional:            true,
			},
			"interval_milliseconds_min_tx": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Minimum transmit interval capability").AddIntegerRangeDescription(4, 9999).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(4, 9999),
				},
			},
			"interval_milliseconds_min_rx": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Minimum receive interval capability").AddIntegerRangeDescription(4, 9999).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(4, 9999),
				},
			},
			"interval_milliseconds_both": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Minimum transmit and receive interval capability").AddIntegerRangeDescription(4, 9999).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(4, 9999),
				},
			},
			"interval_milliseconds_multiplier": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Multiplier value used to compute holddown").AddIntegerRangeDescription(3, 50).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(3, 50),
				},
			},
			"interval_microseconds_min_rx": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Minimum receive interval capability").AddIntegerRangeDescription(3300, 9999000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(3300, 9999000),
				},
			},
			"interval_microseconds_min_tx": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Minimum transmit interval capability").AddIntegerRangeDescription(3300, 9999000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(3300, 9999000),
				},
			},
			"echo": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Use echo adjunct as bfd detection mechanism").String,
				Optional:            true,
			},
			"dampening_half_time": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Half-life time for the penalty").AddIntegerRangeDescription(1, 30).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 30),
				},
			},
			"dampening_unsuppress_time": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Value to unsuppress a session").AddIntegerRangeDescription(1, 18000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 18000),
				},
			},
			"dampening_suppress_time": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Value to start suppressing a session").AddIntegerRangeDescription(1, 18000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 18000),
				},
			},
			"dampening_max_suppressing_time": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Maximum duration to suppress a session").AddIntegerRangeDescription(1, 420).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 420),
				},
			},
		},
	}
}

func (r *BFDTemplateSingleHopResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.data = req.ProviderData.(*IosxeProviderData)
}

func (r *BFDTemplateSingleHopResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan BFDTemplateSingleHop

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

func (r *BFDTemplateSingleHopResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state BFDTemplateSingleHop

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
			state = BFDTemplateSingleHop{Device: state.Device, Id: state.Id}
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
				state.getIdsFromPath()
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

func (r *BFDTemplateSingleHopResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state BFDTemplateSingleHop

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

func (r *BFDTemplateSingleHopResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state BFDTemplateSingleHop

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

func (r *BFDTemplateSingleHopResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}
