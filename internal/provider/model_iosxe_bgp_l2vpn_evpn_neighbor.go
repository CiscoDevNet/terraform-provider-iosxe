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

	"github.com/CiscoDevNet/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type BGPL2VPNEVPNNeighbor struct {
	Device               types.String `tfsdk:"device"`
	Id                   types.String `tfsdk:"id"`
	DeleteMode           types.String `tfsdk:"delete_mode"`
	Asn                  types.String `tfsdk:"asn"`
	Ip                   types.String `tfsdk:"ip"`
	Activate             types.Bool   `tfsdk:"activate"`
	SendCommunity        types.String `tfsdk:"send_community"`
	RouteReflectorClient types.Bool   `tfsdk:"route_reflector_client"`
	SoftReconfiguration  types.String `tfsdk:"soft_reconfiguration"`
}

type BGPL2VPNEVPNNeighborData struct {
	Device               types.String `tfsdk:"device"`
	Id                   types.String `tfsdk:"id"`
	Asn                  types.String `tfsdk:"asn"`
	Ip                   types.String `tfsdk:"ip"`
	Activate             types.Bool   `tfsdk:"activate"`
	SendCommunity        types.String `tfsdk:"send_community"`
	RouteReflectorClient types.Bool   `tfsdk:"route_reflector_client"`
	SoftReconfiguration  types.String `tfsdk:"soft_reconfiguration"`
}

func (data BGPL2VPNEVPNNeighbor) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=%v/address-family/no-vrf/l2vpn=evpn/l2vpn-evpn/neighbor=%s", url.QueryEscape(fmt.Sprintf("%v", data.Asn.ValueString())), url.QueryEscape(fmt.Sprintf("%v", data.Ip.ValueString())))
}

func (data BGPL2VPNEVPNNeighborData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=%v/address-family/no-vrf/l2vpn=evpn/l2vpn-evpn/neighbor=%s", url.QueryEscape(fmt.Sprintf("%v", data.Asn.ValueString())), url.QueryEscape(fmt.Sprintf("%v", data.Ip.ValueString())))
}

// if last path element has a key -> remove it
func (data BGPL2VPNEVPNNeighbor) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data BGPL2VPNEVPNNeighbor) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Ip.IsNull() && !data.Ip.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"id", data.Ip.ValueString())
	}
	if !data.Activate.IsNull() && !data.Activate.IsUnknown() {
		if data.Activate.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"activate", map[string]string{})
		}
	}
	if !data.SendCommunity.IsNull() && !data.SendCommunity.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"send-community.send-community-where", data.SendCommunity.ValueString())
	}
	if !data.RouteReflectorClient.IsNull() && !data.RouteReflectorClient.IsUnknown() {
		if data.RouteReflectorClient.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"route-reflector-client", map[string]string{})
		}
	}
	if !data.SoftReconfiguration.IsNull() && !data.SoftReconfiguration.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"soft-reconfiguration", data.SoftReconfiguration.ValueString())
	}
	return body
}

func (data *BGPL2VPNEVPNNeighbor) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "id"); value.Exists() && !data.Ip.IsNull() {
		data.Ip = types.StringValue(value.String())
	} else {
		data.Ip = types.StringNull()
	}
	if value := res.Get(prefix + "activate"); !data.Activate.IsNull() {
		if value.Exists() {
			data.Activate = types.BoolValue(true)
		} else {
			data.Activate = types.BoolValue(false)
		}
	} else {
		data.Activate = types.BoolNull()
	}
	if value := res.Get(prefix + "send-community.send-community-where"); value.Exists() && !data.SendCommunity.IsNull() {
		data.SendCommunity = types.StringValue(value.String())
	} else {
		data.SendCommunity = types.StringNull()
	}
	if value := res.Get(prefix + "route-reflector-client"); !data.RouteReflectorClient.IsNull() {
		if value.Exists() {
			data.RouteReflectorClient = types.BoolValue(true)
		} else {
			data.RouteReflectorClient = types.BoolValue(false)
		}
	} else {
		data.RouteReflectorClient = types.BoolNull()
	}
	if value := res.Get(prefix + "soft-reconfiguration"); value.Exists() && !data.SoftReconfiguration.IsNull() {
		data.SoftReconfiguration = types.StringValue(value.String())
	} else {
		data.SoftReconfiguration = types.StringNull()
	}
}

func (data *BGPL2VPNEVPNNeighborData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "activate"); value.Exists() {
		data.Activate = types.BoolValue(true)
	} else {
		data.Activate = types.BoolValue(false)
	}
	if value := res.Get(prefix + "send-community.send-community-where"); value.Exists() {
		data.SendCommunity = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "route-reflector-client"); value.Exists() {
		data.RouteReflectorClient = types.BoolValue(true)
	} else {
		data.RouteReflectorClient = types.BoolValue(false)
	}
	if value := res.Get(prefix + "soft-reconfiguration"); value.Exists() {
		data.SoftReconfiguration = types.StringValue(value.String())
	}
}

func (data *BGPL2VPNEVPNNeighbor) getDeletedItems(ctx context.Context, state BGPL2VPNEVPNNeighbor) []string {
	deletedItems := make([]string, 0)
	if !state.SendCommunity.IsNull() && data.SendCommunity.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/send-community/send-community-where", state.getPath()))
	}
	if !state.RouteReflectorClient.IsNull() && data.RouteReflectorClient.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/route-reflector-client", state.getPath()))
	}
	if !state.SoftReconfiguration.IsNull() && data.SoftReconfiguration.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/soft-reconfiguration", state.getPath()))
	}
	return deletedItems
}

func (data *BGPL2VPNEVPNNeighbor) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.Activate.IsNull() && !data.Activate.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/activate", data.getPath()))
	}
	if !data.RouteReflectorClient.IsNull() && !data.RouteReflectorClient.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/route-reflector-client", data.getPath()))
	}
	return emptyLeafsDelete
}

func (data *BGPL2VPNEVPNNeighbor) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.SendCommunity.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/send-community/send-community-where", data.getPath()))
	}
	if !data.RouteReflectorClient.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/route-reflector-client", data.getPath()))
	}
	if !data.SoftReconfiguration.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/soft-reconfiguration", data.getPath()))
	}
	return deletePaths
}
