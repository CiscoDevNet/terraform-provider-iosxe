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

type CryptoIPSecProfile struct {
	Device                                                     types.String `tfsdk:"device"`
	Id                                                         types.String `tfsdk:"id"`
	Name                                                       types.String `tfsdk:"name"`
	SetTransformSet                                            types.List   `tfsdk:"set_transform_set"`
	SetIsakmpProfileIkev2ProfileIkev2ProfileCaseIkev2Profile   types.String `tfsdk:"set_isakmp_profile_ikev2_profile_ikev2_profile_case_ikev2_profile"`
	SetIsakmpProfileIkev2ProfileIsakmpProfileCaseIsakmpProfile types.String `tfsdk:"set_isakmp_profile_ikev2_profile_isakmp_profile_case_isakmp_profile"`
}

type CryptoIPSecProfileData struct {
	Device                                                     types.String `tfsdk:"device"`
	Id                                                         types.String `tfsdk:"id"`
	Name                                                       types.String `tfsdk:"name"`
	SetTransformSet                                            types.List   `tfsdk:"set_transform_set"`
	SetIsakmpProfileIkev2ProfileIkev2ProfileCaseIkev2Profile   types.String `tfsdk:"set_isakmp_profile_ikev2_profile_ikev2_profile_case_ikev2_profile"`
	SetIsakmpProfileIkev2ProfileIsakmpProfileCaseIsakmpProfile types.String `tfsdk:"set_isakmp_profile_ikev2_profile_isakmp_profile_case_isakmp_profile"`
}

func (data CryptoIPSecProfile) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/crypto/Cisco-IOS-XE-crypto:ipsec/profile=%v", url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())))
}

func (data CryptoIPSecProfileData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/crypto/Cisco-IOS-XE-crypto:ipsec/profile=%v", url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())))
}

// if last path element has a key -> remove it
func (data CryptoIPSecProfile) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data CryptoIPSecProfile) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"name", data.Name.ValueString())
	}
	if !data.SetTransformSet.IsNull() && !data.SetTransformSet.IsUnknown() {
		var values []string
		data.SetTransformSet.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"set.transform-set", values)
	}
	if !data.SetIsakmpProfileIkev2ProfileIkev2ProfileCaseIkev2Profile.IsNull() && !data.SetIsakmpProfileIkev2ProfileIkev2ProfileCaseIkev2Profile.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"set.ikev2-profile", data.SetIsakmpProfileIkev2ProfileIkev2ProfileCaseIkev2Profile.ValueString())
	}
	if !data.SetIsakmpProfileIkev2ProfileIsakmpProfileCaseIsakmpProfile.IsNull() && !data.SetIsakmpProfileIkev2ProfileIsakmpProfileCaseIsakmpProfile.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"set.isakmp-profile", data.SetIsakmpProfileIkev2ProfileIsakmpProfileCaseIsakmpProfile.ValueString())
	}
	return body
}

func (data *CryptoIPSecProfile) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(prefix + "set.transform-set"); value.Exists() && !data.SetTransformSet.IsNull() {
		data.SetTransformSet = helpers.GetStringList(value.Array())
	} else {
		data.SetTransformSet = types.ListNull(types.StringType)
	}
	if value := res.Get(prefix + "set.ikev2-profile"); value.Exists() && !data.SetIsakmpProfileIkev2ProfileIkev2ProfileCaseIkev2Profile.IsNull() {
		data.SetIsakmpProfileIkev2ProfileIkev2ProfileCaseIkev2Profile = types.StringValue(value.String())
	} else {
		data.SetIsakmpProfileIkev2ProfileIkev2ProfileCaseIkev2Profile = types.StringNull()
	}
	if value := res.Get(prefix + "set.isakmp-profile"); value.Exists() && !data.SetIsakmpProfileIkev2ProfileIsakmpProfileCaseIsakmpProfile.IsNull() {
		data.SetIsakmpProfileIkev2ProfileIsakmpProfileCaseIsakmpProfile = types.StringValue(value.String())
	} else {
		data.SetIsakmpProfileIkev2ProfileIsakmpProfileCaseIsakmpProfile = types.StringNull()
	}
}

func (data *CryptoIPSecProfile) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "set.transform-set"); value.Exists() {
		data.SetTransformSet = helpers.GetStringList(value.Array())
	} else {
		data.SetTransformSet = types.ListNull(types.StringType)
	}
	if value := res.Get(prefix + "set.ikev2-profile"); value.Exists() {
		data.SetIsakmpProfileIkev2ProfileIkev2ProfileCaseIkev2Profile = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "set.isakmp-profile"); value.Exists() {
		data.SetIsakmpProfileIkev2ProfileIsakmpProfileCaseIsakmpProfile = types.StringValue(value.String())
	}
}

func (data *CryptoIPSecProfileData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "set.transform-set"); value.Exists() {
		data.SetTransformSet = helpers.GetStringList(value.Array())
	} else {
		data.SetTransformSet = types.ListNull(types.StringType)
	}
	if value := res.Get(prefix + "set.ikev2-profile"); value.Exists() {
		data.SetIsakmpProfileIkev2ProfileIkev2ProfileCaseIkev2Profile = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "set.isakmp-profile"); value.Exists() {
		data.SetIsakmpProfileIkev2ProfileIsakmpProfileCaseIsakmpProfile = types.StringValue(value.String())
	}
}

func (data *CryptoIPSecProfile) getDeletedItems(ctx context.Context, state CryptoIPSecProfile) []string {
	deletedItems := make([]string, 0)
	if !state.SetTransformSet.IsNull() {
		if data.SetTransformSet.IsNull() {
			deletedItems = append(deletedItems, fmt.Sprintf("%v/set/transform-set", state.getPath()))
		} else {
			var dataValues, stateValues []string
			data.SetTransformSet.ElementsAs(ctx, &dataValues, false)
			state.SetTransformSet.ElementsAs(ctx, &stateValues, false)
			for _, v := range stateValues {
				found := false
				for _, vv := range dataValues {
					if v == vv {
						found = true
						break
					}
				}
				if !found {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/set/transform-set=%v", state.getPath(), v))
				}
			}
		}
	}
	if !state.SetIsakmpProfileIkev2ProfileIkev2ProfileCaseIkev2Profile.IsNull() && data.SetIsakmpProfileIkev2ProfileIkev2ProfileCaseIkev2Profile.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/set/ikev2-profile", state.getPath()))
	}
	if !state.SetIsakmpProfileIkev2ProfileIsakmpProfileCaseIsakmpProfile.IsNull() && data.SetIsakmpProfileIkev2ProfileIsakmpProfileCaseIsakmpProfile.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/set/isakmp-profile", state.getPath()))
	}
	return deletedItems
}

func (data *CryptoIPSecProfile) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	return emptyLeafsDelete
}

func (data *CryptoIPSecProfile) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.SetTransformSet.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/set/transform-set", data.getPath()))
	}
	if !data.SetIsakmpProfileIkev2ProfileIkev2ProfileCaseIkev2Profile.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/set/ikev2-profile", data.getPath()))
	}
	if !data.SetIsakmpProfileIkev2ProfileIsakmpProfileCaseIsakmpProfile.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/set/isakmp-profile", data.getPath()))
	}
	return deletePaths
}
