package prismacloudcompute

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDsPoliciesComplianceCiImages(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDsPoliciesComplianceCiImagesConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.prismacloudcompute_policies_compliance_container.test", "total"),
				),
			},
		},
	})
}

func testAccDsPoliciesComplianceCiImagesConfig() string {
	return `
data "prismacloudcompute_policies" "test" {
    learningDisabled = true
    rule {
        name = "my rule"
    }
}
`
}
