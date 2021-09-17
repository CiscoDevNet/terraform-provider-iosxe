package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// testAccResourceIOSXERestTemplate test is just a template to generate tests for different features for the IOSXE terraform provider
//
// To use this template, replace __feature__ with appropriate feature name (e.g. vlan) in every possible place.
//
// An example for vlan is given in the same folder as this.
func testAccResourceIOSXE__feature__Rest(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceIOSXE__feature__RestConfig,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
	})
}

const testAccResourceIOSXE__feature__RestConfig = `
resource "iosxe_rest" "test" {
  method = "" # method to be called
  path = "" # path to the feature
  payload = "" # payload if required
}
`
