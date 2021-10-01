package prismacloudcompute

import (
	"fmt"
	"time"

	"github.com/paloaltonetworks/prisma-cloud-compute-go/pcc"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/auth"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRbacRoles() *schema.Resource {
	return &schema.Resource{
		Create: createRbacRole,
		Read:   readRbacRole,
		Update: updateRbacRole,
		Delete: deleteRbacRole,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

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
			"permissions": {
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
						"readwrite": {
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
	parsedRole, err := parseRbacRole(d)
	if err != nil {
		return fmt.Errorf("error parsing role: %s", err)
	}

	if err := auth.UpdateRole(*client, parsedRole); err != nil {
		return fmt.Errorf("error creating role: %s %s", err, parsedRole.Name)
	}

	d.SetId(parsedRole.Name)
	return readRbacRole(d, meta)
}

func readRbacRole(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	roleList, err := auth.ListRoles(*client)
	retrievedRole := roleList[0]
	if err != nil {
		return fmt.Errorf("error reading role: %s", err)
	}

	if err := d.Set("description", retrievedRole.Description); err != nil {
		return fmt.Errorf("error reading %s description: %s", retrievedRole.Description, err)
	}
	if err := d.Set("name", retrievedRole.Name); err != nil {
		return fmt.Errorf("error reading %s name: %s", retrievedRole.Name, err)
	}
	if err := d.Set("perms", flattenRolePermissions(retrievedRole.Permissions)); err != nil {
		return fmt.Errorf("error reading %s perms: %s", retrievedRole.Permissions, err)
	}

	return nil
}

func updateRbacRole(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	parsedRole, err := parseRbacRole(d)
	if err != nil {
		return fmt.Errorf("error parsing rold for update: %s", err)
	}

	if err := auth.UpdateRole(*client, parsedRole); err != nil {
		return fmt.Errorf("error updating role: %s", err)
	}

	return readRbacRole(d, meta)
}

func deleteRbacRole(d *schema.ResourceData, meta interface{}) error {
	// TODO: reset to default role
	client := meta.(*pcc.Client)
	id := d.Id()

	if err := auth.DeleteRole(*client, id); err != nil {
		return fmt.Errorf("failed to update credential: %s", err)
	}

	d.SetId("")
	return nil
}

func parseRbacRole(d *schema.ResourceData) (auth.Role, error) {
	parsedRole := auth.Role{}
	
	if d.Get("description") != nil {
		parsedRole.Description = d.Get("description").(string)
	}
	if d.Get("name") != nil {
		parsedRole.Name = d.Get("name").(string)
	}
	if d.Get("perms") != nil && len(d.Get("perms").([]interface{})) > 0 {
		parsedRole.Permissions = convertRolePermissions(d.Get("perms").([]interface{}))
	}

	return parsedRole, nil
}

func flattenRolePermissions(in []auth.RolePermission) []interface{} {
	ans := make([]interface{}, 0, len(in))
	for _, val := range in {
		m := make(map[string]interface{})
		m["name"] = val.Name
		m["readWrite"] = val.ReadWrite
		ans = append(ans, m)
	}
	return ans
}

func convertRolePermissions(in []interface{}) []auth.RolePermission {
	ans := make([]auth.RolePermission, 0, len(in))
	for _, val := range in {
		valMap := val.(map[string]interface{})
		m := auth.RolePermission{}
		if valMap["name"] != nil {
			m.Name = valMap["name"].(string)
		}
		if valMap["readwrite"] != nil {
			m.ReadWrite = valMap["readwrite"].(bool)
		}
		ans = append(ans, m)
	}
	return ans
}
