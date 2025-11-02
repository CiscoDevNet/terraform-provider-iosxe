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
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-netconf"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &YangResource{}
var _ resource.ResourceWithImportState = &YangResource{}

func NewYangResource() resource.Resource {
	return &YangResource{}
}

type YangResource struct {
	data *IosxeProviderData
}

func (r *YangResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_yang"
}

func (r *YangResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Manages IOS-XE objects via YANG paths. This resource can only manage a single object. It is able to read the state and therefore reconcile configuration drift.",

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
			"path": schema.StringAttribute{
				MarkdownDescription: "A YANG path, e.g. `openconfig-interfaces:interfaces`.",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"delete": schema.BoolAttribute{
				MarkdownDescription: "Delete object during destroy operation. Default value is `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"attributes": schema.MapAttribute{
				MarkdownDescription: "Map of key-value pairs which represents the YANG leafs and its values.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
			},
			"lists": schema.ListNestedAttribute{
				MarkdownDescription: "YANG lists.",
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "YANG list name.",
							Required:            true,
						},
						"key": schema.StringAttribute{
							MarkdownDescription: "YANG list key attribute. In case of multiple keys, those should be separated by a comma (`,`).",
							Required:            true,
						},
						"items": schema.ListAttribute{
							MarkdownDescription: "List of maps of key-value pairs which represents the YANG leafs and its values.",
							Optional:            true,
							ElementType:         types.MapType{ElemType: types.StringType},
						},
						"values": schema.ListAttribute{
							MarkdownDescription: "YANG leaf-list values.",
							Optional:            true,
							ElementType:         types.StringType,
						},
					},
				},
			},
		},
	}
}

func (r *YangResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data Yang

	diag := req.Config.Get(ctx, &data)

	if diag.HasError() {
		return
	}

	for l := range data.Lists {
		key := data.Lists[l].Key.ValueString()
		for i := range data.Lists[l].Items {
			var m map[string]string
			data.Lists[l].Items[i].ElementsAs(ctx, &m, false)
			if _, ok := m[key]; !ok {
				resp.Diagnostics.AddAttributeError(
					path.Root("lists"),
					"Invalid List Configuration",
					fmt.Sprintf("Key '%s' is missing in list item.", key),
				)
			}
		}
	}
}

func (r *YangResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.data = req.ProviderData.(*IosxeProviderData)
}

func (r *YangResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan Yang

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
		if device.Protocol == "restconf" {
			body := plan.toBody(ctx)
			res, err := device.RestconfClient.PatchData(plan.getPathShort(), body)
			if len(res.Errors.Error) > 0 && res.Errors.Error[0].ErrorMessage == "patch to a nonexistent resource" {
				_, err = device.RestconfClient.PutData(plan.getPath(), body)
			}
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PATCH), got error: %s", err))
				return
			}
		} else {
			// Manage NETCONF connection lifecycle
			if device.NetconfClient != nil {
				cleanup, err := helpers.ManageNetconfConnection(ctx, device.NetconfClient, device.ReuseConnection)
				if err != nil {
					resp.Diagnostics.AddError("Connection Error", err.Error())
					return
				}
				defer cleanup()
			}

			body := plan.toBodyXML(ctx)
			if err := helpers.EditConfig(ctx, device.NetconfClient, body, true); err != nil {
				resp.Diagnostics.AddError("Client Error", err.Error())
				return
			}
		}
	}

	plan.Id = plan.Path

	if plan.Attributes.IsUnknown() {
		plan.Attributes = types.MapNull(types.StringType)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.getPath()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *YangResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state Yang

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.getPath()))

	device, ok := r.data.Devices[state.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", state.Device.ValueString()))
		return
	}

	if device.Managed {
		if device.Protocol == "restconf" {
			res, err := device.RestconfClient.GetData(state.getPath())
			if res.StatusCode == 404 {
				state.Attributes = types.MapNull(types.StringType)
				state.Lists = make([]YangList, 0)
			} else {
				if err != nil {
					resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to read object, got error: %s", err))
					return
				}

				state.fromBody(ctx, res.Res)
			}
		} else {
			// Manage NETCONF connection lifecycle
			if device.NetconfClient != nil {
				cleanup, err := helpers.ManageNetconfConnection(ctx, device.NetconfClient, device.ReuseConnection)
				if err != nil {
					resp.Diagnostics.AddError("Connection Error", err.Error())
					return
				}
				defer cleanup()
			}

			filter := helpers.GetXpathFilter(state.getPath())
			res, err := device.NetconfClient.GetConfig(ctx, "running", filter)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to read object, got error: %s", err))
				return
			}

			state.fromBodyXML(ctx, res.Res)
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.getPath()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *YangResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state Yang

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

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.getPath()))

	device, ok := r.data.Devices[plan.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", plan.Device.ValueString()))
		return
	}

	if device.Managed {
		if device.Protocol == "restconf" {
			body := plan.toBody(ctx)
			res, err := device.RestconfClient.PatchData(plan.getPathShort(), body)
			if len(res.Errors.Error) > 0 && res.Errors.Error[0].ErrorMessage == "patch to a nonexistent resource" {
				_, err = device.RestconfClient.PutData(plan.getPath(), body)
			}
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PATCH), got error: %s", err))
				return
			}

			deletedItems := plan.getDeletedItems(ctx, state)
			tflog.Debug(ctx, fmt.Sprintf("List items to delete: %+v", deletedItems))

			for _, i := range deletedItems {
				res, err := device.RestconfClient.DeleteData(i)
				if err != nil && res.StatusCode != 404 {
					resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
					return
				}
			}
		} else {
			// Manage NETCONF connection lifecycle
			if device.NetconfClient != nil {
				cleanup, err := helpers.ManageNetconfConnection(ctx, device.NetconfClient, device.ReuseConnection)
				if err != nil {
					resp.Diagnostics.AddError("Connection Error", err.Error())
					return
				}
				defer cleanup()
			}

			body := plan.toBodyXML(ctx)
			if err := helpers.EditConfig(ctx, device.NetconfClient, body, true); err != nil {
				resp.Diagnostics.AddError("Client Error", err.Error())
				return
			}
		}
	}

	if plan.Attributes.IsUnknown() {
		plan.Attributes = types.MapNull(types.StringType)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.getPath()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *YangResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state Yang

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.getPath()))

	device, ok := r.data.Devices[state.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", state.Device.ValueString()))
		return
	}

	if device.Managed && state.Delete.ValueBool() {
		if device.Protocol == "restconf" {
			res, err := device.RestconfClient.DeleteData(state.getPath())
			if err != nil && res.StatusCode != 404 {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
				return
			}
		} else {
			// Manage NETCONF connection lifecycle
			if device.NetconfClient != nil {
				cleanup, err := helpers.ManageNetconfConnection(ctx, device.NetconfClient, device.ReuseConnection)
				if err != nil {
					resp.Diagnostics.AddError("Connection Error", err.Error())
					return
				}
				defer cleanup()
			}

			body := netconf.Body{}
			body = helpers.RemoveFromXPath(body, state.getPath())
			if err := helpers.EditConfig(ctx, device.NetconfClient, body.Res(), true); err != nil {
				resp.Diagnostics.AddError("Client Error", err.Error())
				return
			}
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.getPath()))

	resp.State.RemoveResource(ctx)
}

func (r *YangResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Import", req.ID))

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("path"), req.ID)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)

	tflog.Debug(ctx, fmt.Sprintf("%s: Import finished successfully", req.ID))
}
