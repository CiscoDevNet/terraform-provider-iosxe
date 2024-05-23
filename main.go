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

package main

import (
	"context"
	"log"

	"github.com/CiscoDevNet/terraform-provider-iosxe/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

// Download the YANG models.
//go:generate go run gen/load_models.go

// Having YANG models as input, generate code (Go and Terraform).
//go:generate go run gen/generator.go

// Format Go code and cleanup imports.
//go:generate go run golang.org/x/tools/cmd/goimports -w internal/provider/

// Format Terraform code.
//go:generate terraform fmt -recursive ./examples/

// Having Terraform code as input, generate documentation.
// Run the docs generation tool, check its repository for more information on how it works and how docs
// can be customized.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

// Update documentation categories.
//go:generate go run gen/doc_category.go

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary
	version string = "dev"

	// goreleaser can also pass the specific commit if you want
	// commit  string = ""
)

func main() {
	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/CiscoDevNet/iosxe",
	}

	err := providerserver.Serve(context.Background(), provider.New(version), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
