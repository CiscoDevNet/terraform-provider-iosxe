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
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type MSDPVRF struct {
	Device       types.String       `tfsdk:"device"`
	Id           types.String       `tfsdk:"id"`
	DeleteMode   types.String       `tfsdk:"delete_mode"`
	Vrf          types.String       `tfsdk:"vrf"`
	OriginatorId types.String       `tfsdk:"originator_id"`
	Passwords    []MSDPVRFPasswords `tfsdk:"passwords"`
	Peers        []MSDPVRFPeers     `tfsdk:"peers"`
}

type MSDPVRFData struct {
	Device       types.String       `tfsdk:"device"`
	Id           types.String       `tfsdk:"id"`
	Vrf          types.String       `tfsdk:"vrf"`
	OriginatorId types.String       `tfsdk:"originator_id"`
	Passwords    []MSDPVRFPasswords `tfsdk:"passwords"`
	Peers        []MSDPVRFPeers     `tfsdk:"peers"`
}
type MSDPVRFPasswords struct {
	Addr       types.String `tfsdk:"addr"`
	Encryption types.Int64  `tfsdk:"encryption"`
	Password   types.String `tfsdk:"password"`
}
type MSDPVRFPeers struct {
	Addr                  types.String `tfsdk:"addr"`
	RemoteAs              types.Int64  `tfsdk:"remote_as"`
	ConnectSourceLoopback types.Int64  `tfsdk:"connect_source_loopback"`
}

func (data MSDPVRF) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/ip/Cisco-IOS-XE-multicast:msdp/vrf=%s", url.QueryEscape(fmt.Sprintf("%v", data.Vrf.ValueString())))
}

func (data MSDPVRFData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/ip/Cisco-IOS-XE-multicast:msdp/vrf=%s", url.QueryEscape(fmt.Sprintf("%v", data.Vrf.ValueString())))
}

// if last path element has a key -> remove it
func (data MSDPVRF) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data MSDPVRF) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Vrf.IsNull() && !data.Vrf.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"name", data.Vrf.ValueString())
	}
	if !data.OriginatorId.IsNull() && !data.OriginatorId.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"originator-id", data.OriginatorId.ValueString())
	}
	if len(data.Passwords) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"password.peer-list", []interface{}{})
		for index, item := range data.Passwords {
			if !item.Addr.IsNull() && !item.Addr.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"password.peer-list"+"."+strconv.Itoa(index)+"."+"addr", item.Addr.ValueString())
			}
			if !item.Encryption.IsNull() && !item.Encryption.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"password.peer-list"+"."+strconv.Itoa(index)+"."+"encryption", strconv.FormatInt(item.Encryption.ValueInt64(), 10))
			}
			if !item.Password.IsNull() && !item.Password.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"password.peer-list"+"."+strconv.Itoa(index)+"."+"password", item.Password.ValueString())
			}
		}
	}
	if len(data.Peers) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"peer", []interface{}{})
		for index, item := range data.Peers {
			if !item.Addr.IsNull() && !item.Addr.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"peer"+"."+strconv.Itoa(index)+"."+"addr", item.Addr.ValueString())
			}
			if !item.RemoteAs.IsNull() && !item.RemoteAs.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"peer"+"."+strconv.Itoa(index)+"."+"remote-as", strconv.FormatInt(item.RemoteAs.ValueInt64(), 10))
			}
			if !item.ConnectSourceLoopback.IsNull() && !item.ConnectSourceLoopback.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"peer"+"."+strconv.Itoa(index)+"."+"connect-source.Loopback", strconv.FormatInt(item.ConnectSourceLoopback.ValueInt64(), 10))
			}
		}
	}
	return body
}

func (data *MSDPVRF) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "name"); value.Exists() && !data.Vrf.IsNull() {
		data.Vrf = types.StringValue(value.String())
	} else {
		data.Vrf = types.StringNull()
	}
	if value := res.Get(prefix + "originator-id"); value.Exists() && !data.OriginatorId.IsNull() {
		data.OriginatorId = types.StringValue(value.String())
	} else {
		data.OriginatorId = types.StringNull()
	}
	for i := range data.Passwords {
		keys := [...]string{"addr"}
		keyValues := [...]string{data.Passwords[i].Addr.ValueString()}

		var r gjson.Result
		res.Get(prefix + "password.peer-list").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("addr"); value.Exists() && !data.Passwords[i].Addr.IsNull() {
			data.Passwords[i].Addr = types.StringValue(value.String())
		} else {
			data.Passwords[i].Addr = types.StringNull()
		}
		if value := r.Get("encryption"); value.Exists() && !data.Passwords[i].Encryption.IsNull() {
			data.Passwords[i].Encryption = types.Int64Value(value.Int())
		} else {
			data.Passwords[i].Encryption = types.Int64Null()
		}
		if value := r.Get("password"); value.Exists() && !data.Passwords[i].Password.IsNull() {
			data.Passwords[i].Password = types.StringValue(value.String())
		} else {
			data.Passwords[i].Password = types.StringNull()
		}
	}
	for i := range data.Peers {
		keys := [...]string{"addr"}
		keyValues := [...]string{data.Peers[i].Addr.ValueString()}

		var r gjson.Result
		res.Get(prefix + "peer").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("addr"); value.Exists() && !data.Peers[i].Addr.IsNull() {
			data.Peers[i].Addr = types.StringValue(value.String())
		} else {
			data.Peers[i].Addr = types.StringNull()
		}
		if value := r.Get("remote-as"); value.Exists() && !data.Peers[i].RemoteAs.IsNull() {
			data.Peers[i].RemoteAs = types.Int64Value(value.Int())
		} else {
			data.Peers[i].RemoteAs = types.Int64Null()
		}
		if value := r.Get("connect-source.Loopback"); value.Exists() && !data.Peers[i].ConnectSourceLoopback.IsNull() {
			data.Peers[i].ConnectSourceLoopback = types.Int64Value(value.Int())
		} else {
			data.Peers[i].ConnectSourceLoopback = types.Int64Null()
		}
	}
}

func (data *MSDPVRF) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "originator-id"); value.Exists() {
		data.OriginatorId = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "password.peer-list"); value.Exists() {
		data.Passwords = make([]MSDPVRFPasswords, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := MSDPVRFPasswords{}
			if cValue := v.Get("addr"); cValue.Exists() {
				item.Addr = types.StringValue(cValue.String())
			}
			if cValue := v.Get("encryption"); cValue.Exists() {
				item.Encryption = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("password"); cValue.Exists() {
				item.Password = types.StringValue(cValue.String())
			}
			data.Passwords = append(data.Passwords, item)
			return true
		})
	}
	if value := res.Get(prefix + "peer"); value.Exists() {
		data.Peers = make([]MSDPVRFPeers, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := MSDPVRFPeers{}
			if cValue := v.Get("addr"); cValue.Exists() {
				item.Addr = types.StringValue(cValue.String())
			}
			if cValue := v.Get("remote-as"); cValue.Exists() {
				item.RemoteAs = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("connect-source.Loopback"); cValue.Exists() {
				item.ConnectSourceLoopback = types.Int64Value(cValue.Int())
			}
			data.Peers = append(data.Peers, item)
			return true
		})
	}
}

func (data *MSDPVRFData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "originator-id"); value.Exists() {
		data.OriginatorId = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "password.peer-list"); value.Exists() {
		data.Passwords = make([]MSDPVRFPasswords, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := MSDPVRFPasswords{}
			if cValue := v.Get("addr"); cValue.Exists() {
				item.Addr = types.StringValue(cValue.String())
			}
			if cValue := v.Get("encryption"); cValue.Exists() {
				item.Encryption = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("password"); cValue.Exists() {
				item.Password = types.StringValue(cValue.String())
			}
			data.Passwords = append(data.Passwords, item)
			return true
		})
	}
	if value := res.Get(prefix + "peer"); value.Exists() {
		data.Peers = make([]MSDPVRFPeers, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := MSDPVRFPeers{}
			if cValue := v.Get("addr"); cValue.Exists() {
				item.Addr = types.StringValue(cValue.String())
			}
			if cValue := v.Get("remote-as"); cValue.Exists() {
				item.RemoteAs = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("connect-source.Loopback"); cValue.Exists() {
				item.ConnectSourceLoopback = types.Int64Value(cValue.Int())
			}
			data.Peers = append(data.Peers, item)
			return true
		})
	}
}

func (data *MSDPVRF) getDeletedItems(ctx context.Context, state MSDPVRF) []string {
	deletedItems := make([]string, 0)
	if !state.OriginatorId.IsNull() && data.OriginatorId.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/originator-id", state.getPath()))
	}
	for i := range state.Passwords {
		stateKeyValues := [...]string{state.Passwords[i].Addr.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.Passwords[i].Addr.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Passwords {
			found = true
			if state.Passwords[i].Addr.ValueString() != data.Passwords[j].Addr.ValueString() {
				found = false
			}
			if found {
				if !state.Passwords[i].Encryption.IsNull() && data.Passwords[j].Encryption.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/password/peer-list=%v/encryption", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				if !state.Passwords[i].Password.IsNull() && data.Passwords[j].Password.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/password/peer-list=%v/password", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				break
			}
		}
		if !found {
			deletedItems = append(deletedItems, fmt.Sprintf("%v/password/peer-list=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	for i := range state.Peers {
		stateKeyValues := [...]string{state.Peers[i].Addr.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.Peers[i].Addr.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Peers {
			found = true
			if state.Peers[i].Addr.ValueString() != data.Peers[j].Addr.ValueString() {
				found = false
			}
			if found {
				if !state.Peers[i].RemoteAs.IsNull() && data.Peers[j].RemoteAs.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/peer=%v/remote-as", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				if !state.Peers[i].ConnectSourceLoopback.IsNull() && data.Peers[j].ConnectSourceLoopback.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/peer=%v/connect-source/Loopback", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				break
			}
		}
		if !found {
			deletedItems = append(deletedItems, fmt.Sprintf("%v/peer=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	return deletedItems
}

func (data *MSDPVRF) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)

	return emptyLeafsDelete
}

func (data *MSDPVRF) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.OriginatorId.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/originator-id", data.getPath()))
	}
	for i := range data.Passwords {
		keyValues := [...]string{data.Passwords[i].Addr.ValueString()}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/password/peer-list=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	return deletePaths
}
