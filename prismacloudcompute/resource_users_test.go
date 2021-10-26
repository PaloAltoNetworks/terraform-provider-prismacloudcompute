package prismacloudcompute

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/auth"
)

func TestUsersConfig(t *testing.T) {
	fmt.Printf("\n\nStart TestAccUserConfig")
	var o user.User
	id := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccUserConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUserExists("prismacloudcompute_users.test", &o),
					testAccCheckUserAttributes(&o, id, true),
				),
			},
			{
				Config: testAccUserConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUserExists("prismacloudcompute_users.test", &o),
					testAccCheckUserAttributes(&o, id, true),
				),
			},
		},
	})
}

func TestUsersNetwork(t *testing.T) {
	var o user.User
	id := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccUserConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUserExists("prismacloudcompute_users.test", &o),
					testAccCheckUserAttributes(&o, id, true),
				),
			},
			{
				Config: testAccUserConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUserExists("prismacloudcompute_users.test", &o),
					testAccCheckUserAttributes(&o, id, true),
				),
			},
		},
	})
}

func TestUsersAuditEvent(t *testing.T) {
	var o user.User
	id := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccUserConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUserExists("prismacloudcompute_users.test", &o),
					testAccCheckUserAttributes(&o, id, true),
				),
			},
			{
				Config: testAccUserConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUserExists("prismacloudcompute_users.test", &o),
					testAccCheckUserAttributes(&o, id, true),
				),
			},
		},
	})
}

func testUsersExists(n string, o *user.User) resource.TestCheckFunc {
return func(s *terraform.State) error {
		// return fmt.Errorf("What is the name: %s", o.UserId)

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Resource not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Object label Id is not set")
		}

		client := testAccProvider.Meta().(*pcc.Client)
		lo, err := users.Get(*client)
		if err != nil {
			return fmt.Errorf("Error in get: %s", err)
		}
		*o = lo

		return nil
	}
}

// func testUsersAttributes(o *user.User, id string, learningDisabled bool) resource.TestCheckFunc {
// 	return func(s *terraform.State) error {
// 		if o.UserId != id {
// 			return fmt.Errorf("\n\nUserId is %s, expected %s", o.UserId, id)
// 		} else {
// 			fmt.Printf("\n\nName is %s", o.UserId)
// 		}

// 		if o.LearningDisabled != learningDisabled {
// 			return fmt.Errorf("LearningDisabled is %t, expected %t", o.LearningDisabled, learningDisabled)
// 		}

// 		return nil
// 	}
// }

// func testUsersDestroy(s *terraform.State) error {
// 	/*	client := testAccProvider.Meta().(*pcc.Client)

// 		for _, rs := range s.RootModule().Resources {

// 			if rs.Type != "prismacloudcompute_users" {
// 				continue
// 			}

// 			if rs.Primary.ID != "" {
// 				name := rs.Primary.ID
// 				if err := user.Delete(client, name); err == nil {
// 					return fmt.Errorf("Object %q still exists", name)
// 				}
// 			}
// 			return nil
// 		}
// 	*/
// 	return nil
// }

// func testUsersConfig(id string) string {
// 	var buf bytes.Buffer
// 	buf.Grow(500)

// 	buf.WriteString(fmt.Sprintf(`
// resource "prismacloudcompute_users" "test" {
//     name = %q
// }`, id))

// 	return buf.String()
// }
