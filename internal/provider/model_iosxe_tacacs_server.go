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

type TACACSServer struct {
	Device      types.String `tfsdk:"device"`
	Id          types.String `tfsdk:"id"`
	DeleteMode  types.String `tfsdk:"delete_mode"`
	Name        types.String `tfsdk:"name"`
	AddressIpv4 types.String `tfsdk:"address_ipv4"`
	Timeout     types.Int64  `tfsdk:"timeout"`
	Encryption  types.String `tfsdk:"encryption"`
	Key         types.String `tfsdk:"key"`
}

type TACACSServerData struct {
	Device      types.String `tfsdk:"device"`
	Id          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	AddressIpv4 types.String `tfsdk:"address_ipv4"`
	Timeout     types.Int64  `tfsdk:"timeout"`
	Encryption  types.String `tfsdk:"encryption"`
	Key         types.String `tfsdk:"key"`
}

func (data TACACSServer) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/tacacs/Cisco-IOS-XE-aaa:server=%v", url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())))
}

func (data TACACSServerData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/tacacs/Cisco-IOS-XE-aaa:server=%v", url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())))
}

// if last path element has a key -> remove it
func (data TACACSServer) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data TACACSServer) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"name", data.Name.ValueString())
	}
	if !data.AddressIpv4.IsNull() && !data.AddressIpv4.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"address.ipv4", data.AddressIpv4.ValueString())
	}
	if !data.Timeout.IsNull() && !data.Timeout.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"timeout", strconv.FormatInt(data.Timeout.ValueInt64(), 10))
	}
	if !data.Encryption.IsNull() && !data.Encryption.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"key.encryption", data.Encryption.ValueString())
	}
	if !data.Key.IsNull() && !data.Key.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"key.key", data.Key.ValueString())
	}
	return body
}

func (data *TACACSServer) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(prefix + "address.ipv4"); value.Exists() && !data.AddressIpv4.IsNull() {
		data.AddressIpv4 = types.StringValue(value.String())
	} else {
		data.AddressIpv4 = types.StringNull()
	}
	if value := res.Get(prefix + "timeout"); value.Exists() && !data.Timeout.IsNull() {
		data.Timeout = types.Int64Value(value.Int())
	} else {
		data.Timeout = types.Int64Null()
	}
	if value := res.Get(prefix + "key.encryption"); value.Exists() && !data.Encryption.IsNull() {
		data.Encryption = types.StringValue(value.String())
	} else {
		data.Encryption = types.StringNull()
	}
}

func (data *TACACSServer) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "address.ipv4"); value.Exists() {
		data.AddressIpv4 = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "timeout"); value.Exists() {
		data.Timeout = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "key.encryption"); value.Exists() {
		data.Encryption = types.StringValue(value.String())
	}
}

func (data *TACACSServerData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "address.ipv4"); value.Exists() {
		data.AddressIpv4 = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "timeout"); value.Exists() {
		data.Timeout = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "key.encryption"); value.Exists() {
		data.Encryption = types.StringValue(value.String())
	}
}

func (data *TACACSServer) getDeletedItems(ctx context.Context, state TACACSServer) []string {
	deletedItems := make([]string, 0)
	if !state.AddressIpv4.IsNull() && data.AddressIpv4.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/address/ipv4", state.getPath()))
	}
	if !state.Timeout.IsNull() && data.Timeout.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/timeout", state.getPath()))
	}
	if !state.Encryption.IsNull() && data.Encryption.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/key/encryption", state.getPath()))
	}
	if !state.Key.IsNull() && data.Key.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/key/key", state.getPath()))
	}
	return deletedItems
}

func (data *TACACSServer) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	return emptyLeafsDelete
}

func (data *TACACSServer) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.AddressIpv4.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/address/ipv4", data.getPath()))
	}
	if !data.Timeout.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/timeout", data.getPath()))
	}
	if !data.Encryption.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/key/encryption", data.getPath()))
	}
	if !data.Key.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/key/key", data.getPath()))
	}
	return deletePaths
}

func (data *TACACSServer) getIdsFromPath() {
	reString := strings.ReplaceAll("Cisco-IOS-XE-native:native/tacacs/Cisco-IOS-XE-aaa:server=%v", "%s", "(.+)")
	reString = strings.ReplaceAll(reString, "%v", "(.+)")
	re := regexp.MustCompile(reString)
	matches := re.FindStringSubmatch(data.Id.ValueString())
	data.Name = types.StringValue(helpers.Must(url.QueryUnescape(matches[1])))
}
