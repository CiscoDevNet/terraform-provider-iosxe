package provider

import (
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// TestAccIosxeCryptoIKEv2ProposalPQCUnsupported verifies the device-side
// failure mode on platforms that run IOS XE but do not implement the PQC YANG
// nodes. It is intentionally separate from the generated positive test because
// it requires a C8000V running IOS XE 17.18.
func TestAccIosxeCryptoIKEv2ProposalPQCUnsupported(t *testing.T) {
	if os.Getenv("IOSXE1718_C8000V") == "" {
		t.Skip("skipping unsupported-platform test; set IOSXE1718_C8000V=1")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
resource "iosxe_crypto_ikev2_proposal" "unsupported_pqc" {
  name = "TF_PQC_UNSUPPORTED_TEST"

  pqc_mlkem768 = true
  pqc_optional = true

  encryption_aes_cbc_256 = true
  integrity_sha256       = true
  group_nineteen         = true
}
`,
				ExpectError: regexp.MustCompile(`(?s)unknown-element.*pqc`),
			},
		},
	})
}
