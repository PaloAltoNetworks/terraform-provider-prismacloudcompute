package prismacloudcompute

import (
	"fmt"
	"time"
	"strings"

	"github.com/paloaltonetworks/prisma-cloud-compute-go/pcc"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/auth"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUsers() *schema.Resource {
	return &schema.Resource{
		Create: createUser,
		Read:   readUser,
		Update: updateUser,
		Delete: deleteUser,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

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
				Description: "Password for authentication.",
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
				Description: "User role.",
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Username for authentication.",
			},			
		},
	}
}

func createUser(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	parsedUser, err := parseUser(d)
	if err != nil {
		return fmt.Errorf("error parsing user: %s", err)
	}

	if err := auth.UpdateUser(*client, parsedUser); err != nil {
		return fmt.Errorf("error creating user: %s %s", err, parsedUser.Username)
	}

	d.SetId(parsedUser.Username)
	return readUser(d, meta)
}

func readUser(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	userList, err := auth.GetUsers(*client)
	retrievedUser := userList[0]
	if err != nil {
		return fmt.Errorf("error reading user: %s", err)
	}

	if err := d.Set("authtype", retrievedUser.AuthType); err != nil {
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

	return nil
}

func updateUser(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	parsedUser, err := parseUser(d)
	if err != nil {
		return fmt.Errorf("error parsing user for update: %s", err)
	}

	if err := auth.UpdateUser(*client, parsedUser); err != nil {
		return fmt.Errorf("error updating user: %s", err)
	}

	return readUser(d, meta)
}

func deleteUser(d *schema.ResourceData, meta interface{}) error {
	// TODO: reset to default user
	client := meta.(*pcc.Client)
	id := d.Id()

	if err := auth.DeleteUser(*client, id); err != nil {
		return fmt.Errorf("failed to update credential: %s", err)
	}

	d.SetId("")
	return nil
}

func parseUser(d *schema.ResourceData) (auth.User, error) {
	parsedUser := auth.User{}
	
	if d.Get("authtype") != nil {
		parsedUser.AuthType = d.Get("authtype").(string)
	}
	if d.Get("password") != nil {
		parsedUser.Password = d.Get("password").(string)
	}
	if d.Get("permissions") != nil && len(d.Get("permissions").([]interface{})) > 0 {
		parsedUser.Permissions = convertUserPermissions(d.Get("permissions").([]interface{}))
	}
	if d.Get("role") != nil {
		parsedUser.Role = d.Get("role").(string)
	}
	if d.Get("username") != nil {
		parsedUser.Username = d.Get("username").(string)
	}

	return parsedUser, nil
}

func flattenUserPermissions(in []auth.UserPermission) []interface{} {
	ans := make([]interface{}, 0, len(in))
	for _, val := range in {
		m := make(map[string]interface{})
		m["collections"] = strings.Join(val.Collections, ",")
		m["project"] = val.Project
		ans = append(ans, m)
	}
	return ans
}

func convertUserPermissions(in []interface{}) []auth.UserPermission {
	ans := make([]auth.UserPermission, 0, len(in))
	for _, val := range in {
		m := auth.UserPermission{}
		valMap := val.(map[string]interface{})
		if valMap["collections"] != nil {
			m.Collections = parseStringArray(valMap["collections"].([]interface {}))
		}
		if valMap["project"] != nil {
			m.Project = valMap["project"].(string)
		}
		ans = append(ans, m)
	}
	return ans
}
