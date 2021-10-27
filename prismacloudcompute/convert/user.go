package convert

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/auth"
)

func SchemaToUser(d *schema.ResourceData) (auth.User, error) {
	parsedUser := auth.User{}

	if val, ok := d.GetOk("authtype"); ok {
		parsedUser.AuthType = val.(string)
	}
	if val, ok := d.GetOk("password"); ok {
		parsedUser.Password = val.(string)
	}
	if val, ok := d.GetOk("permissions"); ok {
		parsedUser.Permissions = schemaToUserPermissions(val.([]interface{}))
	}
	if val, ok := d.GetOk("role"); ok {
		parsedUser.Role = val.(string)
	}
	if val, ok := d.GetOk("username"); ok {
		parsedUser.Username = val.(string)
	}

	return parsedUser, nil
}

func schemaToUserPermissions(in []interface{}) []auth.UserPermission {
	parsedPermissions := make([]auth.UserPermission, 0, len(in))
	for _, val := range in {
		presentPermission := val.(map[string]interface{})
		parsedPermissions = append(parsedPermissions, auth.UserPermission{
			Collections: SchemaToStringSlice(presentPermission["collections"].([]interface{})),
			Project:     presentPermission["project"].(string),
		})
	}
	return parsedPermissions
}

func UserPermissionsToSchema(in []auth.UserPermission) []interface{} {
	ans := make([]interface{}, 0, len(in))
	for _, val := range in {
		m := make(map[string]interface{})
		m["collections"] = strings.Join(val.Collections, ",")
		m["project"] = val.Project
		ans = append(ans, m)
	}
	return ans
}
