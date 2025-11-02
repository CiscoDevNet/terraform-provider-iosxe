package helpers

import (
	"fmt"

	"reflect"
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

// TestConvertRestconfPathToXPath tests the ConvertRestconfPathToXPath function
func TestConvertRestconfPathToXPath(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{
			name:     "format string with %s=%v",
			path:     "Cisco-IOS-XE-native:native/interface/%s=%v",
			expected: "Cisco-IOS-XE-native:native/interface/%s[%v=%v]",
		},
		{
			name:     "format string with multiple placeholders",
			path:     "vrf/%s=%v/address-family/%s=%v",
			expected: "vrf/%s[%v=%v]/address-family/%s[%v=%v]",
		},
		{
			name:     "concrete path with single key",
			path:     "interface/GigabitEthernet=1",
			expected: "interface/GigabitEthernet[%v=1]",
		},
		{
			name:     "concrete path with composite key",
			path:     "vrf=VRF1,address-family=ipv4",
			expected: "vrf[%v=VRF1][%v=address-family=ipv4]",
		},
		{
			name:     "concrete path with multiple segments",
			path:     "native/vrf=VRF1/address-family=ipv4",
			expected: "native/vrf[%v=VRF1]/address-family[%v=ipv4]",
		},
		{
			name:     "path without keys",
			path:     "native/ip/source-route",
			expected: "native/ip/source-route",
		},
		{
			name:     "namespace prefix preserved",
			path:     "Cisco-IOS-XE-native:native/interface/GigabitEthernet=1/ip",
			expected: "Cisco-IOS-XE-native:native/interface/GigabitEthernet[%v=1]/ip",
		},
		{
			name:     "mixed format strings and concrete values",
			path:     "native/interface/%s=%v/vlan=100",
			expected: "native/interface/%s[%v=%v]/vlan[%v=100]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ConvertRestconfPathToXPath(tt.path)
			if result != tt.expected {
				t.Errorf("ConvertRestconfPathToXPath(%q) = %q, want %q",
					tt.path, result, tt.expected)
			}
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

// TestGetXPathFormat tests what format getXPath produces
func TestGetXPathFormat(t *testing.T) {
	// Test ConvertRestconfPathToXPath
	path := ConvertRestconfPathToXPath("Cisco-IOS-XE-native:native/policy/Cisco-IOS-XE-policy:class-map=%v")
	t.Logf("After ConvertRestconfPathToXPath: %s", path)

	// Then sprintf with key
	finalPath := fmt.Sprintf(path, "name", "CM1")
	t.Logf("After fmt.Sprintf: %s", finalPath)
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
