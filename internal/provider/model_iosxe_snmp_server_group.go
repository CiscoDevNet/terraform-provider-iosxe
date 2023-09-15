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

type SNMPServerGroup struct {
	Device     types.String                `tfsdk:"device"`
	Id         types.String                `tfsdk:"id"`
	Name       types.String                `tfsdk:"name"`
	V3Security []SNMPServerGroupV3Security `tfsdk:"v3_security"`
}

type SNMPServerGroupData struct {
	Device     types.String                `tfsdk:"device"`
	Id         types.String                `tfsdk:"id"`
	Name       types.String                `tfsdk:"name"`
	V3Security []SNMPServerGroupV3Security `tfsdk:"v3_security"`
}
type SNMPServerGroupV3Security struct {
	SecurityLevel     types.String `tfsdk:"security_level"`
	ContextNode       types.String `tfsdk:"context_node"`
	MatchNode         types.String `tfsdk:"match_node"`
	ReadNode          types.String `tfsdk:"read_node"`
	WriteNode         types.String `tfsdk:"write_node"`
	NotifyNode        types.String `tfsdk:"notify_node"`
	AccessIpv6Acl     types.String `tfsdk:"access_ipv6_acl"`
	AccessStandardAcl types.Int64  `tfsdk:"access_standard_acl"`
	AccessAclName     types.String `tfsdk:"access_acl_name"`
}

func (data SNMPServerGroup) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/snmp-server/Cisco-IOS-XE-snmp:group=%s", url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())))
}

func (data SNMPServerGroupData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/snmp-server/Cisco-IOS-XE-snmp:group=%s", url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())))
}

// if last path element has a key -> remove it
func (data SNMPServerGroup) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data SNMPServerGroup) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"id", data.Name.ValueString())
	}
	if len(data.V3Security) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"v3.security-level-list", []interface{}{})
		for index, item := range data.V3Security {
			if !item.SecurityLevel.IsNull() && !item.SecurityLevel.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"v3.security-level-list"+"."+strconv.Itoa(index)+"."+"security-level", item.SecurityLevel.ValueString())
			}
			if !item.ContextNode.IsNull() && !item.ContextNode.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"v3.security-level-list"+"."+strconv.Itoa(index)+"."+"context-node", item.ContextNode.ValueString())
			}
			if !item.MatchNode.IsNull() && !item.MatchNode.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"v3.security-level-list"+"."+strconv.Itoa(index)+"."+"match-node", item.MatchNode.ValueString())
			}
			if !item.ReadNode.IsNull() && !item.ReadNode.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"v3.security-level-list"+"."+strconv.Itoa(index)+"."+"read-node", item.ReadNode.ValueString())
			}
			if !item.WriteNode.IsNull() && !item.WriteNode.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"v3.security-level-list"+"."+strconv.Itoa(index)+"."+"write-node", item.WriteNode.ValueString())
			}
			if !item.NotifyNode.IsNull() && !item.NotifyNode.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"v3.security-level-list"+"."+strconv.Itoa(index)+"."+"notify-node", item.NotifyNode.ValueString())
			}
			if !item.AccessIpv6Acl.IsNull() && !item.AccessIpv6Acl.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"v3.security-level-list"+"."+strconv.Itoa(index)+"."+"access-config.ipv6-acl", item.AccessIpv6Acl.ValueString())
			}
			if !item.AccessStandardAcl.IsNull() && !item.AccessStandardAcl.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"v3.security-level-list"+"."+strconv.Itoa(index)+"."+"access-config.standard-acl", strconv.FormatInt(item.AccessStandardAcl.ValueInt64(), 10))
			}
			if !item.AccessAclName.IsNull() && !item.AccessAclName.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"v3.security-level-list"+"."+strconv.Itoa(index)+"."+"access-config.acl-name", item.AccessAclName.ValueString())
			}
		}
	}
	return body
}

func (data *SNMPServerGroup) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "id"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	for i := range data.V3Security {
		keys := [...]string{"security-level"}
		keyValues := [...]string{data.V3Security[i].SecurityLevel.ValueString()}

		var r gjson.Result
		res.Get(prefix + "v3.security-level-list").ForEach(
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
		if value := r.Get("security-level"); value.Exists() && !data.V3Security[i].SecurityLevel.IsNull() {
			data.V3Security[i].SecurityLevel = types.StringValue(value.String())
		} else {
			data.V3Security[i].SecurityLevel = types.StringNull()
		}
		if value := r.Get("context-node"); value.Exists() && !data.V3Security[i].ContextNode.IsNull() {
			data.V3Security[i].ContextNode = types.StringValue(value.String())
		} else {
			data.V3Security[i].ContextNode = types.StringNull()
		}
		if value := r.Get("match-node"); value.Exists() && !data.V3Security[i].MatchNode.IsNull() {
			data.V3Security[i].MatchNode = types.StringValue(value.String())
		} else {
			data.V3Security[i].MatchNode = types.StringNull()
		}
		if value := r.Get("read-node"); value.Exists() && !data.V3Security[i].ReadNode.IsNull() {
			data.V3Security[i].ReadNode = types.StringValue(value.String())
		} else {
			data.V3Security[i].ReadNode = types.StringNull()
		}
		if value := r.Get("write-node"); value.Exists() && !data.V3Security[i].WriteNode.IsNull() {
			data.V3Security[i].WriteNode = types.StringValue(value.String())
		} else {
			data.V3Security[i].WriteNode = types.StringNull()
		}
		if value := r.Get("notify-node"); value.Exists() && !data.V3Security[i].NotifyNode.IsNull() {
			data.V3Security[i].NotifyNode = types.StringValue(value.String())
		} else {
			data.V3Security[i].NotifyNode = types.StringNull()
		}
		if value := r.Get("access-config.ipv6-acl"); value.Exists() && !data.V3Security[i].AccessIpv6Acl.IsNull() {
			data.V3Security[i].AccessIpv6Acl = types.StringValue(value.String())
		} else {
			data.V3Security[i].AccessIpv6Acl = types.StringNull()
		}
		if value := r.Get("access-config.standard-acl"); value.Exists() && !data.V3Security[i].AccessStandardAcl.IsNull() {
			data.V3Security[i].AccessStandardAcl = types.Int64Value(value.Int())
		} else {
			data.V3Security[i].AccessStandardAcl = types.Int64Null()
		}
		if value := r.Get("access-config.acl-name"); value.Exists() && !data.V3Security[i].AccessAclName.IsNull() {
			data.V3Security[i].AccessAclName = types.StringValue(value.String())
		} else {
			data.V3Security[i].AccessAclName = types.StringNull()
		}
	}
}

func (data *SNMPServerGroupData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "v3.security-level-list"); value.Exists() {
		data.V3Security = make([]SNMPServerGroupV3Security, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SNMPServerGroupV3Security{}
			if cValue := v.Get("security-level"); cValue.Exists() {
				item.SecurityLevel = types.StringValue(cValue.String())
			}
			if cValue := v.Get("context-node"); cValue.Exists() {
				item.ContextNode = types.StringValue(cValue.String())
			}
			if cValue := v.Get("match-node"); cValue.Exists() {
				item.MatchNode = types.StringValue(cValue.String())
			}
			if cValue := v.Get("read-node"); cValue.Exists() {
				item.ReadNode = types.StringValue(cValue.String())
			}
			if cValue := v.Get("write-node"); cValue.Exists() {
				item.WriteNode = types.StringValue(cValue.String())
			}
			if cValue := v.Get("notify-node"); cValue.Exists() {
				item.NotifyNode = types.StringValue(cValue.String())
			}
			if cValue := v.Get("access-config.ipv6-acl"); cValue.Exists() {
				item.AccessIpv6Acl = types.StringValue(cValue.String())
			}
			if cValue := v.Get("access-config.standard-acl"); cValue.Exists() {
				item.AccessStandardAcl = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("access-config.acl-name"); cValue.Exists() {
				item.AccessAclName = types.StringValue(cValue.String())
			}
			data.V3Security = append(data.V3Security, item)
			return true
		})
	}
}

func (data *SNMPServerGroup) getDeletedItems(ctx context.Context, state SNMPServerGroup) []string {
	deletedItems := make([]string, 0)
	for i := range state.V3Security {
		stateKeyValues := [...]string{state.V3Security[i].SecurityLevel.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.V3Security[i].SecurityLevel.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.V3Security {
			found = true
			if state.V3Security[i].SecurityLevel.ValueString() != data.V3Security[j].SecurityLevel.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			deletedItems = append(deletedItems, fmt.Sprintf("%v/v3/security-level-list=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	return deletedItems
}

func (data *SNMPServerGroup) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)

	return emptyLeafsDelete
}

func (data *SNMPServerGroup) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	for i := range data.V3Security {
		keyValues := [...]string{data.V3Security[i].SecurityLevel.ValueString()}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/v3/security-level-list=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	return deletePaths
}
