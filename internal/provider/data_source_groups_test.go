package provider

// import (
// 	"testing"

// 	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
// )

// func TestAccDsGroups(t *testing.T) {
// 	resource.Test(t, resource.TestCase{
// 		PreCheck:  func() { testAccPreCheck(t) },
// 		Providers: testAccProviders,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccDsGroupsConfig(),
// 				Check: resource.ComposeAggregateTestCheckFunc(
// 					resource.TestCheckResourceAttrSet("data.prismacloudcompute_groups.test", "total"),
// 				),
// 			},
// 		},
// 	})
// }

// func testAccDsGroupsConfig() string {
// 	return `
// data "prismacloudcompute_groups" "test" {}
// `
// }
