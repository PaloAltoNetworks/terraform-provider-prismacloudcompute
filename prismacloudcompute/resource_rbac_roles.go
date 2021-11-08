package prismacloudcompute

import (
	"fmt"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/prismacloudcompute/convert"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/auth"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/pcc"
)

func resourceRbacRoles() *schema.Resource {
	return &schema.Resource{
		Create: createRbacRole,
		Read:   readRbacRole,
		Update: updateRbacRole,
		Delete: deleteRbacRole,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Role description.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Role name.",
			},
			"permission": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of permissions.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Names roles for the user.",
						},
						"read_write": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Indicates the type of permission.",
						},
					},
				},
			},
		},
	}
}

func createRbacRole(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	parsedRole, err := convert.SchemaToRbacRole(d)
	if err != nil {
		return fmt.Errorf("error creating role '%+v': %s", parsedRole, err)
	}

	if err := auth.CreateRole(*client, parsedRole); err != nil {
		return fmt.Errorf("error creating role '%+v': %s", parsedRole, err)
	}

	d.SetId(parsedRole.Name)
	return readRbacRole(d, meta)
}

func readRbacRole(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	retrievedRole, err := auth.GetRole(*client, d.Id())
	if err != nil {
		return fmt.Errorf("error reading role: %s", err)
	}

	d.Set("description", retrievedRole.Description)
	d.Set("name", retrievedRole.Name)
	if err := d.Set("permission", convert.RbacRolePermissionsToSchema(retrievedRole.Permissions)); err != nil {
		return fmt.Errorf("error reading role: %s", err)
	}

	return nil
}

func updateRbacRole(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	parsedRole, err := convert.SchemaToRbacRole(d)
	if err != nil {
		return fmt.Errorf("error updating role: %s", err)
	}

	if err := auth.UpdateRole(*client, parsedRole); err != nil {
		return fmt.Errorf("error updating role: %s", err)
	}

	return readRbacRole(d, meta)
}

func deleteRbacRole(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	if err := auth.DeleteRole(*client, d.Id()); err != nil {
		return fmt.Errorf("error deleting role '%s': %s", d.Id(), err)
	}

	d.SetId("")
	return nil
}
