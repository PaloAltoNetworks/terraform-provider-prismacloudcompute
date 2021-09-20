package prismacloudcompute

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDsCollections(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDsCollectionsConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.prismacloudcompute_collections.test", "total"),
				),
			},
		},
	})
}

func testAccDsCollectionsConfig() string {
	return `
data "prismacloudcompute_collections" "test" {}
`
}
