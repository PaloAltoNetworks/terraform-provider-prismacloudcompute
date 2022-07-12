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

func TestGroupsConfig(t *testing.T) {
	fmt.Printf("\n\nStart TestAccGroupsConfig")
	var o auth.Group
	id := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccGroupsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGroupsConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGroupsExists("prismacloudcompute_groups.test", &o),
					testAccCheckGroupsAttributes(&o, id, true),
				),
			},
			{
				Config: testAccGroupsConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGroupsExists("prismacloudcompute_groups.test", &o),
					testAccCheckGroupsAttributes(&o, id, true),
				),
			},
		},
	})
}

func TestGroupsNetwork(t *testing.T) {
	var o auth.Group
	id := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccGroupsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGroupsConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGroupsExists("prismacloudcompute_groups.test", &o),
					testAccCheckGroupsAttributes(&o, id, true),
				),
			},
			{
				Config: testAccGroupsConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGroupsExists("prismacloudcompute_groups.test", &o),
					testAccCheckGroupsAttributes(&o, id, true),
				),
			},
		},
	})
}

func TestGroupsAuditEvent(t *testing.T) {
	var o auth.Group
	id := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccGroupsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGroupsConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGroupsExists("prismacloudcompute_groups.test", &o),
					testAccCheckGroupsAttributes(&o, id, true),
				),
			},
			{
				Config: testAccGroupsConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGroupsExists("prismacloudcompute_groups.test", &o),
					testAccCheckGroupsAttributes(&o, id, true),
				),
			},
		},
	})
}

func testAccCheckGroupsExists(n string, o *auth.Group) resource.TestCheckFunc {
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
		lo, err := auth.GetGroup(*client, id)
		if err != nil {
			return fmt.Errorf("Error in get: %s", err)
		}
		o = lo

		return nil
	}
}

func testAccCheckGroupsAttributes(o *auth.Group, id string, oauthGroup bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if o.Id != id {
			return fmt.Errorf("\n\nGroupId is %s, expected %s", o.Id, id)
		} else {
			fmt.Printf("\n\nName is %s", o.Id)
		}

		if o.OauthGroup != oauthGroup {
			return fmt.Errorf("LearningDisabled is %t, expected %t", o.OauthGroup, oauthGroup)
		}

		return nil
	}
}

func testAccGroupsDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*api.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type != "prismacloudcompute_groups" {
			continue
		}

		if rs.Primary.ID != "" {
			name := rs.Primary.ID
			if err := auth.DeleteGroup(*client, name); err == nil {
				return fmt.Errorf("Object %q still exists", name)
			}
		}
		return nil
	}
	return nil
}

func testAccGroupsConfig(id string) string {
	var buf bytes.Buffer
	buf.Grow(500)

	buf.WriteString(fmt.Sprintf(`
resource "prismacloudcompute_groups" "test" {
    name = %q
}`, id))

	return buf.String()
}
