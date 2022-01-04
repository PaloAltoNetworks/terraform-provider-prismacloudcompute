package convert

import (
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/auth"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func SchemaToRbacRole(d *schema.ResourceData) (auth.Role, error) {
	parsedRole := auth.Role{}

	if val, ok := d.GetOk("description"); ok {
		parsedRole.Description = val.(string)
	}
	if val, ok := d.GetOk("name"); ok {
		parsedRole.Name = val.(string)
	}
	if val, ok := d.GetOk("permission"); ok {
		presentPermissions := val.([]interface{})
		parsedPermissions := make([]auth.RolePermission, 0, len(presentPermissions))
		for _, val := range presentPermissions {
			presentPermission := val.(map[string]interface{})
			parsedPermissions = append(parsedPermissions, auth.RolePermission{
				Name:      presentPermission["name"].(string),
				ReadWrite: presentPermission["read_write"].(bool),
			})
		}
		parsedRole.Permissions = parsedPermissions
	}

	return parsedRole, nil
}

func RbacRolePermissionsToSchema(in []auth.RolePermission) []interface{} {
	ans := make([]interface{}, 0, len(in))
	for _, val := range in {
		m := make(map[string]interface{})
		m["name"] = val.Name
		m["read_write"] = val.ReadWrite
		ans = append(ans, m)
	}
	return ans
}
