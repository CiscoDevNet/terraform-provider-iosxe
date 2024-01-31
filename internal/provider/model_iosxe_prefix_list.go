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
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type PrefixList struct {
	Device                types.String                      `tfsdk:"device"`
	Id                    types.String                      `tfsdk:"id"`
	Prefixes              []PrefixListPrefixes              `tfsdk:"prefixes"`
	PrefixListDescription []PrefixListPrefixListDescription `tfsdk:"prefix_list_description"`
}

type PrefixListData struct {
	Device                types.String                      `tfsdk:"device"`
	Id                    types.String                      `tfsdk:"id"`
	Prefixes              []PrefixListPrefixes              `tfsdk:"prefixes"`
	PrefixListDescription []PrefixListPrefixListDescription `tfsdk:"prefix_list_description"`
}
type PrefixListPrefixes struct {
	Name   types.String `tfsdk:"name"`
	Seq    types.Int64  `tfsdk:"seq"`
	Action types.String `tfsdk:"action"`
	Ip     types.String `tfsdk:"ip"`
	Ge     types.Int64  `tfsdk:"ge"`
	Le     types.Int64  `tfsdk:"le"`
}
type PrefixListPrefixListDescription struct {
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
}

func (data PrefixList) getPath() string {
	return "Cisco-IOS-XE-native:native/ip/prefix-lists"
}

func (data PrefixListData) getPath() string {
	return "Cisco-IOS-XE-native:native/ip/prefix-lists"
}

// if last path element has a key -> remove it
func (data PrefixList) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data PrefixList) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if len(data.Prefixes) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"prefixes", []interface{}{})
		for index, item := range data.Prefixes {
			if !item.Name.IsNull() && !item.Name.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"prefixes"+"."+strconv.Itoa(index)+"."+"name", item.Name.ValueString())
			}
			if !item.Seq.IsNull() && !item.Seq.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"prefixes"+"."+strconv.Itoa(index)+"."+"no", strconv.FormatInt(item.Seq.ValueInt64(), 10))
			}
			if !item.Action.IsNull() && !item.Action.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"prefixes"+"."+strconv.Itoa(index)+"."+"action", item.Action.ValueString())
			}
			if !item.Ip.IsNull() && !item.Ip.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"prefixes"+"."+strconv.Itoa(index)+"."+"ip", item.Ip.ValueString())
			}
			if !item.Ge.IsNull() && !item.Ge.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"prefixes"+"."+strconv.Itoa(index)+"."+"ge", strconv.FormatInt(item.Ge.ValueInt64(), 10))
			}
			if !item.Le.IsNull() && !item.Le.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"prefixes"+"."+strconv.Itoa(index)+"."+"le", strconv.FormatInt(item.Le.ValueInt64(), 10))
			}
		}
	}
	if len(data.PrefixListDescription) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"prefix-list-description", []interface{}{})
		for index, item := range data.PrefixListDescription {
			if !item.Name.IsNull() && !item.Name.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"prefix-list-description"+"."+strconv.Itoa(index)+"."+"name", item.Name.ValueString())
			}
			if !item.Description.IsNull() && !item.Description.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"prefix-list-description"+"."+strconv.Itoa(index)+"."+"description", item.Description.ValueString())
			}
		}
	}
	return body
}

func (data *PrefixList) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	for i := range data.Prefixes {
		keys := [...]string{"name", "no"}
		keyValues := [...]string{data.Prefixes[i].Name.ValueString(), strconv.FormatInt(data.Prefixes[i].Seq.ValueInt64(), 10)}

		var r gjson.Result
		res.Get(prefix + "prefixes").ForEach(
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
		if value := r.Get("name"); value.Exists() && !data.Prefixes[i].Name.IsNull() {
			data.Prefixes[i].Name = types.StringValue(value.String())
		} else {
			data.Prefixes[i].Name = types.StringNull()
		}
		if value := r.Get("no"); value.Exists() && !data.Prefixes[i].Seq.IsNull() {
			data.Prefixes[i].Seq = types.Int64Value(value.Int())
		} else {
			data.Prefixes[i].Seq = types.Int64Null()
		}
		if value := r.Get("action"); value.Exists() && !data.Prefixes[i].Action.IsNull() {
			data.Prefixes[i].Action = types.StringValue(value.String())
		} else {
			data.Prefixes[i].Action = types.StringNull()
		}
		if value := r.Get("ip"); value.Exists() && !data.Prefixes[i].Ip.IsNull() {
			data.Prefixes[i].Ip = types.StringValue(value.String())
		} else {
			data.Prefixes[i].Ip = types.StringNull()
		}
		if value := r.Get("ge"); value.Exists() && !data.Prefixes[i].Ge.IsNull() {
			data.Prefixes[i].Ge = types.Int64Value(value.Int())
		} else {
			data.Prefixes[i].Ge = types.Int64Null()
		}
		if value := r.Get("le"); value.Exists() && !data.Prefixes[i].Le.IsNull() {
			data.Prefixes[i].Le = types.Int64Value(value.Int())
		} else {
			data.Prefixes[i].Le = types.Int64Null()
		}
	}
	for i := range data.PrefixListDescription {
		keys := [...]string{"name"}
		keyValues := [...]string{data.PrefixListDescription[i].Name.ValueString()}

		var r gjson.Result
		res.Get(prefix + "prefix-list-description").ForEach(
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
		if value := r.Get("name"); value.Exists() && !data.PrefixListDescription[i].Name.IsNull() {
			data.PrefixListDescription[i].Name = types.StringValue(value.String())
		} else {
			data.PrefixListDescription[i].Name = types.StringNull()
		}
		if value := r.Get("description"); value.Exists() && !data.PrefixListDescription[i].Description.IsNull() {
			data.PrefixListDescription[i].Description = types.StringValue(value.String())
		} else {
			data.PrefixListDescription[i].Description = types.StringNull()
		}
	}
}

func (data *PrefixListData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "prefixes"); value.Exists() {
		data.Prefixes = make([]PrefixListPrefixes, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := PrefixListPrefixes{}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			}
			if cValue := v.Get("no"); cValue.Exists() {
				item.Seq = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("action"); cValue.Exists() {
				item.Action = types.StringValue(cValue.String())
			}
			if cValue := v.Get("ip"); cValue.Exists() {
				item.Ip = types.StringValue(cValue.String())
			}
			if cValue := v.Get("ge"); cValue.Exists() {
				item.Ge = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("le"); cValue.Exists() {
				item.Le = types.Int64Value(cValue.Int())
			}
			data.Prefixes = append(data.Prefixes, item)
			return true
		})
	}
	if value := res.Get(prefix + "prefix-list-description"); value.Exists() {
		data.PrefixListDescription = make([]PrefixListPrefixListDescription, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := PrefixListPrefixListDescription{}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			}
			if cValue := v.Get("description"); cValue.Exists() {
				item.Description = types.StringValue(cValue.String())
			}
			data.PrefixListDescription = append(data.PrefixListDescription, item)
			return true
		})
	}
}

func (data *PrefixList) getDeletedItems(ctx context.Context, state PrefixList) []string {
	deletedItems := make([]string, 0)
	for i := range state.Prefixes {
		stateKeyValues := [...]string{state.Prefixes[i].Name.ValueString(), strconv.FormatInt(state.Prefixes[i].Seq.ValueInt64(), 10)}

		emptyKeys := true
		if !reflect.ValueOf(state.Prefixes[i].Name.ValueString()).IsZero() {
			emptyKeys = false
		}
		if !reflect.ValueOf(state.Prefixes[i].Seq.ValueInt64()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Prefixes {
			found = true
			if state.Prefixes[i].Name.ValueString() != data.Prefixes[j].Name.ValueString() {
				found = false
			}
			if state.Prefixes[i].Seq.ValueInt64() != data.Prefixes[j].Seq.ValueInt64() {
				found = false
			}
			if found {
				if !state.Prefixes[i].Action.IsNull() && data.Prefixes[j].Action.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/prefixes=%v/action", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				if !state.Prefixes[i].Ip.IsNull() && data.Prefixes[j].Ip.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/prefixes=%v/ip", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				if !state.Prefixes[i].Ge.IsNull() && data.Prefixes[j].Ge.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/prefixes=%v/ge", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				if !state.Prefixes[i].Le.IsNull() && data.Prefixes[j].Le.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/prefixes=%v/le", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				break
			}
		}
		if !found {
			deletedItems = append(deletedItems, fmt.Sprintf("%v/prefixes=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	for i := range state.PrefixListDescription {
		stateKeyValues := [...]string{state.PrefixListDescription[i].Name.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.PrefixListDescription[i].Name.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.PrefixListDescription {
			found = true
			if state.PrefixListDescription[i].Name.ValueString() != data.PrefixListDescription[j].Name.ValueString() {
				found = false
			}
			if found {
				if !state.PrefixListDescription[i].Description.IsNull() && data.PrefixListDescription[j].Description.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/prefix-list-description=%v/description", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				break
			}
		}
		if !found {
			deletedItems = append(deletedItems, fmt.Sprintf("%v/prefix-list-description=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	return deletedItems
}

func (data *PrefixList) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)

	return emptyLeafsDelete
}

func (data *PrefixList) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	for i := range data.Prefixes {
		keyValues := [...]string{data.Prefixes[i].Name.ValueString(), strconv.FormatInt(data.Prefixes[i].Seq.ValueInt64(), 10)}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/prefixes=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	for i := range data.PrefixListDescription {
		keyValues := [...]string{data.PrefixListDescription[i].Name.ValueString()}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/prefix-list-description=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	return deletePaths
}
