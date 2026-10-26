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
	"fmt"
	"os"
	"testing"

	"github.com/CiscoDevNet/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// testAccProtoV6ProviderFactories are used to instantiate a provider during
// acceptance testing. The factory function will be invoked for every Terraform
// CLI command executed to create a provider server to which the CLI can
// reattach.
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"iosxe": providerserver.NewProtocol6WithError(New("test")()),
}

func testAccPreCheck(t *testing.T) {
	// You can add code here to run prior to any test case execution, for example assertions
	// about the appropriate environment variables being set are common to see in a pre-check
	// function.
	if v := os.Getenv("IOSXE_USERNAME"); v == "" {
		t.Fatal("IOSXE_USERNAME env variable must be set for acceptance tests")
	}
	if v := os.Getenv("IOSXE_PASSWORD"); v == "" {
		t.Fatal("IOSXE_PASSWORD env variable must be set for acceptance tests")
	}
	if v := os.Getenv("IOSXE_HOST"); v == "" {
		if v := os.Getenv("IOSXE_URL"); v == "" {
			t.Fatal("IOSXE_HOST or IOSXE_URL env variable must be set for acceptance tests")
		}
	}
}

// testAccCheckResourceAttrIPv6 checks that an attribute holds the expected IPv6 address or prefix,
// ignoring differences in case or compression (used for attributes with normalize_ipv6).
func testAccCheckResourceAttrIPv6(name, key, expected string) resource.TestCheckFunc {
	return resource.TestCheckResourceAttrWith(name, key, func(value string) error {
		if helpers.NormalizeIPv6Address(value) != helpers.NormalizeIPv6Address(expected) {
			return fmt.Errorf("attribute '%s' expected %q (normalized %q), got %q", key, expected, helpers.NormalizeIPv6Address(expected), value)
		}
		return nil
	})
}

// testAccCheckImportStateIPv6 returns an ImportStateCheckFunc that verifies the imported
// normalize_ipv6 attributes of resources of type resourceType against their expected values,
// ignoring differences in case or compression. These attributes are excluded from
// ImportStateVerify, as import stores the device form of the value.
func testAccCheckImportStateIPv6(resourceType string, expected map[string]string) resource.ImportStateCheckFunc {
	return func(states []*terraform.InstanceState) error {
		found := false
		for _, s := range states {
			if s.Ephemeral.Type != resourceType {
				continue
			}
			found = true
			for key, want := range expected {
				got, ok := s.Attributes[key]
				if !ok {
					return fmt.Errorf("imported %s: attribute '%s' not found", resourceType, key)
				}
				if helpers.NormalizeIPv6Address(got) != helpers.NormalizeIPv6Address(want) {
					return fmt.Errorf("imported %s: attribute '%s' expected %q (normalized %q), got %q", resourceType, key, want, helpers.NormalizeIPv6Address(want), got)
				}
			}
		}
		if !found {
			return fmt.Errorf("no imported resource of type %s found", resourceType)
		}
		return nil
	}
}
