package prismacloudcompute

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDsCredentials(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDsCredentialsConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.prismacloudcompute_credentials.test", "total"),
				),
			},
		},
	})
}

func testAccDsCredentialsConfig() string {
	return `
data "prismacloudcompute_credentials" "test" {}
`
}
