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
	"errors"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/netascode/go-netconf"
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

// TestNewProviderInitializesClientCache verifies that the constructor wires
// up the per-process client cache. Without this, the cache lookup in
// Configure() would nil-panic.
func TestNewProviderInitializesClientCache(t *testing.T) {
	p, ok := New("test")().(*IosxeProvider)
	if !ok {
		t.Fatal("New(\"test\")() did not return *IosxeProvider")
	}
	if p.clientCache == nil {
		t.Fatal("clientCache map is nil; New() must initialize it so Configure() can read/store entries")
	}
	if len(p.clientCache) != 0 {
		t.Fatalf("clientCache should start empty, got %d entries", len(p.clientCache))
	}
}

// TestGetOrCreateCachedDevice_HitsAndMisses verifies the contract Configure()
// relies on: same key reuses the underlying *netconf.Client and *sync.Mutex
// across calls, different keys produce distinct entries, build() runs exactly
// once per key, and policy fields take the current call's values on every hit.
func TestGetOrCreateCachedDevice_HitsAndMisses(t *testing.T) {
	p := New("test")().(*IosxeProvider)

	buildCalls := map[string]int{}
	build := func(label string) func() (*netconf.Client, error) {
		return func() (*netconf.Client, error) {
			buildCalls[label]++
			// Hand back a fresh *netconf.Client per build so pointer-identity
			// checks below are meaningful. We never Open() these clients.
			return &netconf.Client{}, nil
		}
	}

	// First Configure pass: miss, builds clientA.
	devA1, err := p.getOrCreateCachedDevice("hostA", build("A"), true, true, true)
	if err != nil {
		t.Fatalf("first call returned error: %v", err)
	}
	// Second Configure pass, same key: hit, must reuse client and mutex.
	devA2, err := p.getOrCreateCachedDevice("hostA", build("A"), false, false, false)
	if err != nil {
		t.Fatalf("second call returned error: %v", err)
	}
	if devA1.NetconfClient != devA2.NetconfClient {
		t.Error("same cache key produced different *netconf.Client pointers; cache is not reusing the client")
	}
	if devA1.OpMutex != devA2.OpMutex {
		t.Error("same cache key produced different *sync.Mutex pointers; the OpMutex pointer is the whole point of the cache")
	}
	if buildCalls["A"] != 1 {
		t.Errorf("expected build to run once per key, got %d calls", buildCalls["A"])
	}
	// Policy fields must reflect each call's args, not the cached ones.
	if devA1.Managed != true || devA2.Managed != false {
		t.Errorf("policy fields not refreshed on hit: devA1.Managed=%v devA2.Managed=%v", devA1.Managed, devA2.Managed)
	}
	if devA1.ReuseConnection != true || devA2.ReuseConnection != false {
		t.Errorf("ReuseConnection not refreshed on hit: %v vs %v", devA1.ReuseConnection, devA2.ReuseConnection)
	}

	// Different key: miss, distinct client and mutex.
	devB, err := p.getOrCreateCachedDevice("hostB", build("B"), true, true, true)
	if err != nil {
		t.Fatalf("third call returned error: %v", err)
	}
	if devB.NetconfClient == devA1.NetconfClient {
		t.Error("different cache keys collapsed to the same *netconf.Client")
	}
	if devB.OpMutex == devA1.OpMutex {
		t.Error("different cache keys collapsed to the same *sync.Mutex")
	}
}

// TestGetOrCreateCachedDevice_BuildErrorNotCached verifies a failed build
// doesn't poison the cache: a subsequent call with the same key retries.
func TestGetOrCreateCachedDevice_BuildErrorNotCached(t *testing.T) {
	p := New("test")().(*IosxeProvider)

	buildErr := errors.New("dial failed")
	calls := 0
	failingBuild := func() (*netconf.Client, error) {
		calls++
		return nil, buildErr
	}
	if _, err := p.getOrCreateCachedDevice("hostX", failingBuild, true, true, true); !errors.Is(err, buildErr) {
		t.Fatalf("expected build error to surface, got %v", err)
	}
	if _, err := p.getOrCreateCachedDevice("hostX", failingBuild, true, true, true); !errors.Is(err, buildErr) {
		t.Fatalf("expected build error to surface again, got %v", err)
	}
	if calls != 2 {
		t.Errorf("expected build to be retried after failure, got %d calls", calls)
	}
	if _, ok := p.clientCache["hostX"]; ok {
		t.Error("failed build was cached; the next caller would see a nil client")
	}
}
