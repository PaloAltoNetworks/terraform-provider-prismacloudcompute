package prismacloudcompute

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/auth"
)

func TestRbacRolesConfig(t *testing.T) {
	fmt.Printf("\n\nStart TestAccRbacRolesConfig")
	var o group.Group
	id := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccRbacRolesDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRbacRolesConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRbacRolesExists("prismacloudcompute_rbac_roles.test", &o),
					testAccCheckRbacRolesAttributes(&o, id, true),
				),
			},
			{
				Config: testAccRbacRolesConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRbacRolesExists("prismacloudcompute_rbac_roles.test", &o),
					testAccCheckRbacRolesAttributes(&o, id, true),
				),
			},
		},
	})
}

func TestRbacRolesNetwork(t *testing.T) {
	var o group.Group
	id := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccRbacRolesDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRbacRolesConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRbacRolesExists("prismacloudcompute_rbac_roles.test", &o),
					testAccCheckRbacRolesAttributes(&o, id, true),
				),
			},
			{
				Config: testAccRbacRolesConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRbacRolesExists("prismacloudcompute_rbac_roles.test", &o),
					testAccCheckRbacRolesAttributes(&o, id, true),
				),
			},
		},
	})
}

func TestRbacRolesAuditEvent(t *testing.T) {
	var o group.Group
	id := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccRbacRolesDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRbacRolesConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRbacRolesExists("prismacloudcompute_rbac_roles.test", &o),
					testAccCheckRbacRolesAttributes(&o, id, true),
				),
			},
			{
				Config: testAccRbacRolesConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRbacRolesExists("prismacloudcompute_rbac_roles.test", &o),
					testAccCheckRbacRolesAttributes(&o, id, true),
				),
			},
		},
	})
}

func testRbacRolesExists(n string, o *group.Group) resource.TestCheckFunc {
return func(s *terraform.State) error {
		// return fmt.Errorf("What is the name: %s", o.GroupId)

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Resource not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Object label Id is not set")
		}

		client := testAccProvider.Meta().(*pcc.Client)
		lo, err := RbacRoles.Get(*client)
		if err != nil {
			return fmt.Errorf("Error in get: %s", err)
		}
		*o = lo

		return nil
	}
}

// func testRbacRolesAttributes(o *group.Group, id string, learningDisabled bool) resource.TestCheckFunc {
// 	return func(s *terraform.State) error {
// 		if o.GroupId != id {
// 			return fmt.Errorf("\n\nGroupId is %s, expected %s", o.GroupId, id)
// 		} else {
// 			fmt.Printf("\n\nName is %s", o.GroupId)
// 		}

// 		if o.LearningDisabled != learningDisabled {
// 			return fmt.Errorf("LearningDisabled is %t, expected %t", o.LearningDisabled, learningDisabled)
// 		}

// 		return nil
// 	}
// }

// func testRbacRolesDestroy(s *terraform.State) error {
// 	/*	client := testAccProvider.Meta().(*pcc.Client)

// 		for _, rs := range s.RootModule().Resources {

// 			if rs.Type != "prismacloudcompute_rbac_roles" {
// 				continue
// 			}

// 			if rs.Primary.ID != "" {
// 				name := rs.Primary.ID
// 				if err := group.Delete(client, name); err == nil {
// 					return fmt.Errorf("Object %q still exists", name)
// 				}
// 			}
// 			return nil
// 		}
// 	*/
// 	return nil
// }

// func testRbacRolesConfig(id string) string {
// 	var buf bytes.Buffer
// 	buf.Grow(500)

// 	buf.WriteString(fmt.Sprintf(`
// resource "prismacloudcompute_rbac_roles" "test" {
//     name = %q
// }`, id))

// 	return buf.String()
// }
