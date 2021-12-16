package prismacloudcompute

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDsCustomRule(t *testing.T) {
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDsCustomRule(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.prismacloudcompute_custom_rule.test", "name"),
				),
			},
		},
	})
}

func testAccDsCustomRule(name string) string {
	return fmt.Sprintf(`
	data "prismacloudcompute_custom_rule" "test" {
		name = %q
	}
	`, name)
}
