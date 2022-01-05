package provider

import (
	"fmt"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/auth"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/convert"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
			"id": {
				Description: "The ID of the user.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"authentication_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The user authentication type.",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Password.",
			},
			"permissions": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
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
				Required:    true,
				Description: "Role.",
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Username.",
			},
		},
	}
}

func createUser(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*api.Client)
	parsedUser, err := convert.SchemaToUser(d)
	if err != nil {
		return fmt.Errorf("failed to create user '%+v': %s", parsedUser, err)
	}

	if err := auth.CreateUser(*client, parsedUser); err != nil {
		return fmt.Errorf("failed to create user '%+v': %s", parsedUser, err)
	}

	d.SetId(parsedUser.Username)
	return readUser(d, meta)
}

func readUser(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*api.Client)
	retrievedUser, err := auth.GetUser(*client, d.Id())
	if err != nil {
		return fmt.Errorf("failed to read user: %s", err)
	}

	d.Set("authentication_type", retrievedUser.AuthType)
	d.Set("password", retrievedUser.Password)
	if err := d.Set("permissions", convert.UserPermissionsToSchema(retrievedUser.Permissions)); err != nil {
		return fmt.Errorf("failed to read user: %s", err)
	}
	d.Set("role", retrievedUser.Role)
	d.Set("username", retrievedUser.Username)

	return nil
}

func updateUser(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*api.Client)
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
	client := meta.(*api.Client)
	if err := auth.DeleteUser(*client, d.Id()); err != nil {
		return fmt.Errorf("failed to delete user '%s': %s", d.Id(), err)
	}

	d.SetId("")
	return nil
}
