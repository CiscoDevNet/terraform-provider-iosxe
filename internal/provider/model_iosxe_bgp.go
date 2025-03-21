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

type BGP struct {
	Device             types.String `tfsdk:"device"`
	Id                 types.String `tfsdk:"id"`
	DeleteMode         types.String `tfsdk:"delete_mode"`
	Asn                types.String `tfsdk:"asn"`
	DefaultIpv4Unicast types.Bool   `tfsdk:"default_ipv4_unicast"`
	LogNeighborChanges types.Bool   `tfsdk:"log_neighbor_changes"`
	RouterIdLoopback   types.Int64  `tfsdk:"router_id_loopback"`
}

type BGPData struct {
	Device             types.String `tfsdk:"device"`
	Id                 types.String `tfsdk:"id"`
	Asn                types.String `tfsdk:"asn"`
	DefaultIpv4Unicast types.Bool   `tfsdk:"default_ipv4_unicast"`
	LogNeighborChanges types.Bool   `tfsdk:"log_neighbor_changes"`
	RouterIdLoopback   types.Int64  `tfsdk:"router_id_loopback"`
}

func (data BGP) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=%v", url.QueryEscape(fmt.Sprintf("%v", data.Asn.ValueString())))
}

func (data BGPData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=%v", url.QueryEscape(fmt.Sprintf("%v", data.Asn.ValueString())))
}

// if last path element has a key -> remove it
func (data BGP) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data BGP) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Asn.IsNull() && !data.Asn.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"id", data.Asn.ValueString())
	}
	if !data.DefaultIpv4Unicast.IsNull() && !data.DefaultIpv4Unicast.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"bgp.default.ipv4-unicast", data.DefaultIpv4Unicast.ValueBool())
	}
	if !data.LogNeighborChanges.IsNull() && !data.LogNeighborChanges.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"bgp.log-neighbor-changes", data.LogNeighborChanges.ValueBool())
	}
	if !data.RouterIdLoopback.IsNull() && !data.RouterIdLoopback.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"bgp.router-id.interface.Loopback", strconv.FormatInt(data.RouterIdLoopback.ValueInt64(), 10))
	}
	return body
}

func (data *BGP) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "id"); value.Exists() && !data.Asn.IsNull() {
		data.Asn = types.StringValue(value.String())
	} else {
		data.Asn = types.StringNull()
	}
	if value := res.Get(prefix + "bgp.default.ipv4-unicast"); !data.DefaultIpv4Unicast.IsNull() {
		if value.Exists() {
			data.DefaultIpv4Unicast = types.BoolValue(value.Bool())
		}
	} else {
		data.DefaultIpv4Unicast = types.BoolNull()
	}
	if value := res.Get(prefix + "bgp.log-neighbor-changes"); !data.LogNeighborChanges.IsNull() {
		if value.Exists() {
			data.LogNeighborChanges = types.BoolValue(value.Bool())
		}
	} else {
		data.LogNeighborChanges = types.BoolNull()
	}
	if value := res.Get(prefix + "bgp.router-id.interface.Loopback"); value.Exists() && !data.RouterIdLoopback.IsNull() {
		data.RouterIdLoopback = types.Int64Value(value.Int())
	} else {
		data.RouterIdLoopback = types.Int64Null()
	}
}

func (data *BGP) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "bgp.default.ipv4-unicast"); value.Exists() {
		data.DefaultIpv4Unicast = types.BoolValue(value.Bool())
	} else {
		data.DefaultIpv4Unicast = types.BoolValue(false)
	}
	if value := res.Get(prefix + "bgp.log-neighbor-changes"); value.Exists() {
		data.LogNeighborChanges = types.BoolValue(value.Bool())
	} else {
		data.LogNeighborChanges = types.BoolValue(false)
	}
	if value := res.Get(prefix + "bgp.router-id.interface.Loopback"); value.Exists() {
		data.RouterIdLoopback = types.Int64Value(value.Int())
	}
}

func (data *BGPData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "bgp.default.ipv4-unicast"); value.Exists() {
		data.DefaultIpv4Unicast = types.BoolValue(value.Bool())
	} else {
		data.DefaultIpv4Unicast = types.BoolValue(false)
	}
	if value := res.Get(prefix + "bgp.log-neighbor-changes"); value.Exists() {
		data.LogNeighborChanges = types.BoolValue(value.Bool())
	} else {
		data.LogNeighborChanges = types.BoolValue(false)
	}
	if value := res.Get(prefix + "bgp.router-id.interface.Loopback"); value.Exists() {
		data.RouterIdLoopback = types.Int64Value(value.Int())
	}
}

func (data *BGP) getDeletedItems(ctx context.Context, state BGP) []string {
	deletedItems := make([]string, 0)
	if !state.DefaultIpv4Unicast.IsNull() && data.DefaultIpv4Unicast.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/bgp/default/ipv4-unicast", state.getPath()))
	}
	if !state.LogNeighborChanges.IsNull() && data.LogNeighborChanges.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/bgp/log-neighbor-changes", state.getPath()))
	}
	if !state.RouterIdLoopback.IsNull() && data.RouterIdLoopback.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/bgp/router-id/interface/Loopback", state.getPath()))
	}
	return deletedItems
}

func (data *BGP) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	return emptyLeafsDelete
}

func (data *BGP) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.DefaultIpv4Unicast.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/bgp/default/ipv4-unicast", data.getPath()))
	}
	if !data.LogNeighborChanges.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/bgp/log-neighbor-changes", data.getPath()))
	}
	if !data.RouterIdLoopback.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/bgp/router-id/interface/Loopback", data.getPath()))
	}
	return deletePaths
}

func (data *BGP) getIdsFromPath() {
	reString := strings.ReplaceAll("Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=%v", "%s", "(.+)")
	reString = strings.ReplaceAll(reString, "%v", "(.+)")
	re := regexp.MustCompile(reString)
	matches := re.FindStringSubmatch(data.Id.ValueString())
	data.Asn = types.StringValue(matches[1])
}
