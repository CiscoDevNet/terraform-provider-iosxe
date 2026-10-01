package provider

import (
	"context"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestCryptoIKEv2ProposalPQCToBodyXML(t *testing.T) {
	data := CryptoIKEv2Proposal{
		Name:         types.StringValue("DMVPN-PQC"),
		PqcMlkem768:  types.BoolValue(true),
		PqcMlkem1024: types.BoolValue(true),
		PqcOptional:  types.BoolValue(true),
	}

	body := data.toBodyXML(context.Background(), data)
	for _, element := range []string{"<mlkem768", "<mlkem1024", "<optional"} {
		if !strings.Contains(body, element) {
			t.Fatalf("expected %s in NETCONF body: %s", element, body)
		}
	}
}

func TestCryptoIKEv2ProposalWithoutPQCOmitsPQCXML(t *testing.T) {
	data := CryptoIKEv2Proposal{
		Name: types.StringValue("CLASSICAL-ONLY"),
	}

	body := data.toBodyXML(context.Background(), data)
	if strings.Contains(body, "<pqc") || strings.Contains(body, "<mlkem") {
		t.Fatalf("did not expect PQC elements when no PQC attribute is configured: %s", body)
	}
}

func TestCryptoIPSecProfilePresencePfsToBodyXML(t *testing.T) {
	data := CryptoIPSecProfile{
		Name:   types.StringValue("DMVPN_IPSEC_TEST"),
		SetPfs: types.BoolValue(true),
	}

	body := data.toBodyXML(context.Background(), data)
	if !strings.Contains(body, "<pfs") {
		t.Fatalf("expected presence-only pfs container in NETCONF body: %s", body)
	}
	if strings.Contains(body, "<group") {
		t.Fatalf("did not expect an explicit DH group in NETCONF body: %s", body)
	}
}

func TestCryptoIPSecProfileWithoutPfsOmitsPfsXML(t *testing.T) {
	data := CryptoIPSecProfile{
		Name: types.StringValue("NO_PFS"),
	}

	body := data.toBodyXML(context.Background(), data)
	if strings.Contains(body, "<pfs") {
		t.Fatalf("did not expect a PFS element when no PFS attribute is configured: %s", body)
	}
}
