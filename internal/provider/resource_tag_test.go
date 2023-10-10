package provider

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/tag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccTagConfig(t *testing.T) {
	fmt.Printf("\n\nStart TestAccTagConfig")
	var o tag.Tag
	name := fmt.Sprintf("tf%s", acctest.RandString(6))
	description := fmt.Sprintf("tf%s", acctest.RandString(6))
	vuln := []tag.Vuln{{
		Id:           "CVE-2024-0001",
		PackageName:  "*",
		ResourceType: "*",
	}}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccTagDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTagConfig(name, description, nil),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTagExists("prismacloudcompute_tag.test", &o),
					testAccCheckTagAttributes(&o, name, description, "#A020F0", []tag.Vuln{}),
				),
			},
			{
				Config: testAccTagConfig(name, description, &vuln),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTagExists("prismacloudcompute_tag.test", &o),
					testAccCheckTagAttributes(&o, name, description, "#A020F0", vuln),
				),
			},
		},
	})
}

func testAccCheckTagExists(n string, o *tag.Tag) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Resource not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Object label Name is not set")
		}

		client := testAccProvider.Meta().(*api.Client)
		name := rs.Primary.ID
		lo, err := tag.GetTag(*client, name)
		if err != nil {
			return fmt.Errorf("Error in get: %s", err)
		}

		*o = *lo

		return nil
	}
}

func testAccCheckTagAttributes(o *tag.Tag, name string, description string, color string, vulns []tag.Vuln) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if o.Name != name {
			return fmt.Errorf("\n\nName is %s, expected %s", o.Name, name)
		}

		if o.Description != description {
			return fmt.Errorf("Description is %s, expected %s", o.Description, description)
		}

		if o.Color != color {
			return fmt.Errorf("Color type is %q, expected %q", o.Color, color)
		}

		if !reflect.DeepEqual(o.Vulns, vulns) {
			return fmt.Errorf("Vulns is %#v, expected %#v", o.Vulns, vulns)
		}

		return nil
	}
}

func testAccTagDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*api.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type != "prismacloudcompute_tag" {
			continue
		}

		if rs.Primary.ID != "" {
			name := rs.Primary.ID
			if err := tag.DeleteTag(*client, name); err == nil {
				return fmt.Errorf("Object %q still exists", name)
			}
		}
		return nil
	}

	return nil
}

func testAccTagConfig(name string, description string, vulns *[]tag.Vuln) string {
	var buf bytes.Buffer
	buf.Grow(500)

	if vulns == nil {
		buf.WriteString(fmt.Sprintf(`
resource "prismacloudcompute_tag" "test" {
name = %q
description = %q
}`, name, description))
	} else {
		buf.WriteString(fmt.Sprintf(`
resource "prismacloudcompute_tag" "test" {
name = %q
description = %q
`, name, description))

		for _, vuln := range *vulns {
			buf.WriteString(fmt.Sprintf(`
assignment {
id = %q
}
`, vuln.Id))
		}

		buf.WriteString("}")
	}

	return buf.String()
}
