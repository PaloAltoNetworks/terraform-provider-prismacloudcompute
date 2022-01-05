package provider

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/settings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccRegistryConfig(t *testing.T) {
	fmt.Printf("\n\nStart TestAccRegistryConfig")
	var o settings.RegistrySettings
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccRegistryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRegistryConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRegistryExists("prismacloudcompute_registry.test", o),
					testAccCheckRegistryAttributes(&o, name, "description", "#000000"),
				),
			},
			{
				Config: testAccRegistryConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRegistryExists("prismacloudcompute_registry.test", o),
					testAccCheckRegistryAttributes(&o, name, "description", "#000000"),
				),
			},
		},
	})
}

func TestAccRegistryNetwork(t *testing.T) {
	var o settings.RegistrySettings
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccRegistryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRegistryConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRegistryExists("prismacloudcompute_registry.test", o),
					testAccCheckRegistryAttributes(&o, name, "description", "#000000"),
				),
			},
			{
				Config: testAccRegistryConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRegistryExists("prismacloudcompute_registry.test", o),
					testAccCheckRegistryAttributes(&o, name, "description", "#000000"),
				),
			},
		},
	})
}

func TestAccRegistryAuditEvent(t *testing.T) {
	var o settings.RegistrySettings
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccRegistryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRegistryConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRegistryExists("prismacloudcompute_registry.test", o),
					testAccCheckRegistryAttributes(&o, name, "description", "#000000"),
				),
			},
			{
				Config: testAccRegistryConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRegistryExists("prismacloudcompute_registry.test", o),
					testAccCheckRegistryAttributes(&o, name, "description", "#000000"),
				),
			},
		},
	})
}

func testAccCheckRegistryExists(n string, o settings.RegistrySettings) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// return fmt.Errorf("What is the name: %s", o.Name)

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Resource not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Object label Name is not set")
		}

		client := testAccProvider.Meta().(*api.Client)
		lo, err := settings.GetRegistrySettings(*client)
		if err != nil {
			return fmt.Errorf("Error in get: %s", err)
		}
		o = lo

		return nil
	}
}

func testAccCheckRegistryAttributes(o *settings.RegistrySettings, name string, description string, color string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// if o.Name != name {
		// 	return fmt.Errorf("\n\nName is %s, expected %s", o.Name, name)
		// } else {
		// 	fmt.Printf("\n\nName is %s", o.Name)
		// }

		// if o.Description != description {
		// 	return fmt.Errorf("Description is %s, expected %s", o.Description, description)
		// }

		// if o.Color != color {
		// 	return fmt.Errorf("Color type is %q, expected %q", o.Color, color)
		// }

		return nil
	}
}

func testAccRegistryDestroy(s *terraform.State) error {
	// client := testAccProvider.Meta().(*api.Client)

	// for _, rs := range s.RootModule().Resources {

	// 	if rs.Type != "prismacloudcompute_registry" {
	// 		continue
	// 	}

	// 	if rs.Primary.ID != "" {
	// 		name := rs.Primary.ID
	// 		if err := registry.Delete(*client, name); err == nil {
	// 			return fmt.Errorf("Object %q still exists", name)
	// 		}
	// 	}
	// 	return nil
	// }

	return nil
}

func testAccRegistryConfig(name string) string {
	var buf bytes.Buffer
	buf.Grow(500)

	buf.WriteString(fmt.Sprintf(`
resource "prismacloudcompute_registry" "test" {
    name = %q
}`, name))

	return buf.String()
}
