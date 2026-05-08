package helpers

import (
	"strings"
	"testing"

	"github.com/netascode/go-netconf"
)

func TestRemoveFromXPath_LeafListSiblingConflict(t *testing.T) {
	// Reproduces the bug where RemoveFromXPath with a leaf-list content predicate
	// ([.=value]) overwrites an existing element with the same name, destroying
	// the merge values set by AppendFromXPath (toBodyXML).
	basePath := "/Cisco-IOS-XE-native:native/ipv6/dhcp/Cisco-IOS-XE-dhcp:pool[name='DHCPv6-PD']"

	// Step 1: Simulate toBodyXML — append a leaf-list value
	body := netconf.Body{}
	body = SetFromXPath(body, basePath+"/name", "DHCPv6-PD")
	body = AppendFromXPath(body, basePath+"/dns-server", "2001:4860:4860::8888")

	t.Logf("After AppendFromXPath:\n%s", body.Res())

	if !strings.Contains(body.Res(), "2001:4860:4860::8888") {
		t.Fatal("Setup failed: 8888 not in body after AppendFromXPath")
	}

	// Step 2: Simulate addDeletedItemsXML — remove a different leaf-list value
	body = RemoveFromXPath(body, basePath+"/dns-server[.=2001:4860:4860::4444]")

	result := body.Res()
	t.Logf("After RemoveFromXPath:\n%s", result)

	// The merge value (8888) must still be present
	if !strings.Contains(result, "2001:4860:4860::8888") {
		t.Error("Merge value 8888 was destroyed by RemoveFromXPath")
	}

	// The remove element (4444 with operation="remove") must be present
	if !strings.Contains(result, "2001:4860:4860::4444") {
		t.Error("Remove value 4444 is not present in the XML")
	}
	if !strings.Contains(result, `operation="remove"`) {
		t.Error("operation=\"remove\" attribute is not present")
	}

	// Both should coexist as sibling <dns-server> elements
	count := strings.Count(result, "<dns-server")
	if count < 2 {
		t.Errorf("Expected at least 2 <dns-server> elements, got %d", count)
	}
}

func TestRemoveFromXPath_LeafListNewElement(t *testing.T) {
	// Test that RemoveFromXPath creates correct XML when no existing element
	// with the same name exists (the "new" case).
	body := netconf.Body{}

	body = RemoveFromXPath(body, "/Cisco-IOS-XE-native:native/ipv6/dhcp/Cisco-IOS-XE-dhcp:pool[name='TEST']/dns-server[.=2001:db8::1]")

	result := body.Res()
	t.Logf("Result:\n%s", result)

	// Should produce <dns-server operation="remove">2001:db8::1</dns-server>
	if !strings.Contains(result, ">2001:db8::1<") {
		t.Errorf("Expected text content '2001:db8::1', got: %s", result)
	}
	if !strings.Contains(result, `operation="remove"`) {
		t.Error("Missing operation=\"remove\" attribute")
	}
	// Must NOT contain <.> tags
	if strings.Contains(result, "<.>") || strings.Contains(result, "</.>") {
		t.Error("Invalid <.> element generated — leaf-list content should be text, not a child element")
	}
}

func TestRemoveFromXPath_MultipleLeafListDeletions(t *testing.T) {
	// Test deleting multiple leaf-list values creates proper siblings
	basePath := "/Cisco-IOS-XE-native:native/ipv6/dhcp/Cisco-IOS-XE-dhcp:pool[name='TEST']"

	body := netconf.Body{}
	body = SetFromXPath(body, basePath+"/name", "TEST")
	body = AppendFromXPath(body, basePath+"/dns-server", "2001:db8::1")
	body = RemoveFromXPath(body, basePath+"/dns-server[.=2001:db8::2]")
	body = RemoveFromXPath(body, basePath+"/dns-server[.=2001:db8::3]")

	result := body.Res()
	t.Logf("Result:\n%s", result)

	// Should have 3 dns-server elements: one merge + two removes
	count := strings.Count(result, "<dns-server")
	if count != 3 {
		t.Errorf("Expected 3 <dns-server> elements, got %d", count)
	}

	if !strings.Contains(result, "2001:db8::1") {
		t.Error("Merge value 2001:db8::1 missing")
	}
	if !strings.Contains(result, "2001:db8::2") {
		t.Error("Remove value 2001:db8::2 missing")
	}
	if !strings.Contains(result, "2001:db8::3") {
		t.Error("Remove value 2001:db8::3 missing")
	}
}
