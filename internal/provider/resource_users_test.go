package provider

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/auth"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestUsersConfig(t *testing.T) {
	fmt.Printf("\n\nStart TestAccUserConfig")
	var o auth.User
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
					testAccCheckUserAttributes(&o, id, "sysadmin"),
				),
			},
			{
				Config: testAccUserConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUserExists("prismacloudcompute_users.test", &o),
					testAccCheckUserAttributes(&o, id, "sysadmin"),
				),
			},
		},
	})
}

func TestUsersNetwork(t *testing.T) {
	var o auth.User
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
					testAccCheckUserAttributes(&o, id, "sysadmin"),
				),
			},
			{
				Config: testAccUserConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUserExists("prismacloudcompute_users.test", &o),
					testAccCheckUserAttributes(&o, id, "sysadmin"),
				),
			},
		},
	})
}

func TestUsersAuditEvent(t *testing.T) {
	var o auth.User
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
					testAccCheckUserAttributes(&o, id, "sysadmin"),
				),
			},
			{
				Config: testAccUserConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUserExists("prismacloudcompute_users.test", &o),
					testAccCheckUserAttributes(&o, id, "sysadmin"),
				),
			},
		},
	})
}

func testAccCheckUserExists(n string, o *auth.User) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// return fmt.Errorf("What is the name: %s", o.UserId)

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Resource not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Object label Id is not set")
		}

		client := testAccProvider.Meta().(*api.Client)
		id := rs.Primary.ID
		lo, err := auth.GetUser(*client, id)
		if err != nil {
			return fmt.Errorf("Error in get: %s", err)
		}
		o = lo

		return nil
	}
}

func testAccCheckUserAttributes(o *auth.User, name, role string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if o.Username != name {
			return fmt.Errorf("\n\nUsername is %s, expected %s", o.Username, name)
		} else {
			fmt.Printf("\n\nName is %s", o.Username)
		}

		if o.Role != role {
			return fmt.Errorf("Role is %s, expected %s", o.Role, role)
		}

		return nil
	}
}

func testAccUserDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*api.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type != "prismacloudcompute_users" {
			continue
		}

		if rs.Primary.ID != "" {
			name := rs.Primary.ID
			if err := auth.DeleteUser(*client, name); err == nil {
				return fmt.Errorf("Object %q still exists", name)
			}
		}
		return nil
	}
	return nil
}

func testAccUserConfig(id string) string {
	var buf bytes.Buffer
	buf.Grow(500)

	buf.WriteString(fmt.Sprintf(`
resource "prismacloudcompute_users" "test" {
    Username = %q

}`, id))

	return buf.String()
}
