package helpers

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/netascode/go-netconf"
	"github.com/netascode/xmldot"
)

// TestSplitXPathSegments tests the splitXPathSegments function
func TestSplitXPathSegments(t *testing.T) {
	tests := []struct {
		name     string
		xPath    string
		expected []string
	}{
		{
			name:     "simple path",
			xPath:    "native/interface",
			expected: []string{"native", "interface"},
		},
		{
			name:     "path with single predicate",
			xPath:    "native/interface[name='GigabitEthernet1']",
			expected: []string{"native", "interface[name='GigabitEthernet1']"},
		},
		{
			name:     "path with slash in predicate value",
			xPath:    "interface[name='GigabitEthernet1/0/1']/description",
			expected: []string{"interface[name='GigabitEthernet1/0/1']", "description"},
		},
		{
			name:     "complex path with multiple slashes in predicate",
			xPath:    "native/interface[name='GigabitEthernet1/0/1']/ip/address",
			expected: []string{"native", "interface[name='GigabitEthernet1/0/1']", "ip", "address"},
		},
		{
			name:     "multiple predicates with slashes",
			xPath:    "interface[name='Gi1/0/1'][desc='port 1/0/1']",
			expected: []string{"interface[name='Gi1/0/1'][desc='port 1/0/1']"},
		},
		{
			name:     "nested path with composite keys containing slashes",
			xPath:    "native/interface[name='Gi1/0/1']/vrf[name='VRF1']/address",
			expected: []string{"native", "interface[name='Gi1/0/1']", "vrf[name='VRF1']", "address"},
		},
		{
			name:     "empty path",
			xPath:    "",
			expected: []string{},
		},
		{
			name:     "single segment",
			xPath:    "native",
			expected: []string{"native"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := splitXPathSegments(tt.xPath)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("splitXPathSegments() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestSplitDotSegments tests the splitDotSegments function
func TestSplitDotSegments(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected []string
	}{
		{
			name:     "simple path",
			path:     "native.interface",
			expected: []string{"native", "interface"},
		},
		{
			name:     "path with predicate containing dot",
			path:     "Port-channel[name=10.666].ip.Cisco-IOS-XE-icmp:unreachables",
			expected: []string{"Port-channel[name=10.666]", "ip", "Cisco-IOS-XE-icmp:unreachables"},
		},
		{
			name:     "path with multiple dots in predicate",
			path:     "route[prefix=192.168.1.0].next-hop",
			expected: []string{"route[prefix=192.168.1.0]", "next-hop"},
		},
		{
			name:     "empty path",
			path:     "",
			expected: []string{},
		},
		{
			name:     "single segment",
			path:     "native",
			expected: []string{"native"},
		},
		{
			name:     "path without predicates",
			path:     "native.ip.proxy-arp",
			expected: []string{"native", "ip", "proxy-arp"},
		},
		{
			name:     "nested predicates with dots",
			path:     "interface[name=Gi1/0.1].vrf[name=VRF.1].address",
			expected: []string{"interface[name=Gi1/0.1]", "vrf[name=VRF.1]", "address"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := splitDotSegments(tt.path)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("splitDotSegments() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestAugmentNamespaces_DotInPredicateValue tests that dots inside predicate values
// (e.g., Port-channel subinterface name=10.666) don't corrupt the XML namespace augmentation
func TestAugmentNamespaces_DotInPredicateValue(t *testing.T) {
	body := netconf.Body{}
	// Simulate SetFromXPath for a port-channel subinterface with dot in name
	// The path after SetFromXPath conversion would be:
	// Port-channel[name=10.666].ip.Cisco-IOS-XE-icmp:unreachables
	result := SetFromXPath(body, "/Cisco-IOS-XE-native:native/interface/Port-channel-subinterface/Port-channel[name=10.666]/ip/Cisco-IOS-XE-icmp:unreachables", "false")

	resultXML := result.Res()
	t.Logf("Generated XML:\n%s", resultXML)

	// Should NOT contain corrupted element from split predicate value
	if strings.Contains(resultXML, "<666]>") || strings.Contains(resultXML, "</666]>") {
		t.Errorf("Output contains corrupted element from dot-split predicate value\nGot:\n%s", resultXML)
	}

	// Should contain the ICMP namespace on unreachables
	if !strings.Contains(resultXML, "Cisco-IOS-XE-icmp") {
		t.Errorf("Output should contain Cisco-IOS-XE-icmp namespace\nGot:\n%s", resultXML)
	}

	// Verify unreachables has the correct namespace
	icmpNS := "http://cisco.com/ns/yang/Cisco-IOS-XE-icmp"
	matched, _ := regexp.MatchString(`<unreachables[^>]*xmlns="`+regexp.QuoteMeta(icmpNS)+`"`, resultXML)
	if !matched {
		t.Errorf("unreachables element should have xmlns=%q\nGot:\n%s", icmpNS, resultXML)
	}
}

// TestParseXPathSegment tests the parseXPathSegment function with various XPath formats
func TestParseXPathSegment(t *testing.T) {
	tests := []struct {
		name        string
		segment     string
		wantElement string
		wantKeys    map[string]string
	}{
		{
			name:        "simple element without predicate",
			segment:     "interface",
			wantElement: "interface",
			wantKeys:    nil,
		},
		{
			name:        "element with single predicate single quotes",
			segment:     "interface[name='GigabitEthernet1']",
			wantElement: "interface",
			wantKeys:    map[string]string{"name": "GigabitEthernet1"},
		},
		{
			name:        "element with single predicate double quotes",
			segment:     "interface[name=\"GigabitEthernet1\"]",
			wantElement: "interface",
			wantKeys:    map[string]string{"name": "GigabitEthernet1"},
		},
		{
			name:        "element with multiple separate predicates",
			segment:     "interface[name='GigabitEthernet1'][vrf='VRF1']",
			wantElement: "interface",
			wantKeys:    map[string]string{"name": "GigabitEthernet1", "vrf": "VRF1"},
		},
		{
			name:        "element with multiple predicates using and",
			segment:     "interface[name='GigabitEthernet1' and vrf='VRF1']",
			wantElement: "interface",
			wantKeys:    map[string]string{"name": "GigabitEthernet1", "vrf": "VRF1"},
		},
		{
			name:        "element with three separate predicates",
			segment:     "neighbor[ip='192.168.1.1'][vrf='default'][asn='65000']",
			wantElement: "neighbor",
			wantKeys: map[string]string{
				"ip":  "192.168.1.1",
				"vrf": "default",
				"asn": "65000",
			},
		},
		{
			name:        "element with three predicates using and",
			segment:     "neighbor[ip='192.168.1.1' and vrf='default' and asn='65000']",
			wantElement: "neighbor",
			wantKeys: map[string]string{
				"ip":  "192.168.1.1",
				"vrf": "default",
				"asn": "65000",
			},
		},
		{
			name:        "element with mixed quotes in predicates",
			segment:     "interface[name='GigabitEthernet1'][vrf=\"VRF1\"]",
			wantElement: "interface",
			wantKeys:    map[string]string{"name": "GigabitEthernet1", "vrf": "VRF1"},
		},
		{
			name:        "element with numeric key values",
			segment:     "vlan[id='100']",
			wantElement: "vlan",
			wantKeys:    map[string]string{"id": "100"},
		},
		{
			name:        "element with special characters in value",
			segment:     "interface[name='GigabitEthernet1/0/1']",
			wantElement: "interface",
			wantKeys:    map[string]string{"name": "GigabitEthernet1/0/1"},
		},
		{
			name:        "element with spaces in value",
			segment:     "description[text='test description']",
			wantElement: "description",
			wantKeys:    map[string]string{"text": "test description"},
		},
		{
			name:        "element with empty predicate",
			segment:     "interface[]",
			wantElement: "interface",
			wantKeys:    map[string]string{},
		},
		{
			name:        "element name with hyphen",
			segment:     "access-list[name='TEST']",
			wantElement: "access-list",
			wantKeys:    map[string]string{"name": "TEST"},
		},
		{
			name:        "element name with underscore",
			segment:     "bgp_neighbor[ip='10.0.0.1']",
			wantElement: "bgp_neighbor",
			wantKeys:    map[string]string{"ip": "10.0.0.1"},
		},
		{
			name:        "composite key with multiple separate predicates and mixed quotes",
			segment:     "route[vrf=\"VRF1\"][destination='10.0.0.0/8'][next-hop='192.168.1.1']",
			wantElement: "route",
			wantKeys: map[string]string{
				"vrf":         "VRF1",
				"destination": "10.0.0.0/8",
				"next-hop":    "192.168.1.1",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotElement, gotKeys := parseXPathSegment(tt.segment)

			if gotElement != tt.wantElement {
				t.Errorf("parseXPathSegment() element = %v, want %v", gotElement, tt.wantElement)
			}

			// Convert []KeyValue to map for comparison
			gotKeysMap := make(map[string]string)
			for _, kv := range gotKeys {
				gotKeysMap[kv.Key] = kv.Value
			}

			// Normalize for comparison: nil and empty map should be treated as equal
			wantKeys := tt.wantKeys
			if wantKeys == nil {
				wantKeys = make(map[string]string)
			}

			if !reflect.DeepEqual(gotKeysMap, wantKeys) {
				t.Errorf("parseXPathSegment() keys = %v, want %v", gotKeysMap, wantKeys)
			}
		})
	}
}

// TestSetFromXPath tests the SetFromXPath function with various XPath patterns
func TestSetFromXPath(t *testing.T) {
	tests := []struct {
		name      string
		xPath     string
		wantPaths []string // Expected paths that should be set in the body
	}{
		{
			name:  "simple single element with key",
			xPath: "/interface[name='GigabitEthernet1']",
			wantPaths: []string{
				"interface.name",
			},
		},
		{
			name:  "nested path with single key",
			xPath: "/native/interface[name='GigabitEthernet1']",
			wantPaths: []string{
				"native.interface.name",
			},
		},
		{
			name:  "element with multiple separate predicates",
			xPath: "/interface[name='GigabitEthernet1'][vrf='VRF1']",
			wantPaths: []string{
				"interface.name",
				"interface.vrf",
			},
		},
		{
			name:  "element with multiple predicates using and",
			xPath: "/interface[name='GigabitEthernet1' and vrf='VRF1']",
			wantPaths: []string{
				"interface.name",
				"interface.vrf",
			},
		},
		{
			name:  "nested path with composite keys",
			xPath: "/native/vrf[name='VRF1']/address-family[type='ipv4']",
			wantPaths: []string{
				"native.vrf.name",
				"native.vrf.address-family.type",
			},
		},
		{
			name:  "deep nested path",
			xPath: "/native/router/bgp[asn='65000']/neighbor[ip='192.168.1.1']",
			wantPaths: []string{
				"native.router.bgp.asn",
				"native.router.bgp.neighbor.ip",
			},
		},
		{
			name:  "path without leading slash",
			xPath: "interface[name='GigabitEthernet1']",
			wantPaths: []string{
				"interface.name",
			},
		},
		{
			name:  "path without predicates",
			xPath: "/native/cdp/holdtime",
			wantPaths: []string{
				"native.cdp.holdtime",
			},
		},
		{
			name:  "mixed path with and without predicates",
			xPath: "/native/interface[name='GigabitEthernet1']/ip/address",
			wantPaths: []string{
				"native.interface.name",
				"native.interface.ip.address",
			},
		},
		{
			name:  "three level composite key",
			xPath: "/native/vrf[name='VRF1']/address-family[type='ipv4'][safi='unicast']",
			wantPaths: []string{
				"native.vrf.name",
				"native.vrf.address-family.type",
				"native.vrf.address-family.safi",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Start with empty body
			body := netconf.Body{}

			// Apply SetFromXPath with empty value (just creating structure)
			result := SetFromXPath(body, tt.xPath, "")

			// Verify each expected path exists in the result
			resultStr := result.Res()
			for _, wantPath := range tt.wantPaths {
				// Check if the path was created in the result
				// We verify by checking if xmldot.Get can find the path
				if !xmldotPathExists(resultStr, wantPath) {
					t.Errorf("SetFromXPath() missing expected path %q in result:\n%s", wantPath, resultStr)
				}
			}
		})
	}
}

// TestSetFromXPath_Values tests that SetFromXPath correctly sets the key values
func TestSetFromXPath_Values(t *testing.T) {
	tests := []struct {
		name       string
		xPath      string
		checkPath  string
		checkValue string
	}{
		{
			name:       "single key value",
			xPath:      "/interface[name='GigabitEthernet1']",
			checkPath:  "interface.name",
			checkValue: "GigabitEthernet1",
		},
		{
			name:       "numeric key value",
			xPath:      "/vlan[id='100']",
			checkPath:  "vlan.id",
			checkValue: "100",
		},
		{
			name:       "value with special characters",
			xPath:      "/interface[name='GigabitEthernet1/0/1']",
			checkPath:  "interface.name",
			checkValue: "GigabitEthernet1/0/1",
		},
		{
			name:       "value with spaces",
			xPath:      "/description[text='test description']",
			checkPath:  "description.text",
			checkValue: "test description",
		},
		{
			name:       "nested with multiple keys",
			xPath:      "/native/interface[name='GigabitEthernet1'][vrf='VRF1']",
			checkPath:  "native.interface.vrf",
			checkValue: "VRF1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := netconf.Body{}

			// Log what we're testing
			t.Logf("Testing xPath: %s", tt.xPath)
			t.Logf("Expect value at path %s to be: %s", tt.checkPath, tt.checkValue)

			result := SetFromXPath(body, tt.xPath, "")

			// Check for errors
			if err := result.Err(); err != nil {
				t.Fatalf("SetFromXPath() returned error: %v", err)
			}

			// Get the actual value at the path
			resultXML := result.Res()
			t.Logf("Generated XML:\n%s", resultXML)

			actualValue := xmldotGetValue(resultXML, tt.checkPath)
			if actualValue != tt.checkValue {
				t.Errorf("SetFromXPath() value at %q = %q, want %q", tt.checkPath, actualValue, tt.checkValue)
			}
		})
	}
}

// TestRemoveFromXPath tests the RemoveFromXPath function
func TestRemoveFromXPath(t *testing.T) {
	tests := []struct {
		name          string
		xPath         string
		wantPaths     []string // Paths that should exist in the structure
		operationPath string   // Path where operation="remove" should be set
	}{
		{
			name:  "simple single element with key",
			xPath: "/interface[name='GigabitEthernet1']",
			wantPaths: []string{
				"interface.name",
			},
			operationPath: "interface",
		},
		{
			name:  "nested path with single key",
			xPath: "/native/interface[name='GigabitEthernet1']",
			wantPaths: []string{
				"native.interface.name",
			},
			operationPath: "native.interface",
		},
		{
			name:  "path without predicates",
			xPath: "/native/cdp/holdtime",
			wantPaths: []string{
				"native.cdp.holdtime",
			},
			operationPath: "native.cdp.holdtime",
		},
		{
			name:  "mixed path with and without predicates",
			xPath: "/native/interface[name='GigabitEthernet1']/ip/address",
			wantPaths: []string{
				"native.interface.name",
				"native.interface.ip.address",
			},
			operationPath: "native.interface.ip.address",
		},
		{
			name:  "deep nested path with multiple keys",
			xPath: "/native/router/bgp[asn='65000']/neighbor[ip='192.168.1.1']/description",
			wantPaths: []string{
				"native.router.bgp.asn",
				"native.router.bgp.neighbor.ip",
				"native.router.bgp.neighbor.description",
			},
			operationPath: "native.router.bgp.neighbor.description",
		},
		{
			name:  "interface with slashes in name",
			xPath: "/native/interface[name='GigabitEthernet1/0/1']/shutdown",
			wantPaths: []string{
				"native.interface.name",
				"native.interface.shutdown",
			},
			operationPath: "native.interface.shutdown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Start with empty body
			body := netconf.Body{}

			// Apply RemoveFromXPath
			result := RemoveFromXPath(body, tt.xPath)

			// Verify each expected path exists in the result
			resultStr := result.Res()
			for _, wantPath := range tt.wantPaths {
				if !xmldotPathExists(resultStr, wantPath) {
					t.Errorf("RemoveFromXPath() missing expected path %q in result:\n%s", wantPath, resultStr)
				}
			}

			// Verify operation="remove" is set on the correct element
			operationAttrPath := tt.operationPath + ".@operation"
			operationValue := xmldotGetValue(resultStr, operationAttrPath)
			if operationValue != "remove" {
				t.Errorf("RemoveFromXPath() operation attribute at %q = %q, want %q\nGenerated XML:\n%s",
					operationAttrPath, operationValue, "remove", resultStr)
			}
		})
	}
}

// TestRemoveFromXPath_Values tests that RemoveFromXPath correctly sets key values and operation attribute
func TestRemoveFromXPath_Values(t *testing.T) {
	tests := []struct {
		name          string
		xPath         string
		checkPath     string
		checkValue    string
		operationPath string
	}{
		{
			name:          "single key with operation on element",
			xPath:         "/interface[name='GigabitEthernet1']",
			checkPath:     "interface.name",
			checkValue:    "GigabitEthernet1",
			operationPath: "interface",
		},
		{
			name:          "nested with key and child element",
			xPath:         "/native/interface[name='GigabitEthernet1']/shutdown",
			checkPath:     "native.interface.name",
			checkValue:    "GigabitEthernet1",
			operationPath: "native.interface.shutdown",
		},
		{
			name:          "interface name with slashes",
			xPath:         "/interface[name='GigabitEthernet1/0/1']",
			checkPath:     "interface.name",
			checkValue:    "GigabitEthernet1/0/1",
			operationPath: "interface",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := netconf.Body{}

			t.Logf("Testing xPath: %s", tt.xPath)
			t.Logf("Expect value at path %s to be: %s", tt.checkPath, tt.checkValue)
			t.Logf("Expect operation='remove' at path: %s", tt.operationPath)

			result := RemoveFromXPath(body, tt.xPath)

			if err := result.Err(); err != nil {
				t.Fatalf("RemoveFromXPath() returned error: %v", err)
			}

			resultXML := result.Res()
			t.Logf("Generated XML:\n%s", resultXML)

			// Check the key value is set correctly
			actualValue := xmldotGetValue(resultXML, tt.checkPath)
			if actualValue != tt.checkValue {
				t.Errorf("RemoveFromXPath() value at %q = %q, want %q", tt.checkPath, actualValue, tt.checkValue)
			}

			// Check operation="remove" is set on the correct element
			operationAttrPath := tt.operationPath + ".@operation"
			operationValue := xmldotGetValue(resultXML, operationAttrPath)
			if operationValue != "remove" {
				t.Errorf("RemoveFromXPath() operation at %q = %q, want %q", operationAttrPath, operationValue, "remove")
			}
		})
	}
}

// TestSetFromXPath_WithValue tests that SetFromXPath can set values at the final path
func TestSetFromXPath_WithValue(t *testing.T) {
	tests := []struct {
		name      string
		xPath     string
		value     string
		checkPath string
	}{
		{
			name:      "set value on leaf element",
			xPath:     "/native/interface[name='GigabitEthernet1']/description",
			value:     "Management Interface",
			checkPath: "native.interface.description",
		},
		{
			name:      "set value on deep path",
			xPath:     "/native/cdp/holdtime",
			value:     "180",
			checkPath: "native.cdp.holdtime",
		},
		{
			name:      "set value with multiple keys in path",
			xPath:     "/native/router/bgp[asn='65000']/neighbor[ip='192.168.1.1']/description",
			value:     "Peer Router",
			checkPath: "native.router.bgp.neighbor.description",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := netconf.Body{}

			result := SetFromXPath(body, tt.xPath, tt.value)

			if err := result.Err(); err != nil {
				t.Fatalf("SetFromXPath() returned error: %v", err)
			}

			resultXML := result.Res()
			t.Logf("Generated XML:\n%s", resultXML)

			actualValue := xmldotGetValue(resultXML, tt.checkPath)
			if actualValue != tt.value {
				t.Errorf("SetFromXPath() value at %q = %q, want %q", tt.checkPath, actualValue, tt.value)
			}
		})
	}
}

// TestGetFromXPath tests the GetFromXPath function with filter conversion
func TestGetFromXPath(t *testing.T) {
	tests := []struct {
		name          string
		xml           string
		xPath         string
		expectedValue string
		shouldExist   bool
	}{
		{
			name:          "simple path without predicates",
			xml:           "<native><cdp><holdtime>180</holdtime></cdp></native>",
			xPath:         "/native/cdp/holdtime",
			expectedValue: "180",
			shouldExist:   true,
		},
		{
			name: "filter by key - single interface",
			xml: `<native>
				<interface><name>GigabitEthernet1</name><description>Management</description></interface>
				<interface><name>GigabitEthernet2</name><description>Uplink</description></interface>
			</native>`,
			xPath:         "/native/interface[name='GigabitEthernet1']/description",
			expectedValue: "Management",
			shouldExist:   true,
		},
		{
			name: "filter correctly returns first match",
			xml: `<native>
				<interface><name>GigabitEthernet1</name><description>First</description></interface>
				<interface><name>GigabitEthernet2</name><description>Second</description></interface>
			</native>`,
			xPath:         "/native/interface[name='GigabitEthernet2']/description",
			expectedValue: "Second",
			shouldExist:   true,
		},
		{
			name:          "path with value containing slashes",
			xml:           "<native><interface><name>GigabitEthernet1/0/1</name><shutdown/></interface></native>",
			xPath:         "/native/interface[name='GigabitEthernet1/0/1']/shutdown",
			expectedValue: "",
			shouldExist:   true,
		},
		{
			name:          "path with multiple predicates",
			xml:           "<native><interface><name>Gi1</name><vrf>VRF1</vrf><ip><address>192.168.1.1</address></ip></interface></native>",
			xPath:         "/native/interface[name='Gi1'][vrf='VRF1']/ip/address",
			expectedValue: "192.168.1.1",
			shouldExist:   true,
		},
		{
			name: "filter with multiple predicates - multiple elements",
			xml: `<native>
				<interface><name>Gi1</name><vrf>VRF1</vrf><ip><address>192.168.1.1</address></ip></interface>
				<interface><name>Gi1</name><vrf>VRF2</vrf><ip><address>192.168.2.1</address></ip></interface>
				<interface><name>Gi2</name><vrf>VRF1</vrf><ip><address>192.168.3.1</address></ip></interface>
			</native>`,
			xPath:         "/native/interface[name='Gi1'][vrf='VRF2']/ip/address",
			expectedValue: "192.168.2.1",
			shouldExist:   true,
		},
		{
			name:          "deep nested path with filtering",
			xml:           "<native><router><bgp><asn>65000</asn><neighbor><ip>10.0.0.1</ip><description>Peer</description></neighbor></bgp></router></native>",
			xPath:         "/native/router/bgp[asn='65000']/neighbor[ip='10.0.0.1']/description",
			expectedValue: "Peer",
			shouldExist:   true,
		},
		{
			name:          "non-existent path",
			xml:           "<native><cdp><holdtime>180</holdtime></cdp></native>",
			xPath:         "/native/cdp/run",
			expectedValue: "",
			shouldExist:   false,
		},
		{
			name: "filter with no matching key",
			xml: `<native>
				<interface><name>GigabitEthernet1</name><description>Mgmt</description></interface>
			</native>`,
			xPath:       "/native/interface[name='GigabitEthernet2']/description",
			shouldExist: false,
		},
		{
			name: "single element with predicate accessing key field - radius case",
			xml: `<data>
				<native xmlns="http://cisco.com/ns/yang/Cisco-IOS-XE-native">
					<radius>
						<server xmlns="http://cisco.com/ns/yang/Cisco-IOS-XE-aaa">
							<id>radius_10.10.15.12</id>
						</server>
					</radius>
				</native>
			</data>`,
			xPath:         "/data/Cisco-IOS-XE-native:native/radius/Cisco-IOS-XE-aaa:server[id='radius_10.10.15.12']/id",
			expectedValue: "radius_10.10.15.12",
			shouldExist:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Wrap XML in a root element and get Result, simulating NETCONF response structure
			wrappedXML := "<root>" + tt.xml + "</root>"
			res := xmldot.Get(wrappedXML, "root")
			result := GetFromXPath(res, tt.xPath)

			// Check existence
			if result.Exists() != tt.shouldExist {
				t.Errorf("GetFromXPath() Exists() = %v, want %v", result.Exists(), tt.shouldExist)
			}

			// Check value if it should exist
			if tt.shouldExist {
				actualValue := result.String()
				if actualValue != tt.expectedValue {
					t.Errorf("GetFromXPath() value = %q, want %q", actualValue, tt.expectedValue)
				}
			}

			t.Logf("XPath: %s -> Value: %q (Exists: %v)", tt.xPath, result.String(), result.Exists())
		})
	}

	// Additional test for multiple elements (array behavior)
	t.Run("multiple elements without predicates returns array", func(t *testing.T) {
		xml := `<native>
			<interface>
				<nve>
					<member-in-one-line>
						<member>
							<vni>
								<vni-range>201000</vni-range>
								<vrf>GREEN</vrf>
							</vni>
							<vni>
								<vni-range>201010</vni-range>
								<vrf>BLUE</vrf>
							</vni>
						</member>
					</member-in-one-line>
				</nve>
			</interface>
		</native>`

		wrappedXML := "<root>" + xml + "</root>"
		res := xmldot.Get(wrappedXML, "root")
		result := GetFromXPath(res, "/native/interface/nve/member-in-one-line/member/vni")

		// Should return an array result
		if !result.IsArray() {
			t.Errorf("GetFromXPath() IsArray() = false, want true for multiple elements")
		}

		// Should iterate over both elements
		count := 0
		vniRanges := []string{}
		vrfs := []string{}

		result.ForEach(func(i int, v xmldot.Result) bool {
			count++
			vniRanges = append(vniRanges, v.Get("vni-range").String())
			vrfs = append(vrfs, v.Get("vrf").String())
			return true
		})

		if count != 2 {
			t.Errorf("GetFromXPath() ForEach count = %d, want 2", count)
		}

		expectedVniRanges := []string{"201000", "201010"}
		expectedVrfs := []string{"GREEN", "BLUE"}

		for i := 0; i < len(expectedVniRanges); i++ {
			if i >= len(vniRanges) || vniRanges[i] != expectedVniRanges[i] {
				t.Errorf("GetFromXPath() vni-range[%d] = %q, want %q", i, vniRanges[i], expectedVniRanges[i])
			}
			if i >= len(vrfs) || vrfs[i] != expectedVrfs[i] {
				t.Errorf("GetFromXPath() vrf[%d] = %q, want %q", i, vrfs[i], expectedVrfs[i])
			}
		}

		t.Logf("XPath returned array with %d elements", count)
	})
}

// TestSetWithNamespaces_SpecialChars tests if SetWithNamespaces handles special characters like "/"
func TestSetWithNamespaces_SpecialChars(t *testing.T) {
	body := netconf.Body{}

	// Test with a value containing "/"
	result := setWithNamespaces(body, "interface.name", "GigabitEthernet1/0/1")

	if err := result.Err(); err != nil {
		t.Fatalf("setWithNamespaces() error: %v", err)
	}

	xml := result.Res()
	t.Logf("Generated XML: %s", xml)

	if xml == "" {
		t.Error("Generated XML is empty")
	}

	value := xmldotGetValue(xml, "interface.name")
	if value != "GigabitEthernet1/0/1" {
		t.Errorf("Expected value 'GigabitEthernet1/0/1', got %q", value)
	}
}

// Helper function to check if a path exists in XML using xmldot
func xmldotPathExists(xml, path string) bool {
	result := xmldot.Get(xml, path)
	return result.Exists()
}

// Helper function to get value at a path in XML using xmldot
func xmldotGetValue(xml, path string) string {
	result := xmldot.Get(xml, path)
	return result.String()
}

// TestGetXpathFilter tests the GetXpathFilter function
func TestGetXpathFilter(t *testing.T) {
	tests := []struct {
		name     string
		xPath    string
		expected string
	}{
		{
			name:     "simple path without namespace",
			xPath:    "/native/interface/GigabitEthernet",
			expected: "/native/interface/GigabitEthernet",
		},
		{
			name:     "path with namespace prefix",
			xPath:    "/Cisco-IOS-XE-native:native/interface/GigabitEthernet",
			expected: "/native/interface/GigabitEthernet",
		},
		{
			name:     "path with predicate and namespace",
			xPath:    "/Cisco-IOS-XE-native:native/interface[Cisco-IOS-XE-native:name='GigabitEthernet1']",
			expected: "/native/interface[name='GigabitEthernet1']",
		},
		{
			name:     "path with multiple predicates",
			xPath:    "/native/interface[name='Gi1'][vrf='VRF1']",
			expected: "/native/interface[name='Gi1'][vrf='VRF1']",
		},
		{
			name:     "nested path with namespace prefixes",
			xPath:    "/Cisco-IOS-XE-native:native/interface[Cisco-IOS-XE-native:name='Gi1']/Cisco-IOS-XE-native:ip/Cisco-IOS-XE-native:address",
			expected: "/native/interface[name='Gi1']/ip/address",
		},
		{
			name:     "path without leading slash",
			xPath:    "Cisco-IOS-XE-native:native/interface/GigabitEthernet",
			expected: "/native/interface/GigabitEthernet",
		},
		{
			name:     "path with slash in predicate value",
			xPath:    "/native/interface[name='GigabitEthernet1/0/1']",
			expected: "/native/interface[name='GigabitEthernet1/0/1']",
		},
		{
			name:     "realistic getXPath output",
			xPath:    "/Cisco-IOS-XE-native:native/aaa",
			expected: "/native/aaa",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetXpathFilter(tt.xPath)

			if result.Content != tt.expected {
				t.Errorf("GetXpathFilter(%q).Content = %q, want %q", tt.xPath, result.Content, tt.expected)
			}

			if result.Type != "xpath" {
				t.Errorf("GetXpathFilter(%q).Type = %q, want %q", tt.xPath, result.Type, "xpath")
			}

			t.Logf("XPath: %s -> Filter: %s", tt.xPath, result.Content)
		})
	}
}

// TestSetRawFromXPath_MultiRoot tests multi-root XML fragment creation
func TestSetRawFromXPath_MultiRoot(t *testing.T) {
	tests := []struct {
		name          string
		xPath         string
		values        []string
		expectedCount int
		checkPaths    []string
		checkValues   []string
	}{
		{
			name:          "multiple interface elements at same path",
			xPath:         "/native/interface",
			values:        []string{"<name>Gi1</name><description>First</description>", "<name>Gi2</name><description>Second</description>"},
			expectedCount: 2,
			checkPaths:    []string{"native.interface.0.name", "native.interface.1.name"},
			checkValues:   []string{"Gi1", "Gi2"},
		},
		{
			name:          "multiple list items with keys",
			xPath:         "/native/router/bgp/neighbor",
			values:        []string{"<ip>10.0.0.1</ip><remote-as>65001</remote-as>", "<ip>10.0.0.2</ip><remote-as>65002</remote-as>", "<ip>10.0.0.3</ip><remote-as>65003</remote-as>"},
			expectedCount: 3,
			checkPaths:    []string{"native.router.bgp.neighbor.0.ip", "native.router.bgp.neighbor.1.ip", "native.router.bgp.neighbor.2.ip"},
			checkValues:   []string{"10.0.0.1", "10.0.0.2", "10.0.0.3"},
		},
		{
			name:          "nested list items",
			xPath:         "/native/access-list/extended/rule",
			values:        []string{"<id>10</id><permit>ip</permit>", "<id>20</id><permit>tcp</permit>"},
			expectedCount: 2,
			checkPaths:    []string{"native.access-list.extended.rule.0.id", "native.access-list.extended.rule.1.id"},
			checkValues:   []string{"10", "20"},
		},
		{
			name:          "single element (baseline test)",
			xPath:         "/native/hostname",
			values:        []string{"<value>router1</value>"},
			expectedCount: 1,
			checkPaths:    []string{"native.hostname.value"},
			checkValues:   []string{"router1"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := netconf.Body{}

			// Add each value sequentially
			for i, value := range tt.values {
				t.Logf("Adding element %d with value: %s", i+1, value)
				body = SetRawFromXPath(body, tt.xPath, value)
			}

			// Check for errors
			if err := body.Err(); err != nil {
				t.Fatalf("SetRawFromXPath() returned error: %v", err)
			}

			resultXML := body.Res()
			t.Logf("Generated XML:\n%s", resultXML)

			// Verify each expected value exists at the correct path
			for i, checkPath := range tt.checkPaths {
				actualValue := xmldotGetValue(resultXML, checkPath)
				expectedValue := tt.checkValues[i]

				if actualValue != expectedValue {
					t.Errorf("Value at %q = %q, want %q", checkPath, actualValue, expectedValue)
				} else {
					t.Logf("✓ Verified: %s = %s", checkPath, actualValue)
				}
			}
		})
	}
}

// TestSetRawFromXPath_MultiRoot_NestedLists tests multi-root with nested list structures
func TestSetRawFromXPath_MultiRoot_NestedLists(t *testing.T) {
	body := netconf.Body{}

	// Create first outer list item with nested inner items
	// Build the inner XML manually for the first list item
	innerXML1 := "<id>pool1</id><interface><name>eth0</name><overload>true</overload></interface><interface><name>eth1</name><overload>false</overload></interface>"

	body = SetRawFromXPath(body, "/native/nat/inside/source/list", innerXML1)

	// Create second outer list item with nested inner items
	innerXML2 := "<id>pool2</id><interface><name>eth2</name><overload>true</overload></interface>"

	body = SetRawFromXPath(body, "/native/nat/inside/source/list", innerXML2)

	// Verify structure
	if err := body.Err(); err != nil {
		t.Fatalf("SetRawFromXPath() returned error: %v", err)
	}

	resultXML := body.Res()
	t.Logf("Generated XML:\n%s", resultXML)

	// Verify outer list items
	pool1ID := xmldotGetValue(resultXML, "native.nat.inside.source.list.0.id")
	pool2ID := xmldotGetValue(resultXML, "native.nat.inside.source.list.1.id")

	if pool1ID != "pool1" {
		t.Errorf("First outer list item id = %q, want %q", pool1ID, "pool1")
	}
	if pool2ID != "pool2" {
		t.Errorf("Second outer list item id = %q, want %q", pool2ID, "pool2")
	}

	// Verify inner list items for first outer item
	inner1a := xmldotGetValue(resultXML, "native.nat.inside.source.list.0.interface.0.name")
	inner1b := xmldotGetValue(resultXML, "native.nat.inside.source.list.0.interface.1.name")

	if inner1a != "eth0" {
		t.Errorf("First inner item name = %q, want %q", inner1a, "eth0")
	}
	if inner1b != "eth1" {
		t.Errorf("Second inner item name = %q, want %q", inner1b, "eth1")
	}

	// Verify inner list items for second outer item
	inner2a := xmldotGetValue(resultXML, "native.nat.inside.source.list.1.interface.0.name")
	if inner2a != "eth2" {
		t.Errorf("Third inner item name = %q, want %q", inner2a, "eth2")
	}

	t.Log("✓ All nested list structures verified successfully")
}

// TestSetRawFromXPath_MultiRoot_WithNamespaces tests multi-root with namespace prefixes
func TestSetRawFromXPath_MultiRoot_WithNamespaces(t *testing.T) {
	body := netconf.Body{}

	// Add multiple elements with namespace prefixes in path
	xml1 := "<name>Gi1</name><shutdown/>"
	xml2 := "<name>Gi2</name><description>Uplink</description>"

	body = SetRawFromXPath(body, "/Cisco-IOS-XE-native:native/interface", xml1)
	body = SetRawFromXPath(body, "/Cisco-IOS-XE-native:native/interface", xml2)

	if err := body.Err(); err != nil {
		t.Fatalf("SetRawFromXPath() returned error: %v", err)
	}

	resultXML := body.Res()
	t.Logf("Generated XML:\n%s", resultXML)

	// Verify both interfaces exist
	name1 := xmldotGetValue(resultXML, "native.interface.0.name")
	name2 := xmldotGetValue(resultXML, "native.interface.1.name")

	if name1 != "Gi1" {
		t.Errorf("First interface name = %q, want %q", name1, "Gi1")
	}
	if name2 != "Gi2" {
		t.Errorf("Second interface name = %q, want %q", name2, "Gi2")
	}

	// Verify namespace declaration exists
	if !xmldotPathExists(resultXML, "native.@xmlns") {
		t.Error("Namespace declaration missing")
	}

	t.Log("✓ Multi-root with namespaces verified successfully")
}

// TestSetFromXPath_MultipleRoots tests that SetFromXPath with different root paths creates multi-root XML
func TestSetFromXPath_MultipleRoots(t *testing.T) {
	// This tests the ACL use case where we have both deny and permit in the same cBody
	cBody := netconf.Body{}

	cBody = SetFromXPath(cBody, "sequence", "10")
	cBody = SetFromXPath(cBody, "deny/std-ace/ipv4-address-prefix", "10.0.0.0")
	cBody = SetFromXPath(cBody, "deny/std-ace/mask", "0.0.0.255")
	cBody = SetFromXPath(cBody, "permit/std-ace/ipv4-address-prefix", "192.168.0.0")
	cBody = SetFromXPath(cBody, "permit/std-ace/mask", "0.0.255.255")

	if err := cBody.Err(); err != nil {
		t.Fatalf("SetFromXPath() returned error: %v", err)
	}

	resultXML := cBody.Res()
	t.Logf("Generated XML:\n%s", resultXML)

	// Check if we have both deny and permit as separate elements
	denyPrefix := xmldotGetValue(resultXML, "deny.std-ace.ipv4-address-prefix")
	permitPrefix := xmldotGetValue(resultXML, "permit.std-ace.ipv4-address-prefix")
	sequence := xmldotGetValue(resultXML, "sequence")

	if sequence != "10" {
		t.Errorf("sequence = %q, want %q", sequence, "10")
	}
	if denyPrefix != "10.0.0.0" {
		t.Errorf("deny prefix = %q, want %q", denyPrefix, "10.0.0.0")
	}
	if permitPrefix != "192.168.0.0" {
		t.Errorf("permit prefix = %q, want %q", permitPrefix, "192.168.0.0")
	}

	t.Log("✓ SetFromXPath with multiple roots verified successfully")
}

// TestSetRawFromXPath_WithPredicates tests that XPath predicates don't create malformed element names
// Regression test for issue where namespace augmentation created elements like <standard[name=SACL1]>
func TestSetRawFromXPath_WithPredicates(t *testing.T) {
	// Simulate creating an ACL structure with predicates in the path
	body := netconf.Body{}

	// First, create the name element (this builds the parent structure with predicates)
	xpathWithPredicate := "Cisco-IOS-XE-native:native/ip/access-list/Cisco-IOS-XE-acl:standard[name=SACL1]"
	body = SetFromXPath(body, xpathWithPredicate+"/name", "SACL1")

	// Then add a child structure using SetRawFromXPath
	childXML := "<sequence>10</sequence><remark>Test</remark>"
	body = SetRawFromXPath(body, xpathWithPredicate+"/access-list-seq-rule", childXML)

	xml, err := body.String()
	if err != nil {
		t.Fatalf("Body.String() error: %v", err)
	}

	t.Logf("Generated XML:\n%s", xml)

	// Verify no malformed elements with predicates in the name
	if strings.Contains(xml, "standard[name") {
		t.Errorf("XML contains malformed element with XPath predicate in name: standard[name=...]")
	}

	// Verify the structure is correct with clean element names
	if !strings.Contains(xml, "<standard") {
		t.Errorf("XML missing <standard> element")
	}
	if !strings.Contains(xml, "<name>SACL1</name>") {
		t.Errorf("XML missing name element")
	}
	if !strings.Contains(xml, "<access-list-seq-rule>") {
		t.Errorf("XML missing access-list-seq-rule element")
	}

	t.Log("✓ XPath predicates properly stripped in namespace augmentation")
}

// TestSetFromXPath_NoDuplicateElements verifies that SetFromXPath doesn't create
// duplicate elements when setting a value at a leaf path. This was a bug where
// paths like "match/authorizing-method-priority/greater-than" with value "20"
// would create both an empty <greater-than></greater-than> and <greater-than>20</greater-than>.
func TestSetFromXPath_NoDuplicateElements(t *testing.T) {
	body := netconf.Body{}

	// Test the case that was causing duplicate elements
	body = SetFromXPath(body, "match/authorizing-method-priority/greater-than", 20)

	xml := body.Res()

	// Count occurrences of "<greater-than>"
	openingTagCount := 0
	search := "<greater-than>"
	str := xml
	for {
		idx := strings.Index(str, search)
		if idx == -1 {
			break
		}
		openingTagCount++
		str = str[idx+len(search):]
	}

	if openingTagCount != 1 {
		t.Errorf("Expected exactly 1 <greater-than> element, found %d. XML:\n%s", openingTagCount, xml)
	}

	// Verify the value is set correctly
	result := xmldot.Get(body.Res(), "match.authorizing-method-priority.greater-than")
	if result.Int() != 20 {
		t.Errorf("Expected value 20, got %d", result.Int())
	}

	t.Logf("✓ No duplicate elements, value correctly set to: %d", result.Int())
}

// TestSetFromXPath_BooleanEmptyValue tests that boolean true values (empty strings)
// correctly create empty elements in XML for NETCONF presence containers.
func TestSetFromXPath_BooleanEmptyValue(t *testing.T) {
	body := netconf.Body{}

	// Test setting an empty string (boolean true in NETCONF)
	body = SetFromXPath(body, "match/authorization-status/authorized", "")

	xml := body.Res()
	t.Logf("Generated XML:\n%s", xml)

	// Check that the authorized element exists
	if !strings.Contains(xml, "<authorized>") && !strings.Contains(xml, "<authorized/>") {
		t.Errorf("Expected <authorized> element to be present in XML")
	}

	// Verify the element exists using xmldot
	result := xmldot.Get(body.Res(), "match.authorization-status.authorized")
	if !result.Exists() {
		t.Errorf("Expected match.authorization-status.authorized to exist")
	}

	t.Log("✓ Boolean empty value correctly creates presence container element")
}

// TestSetFromXPath_ClassMapSequence tests the exact sequence used in class-map
// to understand why authorized element is missing.
func TestSetFromXPath_ClassMapSequence(t *testing.T) {
	body := netconf.Body{}
	path := "Cisco-IOS-XE-native:native/policy/Cisco-IOS-XE-policy:class-map[name=CM1]"

	// authorized (boolean empty)
	body = SetFromXPath(body, path+"/match/authorization-status/authorized", "")
	t.Log("After authorized:")
	t.Log(body.Res())

	// aaa-timeout (boolean empty)
	body = SetFromXPath(body, path+"/match/result-type/aaa-timeout", "")
	t.Log("\nAfter aaa-timeout:")
	t.Log(body.Res())

	// activated-service-template list
	cBody := netconf.Body{}
	cBody = SetFromXPath(cBody, "service-name", "CRITICAL_AUTH_ACCESS")
	body = SetRawFromXPath(body, path+"/match/activated-service-template", cBody.Res())
	t.Log("\nAfter activated-service-template:")
	t.Log(body.Res())

	// Check both elements still exist
	authResult := xmldot.Get(body.Res(), "native.policy.class-map.match.authorization-status.authorized")
	if !authResult.Exists() {
		t.Error("❌ authorized element missing after list addition!")
	} else {
		t.Log("✓ authorized element still exists")
	}

	aaaResult := xmldot.Get(body.Res(), "native.policy.class-map.match.result-type.aaa-timeout")
	if !aaaResult.Exists() {
		t.Error("❌ aaa-timeout element missing after list addition!")
	} else {
		t.Log("✓ aaa-timeout element still exists")
	}
}

// TestSetRawFromXPath_DebugParentPath debugs the parent path resolution
func TestSetRawFromXPath_DebugParentPath(t *testing.T) {
	body := netconf.Body{}
	path := "Cisco-IOS-XE-native:native/policy/Cisco-IOS-XE-policy:class-map[name=CM1]"

	// Set some initial content
	body = SetFromXPath(body, path+"/match/authorization-status/authorized", "")
	body = SetFromXPath(body, path+"/match/result-type/aaa-timeout", "")

	t.Log("Initial XML:")
	t.Log(body.Res())

	// Now add the list item
	t.Log("\nAdding activated-service-template...")
	cBody := netconf.Body{}
	cBody = SetFromXPath(cBody, "service-name", "CRITICAL_AUTH_ACCESS")

	// What is the parent path for activated-service-template?
	fullPath := path + "/match/activated-service-template"
	segments := splitXPathSegments(strings.TrimPrefix(fullPath, "/"))
	t.Logf("Segments: %v", segments)
	t.Logf("Number of segments: %d", len(segments))

	parentPathSegments := make([]string, 0, len(segments)-1)
	for _, segment := range segments[:len(segments)-1] {
		elementName, _ := parseXPathSegment(segment)
		parentPathSegments = append(parentPathSegments, elementName)
	}
	parentPath := dotPath(strings.Join(parentPathSegments, "."))
	t.Logf("Parent path: %s", parentPath)

	// What content exists at parent path?
	existingXML := xmldot.Get(body.Res(), parentPath).Raw
	t.Logf("Existing XML at parent: %s", existingXML)

	body = SetRawFromXPath(body, fullPath, cBody.Res())

	t.Log("\nAfter SetRawFromXPath:")
	t.Log(body.Res())
}

// TestBodySetRaw tests how body.SetRaw behaves
func TestBodySetRaw(t *testing.T) {
	body := netconf.Body{}

	// Create initial structure
	body = body.Set("match.foo", "value1")
	t.Log("Initial:")
	t.Log(body.Res())

	// Now try to SetRaw at match with combined content
	existing := xmldot.Get(body.Res(), "match").Raw
	t.Logf("\nExisting at 'match': %s", existing)

	newContent := "<bar>value2</bar>"
	combined := existing + newContent
	t.Logf("Combined: %s", combined)

	body = body.SetRaw("match", combined)
	t.Log("\nAfter SetRaw:")
	t.Log(body.Res())
}

// TestAppendFromXPath tests the AppendFromXPath function
func TestAppendFromXPath(t *testing.T) {
	body := netconf.Body{}
	path := "native/route-map/rule/match/ip/address"

	// Add first item
	body = AppendFromXPath(body, path, "10")
	t.Log("After first append:")
	t.Log(body.Res())

	// Add second item
	body = AppendFromXPath(body, path, "20")
	t.Log("\nAfter second append:")
	t.Log(body.Res())

	// Add third item
	body = AppendFromXPath(body, path, "30")
	t.Log("\nAfter third append:")
	t.Log(body.Res())

	// Verify all three items exist
	result1 := xmldot.Get(body.Res(), "native.route-map.rule.match.ip.address.0")
	result2 := xmldot.Get(body.Res(), "native.route-map.rule.match.ip.address.1")
	result3 := xmldot.Get(body.Res(), "native.route-map.rule.match.ip.address.2")

	if result1.String() != "10" {
		t.Errorf("Expected first item to be '10', got '%s'", result1.String())
	}
	if result2.String() != "20" {
		t.Errorf("Expected second item to be '20', got '%s'", result2.String())
	}
	if result3.String() != "30" {
		t.Errorf("Expected third item to be '30', got '%s'", result3.String())
	}

	t.Log("✓ All three items appended successfully")
}

// TestAppendFromXPath_WithNamespaces tests AppendFromXPath with namespace prefixes
func TestAppendFromXPath_WithNamespaces(t *testing.T) {
	body := netconf.Body{}
	path := "Cisco-IOS-XE-native:native/Cisco-IOS-XE-policy:class-map[name=CM1]/match/dscp"

	// Append multiple dscp values
	body = AppendFromXPath(body, path, 8)
	body = AppendFromXPath(body, path, 16)
	body = AppendFromXPath(body, path, 24)

	t.Log("Generated XML with namespaces:")
	t.Log(body.Res())

	// Verify all items exist
	result1 := xmldot.Get(body.Res(), "native.class-map.match.dscp.0")
	result2 := xmldot.Get(body.Res(), "native.class-map.match.dscp.1")
	result3 := xmldot.Get(body.Res(), "native.class-map.match.dscp.2")

	if result1.Int() != 8 {
		t.Errorf("Expected first dscp to be 8, got %d", result1.Int())
	}
	if result2.Int() != 16 {
		t.Errorf("Expected second dscp to be 16, got %d", result2.Int())
	}
	if result3.Int() != 24 {
		t.Errorf("Expected third dscp to be 24, got %d", result3.Int())
	}

	// Verify namespace declarations
	if !strings.Contains(body.Res(), `xmlns="http://cisco.com/ns/yang/Cisco-IOS-XE-native"`) {
		t.Error("Missing Cisco-IOS-XE-native namespace declaration")
	}
	if !strings.Contains(body.Res(), `xmlns="http://cisco.com/ns/yang/Cisco-IOS-XE-policy"`) {
		t.Error("Missing Cisco-IOS-XE-policy namespace declaration")
	}

	t.Log("✓ List items appended with proper namespaces")
}

// TestAppendFromXPath_EmptyValue tests AppendFromXPath with empty values (boolean presence)
func TestAppendFromXPath_EmptyValue(t *testing.T) {
	body := netconf.Body{}
	path := "native/feature/item"

	// Append empty values (presence containers)
	body = AppendFromXPath(body, path, "")
	body = AppendFromXPath(body, path, "")

	t.Log("Generated XML with empty values:")
	t.Log(body.Res())

	// Verify structure exists
	if !strings.Contains(body.Res(), "<item></item>") && !strings.Contains(body.Res(), "<item/>") {
		t.Error("Expected empty <item> elements")
	}

	t.Log("✓ Empty values (presence containers) appended successfully")
}

// TestRemoveFromXPath_DebugCleanup tests the cleanup mechanism in isolation
func TestRemoveFromXPath_DebugCleanup(t *testing.T) {
	body := netconf.Body{}

	// Add child first
	t.Log("Adding child removal: /logging/trap/severity")
	body = RemoveFromXPath(body, "/logging/trap/severity")
	xml1 := body.Res()
	t.Logf("After child:\n%s\n", xml1)

	// Add parent
	t.Log("Adding parent removal: /logging/trap")
	body = RemoveFromXPath(body, "/logging/trap")
	xml2 := body.Res()
	t.Logf("After parent:\n%s\n", xml2)

	// Check if child element was completely removed
	severityElement := xmldot.Get(xml2, "logging.trap.severity")
	if severityElement.Exists() {
		t.Errorf("Child element should be completely removed, but still exists: %q", severityElement.Raw)
	} else {
		t.Log("✓ Child element was successfully removed")
	}

	// Also check operation attribute doesn't exist
	severityOp := xmldot.Get(xml2, "logging.trap.severity.@operation")
	if severityOp.Exists() {
		t.Errorf("Child operation attribute should not exist, but found: %q", severityOp.String())
	}
}

// TestRemoveFromXPath_ExactAddDeletePathsXMLSequence tests the exact sequence used by generated code
func TestRemoveFromXPath_ExactAddDeletePathsXMLSequence(t *testing.T) {
	body := netconf.Body{}
	basePath := "/Cisco-IOS-XE-native:native/logging"

	// Simulate the exact sequence from addDeletePathsXML
	body = RemoveFromXPath(body, basePath+"/trap/severity")
	body = RemoveFromXPath(body, basePath+"/trap")

	xml := body.Res()
	t.Logf("Generated XML:\n%s\n", xml)

	// Check that severity element is completely gone
	severityElement := xmldot.Get(xml, "native.logging.trap.severity")
	severityOp := xmldot.Get(xml, "native.logging.trap.severity.@operation")

	t.Logf("Severity element exists: %v, Raw: %q", severityElement.Exists(), severityElement.Raw)
	t.Logf("Severity operation exists: %v, value: %q", severityOp.Exists(), severityOp.String())

	// Check if severity element has operation="remove"
	if strings.Contains(xml, `<severity operation="remove">`) {
		t.Error("❌ XML still contains <severity operation=\"remove\">")
	} else if severityElement.Exists() {
		t.Errorf("❌ Severity element exists but shouldn't: Raw=%q", severityElement.Raw)
	} else {
		t.Log("✓ Severity element correctly removed")
	}

	// Verify trap still has operation="remove"
	trapOp := xmldot.Get(xml, "native.logging.trap.@operation")
	if !trapOp.Exists() || trapOp.String() != "remove" {
		t.Errorf("❌ Trap should have operation='remove', but got exists=%v, value=%q", trapOp.Exists(), trapOp.String())
	} else {
		t.Log("✓ Trap has operation='remove'")
	}
}

// TestRemoveFromXPath_RealWorldLoggingExample tests the exact scenario from the user's example
func TestRemoveFromXPath_RealWorldLoggingExample(t *testing.T) {
	body := netconf.Body{}
	basePath := "/Cisco-IOS-XE-native:native/logging"

	// Apply the exact sequence from the user's debug log
	body = RemoveFromXPath(body, basePath+"/monitor-config/common-config/monitor/severity")
	body = RemoveFromXPath(body, basePath+"/buffered/size-value")
	body = RemoveFromXPath(body, basePath+"/buffered/severity-level")
	body = RemoveFromXPath(body, basePath+"/console-config/console")
	body = RemoveFromXPath(body, basePath+"/facility")
	body = RemoveFromXPath(body, basePath+"/history/size")
	body = RemoveFromXPath(body, basePath+"/history/severity-level")
	body = RemoveFromXPath(body, basePath+"/trap")
	body = RemoveFromXPath(body, basePath+"/trap/severity") // Should be skipped
	body = RemoveFromXPath(body, basePath+"/origin-id/type-value")
	body = RemoveFromXPath(body, basePath+"/source-interface-conf/interface-name-non-vrf")
	body = RemoveFromXPath(body, basePath+"/source-interface-conf/source-interface-vrf[vrf='VRF1']")
	body = RemoveFromXPath(body, basePath+"/host/ipv4-host-list[ipv4-host='1.1.1.1']")
	body = RemoveFromXPath(body, basePath+"/host/ipv4-host-transport-list[ipv4-host='1.1.1.1']")
	body = RemoveFromXPath(body, basePath+"/host/ipv4-host-vrf-list[ipv4-host='1.1.1.1'][vrf='VRF1']")
	body = RemoveFromXPath(body, basePath+"/host/ipv4-host-vrf-transport-list[ipv4-host='1.1.1.1'][vrf='VRF1']")
	body = RemoveFromXPath(body, basePath+"/host/ipv6/ipv6-host-list[ipv6-host='2001::1']")
	body = RemoveFromXPath(body, basePath+"/host/ipv6/ipv6-host-transport-list[ipv6-host='2001::1']")
	body = RemoveFromXPath(body, basePath+"/host/ipv6/ipv6-host-vrf-list[ipv6-host='2001::1'][vrf='VRF1']")
	body = RemoveFromXPath(body, basePath+"/host/ipv6/ipv6-host-vrf-transport-list[ipv6-host='2001::1'][vrf='VRF1']")

	resultXML := body.Res()
	t.Logf("Generated XML:\n%s\n", resultXML)

	// Verify that trap has operation="remove"
	trapOp := xmldot.Get(resultXML, "native.logging.trap.@operation")
	if !trapOp.Exists() || trapOp.String() != "remove" {
		t.Errorf("Expected trap to have operation='remove'")
	}

	// Verify that trap/severity does NOT have operation="remove" (redundant)
	trapSeverityOp := xmldot.Get(resultXML, "native.logging.trap.severity.@operation")
	if trapSeverityOp.Exists() {
		t.Errorf("trap/severity should NOT have operation='remove' (redundant), but found: %q", trapSeverityOp.String())
	}

	// Verify namespace is present
	if !strings.Contains(resultXML, `xmlns="http://cisco.com/ns/yang/Cisco-IOS-XE-native"`) {
		t.Error("Missing namespace declaration for Cisco-IOS-XE-native")
	}

	t.Log("✓ Optimized XML successfully removes redundant child operations")
}

// TestRemoveFromXPath_SkipsRedundantChildRemovals tests that RemoveFromXPath skips
// adding child removal operations when a parent is already marked for removal
func TestRemoveFromXPath_SkipsRedundantChildRemovals(t *testing.T) {
	tests := []struct {
		name           string
		removePaths    []string        // Paths to remove in order
		shouldExist    map[string]bool // path -> should have operation="remove"
		shouldNotExist []string        // paths that should not exist
	}{
		{
			name: "child removal skipped when parent removed first",
			removePaths: []string{
				"/logging/trap",
				"/logging/trap/severity",
			},
			shouldExist: map[string]bool{
				"logging.trap.@operation": true,
			},
			shouldNotExist: []string{
				"logging.trap.severity.@operation",
			},
		},
		{
			name: "child removal cleaned when parent added after",
			removePaths: []string{
				"/logging/trap/severity",
				"/logging/trap",
			},
			shouldExist: map[string]bool{
				"logging.trap.@operation": true,
			},
			shouldNotExist: []string{
				"logging.trap.severity.@operation", // Should be removed when parent gets operation="remove"
			},
		},
		{
			name: "multiple children skipped when parent removed",
			removePaths: []string{
				"/logging/history",
				"/logging/history/size",
				"/logging/history/severity-level",
			},
			shouldExist: map[string]bool{
				"logging.history.@operation": true,
			},
			shouldNotExist: []string{
				"logging.history.size.@operation",
				"logging.history.severity-level.@operation",
			},
		},
		{
			name: "nested children at different levels",
			removePaths: []string{
				"/logging/host",
				"/logging/host/ipv6",
				"/logging/host/ipv6/ipv6-host-list",
			},
			shouldExist: map[string]bool{
				"logging.host.@operation": true,
			},
			shouldNotExist: []string{
				"logging.host.ipv6.@operation",
				"logging.host.ipv6.ipv6-host-list.@operation",
			},
		},
		{
			name: "sibling removals both added",
			removePaths: []string{
				"/logging/buffered/size-value",
				"/logging/buffered/severity-level",
			},
			shouldExist: map[string]bool{
				"logging.buffered.size-value.@operation":     true,
				"logging.buffered.severity-level.@operation": true,
			},
		},
		{
			name: "complex real-world scenario",
			removePaths: []string{
				"/logging/trap",
				"/logging/trap/severity",
				"/logging/buffered/size-value",
				"/logging/buffered/severity-level",
				"/logging/host/ipv4-host-list[ipv4-host='1.1.1.1']",
				"/logging/host/ipv4-host-list[ipv4-host='1.1.1.1']/ipv4-host",
			},
			shouldExist: map[string]bool{
				"logging.trap.@operation":                    true,
				"logging.buffered.size-value.@operation":     true,
				"logging.buffered.severity-level.@operation": true,
				"logging.host.ipv4-host-list.@operation":     true,
			},
			shouldNotExist: []string{
				"logging.trap.severity.@operation",
				"logging.host.ipv4-host-list.ipv4-host.@operation",
			},
		},
		{
			name: "child with keys added first, parent added later - keys preserved",
			removePaths: []string{
				"/logging/host/ipv4-host-vrf-list[ipv4-host='1.1.1.1'][vrf='VRF1']/transport",
				"/logging/host/ipv4-host-vrf-list[ipv4-host='1.1.1.1'][vrf='VRF1']",
			},
			shouldExist: map[string]bool{
				"logging.host.ipv4-host-vrf-list.@operation": true,
			},
			shouldNotExist: []string{
				"logging.host.ipv4-host-vrf-list.transport.@operation",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := netconf.Body{}

			// Apply all removal operations
			for _, path := range tt.removePaths {
				body = RemoveFromXPath(body, path)
			}

			resultXML := body.Res()
			t.Logf("Generated XML:\n%s", resultXML)

			// Check paths that should exist
			for path, shouldHaveRemove := range tt.shouldExist {
				result := xmldot.Get(resultXML, path)
				exists := result.Exists()
				value := result.String()

				if shouldHaveRemove && (!exists || value != "remove") {
					t.Errorf("Expected operation='remove' at %q, exists=%v, value=%q", path, exists, value)
				} else if exists {
					t.Logf("✓ Found operation='remove' at %q", path)
				}
			}

			// Check paths that should NOT exist (redundant child removals)
			for _, path := range tt.shouldNotExist {
				result := xmldot.Get(resultXML, path)
				if result.Exists() {
					t.Errorf("Path %q should NOT exist (redundant child removal), but found value: %q", path, result.String())
				} else {
					t.Logf("✓ Correctly skipped redundant removal at %q", path)
				}
			}

			// Verify key elements are preserved (not removed)
			// Check for ipv4-host keys if present in paths
			if strings.Contains(strings.Join(tt.removePaths, " "), "ipv4-host=") {
				// Look for ipv4-host elements that should still exist as keys
				if strings.Contains(resultXML, "<ipv4-host>1.1.1.1</ipv4-host>") {
					t.Logf("✓ Key element <ipv4-host> preserved correctly")
				}
			}
			// Check for vrf keys if present
			if strings.Contains(strings.Join(tt.removePaths, " "), "vrf=") {
				if strings.Contains(resultXML, "<vrf>VRF1</vrf>") {
					t.Logf("✓ Key element <vrf> preserved correctly")
				}
			}
		})
	}
}

// TestCleanupRedundantRemoveOperations tests the cleanup function with various XML scenarios
func TestCleanupRedundantRemoveOperations(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "simple parent-child with redundant child operation",
			input: `<logging xmlns="http://cisco.com/ns/yang/Cisco-IOS-XE-native">
  <trap operation="remove">
    <severity operation="remove"></severity>
  </trap>
</logging>`,
			expected: `<logging xmlns="http://cisco.com/ns/yang/Cisco-IOS-XE-native">
  <trap operation="remove"></trap>
</logging>`,
		},
		{
			name: "nested removals with multiple levels",
			input: `<native xmlns="http://cisco.com/ns/yang/Cisco-IOS-XE-native">
  <logging operation="remove">
    <trap operation="remove">
      <severity operation="remove"></severity>
    </trap>
    <buffered operation="remove">
      <size-value operation="remove"></size-value>
      <severity-level operation="remove"></severity-level>
    </buffered>
  </logging>
</native>`,
			expected: `<native xmlns="http://cisco.com/ns/yang/Cisco-IOS-XE-native">
  <logging operation="remove"></logging>
</native>`,
		},
		{
			name: "preserve keys in removed list items",
			input: `<host operation="remove">
  <ipv4-host-vrf-list operation="remove">
    <ipv4-host>1.1.1.1</ipv4-host>
    <vrf>VRF1</vrf>
    <transport operation="remove"></transport>
  </ipv4-host-vrf-list>
</host>`,
			expected: `<host operation="remove">
  <ipv4-host-vrf-list>
    <ipv4-host>1.1.1.1</ipv4-host>
    <vrf>VRF1</vrf>
  </ipv4-host-vrf-list>
</host>`,
		},
		{
			name: "multiple siblings with removals",
			input: `<logging xmlns="http://cisco.com/ns/yang/Cisco-IOS-XE-native">
  <trap operation="remove">
    <severity operation="remove"></severity>
  </trap>
  <buffered operation="remove">
    <size-value operation="remove"></size-value>
    <severity-level operation="remove"></severity-level>
  </buffered>
  <console operation="remove"></console>
</logging>`,
			expected: `<logging xmlns="http://cisco.com/ns/yang/Cisco-IOS-XE-native">
  <trap operation="remove"></trap>
  <buffered operation="remove"></buffered>
  <console operation="remove"></console>
</logging>`,
		},
		{
			name: "no redundant operations",
			input: `<logging xmlns="http://cisco.com/ns/yang/Cisco-IOS-XE-native">
  <trap operation="remove"></trap>
  <console operation="remove"></console>
</logging>`,
			expected: `<logging xmlns="http://cisco.com/ns/yang/Cisco-IOS-XE-native">
  <trap operation="remove"></trap>
  <console operation="remove"></console>
</logging>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := netconf.NewBody(tt.input)
			result := CleanupRedundantRemoveOperations(body)

			// Normalize whitespace for comparison
			normalizeXML := func(s string) string {
				// Remove leading/trailing whitespace
				s = strings.TrimSpace(s)
				// Collapse multiple spaces/newlines into single space
				re := regexp.MustCompile(`\s+`)
				s = re.ReplaceAllString(s, " ")
				// Remove spaces around < and >
				s = strings.ReplaceAll(s, " <", "<")
				s = strings.ReplaceAll(s, "> ", ">")
				return s
			}

			normalizedResult := normalizeXML(result.Res())
			normalizedExpected := normalizeXML(tt.expected)

			if normalizedResult != normalizedExpected {
				t.Errorf("CleanupRedundantRemoveOperations() mismatch\nGot:\n%s\n\nExpected:\n%s", result.Res(), tt.expected)
			} else {
				t.Logf("✓ Cleanup successful")
			}
		})
	}
}

// TestSetFromXPath_ThenAppendFromXPath_SiblingElements tests the spanning tree MST instance scenario
// where we need to set id as a value and append multiple vlan-ids as siblings
func TestSetFromXPath_ThenAppendFromXPath_SiblingElements(t *testing.T) {
	cBody := netconf.Body{}

	// Simulate what the spanning tree MST instance code does
	// First, set the id value
	cBody = SetFromXPath(cBody, "id", "1")
	t.Log("After SetFromXPath id:")
	t.Log(cBody.Res())

	// Then, append vlan-ids values
	cBody = AppendFromXPath(cBody, "vlan-ids", 10)
	t.Log("\nAfter first AppendFromXPath vlan-ids:")
	t.Log(cBody.Res())

	cBody = AppendFromXPath(cBody, "vlan-ids", 20)
	t.Log("\nAfter second AppendFromXPath vlan-ids:")
	t.Log(cBody.Res())

	resultXML := cBody.Res()

	// Check the structure
	idValue := xmldotGetValue(resultXML, "id")
	vlanIds1 := xmldotGetValue(resultXML, "vlan-ids.0")
	vlanIds2 := xmldotGetValue(resultXML, "vlan-ids.1")

	t.Logf("\nParsed values:")
	t.Logf("  id: %q", idValue)
	t.Logf("  vlan-ids.0: %q", vlanIds1)
	t.Logf("  vlan-ids.1: %q", vlanIds2)

	// Verify id is set correctly
	if idValue != "1" {
		t.Errorf("Expected id='1', got %q", idValue)
	}

	// Verify vlan-ids are siblings, not nested under id
	if vlanIds1 != "10" {
		t.Errorf("Expected vlan-ids.0='10', got %q", vlanIds1)
	}
	if vlanIds2 != "20" {
		t.Errorf("Expected vlan-ids.1='20', got %q", vlanIds2)
	}

	// Verify XML structure - id and vlan-ids should be siblings
	if !strings.Contains(resultXML, "<id>1</id>") {
		t.Error("Expected <id>1</id> element")
	}
	if !strings.Contains(resultXML, "<vlan-ids>10</vlan-ids>") {
		t.Error("Expected <vlan-ids>10</vlan-ids> element")
	}
	if !strings.Contains(resultXML, "<vlan-ids>20</vlan-ids>") {
		t.Error("Expected <vlan-ids>20</vlan-ids> element")
	}

	// Check for the malformed pattern from the bug report
	if strings.Contains(resultXML, "<vlan-ids>10</vlan-ids>1") {
		t.Error("Found malformed XML: vlan-ids followed by id text node (bug reproduced)")
	}

	t.Log("\n✓ All sibling elements verified successfully")
}

// TestCleanupRedundantRemoveOperations_RealWorld tests with actual logging resource XML
func TestCleanupRedundantRemoveOperations_RealWorld(t *testing.T) {
	// Build the same XML that addDeletePathsXML would create
	body := netconf.Body{}
	basePath := "/Cisco-IOS-XE-native:native/logging"

	// Simulate the real order: child first, then parent
	body = RemoveFromXPath(body, basePath+"/trap/severity")
	body = RemoveFromXPath(body, basePath+"/trap")
	body = RemoveFromXPath(body, basePath+"/buffered/size-value")
	body = RemoveFromXPath(body, basePath+"/buffered/severity-level")
	body = RemoveFromXPath(body, basePath+"/buffered") // Add parent removal

	xmlBefore := body.Res()
	t.Logf("Before cleanup:\n%s\n", xmlBefore)

	// Apply cleanup
	body = CleanupRedundantRemoveOperations(body)
	xmlAfter := body.Res()
	t.Logf("After cleanup:\n%s\n", xmlAfter)

	// Verify trap/severity is removed
	if strings.Contains(xmlAfter, `<severity operation="remove">`) {
		t.Error("❌ Still contains <severity operation=\"remove\">")
	} else {
		t.Log("✓ severity operation removed under trap")
	}

	// Verify trap still has operation="remove"
	if !strings.Contains(xmlAfter, `<trap operation="remove">`) {
		t.Error("❌ trap should still have operation=\"remove\"")
	} else {
		t.Log("✓ trap still has operation=\"remove\"")
	}

	// Verify buffered children are removed
	if strings.Contains(xmlAfter, `<size-value operation="remove">`) {
		t.Error("❌ Still contains <size-value operation=\"remove\">")
	} else {
		t.Log("✓ size-value removed under buffered")
	}

	if strings.Contains(xmlAfter, `<severity-level operation="remove">`) {
		t.Error("❌ Still contains <severity-level operation=\"remove\">")
	} else {
		t.Log("✓ severity-level removed under buffered")
	}
}

// TestNamespaceExceptions tests that namespace exceptions are correctly applied
func TestNamespaceExceptions(t *testing.T) {
	tests := []struct {
		name              string
		xPath             string
		expectedNamespace string
		expectedPrefix    string
	}{
		{
			name:              "Standard namespace pattern",
			xPath:             "/Cisco-IOS-XE-native:native/interface",
			expectedNamespace: "http://cisco.com/ns/yang/Cisco-IOS-XE-native",
			expectedPrefix:    "Cisco-IOS-XE-native",
		},
		{
			name:              "Template namespace exception",
			xPath:             "/Cisco-IOS-XE-native:native/template/Cisco-IOS-XE-template:template_details",
			expectedNamespace: "http://cisco.com/ns/yang/ios-xe/template",
			expectedPrefix:    "Cisco-IOS-XE-template",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := netconf.Body{}
			body = SetFromXPath(body, tt.xPath, "")

			resultXML := body.Res()
			t.Logf("Generated XML:\n%s", resultXML)

			// Check if the correct namespace is present in the XML
			if !strings.Contains(resultXML, tt.expectedNamespace) {
				t.Errorf("Expected namespace %q not found in XML", tt.expectedNamespace)
			} else {
				t.Logf("✓ Correct namespace found: %s", tt.expectedNamespace)
			}

			// For template exception, verify both namespaces exist
			if tt.expectedPrefix == "Cisco-IOS-XE-template" {
				standardNamespace := "http://cisco.com/ns/yang/Cisco-IOS-XE-native"
				if !strings.Contains(resultXML, standardNamespace) {
					t.Errorf("Expected standard namespace %q not found in XML", standardNamespace)
				} else {
					t.Logf("✓ Standard namespace also present: %s", standardNamespace)
				}
			}
		})
	}
}

// TestIsGetConfigResponseEmpty tests the IsGetConfigResponseEmpty helper function
func TestIsGetConfigResponseEmpty(t *testing.T) {
	tests := []struct {
		name     string
		xml      string
		expected bool
	}{
		{
			name: "Empty data element",
			xml: `<rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="1">
				<data></data>
			</rpc-reply>`,
			expected: true,
		},
		{
			name: "Data element with whitespace only",
			xml: `<rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="1">
				<data>

				</data>
			</rpc-reply>`,
			expected: true,
		},
		{
			name: "Data element with configuration",
			xml: `<rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="1">
				<data>
					<native xmlns="http://cisco.com/ns/yang/Cisco-IOS-XE-native">
						<hostname>test-device</hostname>
					</native>
				</data>
			</rpc-reply>`,
			expected: false,
		},
		{
			name: "Data element with nested configuration",
			xml: `<rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="1">
				<data>
					<native xmlns="http://cisco.com/ns/yang/Cisco-IOS-XE-native">
						<interface>
							<GigabitEthernet>
								<name>1</name>
								<description>Test Interface</description>
							</GigabitEthernet>
						</interface>
					</native>
				</data>
			</rpc-reply>`,
			expected: false,
		},
		{
			name: "Data element with single attribute",
			xml: `<rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="1">
				<data>
					<native xmlns="http://cisco.com/ns/yang/Cisco-IOS-XE-native">
						<banner>
							<login>
								<banner>Welcome</banner>
							</login>
						</banner>
					</native>
				</data>
			</rpc-reply>`,
			expected: false,
		},
		{
			name: "Empty self-closing data element",
			xml: `<rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="1">
				<data/>
			</rpc-reply>`,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Parse the XML into a netconf.Res structure
			result := xmldot.Get(tt.xml, "rpc-reply")
			res := &netconf.Res{
				Res:       result,
				OK:        true,
				MessageID: "1",
			}

			isEmpty := IsGetConfigResponseEmpty(res)
			if isEmpty != tt.expected {
				t.Errorf("IsGetConfigResponseEmpty() = %v, want %v", isEmpty, tt.expected)
			} else {
				if isEmpty {
					t.Logf("✓ Correctly identified as empty response")
				} else {
					t.Logf("✓ Correctly identified as non-empty response")
				}
			}
		})
	}
}

// TestIsGetConfigResponseEmpty_NilInput tests nil input handling
func TestIsGetConfigResponseEmpty_NilInput(t *testing.T) {
	isEmpty := IsGetConfigResponseEmpty(nil)
	if !isEmpty {
		t.Errorf("IsGetConfigResponseEmpty(nil) = false, want true")
	} else {
		t.Logf("✓ Correctly handles nil input")
	}
}

// TestIsListPath tests the IsListPath helper function
func TestIsListPath(t *testing.T) {
	tests := []struct {
		name     string
		xPath    string
		expected bool
	}{
		{
			name:     "List item with simple predicate",
			xPath:    "/Cisco-IOS-XE-native:native/interface/Vlan[name=10]",
			expected: true,
		},
		{
			name:     "List item with quoted predicate",
			xPath:    "/Cisco-IOS-XE-native:native/interface/GigabitEthernet[name='1/0/1']",
			expected: true,
		},
		{
			name:     "List item with multiple predicates",
			xPath:    "/native/vrf[name='VRF1'][af-name='ipv4']",
			expected: true,
		},
		{
			name:     "List item with formatted predicate",
			xPath:    "/native/router/bgp[id=65000]/neighbor[id='10.0.0.1']",
			expected: true,
		},
		{
			name:     "Container without predicate",
			xPath:    "/Cisco-IOS-XE-native:native/clock",
			expected: false,
		},
		{
			name:     "Singleton without predicate",
			xPath:    "/Cisco-IOS-XE-native:native/hostname",
			expected: false,
		},
		{
			name:     "Deep container path",
			xPath:    "/native/interface/GigabitEthernet/ip/address",
			expected: false,
		},
		{
			name:     "Predicate in middle but not at end (container under list item)",
			xPath:    "/native/router/bgp[id=65000]/neighbor",
			expected: false,
		},
		{
			name:     "Predicate in middle with more nesting",
			xPath:    "/native/router/bgp[id=65000]/neighbor[id='10.0.0.1']/remote-as",
			expected: false,
		},
		{
			name:     "Empty path",
			xPath:    "",
			expected: false,
		},
		{
			name:     "Path with only opening bracket (malformed)",
			xPath:    "/native/interface[name",
			expected: false,
		},
		{
			name:     "Path with only closing bracket (malformed)",
			xPath:    "/native/interface]name",
			expected: false,
		},
		{
			name:     "Complex list path with namespace",
			xPath:    "/Cisco-IOS-XE-native:native/interface/Cisco-IOS-XE-ethernet:Port-channel[name=10]",
			expected: true,
		},
		{
			name:     "List path with trailing whitespace",
			xPath:    "/native/interface/Vlan[name=10]  \n",
			expected: true,
		},
		{
			name:     "Container path with trailing whitespace",
			xPath:    "/native/clock  ",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsListPath(tt.xPath)
			if result != tt.expected {
				t.Errorf("IsListPath(%q) = %v, want %v", tt.xPath, result, tt.expected)
			} else {
				if result {
					t.Logf("✓ Correctly identified as list path")
				} else {
					t.Logf("✓ Correctly identified as container/singleton path")
				}
			}
		})
	}
}

// TestIsGetConfigResponseEmpty_WithIsListPath demonstrates combined usage
func TestIsGetConfigResponseEmpty_WithIsListPath(t *testing.T) {
	tests := []struct {
		name                  string
		xml                   string
		xPath                 string
		shouldRemoveFromState bool
	}{
		{
			name: "Empty response for list item - should remove",
			xml: `<rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="1">
				<data></data>
			</rpc-reply>`,
			xPath:                 "/native/interface/Vlan[name=10]",
			shouldRemoveFromState: true,
		},
		{
			name: "Empty response for container - should NOT remove",
			xml: `<rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="1">
				<data></data>
			</rpc-reply>`,
			xPath:                 "/native/clock",
			shouldRemoveFromState: false,
		},
		{
			name: "Non-empty response for list item - should NOT remove",
			xml: `<rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="1">
				<data>
					<native xmlns="http://cisco.com/ns/yang/Cisco-IOS-XE-native">
						<interface>
							<Vlan>
								<name>10</name>
								<description>Test VLAN</description>
							</Vlan>
						</interface>
					</native>
				</data>
			</rpc-reply>`,
			xPath:                 "/native/interface/Vlan[name=10]",
			shouldRemoveFromState: false,
		},
		{
			name: "Non-empty response for container - should NOT remove",
			xml: `<rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="1">
				<data>
					<native xmlns="http://cisco.com/ns/yang/Cisco-IOS-XE-native">
						<clock>
							<timezone>PST -8 0</timezone>
						</clock>
					</native>
				</data>
			</rpc-reply>`,
			xPath:                 "/native/clock",
			shouldRemoveFromState: false,
		},
		{
			name: "Empty response for container under list item - should NOT remove",
			xml: `<rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="1">
				<data></data>
			</rpc-reply>`,
			xPath:                 "/native/router/bgp[id=65000]/neighbor",
			shouldRemoveFromState: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Parse the XML into a netconf.Res structure
			result := xmldot.Get(tt.xml, "rpc-reply")
			res := &netconf.Res{
				Res:       result,
				OK:        true,
				MessageID: "1",
			}

			// This is the pattern used in the resource template
			isEmpty := IsGetConfigResponseEmpty(res)
			isListPath := IsListPath(tt.xPath)
			shouldRemove := isEmpty && isListPath

			if shouldRemove != tt.shouldRemoveFromState {
				t.Errorf("Combined check = %v, want %v (isEmpty=%v, isListPath=%v)",
					shouldRemove, tt.shouldRemoveFromState, isEmpty, isListPath)
			} else {
				if shouldRemove {
					t.Logf("✓ Correctly determined should remove from state (empty list item)")
				} else {
					t.Logf("✓ Correctly determined should NOT remove from state")
				}
			}
		})
	}
}

// TestSetRawFromXPath_MultipleSiblingsWithNamespaces tests that multiple sibling elements
// each get their own namespace declaration when using namespace prefixes.
// This reproduces the CDP tlv-list issue where the second tlv-list was missing xmlns.
func TestSetRawFromXPath_MultipleSiblingsWithNamespaces(t *testing.T) {
	body := netconf.Body{}

	// Simulate CDP configuration with multiple tlv-list elements
	tlvList1 := `<name>TLIST</name><vtp-mgmt-domain operation="remove"></vtp-mgmt-domain><cos></cos>`
	tlvList2 := `<name>TLIST2</name><vtp-mgmt-domain></vtp-mgmt-domain><cos operation="remove"></cos>`

	// Add first tlv-list with namespace prefix
	body = SetRawFromXPath(body, "/Cisco-IOS-XE-native:native/cdp/Cisco-IOS-XE-cdp:tlv-list", tlvList1)
	// Add second tlv-list with namespace prefix
	body = SetRawFromXPath(body, "/Cisco-IOS-XE-native:native/cdp/Cisco-IOS-XE-cdp:tlv-list", tlvList2)

	if err := body.Err(); err != nil {
		t.Fatalf("SetRawFromXPath() returned error: %v", err)
	}

	resultXML := body.Res()
	t.Logf("Generated XML:\n%s", resultXML)

	// Verify both tlv-list elements exist
	name1 := xmldot.Get(resultXML, "native.cdp.tlv-list.0.name").String()
	name2 := xmldot.Get(resultXML, "native.cdp.tlv-list.1.name").String()

	if name1 != "TLIST" {
		t.Errorf("First tlv-list name = %q, want %q", name1, "TLIST")
	}
	if name2 != "TLIST2" {
		t.Errorf("Second tlv-list name = %q, want %q", name2, "TLIST2")
	}

	// Verify namespace declaration on native element
	nativeXmlns := xmldot.Get(resultXML, "native.@xmlns").String()
	if nativeXmlns != "http://cisco.com/ns/yang/Cisco-IOS-XE-native" {
		t.Errorf("native xmlns = %q, want %q", nativeXmlns, "http://cisco.com/ns/yang/Cisco-IOS-XE-native")
	}

	// Verify namespace declaration on BOTH tlv-list elements
	tlvList1Xmlns := xmldot.Get(resultXML, "native.cdp.tlv-list.0.@xmlns").String()
	tlvList2Xmlns := xmldot.Get(resultXML, "native.cdp.tlv-list.1.@xmlns").String()

	expectedCdpNs := "http://cisco.com/ns/yang/Cisco-IOS-XE-cdp"

	if tlvList1Xmlns != expectedCdpNs {
		t.Errorf("First tlv-list xmlns = %q, want %q", tlvList1Xmlns, expectedCdpNs)
	}
	if tlvList2Xmlns != expectedCdpNs {
		t.Errorf("Second tlv-list xmlns = %q, want %q", tlvList2Xmlns, expectedCdpNs)
	}

	// Verify the XML contains properly formatted elements with namespaces
	if !strings.Contains(resultXML, `<tlv-list xmlns="http://cisco.com/ns/yang/Cisco-IOS-XE-cdp">`) {
		t.Error("Generated XML missing tlv-list elements with namespace declarations")
	}

	// Count how many times the tlv-list xmlns appears - should be 2 (one for each element)
	xmlnsCount := strings.Count(resultXML, `<tlv-list xmlns="http://cisco.com/ns/yang/Cisco-IOS-XE-cdp">`)
	if xmlnsCount != 2 {
		t.Errorf("Found %d tlv-list elements with xmlns, want 2", xmlnsCount)
	}

	t.Log("✓ Multiple sibling elements with namespaces verified successfully")
}

// TestSetRawFromXPath_SingleSegmentMultipleSiblings tests that multiple sibling elements
// are created correctly with single-segment paths (not nested).
// This reproduces the ARP inspection filter vlan issue where multiple vlans within a filter
// were being nested instead of created as siblings.
func TestSetRawFromXPath_SingleSegmentMultipleSiblings(t *testing.T) {
	body := netconf.Body{}

	// Simulate ARP inspection filter with multiple vlan elements
	vlan1 := `<vlan-range>10</vlan-range><static></static>`
	vlan2 := `<vlan-range>100</vlan-range><static operation="remove"></static>`

	// Add first vlan (single-segment path)
	body = SetRawFromXPath(body, "vlan", vlan1)
	// Add second vlan (single-segment path)
	body = SetRawFromXPath(body, "vlan", vlan2)

	if err := body.Err(); err != nil {
		t.Fatalf("SetRawFromXPath() returned error: %v", err)
	}

	resultXML := body.Res()
	t.Logf("Generated XML:\n%s", resultXML)

	// Verify both vlan elements exist as SIBLINGS, not nested
	vlanRange1 := xmldot.Get(resultXML, "vlan.0.vlan-range").String()
	vlanRange2 := xmldot.Get(resultXML, "vlan.1.vlan-range").String()

	if vlanRange1 != "10" {
		t.Errorf("First vlan vlan-range = %q, want %q", vlanRange1, "10")
	}
	if vlanRange2 != "100" {
		t.Errorf("Second vlan vlan-range = %q, want %q", vlanRange2, "100")
	}

	// Verify static attributes
	static1Exists := xmldot.Get(resultXML, "vlan.0.static").Exists()
	static2Op := xmldot.Get(resultXML, "vlan.1.static.@operation").String()

	if !static1Exists {
		t.Error("First vlan static element missing")
	}
	if static2Op != "remove" {
		t.Errorf("Second vlan static operation = %q, want %q", static2Op, "remove")
	}

	// Verify NO nesting - the second vlan should NOT be inside the first vlan
	// Check that vlan.0.vlan does NOT exist (which would indicate nesting)
	nestedVlan := xmldot.Get(resultXML, "vlan.0.vlan")
	if nestedVlan.Exists() {
		t.Error("Found nested vlan inside first vlan - vlans should be siblings, not nested")
	}

	// Count vlan elements - should be exactly 2 at the root level
	vlanCount := xmldot.Get(resultXML, "vlan.#").Int()
	if vlanCount != 2 {
		t.Errorf("Found %d vlan elements, want 2", vlanCount)
	}

	t.Log("✓ Multiple sibling elements with single-segment paths verified successfully")
}

// TestSetRawFromXPath_PreserveOtherRootElements tests that when appending siblings with
// single-segment paths, other root-level elements are preserved.
// This reproduces the ARP filter issue where arpacl was lost when adding multiple vlans.
func TestSetRawFromXPath_PreserveOtherRootElements(t *testing.T) {
	body := netconf.Body{}

	// Simulate ARP filter: first add arpacl, then add multiple vlans
	body = SetFromXPath(body, "arpacl", "FILTER2")

	vlan1 := `<vlan-range>10</vlan-range><static></static>`
	vlan2 := `<vlan-range>100</vlan-range><static operation="remove"></static>`

	// Add first vlan
	body = SetRawFromXPath(body, "vlan", vlan1)
	// Add second vlan - this should preserve the arpacl
	body = SetRawFromXPath(body, "vlan", vlan2)

	if err := body.Err(); err != nil {
		t.Fatalf("SetRawFromXPath() returned error: %v", err)
	}

	resultXML := body.Res()
	t.Logf("Generated XML:\n%s", resultXML)

	// Verify arpacl is still present
	arpacl := xmldot.Get(resultXML, "arpacl").String()
	if arpacl != "FILTER2" {
		t.Errorf("arpacl = %q, want %q (element was lost!)", arpacl, "FILTER2")
	}

	// Verify both vlan elements exist
	vlanCount := xmldot.Get(resultXML, "vlan.#").Int()
	if vlanCount != 2 {
		t.Errorf("Found %d vlan elements, want 2", vlanCount)
	}

	vlanRange1 := xmldot.Get(resultXML, "vlan.0.vlan-range").String()
	vlanRange2 := xmldot.Get(resultXML, "vlan.1.vlan-range").String()

	if vlanRange1 != "10" {
		t.Errorf("First vlan vlan-range = %q, want %q", vlanRange1, "10")
	}
	if vlanRange2 != "100" {
		t.Errorf("Second vlan vlan-range = %q, want %q", vlanRange2, "100")
	}

	// Verify the XML structure: arpacl should come before vlans
	if !strings.Contains(resultXML, "<arpacl>FILTER2</arpacl>") {
		t.Error("arpacl element not found in XML")
	}

	// Check that arpacl appears before the first vlan in the string
	arpclIdx := strings.Index(resultXML, "<arpacl>")
	vlanIdx := strings.Index(resultXML, "<vlan>")
	if arpclIdx == -1 || vlanIdx == -1 || arpclIdx > vlanIdx {
		t.Error("arpacl should appear before vlan elements")
	}

	t.Log("✓ Other root elements preserved successfully when appending siblings")
}

// TestRemoveFromXPath_MultipleSiblings tests that RemoveFromXPath correctly handles
// multiple sibling list elements with different keys (e.g., multiple VLANs with different IDs)
func TestRemoveFromXPath_MultipleSiblings(t *testing.T) {
	// Start with an empty body
	body := netconf.Body{}

	// Simulate deleting priority from multiple VLANs (like spanning-tree vlan removal fix)
	// Use simple paths like other tests - no namespace prefixes needed for testing

	// First VLAN: delete priority for vlan[id='20']
	body = RemoveFromXPath(body, "/native/vlan[id='20']/priority")

	// Second VLAN: delete priority for vlan[id='30']
	body = RemoveFromXPath(body, "/native/vlan[id='30']/priority")

	resultXML := body.Res()
	t.Logf("Result XML:\n%s", resultXML)

	// Verify both vlan elements exist in the XML
	vlanCount := xmldot.Get(resultXML, "native.vlan.#").Int()
	if vlanCount != 2 {
		t.Errorf("Found %d vlan elements, want 2", vlanCount)
	}

	// Check that we have both IDs
	vlan0Id := xmldot.Get(resultXML, "native.vlan.0.id").String()
	vlan1Id := xmldot.Get(resultXML, "native.vlan.1.id").String()

	t.Logf("vlan.0.id = %q, vlan.1.id = %q", vlan0Id, vlan1Id)

	// IDs should be 20 and 30 (order may vary)
	ids := map[string]bool{vlan0Id: true, vlan1Id: true}
	if !ids["20"] {
		t.Errorf("VLAN with id='20' not found in XML")
	}
	if !ids["30"] {
		t.Errorf("VLAN with id='30' not found in XML")
	}

	// Verify both priority elements have operation="remove"
	prio0Op := xmldot.Get(resultXML, "native.vlan.0.priority.@operation").String()
	prio1Op := xmldot.Get(resultXML, "native.vlan.1.priority.@operation").String()

	if prio0Op != "remove" {
		t.Errorf("First vlan priority operation = %q, want %q", prio0Op, "remove")
	}
	if prio1Op != "remove" {
		t.Errorf("Second vlan priority operation = %q, want %q", prio1Op, "remove")
	}

	t.Log("✓ Multiple sibling elements with different keys handled correctly")
}

// TestFindSiblingInfo tests the findSiblingInfo helper function
func TestFindSiblingInfo(t *testing.T) {
	tests := []struct {
		name        string
		setupBody   func() netconf.Body
		parentPath  string
		elementName string
		keys        []KeyValue
		wantResult  SiblingResult
	}{
		{
			name: "no existing elements returns SiblingActionNew",
			setupBody: func() netconf.Body {
				return netconf.NewBody("")
			},
			parentPath:  "spanning-tree",
			elementName: "vlan",
			keys:        []KeyValue{{Key: "id", Value: "10"}},
			wantResult:  SiblingResult{Action: SiblingActionNew, Index: -1},
		},
		{
			name: "finds existing element with matching keys",
			setupBody: func() netconf.Body {
				body := netconf.NewBody("")
				body = body.Set("spanning-tree.vlan.id", "10")
				body = body.Set("spanning-tree.vlan.priority", "4096")
				return body
			},
			parentPath:  "spanning-tree",
			elementName: "vlan",
			keys:        []KeyValue{{Key: "id", Value: "10"}},
			wantResult:  SiblingResult{Action: SiblingActionUpdate, Index: 0},
		},
		{
			name: "returns SiblingActionAppend when no match",
			setupBody: func() netconf.Body {
				body := netconf.NewBody("")
				body = body.Set("spanning-tree.vlan.id", "10")
				return body
			},
			parentPath:  "spanning-tree",
			elementName: "vlan",
			keys:        []KeyValue{{Key: "id", Value: "20"}},
			wantResult:  SiblingResult{Action: SiblingActionAppend, Index: -1},
		},
		{
			name: "finds correct element among multiple siblings",
			setupBody: func() netconf.Body {
				// Use SetRaw to properly create multiple sibling elements
				// (indexed Set paths don't work with xmldot)
				body := netconf.NewBody("")
				body = body.SetRaw("spanning-tree", "<vlan><id>10</id></vlan><vlan><id>20</id></vlan><vlan><id>30</id></vlan>")
				return body
			},
			parentPath:  "spanning-tree",
			elementName: "vlan",
			keys:        []KeyValue{{Key: "id", Value: "20"}},
			wantResult:  SiblingResult{Action: SiblingActionUpdate, Index: 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := tt.setupBody()
			got := findSiblingInfo(body, tt.parentPath, tt.elementName, tt.keys)
			if got.Action != tt.wantResult.Action || got.Index != tt.wantResult.Index {
				t.Errorf("findSiblingInfo() = %+v, want %+v", got, tt.wantResult)
			}
		})
	}
}

// TestTrimNetconfTrailingWhitespace tests the TrimNetconfTrailingWhitespace helper function
// which trims trailing whitespace from multi-line strings from NETCONF responses to prevent Terraform drift
func TestTrimNetconfTrailingWhitespace(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Single line no whitespace",
			input:    "Hello World",
			expected: "Hello World",
		},
		{
			name:     "Single line with trailing spaces",
			input:    "Hello World   ",
			expected: "Hello World",
		},
		{
			name:     "Single line with trailing tabs",
			input:    "Hello World\t\t",
			expected: "Hello World",
		},
		{
			name:     "Multi-line with trailing whitespace on each line",
			input:    "Line 1   \nLine 2\t\t\nLine 3  ",
			expected: "Line 1\nLine 2\nLine 3",
		},
		{
			name:     "Multi-line preserves leading whitespace",
			input:    "  Line 1   \n\tLine 2\t\t\n  Line 3  ",
			expected: "  Line 1\n\tLine 2\n  Line 3",
		},
		{
			name:     "Multi-line with internal empty lines preserved",
			input:    "Line 1   \n   \nLine 3  ",
			expected: "Line 1\n\nLine 3",
		},
		{
			name:     "Preserves internal whitespace",
			input:    "Hello   World   ",
			expected: "Hello   World",
		},
		{
			name:     "Handles carriage return",
			input:    "Line 1  \r\nLine 2\r",
			expected: "Line 1\nLine 2",
		},
		{
			name:     "Removes trailing newline (YAML block scalar)",
			input:    "Line 1\nLine 2\nLine 3\n",
			expected: "Line 1\nLine 2\nLine 3",
		},
		{
			name:     "Removes multiple trailing newlines",
			input:    "Line 1\nLine 2\n\n\n",
			expected: "Line 1\nLine 2",
		},
		{
			name:     "Handles trailing whitespace and trailing newlines together",
			input:    "Line 1  \nLine 2  \n\n",
			expected: "Line 1\nLine 2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TrimNetconfTrailingWhitespace(tt.input)
			if result != tt.expected {
				t.Errorf("TrimNetconfTrailingWhitespace() = %q, want %q", result, tt.expected)
			}
		})
	}
}

// TestRemoveFromXPath_DisabledVlansScenario tests the exact scenario used by disabled_vlans:
// 1. First add regular vlans with priorities (via SetFromXPath)
// 2. Then add disabled_vlans with operation="remove" on the vlan element itself (via RemoveFromXPath)
// This tests that RemoveFromXPath properly appends sibling elements with operation="remove"
// when there are existing siblings.
func TestRemoveFromXPath_DisabledVlansScenario(t *testing.T) {
	body := netconf.Body{}
	basePath := "/Cisco-IOS-XE-native:native/spanning-tree"

	// Step 1: Add regular vlans with priority (mimics toBodyXML for vlans)
	// VLAN 10 with priority 4096
	cBody := netconf.Body{}
	cBody = SetFromXPath(cBody, "id", "10")
	cBody = SetFromXPath(cBody, "priority", "4096")
	body = SetRawFromXPath(body, basePath+"/Cisco-IOS-XE-spanning-tree:vlan", cBody.Res())

	t.Logf("After adding vlan 10:\n%s\n", body.Res())

	// VLAN 20 with priority 8192
	cBody2 := netconf.Body{}
	cBody2 = SetFromXPath(cBody2, "id", "20")
	cBody2 = SetFromXPath(cBody2, "priority", "8192")
	body = SetRawFromXPath(body, basePath+"/Cisco-IOS-XE-spanning-tree:vlan", cBody2.Res())

	t.Logf("After adding vlan 20:\n%s\n", body.Res())

	// Step 2: Add disabled_vlans with operation="remove" on the vlan element
	// Disabled VLAN 40 - should have operation="remove" at the vlan level
	predicates40 := "[id='40']"
	body = RemoveFromXPath(body, basePath+"/Cisco-IOS-XE-spanning-tree:vlan"+predicates40)

	t.Logf("After disabled_vlan 40:\n%s\n", body.Res())

	// Disabled VLAN 50 - should have operation="remove" at the vlan level
	predicates50 := "[id='50']"
	body = RemoveFromXPath(body, basePath+"/Cisco-IOS-XE-spanning-tree:vlan"+predicates50)

	resultXML := body.Res()
	t.Logf("Final XML:\n%s\n", resultXML)

	// Verification
	vlanCount := xmldot.Get(resultXML, "native.spanning-tree.vlan.#").Int()
	t.Logf("Total vlan elements: %d", vlanCount)

	if vlanCount != 4 {
		t.Errorf("Expected 4 vlan elements, got %d", vlanCount)
	}

	// Check each vlan element
	for i := 0; i < int(vlanCount); i++ {
		id := xmldot.Get(resultXML, fmt.Sprintf("native.spanning-tree.vlan.%d.id", i)).String()
		op := xmldot.Get(resultXML, fmt.Sprintf("native.spanning-tree.vlan.%d.@operation", i)).String()
		prio := xmldot.Get(resultXML, fmt.Sprintf("native.spanning-tree.vlan.%d.priority", i)).String()
		t.Logf("  vlan.%d: id=%s, operation=%q, priority=%s", i, id, op, prio)

		// VLANs 10 and 20 should NOT have operation attribute
		if id == "10" || id == "20" {
			if op != "" {
				t.Errorf("Regular vlan %s should not have operation attribute, got %q", id, op)
			}
		}
		// VLANs 40 and 50 should have operation="remove"
		if id == "40" || id == "50" {
			if op != "remove" {
				t.Errorf("Disabled vlan %s should have operation='remove', got %q", id, op)
			}
		}
	}

	// Additional string-based verification (note: the xmlns attribute may appear after operation)
	if !strings.Contains(resultXML, `operation="remove"`) {
		t.Error("Expected at least one element with operation=\"remove\" attribute in output")
	}
}

// TestAugmentNamespaces_IPv6ColonHandling tests that colons in IPv6 addresses within predicates
// are not misinterpreted as YANG namespace separators.
//
// This is a defensive test for the sibling handling code path in buildXPathStructure().
// That code calls augmentNamespaces with the original XPath (including predicates),
// which could contain IPv6 addresses with colons. Without the fix to augmentNamespaces,
// these colons could be misinterpreted as namespace separators.
func TestAugmentNamespaces_IPv6ColonHandling(t *testing.T) {
	tests := []struct {
		name             string
		xPath            string
		setValue         string
		shouldNotContain []string // Strings that should NOT appear in output
		shouldContain    []string // Strings that SHOULD appear in output
	}{
		{
			// Element WITHOUT namespace prefix, but WITH IPv6 in predicate
			name:     "element without namespace prefix but with IPv6 in predicate",
			xPath:    "/Cisco-IOS-XE-native:native/ipv6/route/ipv6-route[prefix='2001:db8::1/64']/next-hop",
			setValue: "10.0.0.1",
			shouldNotContain: []string{
				"http://cisco.com/ns/yang/ipv6-route[prefix", // Malformed namespace URL
			},
			shouldContain: []string{
				"2001:db8::1/64", // IPv6 address preserved intact
			},
		},
		{
			name:     "element without namespace prefix and link-local IPv6",
			xPath:    "/Cisco-IOS-XE-native:native/ipv6/neighbor/neighbor-entry[address='fe80::1']/interface",
			setValue: "GigabitEthernet1",
			shouldNotContain: []string{
				"http://cisco.com/ns/yang/neighbor-entry[address",
			},
			shouldContain: []string{
				"fe80::1",
			},
		},
		{
			name:     "RemoveFromXPath with element without namespace and IPv6 predicate",
			xPath:    "/Cisco-IOS-XE-native:native/ipv6/route/ipv6-route[prefix='2001:db8:abcd::1/128']",
			setValue: "__REMOVE__",
			shouldNotContain: []string{
				"http://cisco.com/ns/yang/ipv6-route[prefix",
			},
			shouldContain: []string{
				"2001:db8:abcd::1/128",
				`operation="remove"`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := netconf.Body{}
			var result netconf.Body

			if tt.setValue == "__REMOVE__" {
				result = RemoveFromXPath(body, tt.xPath)
			} else {
				result = SetFromXPath(body, tt.xPath, tt.setValue)
			}

			resultXML := result.Res()
			t.Logf("XPath: %s", tt.xPath)
			t.Logf("Generated XML:\n%s", resultXML)

			// Check that invalid namespace patterns are NOT present
			for _, badPattern := range tt.shouldNotContain {
				if strings.Contains(resultXML, badPattern) {
					t.Errorf("Output should NOT contain %q (indicates IPv6 colon was misinterpreted as namespace separator)\nGot:\n%s", badPattern, resultXML)
				}
			}

			// Check that expected patterns ARE present
			for _, goodPattern := range tt.shouldContain {
				if !strings.Contains(resultXML, goodPattern) {
					t.Errorf("Output should contain %q\nGot:\n%s", goodPattern, resultXML)
				}
			}

			t.Log("✓ IPv6 address colons correctly handled (not misinterpreted as namespace separators)")
		})
	}
}

// TestTrimNetconfTrailingWhitespace_BannerScenarios tests real-world banner scenarios
// that can cause Terraform drift when comparing NETCONF response to state
func TestTrimNetconfTrailingWhitespace_BannerScenarios(t *testing.T) {
	tests := []struct {
		name         string
		netconfValue string // Value as returned by NETCONF (with trailing whitespace)
		stateValue   string // Value as stored in Terraform state (often from YAML with trailing newline)
		shouldMatch  bool   // After normalization, should they match?
	}{
		{
			name: "Multi-line banner with trailing spaces from NETCONF",
			netconfValue: "***************************************************************************   \n" +
				"*                       EXEC SESSION STARTED                              *   \n" +
				"***************************************************************************   ",
			stateValue: "***************************************************************************\n" +
				"*                       EXEC SESSION STARTED                              *\n" +
				"***************************************************************************",
			shouldMatch: true,
		},
		{
			name: "Multi-line banner: YAML block scalar (trailing newline) vs NETCONF (no trailing newline)",
			netconfValue: "***************************************************************************\n" +
				"*                       EXEC SESSION STARTED                              *\n" +
				"***************************************************************************",
			stateValue: "***************************************************************************\n" +
				"*                       EXEC SESSION STARTED                              *\n" +
				"***************************************************************************\n", // YAML | adds trailing newline
			shouldMatch: true,
		},
		{
			name: "Multi-line banner: NETCONF trailing spaces + YAML trailing newline",
			netconfValue: "***************************************************************************   \n" +
				"*                       EXEC SESSION STARTED                              *   \n" +
				"***************************************************************************   ",
			stateValue: "***************************************************************************\n" +
				"*                       EXEC SESSION STARTED                              *\n" +
				"***************************************************************************\n", // YAML | adds trailing newline
			shouldMatch: true,
		},
		{
			name:         "Single-line banner with trailing space",
			netconfValue: "*** AUTHORIZED ACCESS ONLY ***   ",
			stateValue:   "*** AUTHORIZED ACCESS ONLY ***",
			shouldMatch:  true,
		},
		{
			name: "Banner with mixed trailing whitespace",
			netconfValue: "Line 1  \t\n" +
				"Line 2   \n" +
				"Line 3\t",
			stateValue: "Line 1\n" +
				"Line 2\n" +
				"Line 3",
			shouldMatch: true,
		},
		{
			name:         "Banner already normalized (no changes needed)",
			netconfValue: "Simple banner message",
			stateValue:   "Simple banner message",
			shouldMatch:  true,
		},
		{
			name:         "Different content should not match",
			netconfValue: "Banner A   ",
			stateValue:   "Banner B",
			shouldMatch:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			normalizedNetconf := TrimNetconfTrailingWhitespace(tt.netconfValue)
			normalizedState := TrimNetconfTrailingWhitespace(tt.stateValue)

			matches := normalizedNetconf == normalizedState
			if matches != tt.shouldMatch {
				if tt.shouldMatch {
					t.Errorf("Expected values to match after normalization:\nNETCONF: %q\nState:   %q\nNormalized NETCONF: %q\nNormalized State:   %q",
						tt.netconfValue, tt.stateValue, normalizedNetconf, normalizedState)
				} else {
					t.Errorf("Expected values to NOT match, but they did:\nNormalized: %q", normalizedNetconf)
				}
			}
		})
	}
}
