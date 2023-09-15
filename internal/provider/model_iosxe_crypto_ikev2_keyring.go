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

type CryptoIKEv2Keyring struct {
	Device types.String              `tfsdk:"device"`
	Id     types.String              `tfsdk:"id"`
	Name   types.String              `tfsdk:"name"`
	Peers  []CryptoIKEv2KeyringPeers `tfsdk:"peers"`
}

type CryptoIKEv2KeyringData struct {
	Device types.String              `tfsdk:"device"`
	Id     types.String              `tfsdk:"id"`
	Name   types.String              `tfsdk:"name"`
	Peers  []CryptoIKEv2KeyringPeers `tfsdk:"peers"`
}
type CryptoIKEv2KeyringPeers struct {
	Name                         types.String `tfsdk:"name"`
	Description                  types.String `tfsdk:"description"`
	Hostname                     types.String `tfsdk:"hostname"`
	Ipv4Address                  types.String `tfsdk:"ipv4_address"`
	Ipv4Mask                     types.String `tfsdk:"ipv4_mask"`
	Ipv6Prefix                   types.String `tfsdk:"ipv6_prefix"`
	IdentityKeyId                types.String `tfsdk:"identity_key_id"`
	IdentityAddress              types.String `tfsdk:"identity_address"`
	IdentityEmailName            types.String `tfsdk:"identity_email_name"`
	IdentityEmailDomain          types.String `tfsdk:"identity_email_domain"`
	IdentityFqdnName             types.String `tfsdk:"identity_fqdn_name"`
	IdentityFqdnDomain           types.String `tfsdk:"identity_fqdn_domain"`
	PreSharedKeyLocalEncryption  types.String `tfsdk:"pre_shared_key_local_encryption"`
	PreSharedKeyLocal            types.String `tfsdk:"pre_shared_key_local"`
	PreSharedKeyRemoteEncryption types.String `tfsdk:"pre_shared_key_remote_encryption"`
	PreSharedKeyRemote           types.String `tfsdk:"pre_shared_key_remote"`
	PreSharedKeyEncryption       types.String `tfsdk:"pre_shared_key_encryption"`
	PreSharedKey                 types.String `tfsdk:"pre_shared_key"`
}

func (data CryptoIKEv2Keyring) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/crypto/Cisco-IOS-XE-crypto:ikev2/keyring=%v", url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())))
}

func (data CryptoIKEv2KeyringData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/crypto/Cisco-IOS-XE-crypto:ikev2/keyring=%v", url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())))
}

// if last path element has a key -> remove it
func (data CryptoIKEv2Keyring) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data CryptoIKEv2Keyring) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"name", data.Name.ValueString())
	}
	if len(data.Peers) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"peer", []interface{}{})
		for index, item := range data.Peers {
			if !item.Name.IsNull() && !item.Name.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"peer"+"."+strconv.Itoa(index)+"."+"name", item.Name.ValueString())
			}
			if !item.Description.IsNull() && !item.Description.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"peer"+"."+strconv.Itoa(index)+"."+"description", item.Description.ValueString())
			}
			if !item.Hostname.IsNull() && !item.Hostname.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"peer"+"."+strconv.Itoa(index)+"."+"hostname", item.Hostname.ValueString())
			}
			if !item.Ipv4Address.IsNull() && !item.Ipv4Address.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"peer"+"."+strconv.Itoa(index)+"."+"address.ipv4.ipv4-address", item.Ipv4Address.ValueString())
			}
			if !item.Ipv4Mask.IsNull() && !item.Ipv4Mask.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"peer"+"."+strconv.Itoa(index)+"."+"address.ipv4.ipv4-mask", item.Ipv4Mask.ValueString())
			}
			if !item.Ipv6Prefix.IsNull() && !item.Ipv6Prefix.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"peer"+"."+strconv.Itoa(index)+"."+"address.ipv6-prefix", item.Ipv6Prefix.ValueString())
			}
			if !item.IdentityKeyId.IsNull() && !item.IdentityKeyId.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"peer"+"."+strconv.Itoa(index)+"."+"identity.key-id-number", item.IdentityKeyId.ValueString())
			}
			if !item.IdentityAddress.IsNull() && !item.IdentityAddress.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"peer"+"."+strconv.Itoa(index)+"."+"identity.address-type", item.IdentityAddress.ValueString())
			}
			if !item.IdentityEmailName.IsNull() && !item.IdentityEmailName.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"peer"+"."+strconv.Itoa(index)+"."+"identity.email-option.name", item.IdentityEmailName.ValueString())
			}
			if !item.IdentityEmailDomain.IsNull() && !item.IdentityEmailDomain.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"peer"+"."+strconv.Itoa(index)+"."+"identity.email-option.domain", item.IdentityEmailDomain.ValueString())
			}
			if !item.IdentityFqdnName.IsNull() && !item.IdentityFqdnName.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"peer"+"."+strconv.Itoa(index)+"."+"identity.fqdn-option.name", item.IdentityFqdnName.ValueString())
			}
			if !item.IdentityFqdnDomain.IsNull() && !item.IdentityFqdnDomain.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"peer"+"."+strconv.Itoa(index)+"."+"identity.fqdn-option.domain", item.IdentityFqdnDomain.ValueString())
			}
			if !item.PreSharedKeyLocalEncryption.IsNull() && !item.PreSharedKeyLocalEncryption.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"peer"+"."+strconv.Itoa(index)+"."+"pre-shared-key.local-option.encryption", item.PreSharedKeyLocalEncryption.ValueString())
			}
			if !item.PreSharedKeyLocal.IsNull() && !item.PreSharedKeyLocal.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"peer"+"."+strconv.Itoa(index)+"."+"pre-shared-key.local-option.key", item.PreSharedKeyLocal.ValueString())
			}
			if !item.PreSharedKeyRemoteEncryption.IsNull() && !item.PreSharedKeyRemoteEncryption.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"peer"+"."+strconv.Itoa(index)+"."+"pre-shared-key.remote-option.encryption", item.PreSharedKeyRemoteEncryption.ValueString())
			}
			if !item.PreSharedKeyRemote.IsNull() && !item.PreSharedKeyRemote.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"peer"+"."+strconv.Itoa(index)+"."+"pre-shared-key.remote-option.key", item.PreSharedKeyRemote.ValueString())
			}
			if !item.PreSharedKeyEncryption.IsNull() && !item.PreSharedKeyEncryption.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"peer"+"."+strconv.Itoa(index)+"."+"pre-shared-key.encryption", item.PreSharedKeyEncryption.ValueString())
			}
			if !item.PreSharedKey.IsNull() && !item.PreSharedKey.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"peer"+"."+strconv.Itoa(index)+"."+"pre-shared-key.key", item.PreSharedKey.ValueString())
			}
		}
	}
	return body
}

func (data *CryptoIKEv2Keyring) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	for i := range data.Peers {
		keys := [...]string{"name"}
		keyValues := [...]string{data.Peers[i].Name.ValueString()}

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
		if value := r.Get("name"); value.Exists() && !data.Peers[i].Name.IsNull() {
			data.Peers[i].Name = types.StringValue(value.String())
		} else {
			data.Peers[i].Name = types.StringNull()
		}
		if value := r.Get("description"); value.Exists() && !data.Peers[i].Description.IsNull() {
			data.Peers[i].Description = types.StringValue(value.String())
		} else {
			data.Peers[i].Description = types.StringNull()
		}
		if value := r.Get("hostname"); value.Exists() && !data.Peers[i].Hostname.IsNull() {
			data.Peers[i].Hostname = types.StringValue(value.String())
		} else {
			data.Peers[i].Hostname = types.StringNull()
		}
		if value := r.Get("address.ipv4.ipv4-address"); value.Exists() && !data.Peers[i].Ipv4Address.IsNull() {
			data.Peers[i].Ipv4Address = types.StringValue(value.String())
		} else {
			data.Peers[i].Ipv4Address = types.StringNull()
		}
		if value := r.Get("address.ipv4.ipv4-mask"); value.Exists() && !data.Peers[i].Ipv4Mask.IsNull() {
			data.Peers[i].Ipv4Mask = types.StringValue(value.String())
		} else {
			data.Peers[i].Ipv4Mask = types.StringNull()
		}
		if value := r.Get("address.ipv6-prefix"); value.Exists() && !data.Peers[i].Ipv6Prefix.IsNull() {
			data.Peers[i].Ipv6Prefix = types.StringValue(value.String())
		} else {
			data.Peers[i].Ipv6Prefix = types.StringNull()
		}
		if value := r.Get("identity.key-id-number"); value.Exists() && !data.Peers[i].IdentityKeyId.IsNull() {
			data.Peers[i].IdentityKeyId = types.StringValue(value.String())
		} else {
			data.Peers[i].IdentityKeyId = types.StringNull()
		}
		if value := r.Get("identity.address-type"); value.Exists() && !data.Peers[i].IdentityAddress.IsNull() {
			data.Peers[i].IdentityAddress = types.StringValue(value.String())
		} else {
			data.Peers[i].IdentityAddress = types.StringNull()
		}
		if value := r.Get("identity.email-option.name"); value.Exists() && !data.Peers[i].IdentityEmailName.IsNull() {
			data.Peers[i].IdentityEmailName = types.StringValue(value.String())
		} else {
			data.Peers[i].IdentityEmailName = types.StringNull()
		}
		if value := r.Get("identity.email-option.domain"); value.Exists() && !data.Peers[i].IdentityEmailDomain.IsNull() {
			data.Peers[i].IdentityEmailDomain = types.StringValue(value.String())
		} else {
			data.Peers[i].IdentityEmailDomain = types.StringNull()
		}
		if value := r.Get("identity.fqdn-option.name"); value.Exists() && !data.Peers[i].IdentityFqdnName.IsNull() {
			data.Peers[i].IdentityFqdnName = types.StringValue(value.String())
		} else {
			data.Peers[i].IdentityFqdnName = types.StringNull()
		}
		if value := r.Get("identity.fqdn-option.domain"); value.Exists() && !data.Peers[i].IdentityFqdnDomain.IsNull() {
			data.Peers[i].IdentityFqdnDomain = types.StringValue(value.String())
		} else {
			data.Peers[i].IdentityFqdnDomain = types.StringNull()
		}
		if value := r.Get("pre-shared-key.local-option.encryption"); value.Exists() && !data.Peers[i].PreSharedKeyLocalEncryption.IsNull() {
			data.Peers[i].PreSharedKeyLocalEncryption = types.StringValue(value.String())
		} else {
			data.Peers[i].PreSharedKeyLocalEncryption = types.StringNull()
		}
		if value := r.Get("pre-shared-key.local-option.key"); value.Exists() && !data.Peers[i].PreSharedKeyLocal.IsNull() {
			data.Peers[i].PreSharedKeyLocal = types.StringValue(value.String())
		} else {
			data.Peers[i].PreSharedKeyLocal = types.StringNull()
		}
		if value := r.Get("pre-shared-key.remote-option.encryption"); value.Exists() && !data.Peers[i].PreSharedKeyRemoteEncryption.IsNull() {
			data.Peers[i].PreSharedKeyRemoteEncryption = types.StringValue(value.String())
		} else {
			data.Peers[i].PreSharedKeyRemoteEncryption = types.StringNull()
		}
		if value := r.Get("pre-shared-key.remote-option.key"); value.Exists() && !data.Peers[i].PreSharedKeyRemote.IsNull() {
			data.Peers[i].PreSharedKeyRemote = types.StringValue(value.String())
		} else {
			data.Peers[i].PreSharedKeyRemote = types.StringNull()
		}
		if value := r.Get("pre-shared-key.encryption"); value.Exists() && !data.Peers[i].PreSharedKeyEncryption.IsNull() {
			data.Peers[i].PreSharedKeyEncryption = types.StringValue(value.String())
		} else {
			data.Peers[i].PreSharedKeyEncryption = types.StringNull()
		}
		if value := r.Get("pre-shared-key.key"); value.Exists() && !data.Peers[i].PreSharedKey.IsNull() {
			data.Peers[i].PreSharedKey = types.StringValue(value.String())
		} else {
			data.Peers[i].PreSharedKey = types.StringNull()
		}
	}
}

func (data *CryptoIKEv2KeyringData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "peer"); value.Exists() {
		data.Peers = make([]CryptoIKEv2KeyringPeers, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CryptoIKEv2KeyringPeers{}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			}
			if cValue := v.Get("description"); cValue.Exists() {
				item.Description = types.StringValue(cValue.String())
			}
			if cValue := v.Get("hostname"); cValue.Exists() {
				item.Hostname = types.StringValue(cValue.String())
			}
			if cValue := v.Get("address.ipv4.ipv4-address"); cValue.Exists() {
				item.Ipv4Address = types.StringValue(cValue.String())
			}
			if cValue := v.Get("address.ipv4.ipv4-mask"); cValue.Exists() {
				item.Ipv4Mask = types.StringValue(cValue.String())
			}
			if cValue := v.Get("address.ipv6-prefix"); cValue.Exists() {
				item.Ipv6Prefix = types.StringValue(cValue.String())
			}
			if cValue := v.Get("identity.key-id-number"); cValue.Exists() {
				item.IdentityKeyId = types.StringValue(cValue.String())
			}
			if cValue := v.Get("identity.address-type"); cValue.Exists() {
				item.IdentityAddress = types.StringValue(cValue.String())
			}
			if cValue := v.Get("identity.email-option.name"); cValue.Exists() {
				item.IdentityEmailName = types.StringValue(cValue.String())
			}
			if cValue := v.Get("identity.email-option.domain"); cValue.Exists() {
				item.IdentityEmailDomain = types.StringValue(cValue.String())
			}
			if cValue := v.Get("identity.fqdn-option.name"); cValue.Exists() {
				item.IdentityFqdnName = types.StringValue(cValue.String())
			}
			if cValue := v.Get("identity.fqdn-option.domain"); cValue.Exists() {
				item.IdentityFqdnDomain = types.StringValue(cValue.String())
			}
			if cValue := v.Get("pre-shared-key.local-option.encryption"); cValue.Exists() {
				item.PreSharedKeyLocalEncryption = types.StringValue(cValue.String())
			}
			if cValue := v.Get("pre-shared-key.local-option.key"); cValue.Exists() {
				item.PreSharedKeyLocal = types.StringValue(cValue.String())
			}
			if cValue := v.Get("pre-shared-key.remote-option.encryption"); cValue.Exists() {
				item.PreSharedKeyRemoteEncryption = types.StringValue(cValue.String())
			}
			if cValue := v.Get("pre-shared-key.remote-option.key"); cValue.Exists() {
				item.PreSharedKeyRemote = types.StringValue(cValue.String())
			}
			if cValue := v.Get("pre-shared-key.encryption"); cValue.Exists() {
				item.PreSharedKeyEncryption = types.StringValue(cValue.String())
			}
			if cValue := v.Get("pre-shared-key.key"); cValue.Exists() {
				item.PreSharedKey = types.StringValue(cValue.String())
			}
			data.Peers = append(data.Peers, item)
			return true
		})
	}
}

func (data *CryptoIKEv2Keyring) getDeletedItems(ctx context.Context, state CryptoIKEv2Keyring) []string {
	deletedItems := make([]string, 0)
	for i := range state.Peers {
		stateKeyValues := [...]string{state.Peers[i].Name.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.Peers[i].Name.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Peers {
			found = true
			if state.Peers[i].Name.ValueString() != data.Peers[j].Name.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			deletedItems = append(deletedItems, fmt.Sprintf("%v/peer=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	return deletedItems
}

func (data *CryptoIKEv2Keyring) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)

	return emptyLeafsDelete
}

func (data *CryptoIKEv2Keyring) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	for i := range data.Peers {
		keyValues := [...]string{data.Peers[i].Name.ValueString()}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/peer=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	return deletePaths
}
