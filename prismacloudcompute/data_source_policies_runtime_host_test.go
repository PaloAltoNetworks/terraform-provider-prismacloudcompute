package prismacloudcompute

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDsPolicies(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDsPoliciesConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.prismacloudcompute_policies_runtime_host.test", "total"),
				),
			},
		},
	})
}

func testAccDsPoliciesConfig() string {
	return `
data "prismacloudcompute_policies" "test" {
    learningDisabled = true
    rule {
        name = "my rule"
    }
}
`
}
