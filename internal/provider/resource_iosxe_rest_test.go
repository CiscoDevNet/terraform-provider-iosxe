package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceIOSXE__vlan__Rest(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceIOSXE__vlan__RestConfig,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
	})
}

const testAccResourceIOSXE__vlan__RestConfig = `
resource "iosxe_rest" "test" {
  method = "GET"
  path = "/data/Cisco-IOS-XE-native:native/vlan"
}
`
