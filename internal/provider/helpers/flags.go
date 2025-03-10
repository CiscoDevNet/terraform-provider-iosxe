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

package helpers

import (
	"context"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func IsFlagImporting(ctx context.Context, req resource.ReadRequest) (bool, diag.Diagnostics) {
	v, diags := req.Private.GetKey(ctx, "importing")

	return slices.Equal(v, []byte("1")), diags
}

// SetFlagImporting checks the respDiags and if they are error-free it sets the `importing` as a private flag inside
// SetKeyer. It appends its own results to respDiags.
//
// The caller must include in respDiags the result of state modification in the first place, to ensure consistency.
// The SetKeyer is something like resp.Private.
func SetFlagImporting(ctx context.Context, importing bool, sk SetKeyer, respDiags *diag.Diagnostics) {
	if respDiags.HasError() {
		return
	}

	b := []byte("0")
	if importing {
		b = []byte("1")
	}

	diags := sk.SetKey(ctx, "importing", b)
	respDiags.Append(diags...)
}

// SetKeyer is something like ReadResponse.Private or ImportStateResponse.Private.
type SetKeyer interface {
	SetKey(ctx context.Context, key string, value []byte) diag.Diagnostics
}

var (
	rr resource.ReadResponse
	ir resource.ImportStateResponse

	// ensure interface match
	_ SetKeyer = rr.Private
	_ SetKeyer = ir.Private
)
