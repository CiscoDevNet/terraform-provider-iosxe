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

// Hand-written support for iosxe_wireless_ap_tag_source_priority (not generated).
//
// The C9800 only lets priorities 2 and 3 of the AP tag source table be changed; priorities 0
// (static), 1 (location) and 4 (default) are fixed. It also validates every single edit-config
// on its own and rejects one that, even transiently, assigns the same source to two priorities
// ("Tag Source is already assigned to Priority N"), whether the swap is sent as a merge, a
// container replace or a delete plus create in one payload. Swapping the sources of priorities
// 2 and 3 therefore takes two edit-config operations: first remove the priorities whose source
// changes, then write the desired priorities.
//
// This file also holds ValidateConfig for the resource, so that it survives code generation.

import (
	"context"
	"fmt"

	"github.com/CiscoDevNet/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-netconf"
	"github.com/netascode/xmldot"
)

// wirelessAPTagSourceFactoryPriorities is the factory order of the configurable priorities. It is
// restored when the resource is destroyed.
var wirelessAPTagSourceFactoryPriorities = map[int64]string{
	2: "tag-source-filter",
	3: "tag-source-ap-pnp",
}

var _ resource.ResourceWithValidateConfig = &WirelessAPTagSourcePriorityResource{}

// wirelessAPTagSourceConfigurablePriorities lists the configurable priorities in write order.
var wirelessAPTagSourceConfigurablePriorities = []int64{2, 3}

// onlyConfigurableTagSourcePriorities drops the fixed priorities (0, 1 and 4) read from the device.
func onlyConfigurableTagSourcePriorities(items []WirelessAPTagSourcePriorityPriorities) []WirelessAPTagSourcePriorityPriorities {
	if items == nil {
		return nil
	}
	result := make([]WirelessAPTagSourcePriorityPriorities, 0, len(items))
	for _, item := range items {
		if _, ok := wirelessAPTagSourceFactoryPriorities[item.Priority.ValueInt64()]; ok {
			result = append(result, item)
		}
	}
	return result
}

// tagSourcePriorityTarget returns the complete table of the configurable priorities to write. The
// device requires both priorities to hold a source and every source to be used once, so a
// priority that is not configured gets the source left over by the configured ones, which is
// its factory source when the configured priorities keep the factory order. ValidateConfig has
// already rejected a list that configures a priority or a source twice, so the result is always a
// complete table with two distinct sources.
func tagSourcePriorityTarget(items []WirelessAPTagSourcePriorityPriorities) map[int64]string {
	target := make(map[int64]string, len(wirelessAPTagSourceConfigurablePriorities))
	used := make(map[string]bool, len(wirelessAPTagSourceConfigurablePriorities))
	for _, item := range items {
		if item.Priority.IsNull() || item.Priority.IsUnknown() || item.Source.IsNull() || item.Source.IsUnknown() {
			continue
		}
		target[item.Priority.ValueInt64()] = item.Source.ValueString()
		used[item.Source.ValueString()] = true
	}
	for _, priority := range wirelessAPTagSourceConfigurablePriorities {
		if _, ok := target[priority]; ok {
			continue
		}
		candidates := []string{wirelessAPTagSourceFactoryPriorities[priority]}
		for _, p := range wirelessAPTagSourceConfigurablePriorities {
			candidates = append(candidates, wirelessAPTagSourceFactoryPriorities[p])
		}
		for _, source := range candidates {
			if !used[source] {
				target[priority] = source
				used[source] = true
				break
			}
		}
	}
	return target
}

// tagSourcePriorityConfigErrors checks a configured `priorities` list. The device needs every
// source to be used once, so a priority or a source may appear at most once. Unknown values are
// skipped; they are checked again once known. The checks run at plan time, so that an invalid
// table is rejected before any edit-config is sent.
func tagSourcePriorityConfigErrors(items []WirelessAPTagSourcePriorityPriorities) []string {
	var errs []string
	priorities := make(map[int64]bool, len(items))
	sources := make(map[string]bool, len(items))
	for _, item := range items {
		if !item.Priority.IsNull() && !item.Priority.IsUnknown() {
			priority := item.Priority.ValueInt64()
			if priorities[priority] {
				errs = append(errs, fmt.Sprintf("Priority %d is configured more than once in `priorities`.", priority))
			}
			priorities[priority] = true
		}
		if !item.Source.IsNull() && !item.Source.IsUnknown() {
			source := item.Source.ValueString()
			if sources[source] {
				errs = append(errs, fmt.Sprintf("Source %q is assigned to more than one priority in `priorities`; the device allows each source only once.", source))
			}
			sources[source] = true
		}
	}
	return errs
}

// ValidateConfig rejects a `priorities` list the device would refuse (see
// tagSourcePriorityConfigErrors). It lives in this hand-written file rather than in the generated
// resource, so that it survives code generation.
func (r *WirelessAPTagSourcePriorityResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var list types.List
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("priorities"), &list)...)
	if resp.Diagnostics.HasError() || list.IsNull() || list.IsUnknown() {
		return
	}
	for _, element := range list.Elements() {
		if element.IsUnknown() {
			return
		}
	}
	var items []WirelessAPTagSourcePriorityPriorities
	resp.Diagnostics.Append(list.ElementsAs(ctx, &items, false)...)
	if resp.Diagnostics.HasError() {
		return
	}
	for _, msg := range tagSourcePriorityConfigErrors(items) {
		resp.Diagnostics.AddAttributeError(path.Root("priorities"), "Invalid tag source priorities", msg)
	}
}

// tagSourcePrioritiesToRemove returns the configurable priorities whose current source differs
// from the target. They are removed in a separate edit-config before the target is written.
func tagSourcePrioritiesToRemove(current, target map[int64]string) []int64 {
	var result []int64
	for _, priority := range wirelessAPTagSourceConfigurablePriorities {
		if source, ok := current[priority]; ok && source != target[priority] {
			result = append(result, priority)
		}
	}
	return result
}

// readTagSourcePriorities reads the current priority -> source table from the device.
//
// It reads the running datastore. With AutoCommit disabled, every edit-config of this resource is
// only staged in the candidate datastore until the provider commits, so the table read here does
// not reflect edits staged earlier in the same run (it is stale until the commit). Both steps of
// applyTagSourcePriorities, and the restore on failure, are then staged in candidate and validated
// by the device together at commit time.
func (data WirelessAPTagSourcePriority) readTagSourcePriorities(ctx context.Context, device *IosxeProviderDataDevice) (map[int64]string, error) {
	res, err := device.NetconfClient.GetConfig(ctx, "running", helpers.GetXpathFilter(data.getXPath()))
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve object (%s), got error: %s", data.getPath(), err)
	}
	current := make(map[int64]string)
	helpers.GetFromXPath(res.Res, "data"+data.getXPath()+"/tag-source-priority-config").ForEach(
		func(_ int, v xmldot.Result) bool {
			current[v.Get("priority").Int()] = v.Get("tag-src").String()
			return true
		},
	)
	return current, nil
}

// tagSourcePriorityWrite returns the edit-config body that writes the given priorities from table.
func (data WirelessAPTagSourcePriority) tagSourcePriorityWrite(ctx context.Context, priorities []int64, table map[int64]string) string {
	items := make([]WirelessAPTagSourcePriorityPriorities, 0, len(priorities))
	for _, priority := range priorities {
		if source, ok := table[priority]; ok {
			items = append(items, WirelessAPTagSourcePriorityPriorities{
				Priority: types.Int64Value(priority),
				Source:   types.StringValue(source),
			})
		}
	}
	if len(items) == 0 {
		return ""
	}
	write := WirelessAPTagSourcePriority{Device: data.Device, Priorities: items}
	return write.toBodyXML(ctx, write)
}

// applyTagSourcePriorities writes the target table to the device. Priorities whose source changes
// are removed in a first edit-config, so that no single edit assigns a source to two priorities.
//
// If the second edit-config fails after the first one removed priorities, the device would be left
// without them. A best-effort restore of the removed priorities to their previous sources is then
// attempted, and the returned error states whether the previous table was restored.
func (data WirelessAPTagSourcePriority) applyTagSourcePriorities(ctx context.Context, device *IosxeProviderDataDevice, target map[int64]string) error {
	current, err := data.readTagSourcePriorities(ctx, device)
	if err != nil {
		return err
	}

	remove := tagSourcePrioritiesToRemove(current, target)
	if len(remove) > 0 {
		b := netconf.Body{}
		for _, priority := range remove {
			b = helpers.RemoveFromXPath(b, fmt.Sprintf(data.getXPath()+"/tag-source-priority-config[priority='%d']", priority))
		}
		tflog.Debug(ctx, fmt.Sprintf("%s: Removing tag source priorities %v before writing %v", data.getPath(), remove, target))
		if err := helpers.EditConfig(ctx, device.NetconfClient, b.Res(), device.AutoCommit); err != nil {
			return fmt.Errorf("failed to remove tag source priorities %v (no change was made): %s", remove, err)
		}
	}

	var write []int64
	for _, priority := range wirelessAPTagSourceConfigurablePriorities {
		if source, ok := target[priority]; ok && (current[priority] != source || len(remove) > 0) {
			write = append(write, priority)
		}
	}
	body := data.tagSourcePriorityWrite(ctx, write, target)
	if body == "" {
		return nil
	}
	err = helpers.EditConfig(ctx, device.NetconfClient, body, device.AutoCommit)
	if err == nil || len(remove) == 0 {
		return err
	}

	// The first step removed priorities; put their previous sources back. The remaining priorities
	// still hold their previous sources, so this restores the previous table without assigning a
	// source twice.
	tflog.Debug(ctx, fmt.Sprintf("%s: Writing %v failed, restoring tag source priorities %v to %v", data.getPath(), target, remove, current))
	restoreErr := helpers.EditConfig(ctx, device.NetconfClient, data.tagSourcePriorityWrite(ctx, remove, current), device.AutoCommit)
	if restoreErr != nil {
		return fmt.Errorf("failed to write tag source priorities %v: %s. Priorities %v had already been removed, and restoring their previous sources also failed: %s. The device is left without tag source priorities %v; re-apply, or configure them with `ap tag-source-priority`", target, err, remove, restoreErr, remove)
	}
	return fmt.Errorf("failed to write tag source priorities %v: %s. The previous table %v was restored", target, err, tagSourceConfigurable(current))
}

// tagSourceConfigurable returns only the configurable priorities of a table, for error messages.
func tagSourceConfigurable(table map[int64]string) map[int64]string {
	result := make(map[int64]string, len(wirelessAPTagSourceConfigurablePriorities))
	for _, priority := range wirelessAPTagSourceConfigurablePriorities {
		if source, ok := table[priority]; ok {
			result[priority] = source
		}
	}
	return result
}
