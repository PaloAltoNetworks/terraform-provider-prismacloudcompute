package convert

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/auth"
)

func SchemaToGroup(d *schema.ResourceData) (auth.Group, error) {
	parsedGroup := auth.Group{}

	if val, ok := d.GetOk("groupid"); ok {
		parsedGroup.Id = val.(string)
	}
	if val, ok := d.GetOk("ldapgroup"); ok {
		parsedGroup.LdapGroup = val.(bool)
	}
	if val, ok := d.GetOk("groupname"); ok {
		parsedGroup.Name = val.(string)
	}
	if val, ok := d.GetOk("oauthgroup"); ok {
		parsedGroup.OauthGroup = val.(bool)
	}
	if val, ok := d.GetOk("oidcgroup"); ok {
		parsedGroup.OidcGroup = val.(bool)
	}
	if val, ok := d.GetOk("permissions"); ok {
		parsedGroup.Permissions = schemaToGroupPermissions(val.([]interface{}))
	}
	if val, ok := d.GetOk("role"); ok {
		parsedGroup.Role = val.(string)
	}
	if val, ok := d.GetOk("samlgroup"); ok {
		parsedGroup.SamlGroup = val.(bool)
	}
	if val, ok := d.GetOk("users"); ok && len(val.([]interface{})) > 0 {
		parsedGroup.Users = schemaToGroupUsers(val.([]interface{}))
	}

	return parsedGroup, nil
}

func schemaToGroupPermissions(in []interface{}) []auth.GroupPermission {
	parsedPermissions := make([]auth.GroupPermission, 0, len(in))
	for _, val := range in {
		presentPermission := val.(map[string]interface{})
		parsedPermissions = append(parsedPermissions, auth.GroupPermission{
			Collections: SchemaToStringSlice(presentPermission["collections"].([]interface{})),
			Project:     presentPermission["project"].(string),
		})
	}
	return parsedPermissions
}

func schemaToGroupUsers(in []interface{}) []auth.GroupUser {
	parsedUsers := make([]auth.GroupUser, 0, len(in))
	for _, val := range in {
		presentUser := val.(map[string]interface{})
		parsedUsers = append(parsedUsers, auth.GroupUser{
			Username: presentUser["username"].(string),
		})
	}
	return parsedUsers
}

func GroupPermissionsToSchema(in []auth.GroupPermission) []interface{} {
	ans := make([]interface{}, 0, len(in))
	for _, val := range in {
		m := make(map[string]interface{})
		m["collections"] = strings.Join(val.Collections, ",")
		m["project"] = val.Project
		ans = append(ans, m)
	}
	return ans
}

func GroupUsersToSchema(in []auth.GroupUser) []interface{} {
	ans := make([]interface{}, 0, len(in))
	for _, val := range in {
		m := make(map[string]interface{})
		m["username"] = val.Username
		ans = append(ans, m)
	}
	return ans
}
