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

func TestCredentialsConfig(t *testing.T) {
	fmt.Printf("\n\nStart TestAccCredentialsConfig")
	var o auth.Credential
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCredentialsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCredentialsConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCredentialsExists("prismacloudcompute_credentials.test", &o),
					testAccCheckCredentialsAttributes(&o, name, true),
				),
			},
			{
				Config: testAccCredentialsConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCredentialsExists("prismacloudcompute_credentials.test", &o),
					testAccCheckCredentialsAttributes(&o, name, true),
				),
			},
		},
	})
}

func TestCredentialsNetwork(t *testing.T) {
	var o auth.Credential
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCredentialsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCredentialsConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCredentialsExists("prismacloudcompute_credentials.test", &o),
					testAccCheckCredentialsAttributes(&o, name, true),
				),
			},
			{
				Config: testAccCredentialsConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCredentialsExists("prismacloudcompute_credentials.test", &o),
					testAccCheckCredentialsAttributes(&o, name, true),
				),
			},
		},
	})
}

func TestCredentialsAuditEvent(t *testing.T) {
	var o auth.Credential
	id := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCredentialsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCredentialsConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCredentialsExists("prismacloudcompute_credentials.test", &o),
					testAccCheckCredentialsAttributes(&o, id, true),
				),
			},
			{
				Config: testAccCredentialsConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCredentialsExists("prismacloudcompute_credentials.test", &o),
					testAccCheckCredentialsAttributes(&o, id, true),
				),
			},
		},
	})
}

func testAccCheckCredentialsExists(n string, o *auth.Credential) resource.TestCheckFunc {
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
		lo, err := auth.GetCredential(*client, id)
		if err != nil {
			return fmt.Errorf("Error in get: %s", err)
		}
		o = lo

		return nil
	}
}

func testAccCheckCredentialsAttributes(o *auth.Credential, id string, skipVerify bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if o.Id != id {
			return fmt.Errorf("\n\nCredentialId is %s, expected %s", o.Id, id)
		} else {
			fmt.Printf("\n\nId is %s", o.Id)
		}

		if o.SkipVerify != skipVerify {
			return fmt.Errorf("SkipVerify is %t, expected %t", o.SkipVerify, skipVerify)
		}

		return nil
	}
}

func testAccCredentialsDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*api.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type != "prismacloudcompute_credentials" {
			continue
		}

		if rs.Primary.ID != "" {
			id := rs.Primary.ID
			if err := auth.DeleteCredential(*client, id); err == nil {
				return fmt.Errorf("Object %q still exists", id)
			}
		}
		return nil
	}
	return nil
}

func testAccCredentialsConfig(id string) string {
	var buf bytes.Buffer
	buf.Grow(500)

	buf.WriteString(fmt.Sprintf(`
resource "prismacloudcompute_credentials" "test" {
    id = %q
}`, id))

	return buf.String()
}
