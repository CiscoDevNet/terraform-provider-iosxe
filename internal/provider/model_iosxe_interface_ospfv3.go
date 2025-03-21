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

type InterfaceOSPFv3 struct {
	Device                       types.String `tfsdk:"device"`
	Id                           types.String `tfsdk:"id"`
	DeleteMode                   types.String `tfsdk:"delete_mode"`
	Type                         types.String `tfsdk:"type"`
	Name                         types.String `tfsdk:"name"`
	NetworkTypeBroadcast         types.Bool   `tfsdk:"network_type_broadcast"`
	NetworkTypeNonBroadcast      types.Bool   `tfsdk:"network_type_non_broadcast"`
	NetworkTypePointToMultipoint types.Bool   `tfsdk:"network_type_point_to_multipoint"`
	NetworkTypePointToPoint      types.Bool   `tfsdk:"network_type_point_to_point"`
	Cost                         types.Int64  `tfsdk:"cost"`
}

type InterfaceOSPFv3Data struct {
	Device                       types.String `tfsdk:"device"`
	Id                           types.String `tfsdk:"id"`
	Type                         types.String `tfsdk:"type"`
	Name                         types.String `tfsdk:"name"`
	NetworkTypeBroadcast         types.Bool   `tfsdk:"network_type_broadcast"`
	NetworkTypeNonBroadcast      types.Bool   `tfsdk:"network_type_non_broadcast"`
	NetworkTypePointToMultipoint types.Bool   `tfsdk:"network_type_point_to_multipoint"`
	NetworkTypePointToPoint      types.Bool   `tfsdk:"network_type_point_to_point"`
	Cost                         types.Int64  `tfsdk:"cost"`
}

func (data InterfaceOSPFv3) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/interface/%s=%v/Cisco-IOS-XE-ospfv3:ospfv3", url.QueryEscape(fmt.Sprintf("%v", data.Type.ValueString())), url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())))
}

func (data InterfaceOSPFv3Data) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/interface/%s=%v/Cisco-IOS-XE-ospfv3:ospfv3", url.QueryEscape(fmt.Sprintf("%v", data.Type.ValueString())), url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())))
}

// if last path element has a key -> remove it
func (data InterfaceOSPFv3) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data InterfaceOSPFv3) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.NetworkTypeBroadcast.IsNull() && !data.NetworkTypeBroadcast.IsUnknown() {
		if data.NetworkTypeBroadcast.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"network-type.broadcast", map[string]string{})
		}
	}
	if !data.NetworkTypeNonBroadcast.IsNull() && !data.NetworkTypeNonBroadcast.IsUnknown() {
		if data.NetworkTypeNonBroadcast.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"network-type.non-broadcast", map[string]string{})
		}
	}
	if !data.NetworkTypePointToMultipoint.IsNull() && !data.NetworkTypePointToMultipoint.IsUnknown() {
		if data.NetworkTypePointToMultipoint.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"network-type.point-to-multipoint", map[string]string{})
		}
	}
	if !data.NetworkTypePointToPoint.IsNull() && !data.NetworkTypePointToPoint.IsUnknown() {
		if data.NetworkTypePointToPoint.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"network-type.point-to-point", map[string]string{})
		}
	}
	if !data.Cost.IsNull() && !data.Cost.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"cost-config.value", strconv.FormatInt(data.Cost.ValueInt64(), 10))
	}
	return body
}

func (data *InterfaceOSPFv3) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "network-type.broadcast"); !data.NetworkTypeBroadcast.IsNull() {
		if value.Exists() {
			data.NetworkTypeBroadcast = types.BoolValue(true)
		} else {
			data.NetworkTypeBroadcast = types.BoolValue(false)
		}
	} else {
		data.NetworkTypeBroadcast = types.BoolNull()
	}
	if value := res.Get(prefix + "network-type.non-broadcast"); !data.NetworkTypeNonBroadcast.IsNull() {
		if value.Exists() {
			data.NetworkTypeNonBroadcast = types.BoolValue(true)
		} else {
			data.NetworkTypeNonBroadcast = types.BoolValue(false)
		}
	} else {
		data.NetworkTypeNonBroadcast = types.BoolNull()
	}
	if value := res.Get(prefix + "network-type.point-to-multipoint"); !data.NetworkTypePointToMultipoint.IsNull() {
		if value.Exists() {
			data.NetworkTypePointToMultipoint = types.BoolValue(true)
		} else {
			data.NetworkTypePointToMultipoint = types.BoolValue(false)
		}
	} else {
		data.NetworkTypePointToMultipoint = types.BoolNull()
	}
	if value := res.Get(prefix + "network-type.point-to-point"); !data.NetworkTypePointToPoint.IsNull() {
		if value.Exists() {
			data.NetworkTypePointToPoint = types.BoolValue(true)
		} else {
			data.NetworkTypePointToPoint = types.BoolValue(false)
		}
	} else {
		data.NetworkTypePointToPoint = types.BoolNull()
	}
	if value := res.Get(prefix + "cost-config.value"); value.Exists() && !data.Cost.IsNull() {
		data.Cost = types.Int64Value(value.Int())
	} else {
		data.Cost = types.Int64Null()
	}
}

func (data *InterfaceOSPFv3) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "network-type.broadcast"); value.Exists() {
		data.NetworkTypeBroadcast = types.BoolValue(true)
	} else {
		data.NetworkTypeBroadcast = types.BoolValue(false)
	}
	if value := res.Get(prefix + "network-type.non-broadcast"); value.Exists() {
		data.NetworkTypeNonBroadcast = types.BoolValue(true)
	} else {
		data.NetworkTypeNonBroadcast = types.BoolValue(false)
	}
	if value := res.Get(prefix + "network-type.point-to-multipoint"); value.Exists() {
		data.NetworkTypePointToMultipoint = types.BoolValue(true)
	} else {
		data.NetworkTypePointToMultipoint = types.BoolValue(false)
	}
	if value := res.Get(prefix + "network-type.point-to-point"); value.Exists() {
		data.NetworkTypePointToPoint = types.BoolValue(true)
	} else {
		data.NetworkTypePointToPoint = types.BoolValue(false)
	}
	if value := res.Get(prefix + "cost-config.value"); value.Exists() {
		data.Cost = types.Int64Value(value.Int())
	}
}

func (data *InterfaceOSPFv3Data) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "network-type.broadcast"); value.Exists() {
		data.NetworkTypeBroadcast = types.BoolValue(true)
	} else {
		data.NetworkTypeBroadcast = types.BoolValue(false)
	}
	if value := res.Get(prefix + "network-type.non-broadcast"); value.Exists() {
		data.NetworkTypeNonBroadcast = types.BoolValue(true)
	} else {
		data.NetworkTypeNonBroadcast = types.BoolValue(false)
	}
	if value := res.Get(prefix + "network-type.point-to-multipoint"); value.Exists() {
		data.NetworkTypePointToMultipoint = types.BoolValue(true)
	} else {
		data.NetworkTypePointToMultipoint = types.BoolValue(false)
	}
	if value := res.Get(prefix + "network-type.point-to-point"); value.Exists() {
		data.NetworkTypePointToPoint = types.BoolValue(true)
	} else {
		data.NetworkTypePointToPoint = types.BoolValue(false)
	}
	if value := res.Get(prefix + "cost-config.value"); value.Exists() {
		data.Cost = types.Int64Value(value.Int())
	}
}

func (data *InterfaceOSPFv3) getDeletedItems(ctx context.Context, state InterfaceOSPFv3) []string {
	deletedItems := make([]string, 0)
	if !state.NetworkTypeBroadcast.IsNull() && data.NetworkTypeBroadcast.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/network-type/broadcast", state.getPath()))
	}
	if !state.NetworkTypeNonBroadcast.IsNull() && data.NetworkTypeNonBroadcast.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/network-type/non-broadcast", state.getPath()))
	}
	if !state.NetworkTypePointToMultipoint.IsNull() && data.NetworkTypePointToMultipoint.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/network-type/point-to-multipoint", state.getPath()))
	}
	if !state.NetworkTypePointToPoint.IsNull() && data.NetworkTypePointToPoint.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/network-type/point-to-point", state.getPath()))
	}
	if !state.Cost.IsNull() && data.Cost.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/cost-config/value", state.getPath()))
	}
	return deletedItems
}

func (data *InterfaceOSPFv3) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.NetworkTypeBroadcast.IsNull() && !data.NetworkTypeBroadcast.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/network-type/broadcast", data.getPath()))
	}
	if !data.NetworkTypeNonBroadcast.IsNull() && !data.NetworkTypeNonBroadcast.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/network-type/non-broadcast", data.getPath()))
	}
	if !data.NetworkTypePointToMultipoint.IsNull() && !data.NetworkTypePointToMultipoint.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/network-type/point-to-multipoint", data.getPath()))
	}
	if !data.NetworkTypePointToPoint.IsNull() && !data.NetworkTypePointToPoint.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/network-type/point-to-point", data.getPath()))
	}
	return emptyLeafsDelete
}

func (data *InterfaceOSPFv3) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.NetworkTypeBroadcast.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/network-type/broadcast", data.getPath()))
	}
	if !data.NetworkTypeNonBroadcast.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/network-type/non-broadcast", data.getPath()))
	}
	if !data.NetworkTypePointToMultipoint.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/network-type/point-to-multipoint", data.getPath()))
	}
	if !data.NetworkTypePointToPoint.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/network-type/point-to-point", data.getPath()))
	}
	if !data.Cost.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/cost-config/value", data.getPath()))
	}
	return deletePaths
}

func (data *InterfaceOSPFv3) getIdsFromPath() {
	reString := strings.ReplaceAll("Cisco-IOS-XE-native:native/interface/%s=%v/Cisco-IOS-XE-ospfv3:ospfv3", "%s", "(.+)")
	reString = strings.ReplaceAll(reString, "%v", "(.+)")
	re := regexp.MustCompile(reString)
	matches := re.FindStringSubmatch(data.Id.ValueString())
	data.Type = types.StringValue(helpers.Must(url.QueryUnescape(matches[1])))
	data.Name = types.StringValue(helpers.Must(url.QueryUnescape(matches[2])))
}
