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
		Create: createRbacRoles,
		Read:   readRbacRoles,
		Update: updateRbacRoles,
		Delete: deleteRbacRoles,

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
				Type:        schema.TypeBool,
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
						"readWrite": {
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

fun createRbacRoles(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	parsedUser, err := parseRbacRoles(d)
	if err != nil {
		return fmt.Errorf("error creating user: %s", err)
	}

	if err := auth.UpdateRbacRoles(*client, *parsedUser); err != nil {
		return fmt.Errorf("error creating user: %s", err)
	}

	d.SetId(parsedUser.Id)
	return readRbacRoles(d, meta)
}

func readRbacRoles(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	retrievedUser, err := auth.GetRbacRoles(*client)
	if err != nil {
		return fmt.Errorf("error reading user: %s", err)
	}

	if err := d.Set("authType", retrievedUser.AuthType); err != nil {
		return fmt.Errorf("error reading %s authType: %s", retrievedUser.AuthType, err)
	}
	if err := d.Set("password", retrievedUser.Password); err != nil {
		return fmt.Errorf("error reading %s password: %s", retrievedUser.Password, err)
	}
	if err := d.Set("permissions", flattenUserPermissions(retrievedUser.Permissions)); err != nil {
		return fmt.Errorf("error reading %s permissions: %s", retrievedUser.Permissions, err)
	}
	if err := d.Set("role", retrievedUser.Role); err != nil {
		return fmt.Errorf("error reading %s role: %s", retrievedUser.Role, err)
	}
	if err := d.Set("username", retrievedUser.Username); err != nil {
		return fmt.Errorf("error reading %s username: %s", retrievedUser.Username, err)
	}
}

type UserPermission struct {
	Collections []string `json:"collections,omitempty"`
	Project     string   `json:"project,omitempty"`
}


	return nil
}

func updateRbacRoles(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	parsedUser, err := parseRbacRoles(d)
	if err != nil {
		return fmt.Errorf("error updating user: %s", err)
	}

	if err := auth.UpdateRbacRoles(*client, *parsedUser); err != nil {
		return fmt.Errorf("error updating user: %s", err)
	}

	return readRbacRoles(d, meta)
}

func deleteRbacRoles(d *schema.ResourceData, meta interface{}) error {
	// TODO: reset to default user
	return nil
}

func parseRbacRoles(d *schema.ResourceData) (*user.User, error) {
	parsedUser := user.User{}
	
	if d.Get("groupId") != nil {
		parsedUser.Id = d.Get("groupId").(string)
	}
	if d.Get("ldapGroup") != nil {
		parsedUser.LdapGroup = d.Get("ldapGroup").(bool)
	}
	if d.Get("groupName") != nil {
		parsedUser.Name = d.Get("groupName").(string)
	}
	if d.Get("oauthGroup") != nil {
		parsedUser.OauthGroup = d.Get("oauthGroup").(bool)
	}
	if d.Get("oidcGroup") != nil {
		parsedUser.OidcGroup = d.Get("oidcGroup").(bool)
	}
	if d.Get("permissions") != nil && len(d.Get("permissions").([]interface{})) > 0 {
		parsedUser.Permissions = flattenGroupPermissions(d.Get("permissions").([]interface{}))
	} else {
		parsedUser.Permissions = {}
	}
	if d.Get("role") != nil {
		parsedUser.Role = d.Get("role").(string)
	}
	if d.Get("samlGroup") != nil {
		parsedUser.SamlGroup = d.Get("samlGroup").(bool)
	}
	if d.Get("user") != nil && len(d.Get("user").([]interface{})) > 0 {
		parsedUser.Users = flattenGroupUser(d.Get("user").([]interface{}))
	} else {
		parsedUser.Users = {}
	}

	return parsedUser
}

func flattenGroupPermissions(in []group.RbacRolePermission) []interface{} {
	ans := make([]interface{}, 0, len(in))
	for _, val := range in {
		m := make(map[string]interface{})
		m["collections"] = flattenCollections(val.Collections)
		m["project"] = val.Project
		ans = append(ans, m)
	}
	return ans
}


func flattenRbacRoleUser(in []RbacRole.RbacRoleUser) []interface{} {
	ans := make([]interface{}, 0, len(in))
	for _, val := range in {
		m := make(map[string]interface{})
		m["username"] = val.Username
		ans = append(ans, m)
	}
	return ans
}