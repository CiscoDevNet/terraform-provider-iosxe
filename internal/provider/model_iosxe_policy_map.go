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

type PolicyMap struct {
	Device      types.String       `tfsdk:"device"`
	Id          types.String       `tfsdk:"id"`
	Name        types.String       `tfsdk:"name"`
	Type        types.String       `tfsdk:"type"`
	Subscriber  types.Bool         `tfsdk:"subscriber"`
	Description types.String       `tfsdk:"description"`
	Classes     []PolicyMapClasses `tfsdk:"classes"`
}

type PolicyMapData struct {
	Device      types.String       `tfsdk:"device"`
	Id          types.String       `tfsdk:"id"`
	Name        types.String       `tfsdk:"name"`
	Type        types.String       `tfsdk:"type"`
	Subscriber  types.Bool         `tfsdk:"subscriber"`
	Description types.String       `tfsdk:"description"`
	Classes     []PolicyMapClasses `tfsdk:"classes"`
}
type PolicyMapClasses struct {
	Name    types.String              `tfsdk:"name"`
	Actions []PolicyMapClassesActions `tfsdk:"actions"`
}
type PolicyMapClassesActions struct {
	Type                                 types.String `tfsdk:"type"`
	BandwidthBits                        types.Int64  `tfsdk:"bandwidth_bits"`
	BandwidthPercent                     types.Int64  `tfsdk:"bandwidth_percent"`
	BandwidthRemainingOption             types.String `tfsdk:"bandwidth_remaining_option"`
	BandwidthRemainingPercent            types.Int64  `tfsdk:"bandwidth_remaining_percent"`
	BandwidthRemainingRatio              types.Int64  `tfsdk:"bandwidth_remaining_ratio"`
	PriorityLevel                        types.Int64  `tfsdk:"priority_level"`
	PriorityBurst                        types.Int64  `tfsdk:"priority_burst"`
	QueueLimit                           types.Int64  `tfsdk:"queue_limit"`
	QueueLimitType                       types.String `tfsdk:"queue_limit_type"`
	ShapeAverageBitRate                  types.Int64  `tfsdk:"shape_average_bit_rate"`
	ShapeAverageBitsPerIntervalSustained types.Int64  `tfsdk:"shape_average_bits_per_interval_sustained"`
	ShapeAverageBitsPerIntervalExcess    types.Int64  `tfsdk:"shape_average_bits_per_interval_excess"`
	ShapeAveragePercent                  types.Int64  `tfsdk:"shape_average_percent"`
	ShapeAverageBurstSizeSustained       types.Int64  `tfsdk:"shape_average_burst_size_sustained"`
	ShapeAverageMs                       types.Bool   `tfsdk:"shape_average_ms"`
}

func (data PolicyMap) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/policy/Cisco-IOS-XE-policy:policy-map=%v", url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())))
}

func (data PolicyMapData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/policy/Cisco-IOS-XE-policy:policy-map=%v", url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())))
}

// if last path element has a key -> remove it
func (data PolicyMap) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data PolicyMap) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"name", data.Name.ValueString())
	}
	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"type", data.Type.ValueString())
	}
	if !data.Subscriber.IsNull() && !data.Subscriber.IsUnknown() {
		if data.Subscriber.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"subscriber", map[string]string{})
		}
	}
	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"description", data.Description.ValueString())
	}
	if len(data.Classes) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"class", []interface{}{})
		for index, item := range data.Classes {
			if !item.Name.IsNull() && !item.Name.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"class"+"."+strconv.Itoa(index)+"."+"name", item.Name.ValueString())
			}
			if len(item.Actions) > 0 {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"class"+"."+strconv.Itoa(index)+"."+"action-list", []interface{}{})
				for cindex, citem := range item.Actions {
					if !citem.Type.IsNull() && !citem.Type.IsUnknown() {
						body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"class"+"."+strconv.Itoa(index)+"."+"action-list"+"."+strconv.Itoa(cindex)+"."+"action-type", citem.Type.ValueString())
					}
					if !citem.BandwidthBits.IsNull() && !citem.BandwidthBits.IsUnknown() {
						body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"class"+"."+strconv.Itoa(index)+"."+"action-list"+"."+strconv.Itoa(cindex)+"."+"bandwidth.bits", strconv.FormatInt(citem.BandwidthBits.ValueInt64(), 10))
					}
					if !citem.BandwidthPercent.IsNull() && !citem.BandwidthPercent.IsUnknown() {
						body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"class"+"."+strconv.Itoa(index)+"."+"action-list"+"."+strconv.Itoa(cindex)+"."+"bandwidth.percent", strconv.FormatInt(citem.BandwidthPercent.ValueInt64(), 10))
					}
					if !citem.BandwidthRemainingOption.IsNull() && !citem.BandwidthRemainingOption.IsUnknown() {
						body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"class"+"."+strconv.Itoa(index)+"."+"action-list"+"."+strconv.Itoa(cindex)+"."+"bandwidth.remaining.rem-option", citem.BandwidthRemainingOption.ValueString())
					}
					if !citem.BandwidthRemainingPercent.IsNull() && !citem.BandwidthRemainingPercent.IsUnknown() {
						body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"class"+"."+strconv.Itoa(index)+"."+"action-list"+"."+strconv.Itoa(cindex)+"."+"bandwidth.remaining.percent", strconv.FormatInt(citem.BandwidthRemainingPercent.ValueInt64(), 10))
					}
					if !citem.BandwidthRemainingRatio.IsNull() && !citem.BandwidthRemainingRatio.IsUnknown() {
						body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"class"+"."+strconv.Itoa(index)+"."+"action-list"+"."+strconv.Itoa(cindex)+"."+"bandwidth.remaining.ratio", strconv.FormatInt(citem.BandwidthRemainingRatio.ValueInt64(), 10))
					}
					if !citem.PriorityLevel.IsNull() && !citem.PriorityLevel.IsUnknown() {
						body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"class"+"."+strconv.Itoa(index)+"."+"action-list"+"."+strconv.Itoa(cindex)+"."+"priority.level", strconv.FormatInt(citem.PriorityLevel.ValueInt64(), 10))
					}
					if !citem.PriorityBurst.IsNull() && !citem.PriorityBurst.IsUnknown() {
						body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"class"+"."+strconv.Itoa(index)+"."+"action-list"+"."+strconv.Itoa(cindex)+"."+"priority.burst", strconv.FormatInt(citem.PriorityBurst.ValueInt64(), 10))
					}
					if !citem.QueueLimit.IsNull() && !citem.QueueLimit.IsUnknown() {
						body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"class"+"."+strconv.Itoa(index)+"."+"action-list"+"."+strconv.Itoa(cindex)+"."+"queue-limit.queue-limit-value", strconv.FormatInt(citem.QueueLimit.ValueInt64(), 10))
					}
					if !citem.QueueLimitType.IsNull() && !citem.QueueLimitType.IsUnknown() {
						body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"class"+"."+strconv.Itoa(index)+"."+"action-list"+"."+strconv.Itoa(cindex)+"."+"queue-limit.queue-limit-type", citem.QueueLimitType.ValueString())
					}
					if !citem.ShapeAverageBitRate.IsNull() && !citem.ShapeAverageBitRate.IsUnknown() {
						body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"class"+"."+strconv.Itoa(index)+"."+"action-list"+"."+strconv.Itoa(cindex)+"."+"shape.average.bit-rate", strconv.FormatInt(citem.ShapeAverageBitRate.ValueInt64(), 10))
					}
					if !citem.ShapeAverageBitsPerIntervalSustained.IsNull() && !citem.ShapeAverageBitsPerIntervalSustained.IsUnknown() {
						body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"class"+"."+strconv.Itoa(index)+"."+"action-list"+"."+strconv.Itoa(cindex)+"."+"shape.average.bits-per-interval-sustained", strconv.FormatInt(citem.ShapeAverageBitsPerIntervalSustained.ValueInt64(), 10))
					}
					if !citem.ShapeAverageBitsPerIntervalExcess.IsNull() && !citem.ShapeAverageBitsPerIntervalExcess.IsUnknown() {
						body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"class"+"."+strconv.Itoa(index)+"."+"action-list"+"."+strconv.Itoa(cindex)+"."+"shape.average.bits-per-interval-excess", strconv.FormatInt(citem.ShapeAverageBitsPerIntervalExcess.ValueInt64(), 10))
					}
					if !citem.ShapeAveragePercent.IsNull() && !citem.ShapeAveragePercent.IsUnknown() {
						body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"class"+"."+strconv.Itoa(index)+"."+"action-list"+"."+strconv.Itoa(cindex)+"."+"shape.average.percent", strconv.FormatInt(citem.ShapeAveragePercent.ValueInt64(), 10))
					}
					if !citem.ShapeAverageBurstSizeSustained.IsNull() && !citem.ShapeAverageBurstSizeSustained.IsUnknown() {
						body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"class"+"."+strconv.Itoa(index)+"."+"action-list"+"."+strconv.Itoa(cindex)+"."+"shape.average.burst-size-sustained", strconv.FormatInt(citem.ShapeAverageBurstSizeSustained.ValueInt64(), 10))
					}
					if !citem.ShapeAverageMs.IsNull() && !citem.ShapeAverageMs.IsUnknown() {
						if citem.ShapeAverageMs.ValueBool() {
							body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"class"+"."+strconv.Itoa(index)+"."+"action-list"+"."+strconv.Itoa(cindex)+"."+"shape.average.ms", map[string]string{})
						}
					}
				}
			}
		}
	}
	return body
}

func (data *PolicyMap) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(prefix + "type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get(prefix + "subscriber"); !data.Subscriber.IsNull() {
		if value.Exists() {
			data.Subscriber = types.BoolValue(true)
		} else {
			data.Subscriber = types.BoolValue(false)
		}
	} else {
		data.Subscriber = types.BoolNull()
	}
	if value := res.Get(prefix + "description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	for i := range data.Classes {
		keys := [...]string{"name"}
		keyValues := [...]string{data.Classes[i].Name.ValueString()}

		var r gjson.Result
		res.Get(prefix + "class").ForEach(
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
		if value := r.Get("name"); value.Exists() && !data.Classes[i].Name.IsNull() {
			data.Classes[i].Name = types.StringValue(value.String())
		} else {
			data.Classes[i].Name = types.StringNull()
		}
		for ci := range data.Classes[i].Actions {
			keys := [...]string{"action-type"}
			keyValues := [...]string{data.Classes[i].Actions[ci].Type.ValueString()}

			var cr gjson.Result
			r.Get("action-list").ForEach(
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
						cr = v
						return false
					}
					return true
				},
			)
			if value := cr.Get("action-type"); value.Exists() && !data.Classes[i].Actions[ci].Type.IsNull() {
				data.Classes[i].Actions[ci].Type = types.StringValue(value.String())
			} else {
				data.Classes[i].Actions[ci].Type = types.StringNull()
			}
			if value := cr.Get("bandwidth.bits"); value.Exists() && !data.Classes[i].Actions[ci].BandwidthBits.IsNull() {
				data.Classes[i].Actions[ci].BandwidthBits = types.Int64Value(value.Int())
			} else {
				data.Classes[i].Actions[ci].BandwidthBits = types.Int64Null()
			}
			if value := cr.Get("bandwidth.percent"); value.Exists() && !data.Classes[i].Actions[ci].BandwidthPercent.IsNull() {
				data.Classes[i].Actions[ci].BandwidthPercent = types.Int64Value(value.Int())
			} else {
				data.Classes[i].Actions[ci].BandwidthPercent = types.Int64Null()
			}
			if value := cr.Get("bandwidth.remaining.rem-option"); value.Exists() && !data.Classes[i].Actions[ci].BandwidthRemainingOption.IsNull() {
				data.Classes[i].Actions[ci].BandwidthRemainingOption = types.StringValue(value.String())
			} else {
				data.Classes[i].Actions[ci].BandwidthRemainingOption = types.StringNull()
			}
			if value := cr.Get("bandwidth.remaining.percent"); value.Exists() && !data.Classes[i].Actions[ci].BandwidthRemainingPercent.IsNull() {
				data.Classes[i].Actions[ci].BandwidthRemainingPercent = types.Int64Value(value.Int())
			} else {
				data.Classes[i].Actions[ci].BandwidthRemainingPercent = types.Int64Null()
			}
			if value := cr.Get("bandwidth.remaining.ratio"); value.Exists() && !data.Classes[i].Actions[ci].BandwidthRemainingRatio.IsNull() {
				data.Classes[i].Actions[ci].BandwidthRemainingRatio = types.Int64Value(value.Int())
			} else {
				data.Classes[i].Actions[ci].BandwidthRemainingRatio = types.Int64Null()
			}
			if value := cr.Get("priority.level"); value.Exists() && !data.Classes[i].Actions[ci].PriorityLevel.IsNull() {
				data.Classes[i].Actions[ci].PriorityLevel = types.Int64Value(value.Int())
			} else {
				data.Classes[i].Actions[ci].PriorityLevel = types.Int64Null()
			}
			if value := cr.Get("priority.burst"); value.Exists() && !data.Classes[i].Actions[ci].PriorityBurst.IsNull() {
				data.Classes[i].Actions[ci].PriorityBurst = types.Int64Value(value.Int())
			} else {
				data.Classes[i].Actions[ci].PriorityBurst = types.Int64Null()
			}
			if value := cr.Get("queue-limit.queue-limit-value"); value.Exists() && !data.Classes[i].Actions[ci].QueueLimit.IsNull() {
				data.Classes[i].Actions[ci].QueueLimit = types.Int64Value(value.Int())
			} else {
				data.Classes[i].Actions[ci].QueueLimit = types.Int64Null()
			}
			if value := cr.Get("queue-limit.queue-limit-type"); value.Exists() && !data.Classes[i].Actions[ci].QueueLimitType.IsNull() {
				data.Classes[i].Actions[ci].QueueLimitType = types.StringValue(value.String())
			} else {
				data.Classes[i].Actions[ci].QueueLimitType = types.StringNull()
			}
			if value := cr.Get("shape.average.bit-rate"); value.Exists() && !data.Classes[i].Actions[ci].ShapeAverageBitRate.IsNull() {
				data.Classes[i].Actions[ci].ShapeAverageBitRate = types.Int64Value(value.Int())
			} else {
				data.Classes[i].Actions[ci].ShapeAverageBitRate = types.Int64Null()
			}
			if value := cr.Get("shape.average.bits-per-interval-sustained"); value.Exists() && !data.Classes[i].Actions[ci].ShapeAverageBitsPerIntervalSustained.IsNull() {
				data.Classes[i].Actions[ci].ShapeAverageBitsPerIntervalSustained = types.Int64Value(value.Int())
			} else {
				data.Classes[i].Actions[ci].ShapeAverageBitsPerIntervalSustained = types.Int64Null()
			}
			if value := cr.Get("shape.average.bits-per-interval-excess"); value.Exists() && !data.Classes[i].Actions[ci].ShapeAverageBitsPerIntervalExcess.IsNull() {
				data.Classes[i].Actions[ci].ShapeAverageBitsPerIntervalExcess = types.Int64Value(value.Int())
			} else {
				data.Classes[i].Actions[ci].ShapeAverageBitsPerIntervalExcess = types.Int64Null()
			}
			if value := cr.Get("shape.average.percent"); value.Exists() && !data.Classes[i].Actions[ci].ShapeAveragePercent.IsNull() {
				data.Classes[i].Actions[ci].ShapeAveragePercent = types.Int64Value(value.Int())
			} else {
				data.Classes[i].Actions[ci].ShapeAveragePercent = types.Int64Null()
			}
			if value := cr.Get("shape.average.burst-size-sustained"); value.Exists() && !data.Classes[i].Actions[ci].ShapeAverageBurstSizeSustained.IsNull() {
				data.Classes[i].Actions[ci].ShapeAverageBurstSizeSustained = types.Int64Value(value.Int())
			} else {
				data.Classes[i].Actions[ci].ShapeAverageBurstSizeSustained = types.Int64Null()
			}
			if value := cr.Get("shape.average.ms"); !data.Classes[i].Actions[ci].ShapeAverageMs.IsNull() {
				if value.Exists() {
					data.Classes[i].Actions[ci].ShapeAverageMs = types.BoolValue(true)
				} else {
					data.Classes[i].Actions[ci].ShapeAverageMs = types.BoolValue(false)
				}
			} else {
				data.Classes[i].Actions[ci].ShapeAverageMs = types.BoolNull()
			}
		}
	}
}

func (data *PolicyMap) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "subscriber"); value.Exists() {
		data.Subscriber = types.BoolValue(true)
	} else {
		data.Subscriber = types.BoolValue(false)
	}
	if value := res.Get(prefix + "description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "class"); value.Exists() {
		data.Classes = make([]PolicyMapClasses, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := PolicyMapClasses{}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			}
			if cValue := v.Get("action-list"); cValue.Exists() {
				item.Actions = make([]PolicyMapClassesActions, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := PolicyMapClassesActions{}
					if ccValue := cv.Get("action-type"); ccValue.Exists() {
						cItem.Type = types.StringValue(ccValue.String())
					}
					if ccValue := cv.Get("bandwidth.bits"); ccValue.Exists() {
						cItem.BandwidthBits = types.Int64Value(ccValue.Int())
					}
					if ccValue := cv.Get("bandwidth.percent"); ccValue.Exists() {
						cItem.BandwidthPercent = types.Int64Value(ccValue.Int())
					}
					if ccValue := cv.Get("bandwidth.remaining.rem-option"); ccValue.Exists() {
						cItem.BandwidthRemainingOption = types.StringValue(ccValue.String())
					}
					if ccValue := cv.Get("bandwidth.remaining.percent"); ccValue.Exists() {
						cItem.BandwidthRemainingPercent = types.Int64Value(ccValue.Int())
					}
					if ccValue := cv.Get("bandwidth.remaining.ratio"); ccValue.Exists() {
						cItem.BandwidthRemainingRatio = types.Int64Value(ccValue.Int())
					}
					if ccValue := cv.Get("priority.level"); ccValue.Exists() {
						cItem.PriorityLevel = types.Int64Value(ccValue.Int())
					}
					if ccValue := cv.Get("priority.burst"); ccValue.Exists() {
						cItem.PriorityBurst = types.Int64Value(ccValue.Int())
					}
					if ccValue := cv.Get("queue-limit.queue-limit-value"); ccValue.Exists() {
						cItem.QueueLimit = types.Int64Value(ccValue.Int())
					}
					if ccValue := cv.Get("queue-limit.queue-limit-type"); ccValue.Exists() {
						cItem.QueueLimitType = types.StringValue(ccValue.String())
					}
					if ccValue := cv.Get("shape.average.bit-rate"); ccValue.Exists() {
						cItem.ShapeAverageBitRate = types.Int64Value(ccValue.Int())
					}
					if ccValue := cv.Get("shape.average.bits-per-interval-sustained"); ccValue.Exists() {
						cItem.ShapeAverageBitsPerIntervalSustained = types.Int64Value(ccValue.Int())
					}
					if ccValue := cv.Get("shape.average.bits-per-interval-excess"); ccValue.Exists() {
						cItem.ShapeAverageBitsPerIntervalExcess = types.Int64Value(ccValue.Int())
					}
					if ccValue := cv.Get("shape.average.percent"); ccValue.Exists() {
						cItem.ShapeAveragePercent = types.Int64Value(ccValue.Int())
					}
					if ccValue := cv.Get("shape.average.burst-size-sustained"); ccValue.Exists() {
						cItem.ShapeAverageBurstSizeSustained = types.Int64Value(ccValue.Int())
					}
					if ccValue := cv.Get("shape.average.ms"); ccValue.Exists() {
						cItem.ShapeAverageMs = types.BoolValue(true)
					} else {
						cItem.ShapeAverageMs = types.BoolValue(false)
					}
					item.Actions = append(item.Actions, cItem)
					return true
				})
			}
			data.Classes = append(data.Classes, item)
			return true
		})
	}
}

func (data *PolicyMapData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "subscriber"); value.Exists() {
		data.Subscriber = types.BoolValue(true)
	} else {
		data.Subscriber = types.BoolValue(false)
	}
	if value := res.Get(prefix + "description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "class"); value.Exists() {
		data.Classes = make([]PolicyMapClasses, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := PolicyMapClasses{}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			}
			if cValue := v.Get("action-list"); cValue.Exists() {
				item.Actions = make([]PolicyMapClassesActions, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := PolicyMapClassesActions{}
					if ccValue := cv.Get("action-type"); ccValue.Exists() {
						cItem.Type = types.StringValue(ccValue.String())
					}
					if ccValue := cv.Get("bandwidth.bits"); ccValue.Exists() {
						cItem.BandwidthBits = types.Int64Value(ccValue.Int())
					}
					if ccValue := cv.Get("bandwidth.percent"); ccValue.Exists() {
						cItem.BandwidthPercent = types.Int64Value(ccValue.Int())
					}
					if ccValue := cv.Get("bandwidth.remaining.rem-option"); ccValue.Exists() {
						cItem.BandwidthRemainingOption = types.StringValue(ccValue.String())
					}
					if ccValue := cv.Get("bandwidth.remaining.percent"); ccValue.Exists() {
						cItem.BandwidthRemainingPercent = types.Int64Value(ccValue.Int())
					}
					if ccValue := cv.Get("bandwidth.remaining.ratio"); ccValue.Exists() {
						cItem.BandwidthRemainingRatio = types.Int64Value(ccValue.Int())
					}
					if ccValue := cv.Get("priority.level"); ccValue.Exists() {
						cItem.PriorityLevel = types.Int64Value(ccValue.Int())
					}
					if ccValue := cv.Get("priority.burst"); ccValue.Exists() {
						cItem.PriorityBurst = types.Int64Value(ccValue.Int())
					}
					if ccValue := cv.Get("queue-limit.queue-limit-value"); ccValue.Exists() {
						cItem.QueueLimit = types.Int64Value(ccValue.Int())
					}
					if ccValue := cv.Get("queue-limit.queue-limit-type"); ccValue.Exists() {
						cItem.QueueLimitType = types.StringValue(ccValue.String())
					}
					if ccValue := cv.Get("shape.average.bit-rate"); ccValue.Exists() {
						cItem.ShapeAverageBitRate = types.Int64Value(ccValue.Int())
					}
					if ccValue := cv.Get("shape.average.bits-per-interval-sustained"); ccValue.Exists() {
						cItem.ShapeAverageBitsPerIntervalSustained = types.Int64Value(ccValue.Int())
					}
					if ccValue := cv.Get("shape.average.bits-per-interval-excess"); ccValue.Exists() {
						cItem.ShapeAverageBitsPerIntervalExcess = types.Int64Value(ccValue.Int())
					}
					if ccValue := cv.Get("shape.average.percent"); ccValue.Exists() {
						cItem.ShapeAveragePercent = types.Int64Value(ccValue.Int())
					}
					if ccValue := cv.Get("shape.average.burst-size-sustained"); ccValue.Exists() {
						cItem.ShapeAverageBurstSizeSustained = types.Int64Value(ccValue.Int())
					}
					if ccValue := cv.Get("shape.average.ms"); ccValue.Exists() {
						cItem.ShapeAverageMs = types.BoolValue(true)
					} else {
						cItem.ShapeAverageMs = types.BoolValue(false)
					}
					item.Actions = append(item.Actions, cItem)
					return true
				})
			}
			data.Classes = append(data.Classes, item)
			return true
		})
	}
}

func (data *PolicyMap) getDeletedItems(ctx context.Context, state PolicyMap) []string {
	deletedItems := make([]string, 0)
	if !state.Type.IsNull() && data.Type.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/type", state.getPath()))
	}
	if !state.Subscriber.IsNull() && data.Subscriber.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/subscriber", state.getPath()))
	}
	if !state.Description.IsNull() && data.Description.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/description", state.getPath()))
	}
	for i := range state.Classes {
		stateKeyValues := [...]string{state.Classes[i].Name.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.Classes[i].Name.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Classes {
			found = true
			if state.Classes[i].Name.ValueString() != data.Classes[j].Name.ValueString() {
				found = false
			}
			if found {
				for ci := range state.Classes[i].Actions {
					cstateKeyValues := [...]string{state.Classes[i].Actions[ci].Type.ValueString()}

					cemptyKeys := true
					if !reflect.ValueOf(state.Classes[i].Actions[ci].Type.ValueString()).IsZero() {
						cemptyKeys = false
					}
					if cemptyKeys {
						continue
					}

					found := false
					for cj := range data.Classes[j].Actions {
						found = true
						if state.Classes[i].Actions[ci].Type.ValueString() != data.Classes[j].Actions[cj].Type.ValueString() {
							found = false
						}
						if found {
							if !state.Classes[i].Actions[ci].BandwidthBits.IsNull() && data.Classes[j].Actions[cj].BandwidthBits.IsNull() {
								deletedItems = append(deletedItems, fmt.Sprintf("%v/class=%v/action-list=%v/bandwidth/bits", state.getPath(), strings.Join(stateKeyValues[:], ","), strings.Join(cstateKeyValues[:], ",")))
							}
							if !state.Classes[i].Actions[ci].BandwidthPercent.IsNull() && data.Classes[j].Actions[cj].BandwidthPercent.IsNull() {
								deletedItems = append(deletedItems, fmt.Sprintf("%v/class=%v/action-list=%v/bandwidth/percent", state.getPath(), strings.Join(stateKeyValues[:], ","), strings.Join(cstateKeyValues[:], ",")))
							}
							if !state.Classes[i].Actions[ci].BandwidthRemainingOption.IsNull() && data.Classes[j].Actions[cj].BandwidthRemainingOption.IsNull() {
								deletedItems = append(deletedItems, fmt.Sprintf("%v/class=%v/action-list=%v/bandwidth/remaining/rem-option", state.getPath(), strings.Join(stateKeyValues[:], ","), strings.Join(cstateKeyValues[:], ",")))
							}
							if !state.Classes[i].Actions[ci].BandwidthRemainingPercent.IsNull() && data.Classes[j].Actions[cj].BandwidthRemainingPercent.IsNull() {
								deletedItems = append(deletedItems, fmt.Sprintf("%v/class=%v/action-list=%v/bandwidth/remaining/percent", state.getPath(), strings.Join(stateKeyValues[:], ","), strings.Join(cstateKeyValues[:], ",")))
							}
							if !state.Classes[i].Actions[ci].BandwidthRemainingRatio.IsNull() && data.Classes[j].Actions[cj].BandwidthRemainingRatio.IsNull() {
								deletedItems = append(deletedItems, fmt.Sprintf("%v/class=%v/action-list=%v/bandwidth/remaining/ratio", state.getPath(), strings.Join(stateKeyValues[:], ","), strings.Join(cstateKeyValues[:], ",")))
							}
							if !state.Classes[i].Actions[ci].PriorityLevel.IsNull() && data.Classes[j].Actions[cj].PriorityLevel.IsNull() {
								deletedItems = append(deletedItems, fmt.Sprintf("%v/class=%v/action-list=%v/priority/level", state.getPath(), strings.Join(stateKeyValues[:], ","), strings.Join(cstateKeyValues[:], ",")))
							}
							if !state.Classes[i].Actions[ci].PriorityBurst.IsNull() && data.Classes[j].Actions[cj].PriorityBurst.IsNull() {
								deletedItems = append(deletedItems, fmt.Sprintf("%v/class=%v/action-list=%v/priority/burst", state.getPath(), strings.Join(stateKeyValues[:], ","), strings.Join(cstateKeyValues[:], ",")))
							}
							if !state.Classes[i].Actions[ci].QueueLimit.IsNull() && data.Classes[j].Actions[cj].QueueLimit.IsNull() {
								deletedItems = append(deletedItems, fmt.Sprintf("%v/class=%v/action-list=%v/queue-limit/queue-limit-value", state.getPath(), strings.Join(stateKeyValues[:], ","), strings.Join(cstateKeyValues[:], ",")))
							}
							if !state.Classes[i].Actions[ci].QueueLimitType.IsNull() && data.Classes[j].Actions[cj].QueueLimitType.IsNull() {
								deletedItems = append(deletedItems, fmt.Sprintf("%v/class=%v/action-list=%v/queue-limit/queue-limit-type", state.getPath(), strings.Join(stateKeyValues[:], ","), strings.Join(cstateKeyValues[:], ",")))
							}
							if !state.Classes[i].Actions[ci].ShapeAverageBitRate.IsNull() && data.Classes[j].Actions[cj].ShapeAverageBitRate.IsNull() {
								deletedItems = append(deletedItems, fmt.Sprintf("%v/class=%v/action-list=%v/shape/average/bit-rate", state.getPath(), strings.Join(stateKeyValues[:], ","), strings.Join(cstateKeyValues[:], ",")))
							}
							if !state.Classes[i].Actions[ci].ShapeAverageBitsPerIntervalSustained.IsNull() && data.Classes[j].Actions[cj].ShapeAverageBitsPerIntervalSustained.IsNull() {
								deletedItems = append(deletedItems, fmt.Sprintf("%v/class=%v/action-list=%v/shape/average/bits-per-interval-sustained", state.getPath(), strings.Join(stateKeyValues[:], ","), strings.Join(cstateKeyValues[:], ",")))
							}
							if !state.Classes[i].Actions[ci].ShapeAverageBitsPerIntervalExcess.IsNull() && data.Classes[j].Actions[cj].ShapeAverageBitsPerIntervalExcess.IsNull() {
								deletedItems = append(deletedItems, fmt.Sprintf("%v/class=%v/action-list=%v/shape/average/bits-per-interval-excess", state.getPath(), strings.Join(stateKeyValues[:], ","), strings.Join(cstateKeyValues[:], ",")))
							}
							if !state.Classes[i].Actions[ci].ShapeAveragePercent.IsNull() && data.Classes[j].Actions[cj].ShapeAveragePercent.IsNull() {
								deletedItems = append(deletedItems, fmt.Sprintf("%v/class=%v/action-list=%v/shape/average/percent", state.getPath(), strings.Join(stateKeyValues[:], ","), strings.Join(cstateKeyValues[:], ",")))
							}
							if !state.Classes[i].Actions[ci].ShapeAverageBurstSizeSustained.IsNull() && data.Classes[j].Actions[cj].ShapeAverageBurstSizeSustained.IsNull() {
								deletedItems = append(deletedItems, fmt.Sprintf("%v/class=%v/action-list=%v/shape/average/burst-size-sustained", state.getPath(), strings.Join(stateKeyValues[:], ","), strings.Join(cstateKeyValues[:], ",")))
							}
							if !state.Classes[i].Actions[ci].ShapeAverageMs.IsNull() && data.Classes[j].Actions[cj].ShapeAverageMs.IsNull() {
								deletedItems = append(deletedItems, fmt.Sprintf("%v/class=%v/action-list=%v/shape/average/ms", state.getPath(), strings.Join(stateKeyValues[:], ","), strings.Join(cstateKeyValues[:], ",")))
							}
							break
						}
					}
					if !found {
						deletedItems = append(deletedItems, fmt.Sprintf("%v/class=%v/action-list=%v", state.getPath(), strings.Join(stateKeyValues[:], ","), strings.Join(cstateKeyValues[:], ",")))
					}
				}
				break
			}
		}
		if !found {
			deletedItems = append(deletedItems, fmt.Sprintf("%v/class=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	return deletedItems
}

func (data *PolicyMap) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.Subscriber.IsNull() && !data.Subscriber.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/subscriber", data.getPath()))
	}

	for i := range data.Classes {
		keyValues := [...]string{data.Classes[i].Name.ValueString()}

		for ci := range data.Classes[i].Actions {
			ckeyValues := [...]string{data.Classes[i].Actions[ci].Type.ValueString()}
			if !data.Classes[i].Actions[ci].ShapeAverageMs.IsNull() && !data.Classes[i].Actions[ci].ShapeAverageMs.ValueBool() {
				emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/class=%v/action-list=%v/shape/average/ms", data.getPath(), strings.Join(keyValues[:], ","), strings.Join(ckeyValues[:], ",")))
			}
		}
	}
	return emptyLeafsDelete
}

func (data *PolicyMap) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.Type.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/type", data.getPath()))
	}
	if !data.Subscriber.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/subscriber", data.getPath()))
	}
	if !data.Description.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/description", data.getPath()))
	}
	for i := range data.Classes {
		keyValues := [...]string{data.Classes[i].Name.ValueString()}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/class=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	return deletePaths
}

func (data *PolicyMap) getIdsFromPath() {
	reString := strings.ReplaceAll("Cisco-IOS-XE-native:native/policy/Cisco-IOS-XE-policy:policy-map=%v", "%s", "(.+)")
	reString = strings.ReplaceAll(reString, "%v", "(.+)")
	re := regexp.MustCompile(reString)
	matches := re.FindStringSubmatch(data.Id.ValueString())
	data.Name = types.StringValue(helpers.Must(url.QueryUnescape(matches[1])))
}
