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
	"strings"

	"github.com/CiscoDevNet/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-netconf"
	"github.com/netascode/xmldot"
)

type Yang struct {
	Device     types.String `tfsdk:"device"`
	Id         types.String `tfsdk:"id"`
	Path       types.String `tfsdk:"path"`
	Delete     types.Bool   `tfsdk:"delete"`
	Attributes types.Map    `tfsdk:"attributes"`
	Lists      []YangList   `tfsdk:"lists"`
}

type YangList struct {
	Name   types.String `tfsdk:"name"`
	Key    types.String `tfsdk:"key"`
	Items  []types.Map  `tfsdk:"items"`
	Values types.List   `tfsdk:"values"`
}

type YangDataSourceModel struct {
	Device     types.String `tfsdk:"device"`
	Id         types.String `tfsdk:"id"`
	Path       types.String `tfsdk:"path"`
	Attributes types.Map    `tfsdk:"attributes"`
}

func (data Yang) toBodyXML(ctx context.Context) string {
	body := netconf.Body{}

	var attributes map[string]string
	data.Attributes.ElementsAs(ctx, &attributes, false)

	for attr, value := range attributes {
		body = helpers.SetFromXPath(body, data.Path.ValueString()+"/"+attr, value)
	}
	for i := range data.Lists {
		if len(data.Lists[i].Items) > 0 {
			for ii := range data.Lists[i].Items {
				var listAttributes map[string]string
				data.Lists[i].Items[ii].ElementsAs(ctx, &listAttributes, false)
				attrs := netconf.Body{}

				// Get key(s) for this list
				keys := strings.Split(data.Lists[i].Key.ValueString(), ",")

				// First, set all key attributes in order
				for _, key := range keys {
					if value, ok := listAttributes[key]; ok {
						attrs = helpers.SetFromXPath(attrs, key, value)
					}
				}

				// Then, set all non-key attributes
				for attr, value := range listAttributes {
					// Skip if this is a key attribute (already set above)
					isKey := false
					for _, key := range keys {
						if attr == key {
							isKey = true
							break
						}
					}
					if !isKey {
						attrs = helpers.SetFromXPath(attrs, attr, value)
					}
				}

				body = helpers.SetRawFromXPath(body, data.Path.ValueString()+"/"+data.Lists[i].Name.ValueString(), attrs.Res())
			}
		} else if len(data.Lists[i].Values.Elements()) > 0 {
			var values []string
			data.Lists[i].Values.ElementsAs(ctx, &values, false)
			body = helpers.SetFromXPath(body, data.Path.ValueString()+"/"+data.Lists[i].Name.ValueString(), values)
		}
	}

	return body.Res()
}

func (data *Yang) fromBodyXML(ctx context.Context, res xmldot.Result) {

	// Parse attributes
	attributes := data.Attributes.Elements()
	for attr := range attributes {
		value := helpers.GetFromXPath(res, "data"+data.Path.ValueString()+"/"+attr)
		if !value.Exists() || value.String() == "" {
			attributes[attr] = types.StringValue("")
		} else {
			attributes[attr] = types.StringValue(value.String())
		}
	}
	data.Attributes = types.MapValueMust(types.StringType, attributes)

	// Parse lists
	for i := range data.Lists {
		keys := strings.Split(data.Lists[i].Key.ValueString(), ",")
		listName := data.Lists[i].Name.ValueString()

		if len(data.Lists[i].Items) > 0 {
			// Complex list items with multiple attributes
			for ii := range data.Lists[i].Items {
				// Get key values from plan
				var keyValues []string
				for _, key := range keys {
					v, _ := data.Lists[i].Items[ii].Elements()[key].ToTerraformValue(ctx)
					var keyValue string
					v.As(&keyValue)
					keyValues = append(keyValues, keyValue)
				}

				// Build XPath to find the specific list item by key(s)
				xpathPredicates := ""
				for ik, key := range keys {
					xpathPredicates += "[" + key + "='" + keyValues[ik] + "']"
				}
				itemXPath := listName + xpathPredicates

				// Find the matching list item in XML response
				itemResult := helpers.GetFromXPath(res, "data"+data.Path.ValueString()+"/"+itemXPath)

				// Parse attributes from the matched item
				itemAttributes := data.Lists[i].Items[ii].Elements()
				for attr := range itemAttributes {
					value := helpers.GetFromXPath(itemResult, attr)
					if !value.Exists() || value.String() == "" {
						itemAttributes[attr] = types.StringValue("")
					} else {
						itemAttributes[attr] = types.StringValue(value.String())
					}
				}
				data.Lists[i].Items[ii] = types.MapValueMust(types.StringType, itemAttributes)
			}
		} else if len(data.Lists[i].Values.Elements()) > 0 {
			// Simple leaf-list values
			listResult := helpers.GetFromXPath(res, "data"+data.Path.ValueString()+"/"+listName)
			if listResult.IsArray() {
				values := make([]attr.Value, 0)
				for _, v := range listResult.Array() {
					values = append(values, types.StringValue(v.String()))
				}
				data.Lists[i].Values = types.ListValueMust(data.Lists[i].Values.ElementType(ctx), values)
			}
		}
	}
}
