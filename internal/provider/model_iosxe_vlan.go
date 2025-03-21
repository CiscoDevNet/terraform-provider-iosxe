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
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type VLAN struct {
	Device                 types.String `tfsdk:"device"`
	Id                     types.String `tfsdk:"id"`
	VlanId                 types.Int64  `tfsdk:"vlan_id"`
	RemoteSpan             types.Bool   `tfsdk:"remote_span"`
	PrivateVlanPrimary     types.Bool   `tfsdk:"private_vlan_primary"`
	PrivateVlanAssociation types.String `tfsdk:"private_vlan_association"`
	PrivateVlanCommunity   types.Bool   `tfsdk:"private_vlan_community"`
	PrivateVlanIsolated    types.Bool   `tfsdk:"private_vlan_isolated"`
	Name                   types.String `tfsdk:"name"`
	Shutdown               types.Bool   `tfsdk:"shutdown"`
}

type VLANData struct {
	Device                 types.String `tfsdk:"device"`
	Id                     types.String `tfsdk:"id"`
	VlanId                 types.Int64  `tfsdk:"vlan_id"`
	RemoteSpan             types.Bool   `tfsdk:"remote_span"`
	PrivateVlanPrimary     types.Bool   `tfsdk:"private_vlan_primary"`
	PrivateVlanAssociation types.String `tfsdk:"private_vlan_association"`
	PrivateVlanCommunity   types.Bool   `tfsdk:"private_vlan_community"`
	PrivateVlanIsolated    types.Bool   `tfsdk:"private_vlan_isolated"`
	Name                   types.String `tfsdk:"name"`
	Shutdown               types.Bool   `tfsdk:"shutdown"`
}

func (data VLAN) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/vlan/Cisco-IOS-XE-vlan:vlan-list=%v", url.QueryEscape(fmt.Sprintf("%v", data.VlanId.ValueInt64())))
}

func (data VLANData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/vlan/Cisco-IOS-XE-vlan:vlan-list=%v", url.QueryEscape(fmt.Sprintf("%v", data.VlanId.ValueInt64())))
}

// if last path element has a key -> remove it
func (data VLAN) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data VLAN) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.VlanId.IsNull() && !data.VlanId.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"id", strconv.FormatInt(data.VlanId.ValueInt64(), 10))
	}
	if !data.RemoteSpan.IsNull() && !data.RemoteSpan.IsUnknown() {
		if data.RemoteSpan.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"remote-span", map[string]string{})
		}
	}
	if !data.PrivateVlanPrimary.IsNull() && !data.PrivateVlanPrimary.IsUnknown() {
		if data.PrivateVlanPrimary.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"private-vlan.primary", map[string]string{})
		}
	}
	if !data.PrivateVlanAssociation.IsNull() && !data.PrivateVlanAssociation.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"private-vlan.association", data.PrivateVlanAssociation.ValueString())
	}
	if !data.PrivateVlanCommunity.IsNull() && !data.PrivateVlanCommunity.IsUnknown() {
		if data.PrivateVlanCommunity.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"private-vlan.community", map[string]string{})
		}
	}
	if !data.PrivateVlanIsolated.IsNull() && !data.PrivateVlanIsolated.IsUnknown() {
		if data.PrivateVlanIsolated.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"private-vlan.isolated", map[string]string{})
		}
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"name", data.Name.ValueString())
	}
	if !data.Shutdown.IsNull() && !data.Shutdown.IsUnknown() {
		if data.Shutdown.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"shutdown", map[string]string{})
		}
	}
	return body
}

func (data *VLAN) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "id"); value.Exists() && !data.VlanId.IsNull() {
		data.VlanId = types.Int64Value(value.Int())
	} else {
		data.VlanId = types.Int64Null()
	}
	if value := res.Get(prefix + "remote-span"); !data.RemoteSpan.IsNull() {
		if value.Exists() {
			data.RemoteSpan = types.BoolValue(true)
		} else {
			data.RemoteSpan = types.BoolValue(false)
		}
	} else {
		data.RemoteSpan = types.BoolNull()
	}
	if value := res.Get(prefix + "private-vlan.primary"); !data.PrivateVlanPrimary.IsNull() {
		if value.Exists() {
			data.PrivateVlanPrimary = types.BoolValue(true)
		} else {
			data.PrivateVlanPrimary = types.BoolValue(false)
		}
	} else {
		data.PrivateVlanPrimary = types.BoolNull()
	}
	if value := res.Get(prefix + "private-vlan.association"); value.Exists() && !data.PrivateVlanAssociation.IsNull() {
		data.PrivateVlanAssociation = types.StringValue(value.String())
	} else {
		data.PrivateVlanAssociation = types.StringNull()
	}
	if value := res.Get(prefix + "private-vlan.community"); !data.PrivateVlanCommunity.IsNull() {
		if value.Exists() {
			data.PrivateVlanCommunity = types.BoolValue(true)
		} else {
			data.PrivateVlanCommunity = types.BoolValue(false)
		}
	} else {
		data.PrivateVlanCommunity = types.BoolNull()
	}
	if value := res.Get(prefix + "private-vlan.isolated"); !data.PrivateVlanIsolated.IsNull() {
		if value.Exists() {
			data.PrivateVlanIsolated = types.BoolValue(true)
		} else {
			data.PrivateVlanIsolated = types.BoolValue(false)
		}
	} else {
		data.PrivateVlanIsolated = types.BoolNull()
	}
	if value := res.Get(prefix + "name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(prefix + "shutdown"); !data.Shutdown.IsNull() {
		if value.Exists() {
			data.Shutdown = types.BoolValue(true)
		} else {
			data.Shutdown = types.BoolValue(false)
		}
	} else {
		data.Shutdown = types.BoolNull()
	}
}

func (data *VLAN) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "remote-span"); value.Exists() {
		data.RemoteSpan = types.BoolValue(true)
	} else {
		data.RemoteSpan = types.BoolValue(false)
	}
	if value := res.Get(prefix + "private-vlan.primary"); value.Exists() {
		data.PrivateVlanPrimary = types.BoolValue(true)
	} else {
		data.PrivateVlanPrimary = types.BoolValue(false)
	}
	if value := res.Get(prefix + "private-vlan.association"); value.Exists() {
		data.PrivateVlanAssociation = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "private-vlan.community"); value.Exists() {
		data.PrivateVlanCommunity = types.BoolValue(true)
	} else {
		data.PrivateVlanCommunity = types.BoolValue(false)
	}
	if value := res.Get(prefix + "private-vlan.isolated"); value.Exists() {
		data.PrivateVlanIsolated = types.BoolValue(true)
	} else {
		data.PrivateVlanIsolated = types.BoolValue(false)
	}
	if value := res.Get(prefix + "name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "shutdown"); value.Exists() {
		data.Shutdown = types.BoolValue(true)
	} else {
		data.Shutdown = types.BoolValue(false)
	}
}

func (data *VLANData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "remote-span"); value.Exists() {
		data.RemoteSpan = types.BoolValue(true)
	} else {
		data.RemoteSpan = types.BoolValue(false)
	}
	if value := res.Get(prefix + "private-vlan.primary"); value.Exists() {
		data.PrivateVlanPrimary = types.BoolValue(true)
	} else {
		data.PrivateVlanPrimary = types.BoolValue(false)
	}
	if value := res.Get(prefix + "private-vlan.association"); value.Exists() {
		data.PrivateVlanAssociation = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "private-vlan.community"); value.Exists() {
		data.PrivateVlanCommunity = types.BoolValue(true)
	} else {
		data.PrivateVlanCommunity = types.BoolValue(false)
	}
	if value := res.Get(prefix + "private-vlan.isolated"); value.Exists() {
		data.PrivateVlanIsolated = types.BoolValue(true)
	} else {
		data.PrivateVlanIsolated = types.BoolValue(false)
	}
	if value := res.Get(prefix + "name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "shutdown"); value.Exists() {
		data.Shutdown = types.BoolValue(true)
	} else {
		data.Shutdown = types.BoolValue(false)
	}
}

func (data *VLAN) getDeletedItems(ctx context.Context, state VLAN) []string {
	deletedItems := make([]string, 0)
	if !state.RemoteSpan.IsNull() && data.RemoteSpan.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/remote-span", state.getPath()))
	}
	if !state.PrivateVlanPrimary.IsNull() && data.PrivateVlanPrimary.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/private-vlan/primary", state.getPath()))
	}
	if !state.PrivateVlanAssociation.IsNull() && data.PrivateVlanAssociation.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/private-vlan/association", state.getPath()))
	}
	if !state.PrivateVlanCommunity.IsNull() && data.PrivateVlanCommunity.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/private-vlan/community", state.getPath()))
	}
	if !state.PrivateVlanIsolated.IsNull() && data.PrivateVlanIsolated.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/private-vlan/isolated", state.getPath()))
	}
	if !state.Name.IsNull() && data.Name.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/name", state.getPath()))
	}
	if !state.Shutdown.IsNull() && data.Shutdown.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/shutdown", state.getPath()))
	}
	return deletedItems
}

func (data *VLAN) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.RemoteSpan.IsNull() && !data.RemoteSpan.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/remote-span", data.getPath()))
	}
	if !data.PrivateVlanPrimary.IsNull() && !data.PrivateVlanPrimary.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/private-vlan/primary", data.getPath()))
	}
	if !data.PrivateVlanCommunity.IsNull() && !data.PrivateVlanCommunity.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/private-vlan/community", data.getPath()))
	}
	if !data.PrivateVlanIsolated.IsNull() && !data.PrivateVlanIsolated.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/private-vlan/isolated", data.getPath()))
	}
	if !data.Shutdown.IsNull() && !data.Shutdown.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/shutdown", data.getPath()))
	}
	return emptyLeafsDelete
}

func (data *VLAN) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.RemoteSpan.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/remote-span", data.getPath()))
	}
	if !data.PrivateVlanPrimary.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/private-vlan/primary", data.getPath()))
	}
	if !data.PrivateVlanAssociation.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/private-vlan/association", data.getPath()))
	}
	if !data.PrivateVlanCommunity.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/private-vlan/community", data.getPath()))
	}
	if !data.PrivateVlanIsolated.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/private-vlan/isolated", data.getPath()))
	}
	if !data.Name.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/name", data.getPath()))
	}
	if !data.Shutdown.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/shutdown", data.getPath()))
	}
	return deletePaths
}

func (data *VLAN) getIdsFromPath() {
	reString := strings.ReplaceAll("Cisco-IOS-XE-native:native/vlan/Cisco-IOS-XE-vlan:vlan-list=%v", "%s", "(.+)")
	reString = strings.ReplaceAll(reString, "%v", "(.+)")
	re := regexp.MustCompile(reString)
	matches := re.FindStringSubmatch(data.Id.ValueString())
	data.VlanId = types.Int64Value(helpers.Must(strconv.ParseInt(matches[1], 10, 0)))
}
