package prismacloudcompute

import (
	"fmt"
	"time"

	"github.com/paloaltonetworks/prisma-cloud-compute-go/pcc"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/auth"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGroups() *schema.Resource {
	return &schema.Resource{
		Create: createGroup,
		Read:   readGroup,
		Update: updateGroup,
		Delete: deleteGroup,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"groupid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Group ID.",
			},
			"ldapgroup": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates if the group is an LDAP group (true) or not (false).",
			},
			"groupname": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Group name.",
			},
			"oauthgroup": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates if the group is an OAuth group (true) or not (false).",
			},
			"oidcgroup": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates if the group is an OpenID Connect group (true) or not (false).",
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
				Description: "Role of the group.",
			},
			"samlgroup": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates if the group is a SAML group (true) or not (false).",
			},			
			"user": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Users in the group.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"username": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Name of a user.",
						},
					},
				},
			},
		},
	}
}

func createGroup(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	parsedGroup, err := parseGroups(d)
	if err != nil {
		return fmt.Errorf("error creating user: %s", err)
	}

	if err := auth.UpdateGroup(*client, parsedGroup); err != nil {
		return fmt.Errorf("error creating user: %s", err)
	}

	d.SetId(parsedGroup.Id)
	return readGroup(d, meta)
}

func readGroup(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	retrievedGroup, err := auth.GetGroups(*client)
	if err != nil {
		return fmt.Errorf("error reading user: %s", err)
	}
	
	firstGroup := retrievedGroup[0]

	if err := d.Set("groupId", firstGroup.Id); err != nil {
		return fmt.Errorf("error reading %s groupId: %s", firstGroup.Id, err)
	}
	if err := d.Set("ldapGroup", firstGroup.LdapGroup); err != nil {
		return fmt.Errorf("error reading %s ldapGroup: %s", firstGroup.LdapGroup, err)
	}
	if err := d.Set("groupName", firstGroup.Name); err != nil {
		return fmt.Errorf("error reading %s groupName: %s", firstGroup.Name, err)
	}
	if err := d.Set("oauthGroup", firstGroup.OauthGroup); err != nil {
		return fmt.Errorf("error reading %s oauthGroup: %s", firstGroup.OauthGroup, err)
	}
	if err := d.Set("oidcGroup", firstGroup.OidcGroup); err != nil {
		return fmt.Errorf("error reading %s oidcGroup: %s", firstGroup.OidcGroup, err)
	}
	if err := d.Set("permissions", flattenGroupPermissions(firstGroup.Permissions)); err != nil {
		return fmt.Errorf("error reading %s permissions: %s", firstGroup.Permissions, err)
	}
	if err := d.Set("role", firstGroup.Role); err != nil {
		return fmt.Errorf("error reading %s role: %s", firstGroup.Role, err)
	}
	if err := d.Set("samlGroup", firstGroup.SamlGroup); err != nil {
		return fmt.Errorf("error reading %s samlGroup: %s", firstGroup.SamlGroup, err)
	}
	if err := d.Set("user", flattenGroupUser(firstGroup.Users)); err != nil {
	}
	return nil
}

func updateGroup(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	parsedGroup, err := parseGroups(d)
	if err != nil {
		return fmt.Errorf("error updating user: %s", err)
	}

	if err := auth.UpdateGroup(*client, parsedGroup); err != nil {
		return fmt.Errorf("error updating user: %s", err)
	}

	return readGroup(d, meta)
}

func deleteGroup(d *schema.ResourceData, meta interface{}) error {
	// TODO: reset to default group
	return nil
}

func parseGroups(d *schema.ResourceData) (auth.Group, error) {
	parsedGroup := auth.Group{}
	
	if d.Get("groupId") != nil {
		parsedGroup.Id = d.Get("groupId").(string)
	}
	if d.Get("ldapGroup") != nil {
		parsedGroup.LdapGroup = d.Get("ldapGroup").(bool)
	}
	if d.Get("groupName") != nil {
		parsedGroup.Name = d.Get("groupName").(string)
	}
	if d.Get("oauthGroup") != nil {
		parsedGroup.OauthGroup = d.Get("oauthGroup").(bool)
	}
	if d.Get("oidcGroup") != nil {
		parsedGroup.OidcGroup = d.Get("oidcGroup").(bool)
	}
	if d.Get("permissions") != nil && len(d.Get("permissions").([]interface{})) > 0 {
		parsedGroup.Permissions = convertGroupPermissions(d.Get("permissions").([]interface{}))
	}
	if d.Get("role") != nil {
		parsedGroup.Role = d.Get("role").(string)
	}
	if d.Get("samlGroup") != nil {
		parsedGroup.SamlGroup = d.Get("samlGroup").(bool)
	}
	if d.Get("user") != nil && len(d.Get("user").([]interface{})) > 0 {
		parsedGroup.Users = convertGroupUser(d.Get("user").([]interface{}))
	}

	return parsedGroup, nil
}

func flattenGroupPermissions(in []auth.GroupPermission) []interface{} {
	ans := make([]interface{}, 0, len(in))
	for _, val := range in {
		m := make(map[string]interface{})
		m["collections"] = val.Collections
		m["project"] = val.Project
		ans = append(ans, m)
	}
	return ans
}

func convertGroupPermissions(in []interface{}) []auth.GroupPermission {
	ans := make([]auth.GroupPermission, 0, len(in))
	for _, val := range in {
		valMap := val.(map[string]interface{})
		m := auth.GroupPermission{}
		m.Collections = parseStringArray(valMap["collections"].([]interface {}))
		m.Project = valMap["project"].(string)
		ans = append(ans, m)
	}
	return ans
}

func flattenGroupUser(in []auth.GroupUser) []interface{} {
	ans := make([]interface{}, 0, len(in))
	for _, val := range in {
		m := make(map[string]interface{})
		m["username"] = val.Username
		ans = append(ans, m)
	}
	return ans
}

func convertGroupUser(in []interface{}) []auth.GroupUser {
	ans := make([]auth.GroupUser, 0, len(in))
	for _, val := range in {
		valMap := val.(map[string]interface{})
		m := auth.GroupUser{}
		m.Username = valMap["username"].(string)
		ans = append(ans, m)
	}
	return ans
}
