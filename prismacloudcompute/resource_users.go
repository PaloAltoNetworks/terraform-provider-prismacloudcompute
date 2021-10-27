package prismacloudcompute

import (
	"fmt"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/prismacloudcompute/convert"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/auth"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/pcc"
)

func resourceUsers() *schema.Resource {
	return &schema.Resource{
		Create: createUser,
		Read:   readUser,
		Update: updateUser,
		Delete: deleteUser,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"authtype": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The user authentication type.",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Password.",
			},
			"permissions": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of permissions.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"collections": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Specifies the set of Defenders in-scope for working on a scan job.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"project": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Names of projects which the user can access.",
						},
					},
				},
			},
			"role": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Role.",
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Username.",
			},
		},
	}
}

func createUser(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	parsedUser, err := convert.SchemaToUser(d)
	if err != nil {
		return fmt.Errorf("failed to create user '%+v': %s", parsedUser, err)
	}

	if err := auth.UpdateUser(*client, parsedUser); err != nil {
		return fmt.Errorf("failed to create user '%+v': %s", parsedUser, err)
	}

	d.SetId(parsedUser.Username)
	return readUser(d, meta)
}

func readUser(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	retrievedUser, err := auth.GetUser(*client, d.Id())
	if err != nil {
		return fmt.Errorf("failed to read user: %s", err)
	}

	d.Set("authtype", retrievedUser.AuthType)
	d.Set("password", retrievedUser.Password)
	if err := d.Set("permissions", convert.UserPermissionsToSchema(retrievedUser.Permissions)); err != nil {
		return fmt.Errorf("failed to read user: %s", retrievedUser.Permissions, err)
	}
	d.Set("role", retrievedUser.Role)
	d.Set("username", retrievedUser.Username)

	return nil
}

func updateUser(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	parsedUser, err := convert.SchemaToUser(d)
	if err != nil {
		return fmt.Errorf("failed to update user: %s", err)
	}

	if err := auth.UpdateUser(*client, parsedUser); err != nil {
		return fmt.Errorf("failed to update user: %s", err)
	}

	return readUser(d, meta)
}

func deleteUser(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	if err := auth.DeleteUser(*client, d.Id()); err != nil {
		return fmt.Errorf("failed to delete user '%s': %s", d.Id(), err)
	}

	d.SetId("")
	return nil
}
