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

func TestRbacRolesConfig(t *testing.T) {
	fmt.Printf("\n\nStart TestAccRbacRolesConfig")
	var o auth.Role
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
					testAccCheckRbacRolesAttributes(&o, id, "plop"),
				),
			},
			{
				Config: testAccRbacRolesConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRbacRolesExists("prismacloudcompute_rbac_roles.test", &o),
					testAccCheckRbacRolesAttributes(&o, id, "plop"),
				),
			},
		},
	})
}

func TestRbacRolesNetwork(t *testing.T) {
	var o auth.Role
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
					testAccCheckRbacRolesAttributes(&o, id, "plop"),
				),
			},
			{
				Config: testAccRbacRolesConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRbacRolesExists("prismacloudcompute_rbac_roles.test", &o),
					testAccCheckRbacRolesAttributes(&o, id, "plop"),
				),
			},
		},
	})
}

func TestRbacRolesAuditEvent(t *testing.T) {
	var o auth.Role
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
					testAccCheckRbacRolesAttributes(&o, id, "plop"),
				),
			},
			{
				Config: testAccRbacRolesConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRbacRolesExists("prismacloudcompute_rbac_roles.test", &o),
					testAccCheckRbacRolesAttributes(&o, id, "plop"),
				),
			},
		},
	})
}

func testAccCheckRbacRolesExists(n string, o *auth.Role) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// return fmt.Errorf("What is the name: %s", o.GroupId)

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Resource not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Object label Id is not set")
		}

		client := testAccProvider.Meta().(*api.Client)
		id := rs.Primary.ID
		lo, err := auth.GetRole(*client, id)
		if err != nil {
			return fmt.Errorf("Error in get: %s", err)
		}
		o = lo

		return nil
	}
}

func testAccCheckRbacRolesAttributes(o *auth.Role, name, description string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if o.Name != name {
			return fmt.Errorf("\n\nRole name is %s, expected %s", o.Name, name)
		} else {
			fmt.Printf("\n\nName is %s", o.Name)
		}

		if o.Description != description {
			return fmt.Errorf("Description is %s, expected %s", o.Description, description)
		}

		return nil
	}
}

func testAccRbacRolesDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*api.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type != "prismacloudcompute_rbac_roles" {
			continue
		}

		if rs.Primary.ID != "" {
			name := rs.Primary.ID
			if err := auth.DeleteRole(*client, name); err == nil {
				return fmt.Errorf("Object %q still exists", name)
			}
		}
		return nil
	}
	return nil
}

func testAccRbacRolesConfig(id string) string {
	var buf bytes.Buffer
	buf.Grow(500)

	buf.WriteString(fmt.Sprintf(`
resource "prismacloudcompute_rbac_roles" "test" {
    name = %q
}`, id))

	return buf.String()
}
