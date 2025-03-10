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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-restconf"
)

func NewBGPIPv4UnicastVRFNeighborResource() resource.Resource {
	return &BGPIPv4UnicastVRFNeighborResource{}
}

type BGPIPv4UnicastVRFNeighborResource struct {
	data *IosxeProviderData
}

func (r *BGPIPv4UnicastVRFNeighborResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_bgp_ipv4_unicast_vrf_neighbor"
}

func (r *BGPIPv4UnicastVRFNeighborResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the BGP IPv4 Unicast VRF Neighbor configuration.",

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
			"asn": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"vrf": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"ip": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"remote_as": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specify a BGP peer-group remote-as").String,
				Optional:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Neighbor specific description").String,
				Optional:            true,
			},
			"shutdown": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Administratively shut down this neighbor").String,
				Optional:            true,
			},
			"cluster_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
			},
			"log_neighbor_changes_disable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("disable").String,
				Optional:            true,
			},
			"password_type": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Encryption type (0 to disable encryption, 7 for proprietary)").AddIntegerRangeDescription(0, 7).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 7),
				},
			},
			"password": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 52),
					stringvalidator.RegexMatches(regexp.MustCompile(`.*`), ""),
				},
			},
			"timers_keepalive_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddIntegerRangeDescription(0, 65535).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 65535),
				},
			},
			"timers_holdtime": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddIntegerRangeDescription(0, 65535).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 65535),
				},
			},
			"timers_minimum_neighbor_hold": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddIntegerRangeDescription(0, 65535).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 65535),
				},
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the BGP version to match a neighbor").AddIntegerRangeDescription(4, 4).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(4, 4),
				},
			},
			"fall_over_default_route_map": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
			},
			"fall_over_bfd": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Use BFD to detect failure").String,
				Optional:            true,
			},
			"fall_over_bfd_multi_hop": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Force BFD multi-hop to detect failure").String,
				Optional:            true,
			},
			"fall_over_bfd_single_hop": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Force BFD single-hop to detect failure").String,
				Optional:            true,
			},
			"fall_over_bfd_check_control_plane_failure": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Retrieve control plane dependent failure info from BFD for BGP GR/NSR operation").String,
				Optional:            true,
			},
			"fall_over_bfd_strict_mode": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable BFD strict-mode").String,
				Optional:            true,
			},
			"fall_over_maximum_metric_route_map": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
			},
			"disable_connected_check": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("one-hop away EBGP peer using loopback address").String,
				Optional:            true,
			},
			"ttl_security_hops": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("IP hops").AddIntegerRangeDescription(1, 254).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 254),
				},
			},
			"local_as": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
			},
			"local_as_no_prepend": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Do not prepend local-as to updates from ebgp peers").String,
				Optional:            true,
			},
			"local_as_replace_as": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Replace real AS with local AS in the EBGP updates").String,
				Optional:            true,
			},
			"local_as_dual_as": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Accept either real AS or local AS from the ebgp peer").String,
				Optional:            true,
			},
			"update_source_loopback": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Loopback interface").String,
				Optional:            true,
			},
			"activate": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable the address family for this neighbor").AddDefaultValueDescription("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"send_community": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("both", "extended", "standard").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("both", "extended", "standard"),
				},
			},
			"route_reflector_client": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure a neighbor as Route Reflector client").String,
				Optional:            true,
			},
			"soft_reconfiguration": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Per neighbor soft reconfiguration").AddStringEnumDescription("inbound").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("inbound"),
				},
			},
			"default_originate": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Originate default route to this neighbor").String,
				Optional:            true,
			},
			"default_originate_route_map": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Route-map to specify criteria to originate default").String,
				Optional:            true,
			},
			"route_maps": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Apply route map to neighbor").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"in_out": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("in", "out").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("in", "out"),
							},
						},
						"route_map_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							Required:            true,
						},
					},
				},
			},
			"ebgp_multihop": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Allow EBGP neighbors not on directly connected networks. For single-hop ebgp peers, delete ebgp-multihop directly.").String,
				Optional:            true,
			},
			"ebgp_multihop_max_hop": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddIntegerRangeDescription(2, 255).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(2, 255),
				},
			},
			"ha_mode_graceful_restart": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("graceful-restart for this peer").String,
				Optional:            true,
			},
			"next_hop_self": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Disable the next hop calculation for this neighbor").String,
				Optional:            true,
			},
			"next_hop_self_all": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable next-hop-self for both eBGP and iBGP received paths").String,
				Optional:            true,
			},
			"advertisement_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Minimum interval between sending BGP routing updates").AddIntegerRangeDescription(0, 600).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 600),
				},
			},
		},
	}
}

func (r *BGPIPv4UnicastVRFNeighborResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.data = req.ProviderData.(*IosxeProviderData)
}

func (r *BGPIPv4UnicastVRFNeighborResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan BGPIPv4UnicastVRFNeighbor

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

func (r *BGPIPv4UnicastVRFNeighborResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state BGPIPv4UnicastVRFNeighbor

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
			state = BGPIPv4UnicastVRFNeighbor{Device: state.Device, Id: state.Id}
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

func (r *BGPIPv4UnicastVRFNeighborResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state BGPIPv4UnicastVRFNeighbor

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

func (r *BGPIPv4UnicastVRFNeighborResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state BGPIPv4UnicastVRFNeighbor

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

func (r *BGPIPv4UnicastVRFNeighborResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}
