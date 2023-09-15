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

package provider

import (
	"context"
	"regexp"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-restconf"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type Restconf struct {
	Device     types.String   `tfsdk:"device"`
	Id         types.String   `tfsdk:"id"`
	Path       types.String   `tfsdk:"path"`
	Delete     types.Bool     `tfsdk:"delete"`
	Attributes types.Map      `tfsdk:"attributes"`
	Lists      []RestconfList `tfsdk:"lists"`
}

type RestconfList struct {
	Name   types.String `tfsdk:"name"`
	Key    types.String `tfsdk:"key"`
	Items  []types.Map  `tfsdk:"items"`
	Values types.List   `tfsdk:"values"`
}

type RestconfDataSourceModel struct {
	Device     types.String `tfsdk:"device"`
	Id         types.String `tfsdk:"id"`
	Path       types.String `tfsdk:"path"`
	Attributes types.Map    `tfsdk:"attributes"`
}

func (data Restconf) getPath() string {
	return data.Path.ValueString()
}

// if last path element has a key -> remove it
func (data Restconf) getPathShort() string {
	path := data.Path.ValueString()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data Restconf) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.Path.ValueString()) + `":{}}`

	var attributes map[string]string
	data.Attributes.ElementsAs(ctx, &attributes, false)

	for attr, value := range attributes {
		attr = strings.ReplaceAll(attr, "/", ".")
		body, _ = sjson.Set(body, helpers.LastElement(data.Path.ValueString())+"."+attr, value)
	}
	for i := range data.Lists {
		listName := strings.ReplaceAll(data.Lists[i].Name.ValueString(), "/", ".")
		if len(data.Lists[i].Items) > 0 {
			body, _ = sjson.Set(body, helpers.LastElement(data.Path.ValueString())+"."+listName, []interface{}{})
			for ii := range data.Lists[i].Items {
				var listAttributes map[string]string
				data.Lists[i].Items[ii].ElementsAs(ctx, &listAttributes, false)
				attrs := restconf.Body{}
				for attr, value := range listAttributes {
					attr = strings.ReplaceAll(attr, "/", ".")
					attrs = attrs.Set(attr, value)
				}
				body, _ = sjson.SetRaw(body, helpers.LastElement(data.Path.ValueString())+"."+listName+".-1", attrs.Str)
			}
		} else if len(data.Lists[i].Values.Elements()) > 0 {
			var values []string
			data.Lists[i].Values.ElementsAs(ctx, &values, false)
			body, _ = sjson.Set(body, helpers.LastElement(data.Path.ValueString())+"."+listName, values)
		}
	}

	return body
}

func (data *Restconf) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	attributes := data.Attributes.Elements()
	for attr := range attributes {
		attrPath := strings.ReplaceAll(attr, "/", ".")
		value := res.Get(prefix + attrPath)
		if !value.Exists() ||
			(value.IsObject() && len(value.Map()) == 0) ||
			value.Raw == "[null]" {

			attributes[attr] = types.StringValue("")
		} else {
			attributes[attr] = types.StringValue(value.String())
		}
	}
	data.Attributes = types.MapValueMust(types.StringType, attributes)

	for i := range data.Lists {
		keys := strings.Split(data.Lists[i].Key.ValueString(), ",")
		namePath := strings.ReplaceAll(data.Lists[i].Name.ValueString(), "/", ".")
		if len(data.Lists[i].Items) > 0 {
			for ii := range data.Lists[i].Items {
				var keyValues []string
				for _, key := range keys {
					v, _ := data.Lists[i].Items[ii].Elements()[key].ToTerraformValue(ctx)
					var keyValue string
					v.As(&keyValue)
					keyValues = append(keyValues, keyValue)
				}

				// find item by key(s)
				var r gjson.Result
				res.Get(prefix + namePath).ForEach(
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

				attributes := data.Lists[i].Items[ii].Elements()
				for attr := range attributes {
					attrPath := strings.ReplaceAll(attr, "/", ".")
					value := r.Get(attrPath)
					if !value.Exists() ||
						(value.IsObject() && len(value.Map()) == 0) ||
						value.Raw == "[null]" {

						attributes[attr] = types.StringValue("")
					} else {
						attributes[attr] = types.StringValue(value.String())
					}
				}
				data.Lists[i].Items[ii] = types.MapValueMust(types.StringType, attributes)
			}
		} else if len(data.Lists[i].Values.Elements()) > 0 {
			values := res.Get(prefix + namePath)
			if values.IsArray() {
				data.Lists[i].Values = types.ListValueMust(data.Lists[i].Values.ElementType(ctx), helpers.GetValueSlice(values.Array()))
			}
		}
	}
}

func (data *Restconf) getDeletedItems(ctx context.Context, state Restconf) []string {
	deletedItems := make([]string, 0)
	for l := range state.Lists {
		name := state.Lists[l].Name.ValueString()
		namePath := strings.ReplaceAll(name, "/", ".")
		keys := strings.Split(state.Lists[l].Key.ValueString(), ",")
		var dataList RestconfList
		for _, dl := range data.Lists {
			if dl.Name.ValueString() == name {
				dataList = dl
			}
		}
		if len(state.Lists[l].Items) > 0 {
			// check if state item is also included in plan, if not delete item
			for i := range state.Lists[l].Items {
				var slia map[string]string
				state.Lists[l].Items[i].ElementsAs(ctx, &slia, false)

				// if state key values are empty move on to next item
				emptyKey := false
				for _, key := range keys {
					if slia[key] == "" {
						emptyKey = true
						break
					}
				}
				if emptyKey {
					continue
				}

				// find data (plan) item with matching key values
				found := false
				for dli := range dataList.Items {
					var dlia map[string]string
					dataList.Items[dli].ElementsAs(ctx, &dlia, false)
					for _, key := range keys {
						if dlia[key] == slia[key] {
							found = true
							continue
						}
						found = false
						break
					}
					if found {
						break
					}
				}

				// if no matching item in plan found -> delete
				if !found {
					keyValues := make([]string, len(keys))
					for k, key := range keys {
						keyValues[k] = slia[key]
					}
					deletedItems = append(deletedItems, state.getPath()+"/"+namePath+"="+strings.Join(keyValues, ","))
				}
			}
		} else if len(state.Lists[l].Values.Elements()) > 0 {
			var slv []string
			state.Lists[l].Values.ElementsAs(ctx, &slv, false)
			// check if state value is also included in plan, if not delete value from list
			for _, stateValue := range slv {
				found := false
				var dlv []string
				dataList.Values.ElementsAs(ctx, &dlv, false)
				for _, dataValue := range dlv {
					if stateValue == dataValue {
						found = true
						break
					}
				}
				if !found {
					deletedItems = append(deletedItems, state.getPath()+"/"+namePath+"="+stateValue)
				}
			}
		}
	}
	return deletedItems
}
