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
	"strings"

	"github.com/CiscoDevNet/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type BGPAddressFamilyL2VPN struct {
	Device     types.String `tfsdk:"device"`
	Id         types.String `tfsdk:"id"`
	DeleteMode types.String `tfsdk:"delete_mode"`
	Asn        types.String `tfsdk:"asn"`
	AfName     types.String `tfsdk:"af_name"`
}

type BGPAddressFamilyL2VPNData struct {
	Device types.String `tfsdk:"device"`
	Id     types.String `tfsdk:"id"`
	Asn    types.String `tfsdk:"asn"`
	AfName types.String `tfsdk:"af_name"`
}

func (data BGPAddressFamilyL2VPN) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=%v/address-family/no-vrf/l2vpn=%s", url.QueryEscape(fmt.Sprintf("%v", data.Asn.ValueString())), url.QueryEscape(fmt.Sprintf("%v", data.AfName.ValueString())))
}

func (data BGPAddressFamilyL2VPNData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=%v/address-family/no-vrf/l2vpn=%s", url.QueryEscape(fmt.Sprintf("%v", data.Asn.ValueString())), url.QueryEscape(fmt.Sprintf("%v", data.AfName.ValueString())))
}

// if last path element has a key -> remove it
func (data BGPAddressFamilyL2VPN) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data BGPAddressFamilyL2VPN) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.AfName.IsNull() && !data.AfName.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"af-name", data.AfName.ValueString())
	}
	return body
}

func (data *BGPAddressFamilyL2VPN) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "af-name"); value.Exists() && !data.AfName.IsNull() {
		data.AfName = types.StringValue(value.String())
	} else {
		data.AfName = types.StringNull()
	}
}

func (data *BGPAddressFamilyL2VPN) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
}

func (data *BGPAddressFamilyL2VPNData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
}

func (data *BGPAddressFamilyL2VPN) getDeletedItems(ctx context.Context, state BGPAddressFamilyL2VPN) []string {
	deletedItems := make([]string, 0)
	return deletedItems
}

func (data *BGPAddressFamilyL2VPN) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	return emptyLeafsDelete
}

func (data *BGPAddressFamilyL2VPN) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	return deletePaths
}

func (data *BGPAddressFamilyL2VPN) getIdsFromPath() {
	reString := strings.ReplaceAll("Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=%v/address-family/no-vrf/l2vpn=%s", "%s", "(.+)")
	reString = strings.ReplaceAll(reString, "%v", "(.+)")
	re := regexp.MustCompile(reString)
	matches := re.FindStringSubmatch(data.Id.ValueString())
	data.Asn = types.StringValue(matches[1])
	data.AfName = types.StringValue(matches[2])
}
