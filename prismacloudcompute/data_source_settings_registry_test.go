package prismacloudcompute

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDsRegistry(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDsRegistryConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.prismacloudcompute_registry.test", "total"),
				),
			},
		},
	})
}

func testAccDsRegistryConfig() string {
	return `
data "prismacloudcompute_registry" "test" {}
`
}
