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

type MDTSubscription struct {
	Device               types.String               `tfsdk:"device"`
	Id                   types.String               `tfsdk:"id"`
	SubscriptionId       types.Int64                `tfsdk:"subscription_id"`
	Stream               types.String               `tfsdk:"stream"`
	Encoding             types.String               `tfsdk:"encoding"`
	SourceVrf            types.String               `tfsdk:"source_vrf"`
	SourceAddress        types.String               `tfsdk:"source_address"`
	UpdatePolicyPeriodic types.Int64                `tfsdk:"update_policy_periodic"`
	UpdatePolicyOnChange types.Bool                 `tfsdk:"update_policy_on_change"`
	FilterXpath          types.String               `tfsdk:"filter_xpath"`
	Receivers            []MDTSubscriptionReceivers `tfsdk:"receivers"`
}

type MDTSubscriptionData struct {
	Device               types.String               `tfsdk:"device"`
	Id                   types.String               `tfsdk:"id"`
	SubscriptionId       types.Int64                `tfsdk:"subscription_id"`
	Stream               types.String               `tfsdk:"stream"`
	Encoding             types.String               `tfsdk:"encoding"`
	SourceVrf            types.String               `tfsdk:"source_vrf"`
	SourceAddress        types.String               `tfsdk:"source_address"`
	UpdatePolicyPeriodic types.Int64                `tfsdk:"update_policy_periodic"`
	UpdatePolicyOnChange types.Bool                 `tfsdk:"update_policy_on_change"`
	FilterXpath          types.String               `tfsdk:"filter_xpath"`
	Receivers            []MDTSubscriptionReceivers `tfsdk:"receivers"`
}
type MDTSubscriptionReceivers struct {
	Address  types.String `tfsdk:"address"`
	Port     types.Int64  `tfsdk:"port"`
	Protocol types.String `tfsdk:"protocol"`
}

func (data MDTSubscription) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-mdt-cfg:mdt-config-data/mdt-subscription=%s", url.QueryEscape(fmt.Sprintf("%v", data.SubscriptionId.ValueInt64())))
}

func (data MDTSubscriptionData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-mdt-cfg:mdt-config-data/mdt-subscription=%s", url.QueryEscape(fmt.Sprintf("%v", data.SubscriptionId.ValueInt64())))
}

// if last path element has a key -> remove it
func (data MDTSubscription) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data MDTSubscription) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.SubscriptionId.IsNull() && !data.SubscriptionId.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"subscription-id", strconv.FormatInt(data.SubscriptionId.ValueInt64(), 10))
	}
	if !data.Stream.IsNull() && !data.Stream.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"base.stream", data.Stream.ValueString())
	}
	if !data.Encoding.IsNull() && !data.Encoding.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"base.encoding", data.Encoding.ValueString())
	}
	if !data.SourceVrf.IsNull() && !data.SourceVrf.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"base.source-vrf", data.SourceVrf.ValueString())
	}
	if !data.SourceAddress.IsNull() && !data.SourceAddress.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"base.source-address", data.SourceAddress.ValueString())
	}
	if !data.UpdatePolicyPeriodic.IsNull() && !data.UpdatePolicyPeriodic.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"base.period", strconv.FormatInt(data.UpdatePolicyPeriodic.ValueInt64(), 10))
	}
	if !data.UpdatePolicyOnChange.IsNull() && !data.UpdatePolicyOnChange.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"base.no-synch-on-start", data.UpdatePolicyOnChange.ValueBool())
	}
	if !data.FilterXpath.IsNull() && !data.FilterXpath.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"base.xpath", data.FilterXpath.ValueString())
	}
	if len(data.Receivers) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"mdt-receivers", []interface{}{})
		for index, item := range data.Receivers {
			if !item.Address.IsNull() && !item.Address.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"mdt-receivers"+"."+strconv.Itoa(index)+"."+"address", item.Address.ValueString())
			}
			if !item.Port.IsNull() && !item.Port.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"mdt-receivers"+"."+strconv.Itoa(index)+"."+"port", strconv.FormatInt(item.Port.ValueInt64(), 10))
			}
			if !item.Protocol.IsNull() && !item.Protocol.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"mdt-receivers"+"."+strconv.Itoa(index)+"."+"protocol", item.Protocol.ValueString())
			}
		}
	}
	return body
}

func (data *MDTSubscription) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "subscription-id"); value.Exists() && !data.SubscriptionId.IsNull() {
		data.SubscriptionId = types.Int64Value(value.Int())
	} else {
		data.SubscriptionId = types.Int64Null()
	}
	if value := res.Get(prefix + "base.stream"); value.Exists() && !data.Stream.IsNull() {
		data.Stream = types.StringValue(value.String())
	} else {
		data.Stream = types.StringNull()
	}
	if value := res.Get(prefix + "base.encoding"); value.Exists() && !data.Encoding.IsNull() {
		data.Encoding = types.StringValue(value.String())
	} else {
		data.Encoding = types.StringNull()
	}
	if value := res.Get(prefix + "base.source-vrf"); value.Exists() && !data.SourceVrf.IsNull() {
		data.SourceVrf = types.StringValue(value.String())
	} else {
		data.SourceVrf = types.StringNull()
	}
	if value := res.Get(prefix + "base.source-address"); value.Exists() && !data.SourceAddress.IsNull() {
		data.SourceAddress = types.StringValue(value.String())
	} else {
		data.SourceAddress = types.StringNull()
	}
	if value := res.Get(prefix + "base.period"); value.Exists() && !data.UpdatePolicyPeriodic.IsNull() {
		data.UpdatePolicyPeriodic = types.Int64Value(value.Int())
	} else {
		data.UpdatePolicyPeriodic = types.Int64Null()
	}
	if value := res.Get(prefix + "base.no-synch-on-start"); !data.UpdatePolicyOnChange.IsNull() {
		if value.Exists() {
			data.UpdatePolicyOnChange = types.BoolValue(value.Bool())
		}
	} else {
		data.UpdatePolicyOnChange = types.BoolNull()
	}
	if value := res.Get(prefix + "base.xpath"); value.Exists() && !data.FilterXpath.IsNull() {
		data.FilterXpath = types.StringValue(value.String())
	} else {
		data.FilterXpath = types.StringNull()
	}
	for i := range data.Receivers {
		keys := [...]string{"address", "port"}
		keyValues := [...]string{data.Receivers[i].Address.ValueString(), strconv.FormatInt(data.Receivers[i].Port.ValueInt64(), 10)}

		var r gjson.Result
		res.Get(prefix + "mdt-receivers").ForEach(
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
		if value := r.Get("address"); value.Exists() && !data.Receivers[i].Address.IsNull() {
			data.Receivers[i].Address = types.StringValue(value.String())
		} else {
			data.Receivers[i].Address = types.StringNull()
		}
		if value := r.Get("port"); value.Exists() && !data.Receivers[i].Port.IsNull() {
			data.Receivers[i].Port = types.Int64Value(value.Int())
		} else {
			data.Receivers[i].Port = types.Int64Null()
		}
		if value := r.Get("protocol"); value.Exists() && !data.Receivers[i].Protocol.IsNull() {
			data.Receivers[i].Protocol = types.StringValue(value.String())
		} else {
			data.Receivers[i].Protocol = types.StringNull()
		}
	}
}

func (data *MDTSubscriptionData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "base.stream"); value.Exists() {
		data.Stream = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "base.encoding"); value.Exists() {
		data.Encoding = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "base.source-vrf"); value.Exists() {
		data.SourceVrf = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "base.source-address"); value.Exists() {
		data.SourceAddress = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "base.period"); value.Exists() {
		data.UpdatePolicyPeriodic = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "base.no-synch-on-start"); value.Exists() {
		data.UpdatePolicyOnChange = types.BoolValue(value.Bool())
	} else {
		data.UpdatePolicyOnChange = types.BoolValue(false)
	}
	if value := res.Get(prefix + "base.xpath"); value.Exists() {
		data.FilterXpath = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "mdt-receivers"); value.Exists() {
		data.Receivers = make([]MDTSubscriptionReceivers, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := MDTSubscriptionReceivers{}
			if cValue := v.Get("address"); cValue.Exists() {
				item.Address = types.StringValue(cValue.String())
			}
			if cValue := v.Get("port"); cValue.Exists() {
				item.Port = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("protocol"); cValue.Exists() {
				item.Protocol = types.StringValue(cValue.String())
			}
			data.Receivers = append(data.Receivers, item)
			return true
		})
	}
}

func (data *MDTSubscription) getDeletedItems(ctx context.Context, state MDTSubscription) []string {
	deletedItems := make([]string, 0)
	for i := range state.Receivers {
		stateKeyValues := [...]string{state.Receivers[i].Address.ValueString(), strconv.FormatInt(state.Receivers[i].Port.ValueInt64(), 10)}

		emptyKeys := true
		if !reflect.ValueOf(state.Receivers[i].Address.ValueString()).IsZero() {
			emptyKeys = false
		}
		if !reflect.ValueOf(state.Receivers[i].Port.ValueInt64()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Receivers {
			found = true
			if state.Receivers[i].Address.ValueString() != data.Receivers[j].Address.ValueString() {
				found = false
			}
			if state.Receivers[i].Port.ValueInt64() != data.Receivers[j].Port.ValueInt64() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			deletedItems = append(deletedItems, fmt.Sprintf("%v/mdt-receivers=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	return deletedItems
}

func (data *MDTSubscription) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)

	return emptyLeafsDelete
}

func (data *MDTSubscription) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.Stream.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/base/stream", data.getPath()))
	}
	if !data.Encoding.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/base/encoding", data.getPath()))
	}
	if !data.SourceVrf.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/base/source-vrf", data.getPath()))
	}
	if !data.SourceAddress.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/base/source-address", data.getPath()))
	}
	if !data.UpdatePolicyPeriodic.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/base/period", data.getPath()))
	}
	if !data.UpdatePolicyOnChange.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/base/no-synch-on-start", data.getPath()))
	}
	if !data.FilterXpath.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/base/xpath", data.getPath()))
	}
	for i := range data.Receivers {
		keyValues := [...]string{data.Receivers[i].Address.ValueString(), strconv.FormatInt(data.Receivers[i].Port.ValueInt64(), 10)}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/mdt-receivers=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	return deletePaths
}
