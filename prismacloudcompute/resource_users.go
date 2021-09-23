package prismacloudcompute

import (
	"fmt"
	"time"

	"github.com/paloaltonetworks/prisma-cloud-compute-go/pcc"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/auth"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUsers() *schema.Resource {
	return &schema.Resource{
		Create: createUsers,
		Read:   readUsers,
		Update: updateUsers,
		Delete: deleteUsers,

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

func createUsers(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	parsedUser, err := parseUsers(d)
	if err != nil {
		return fmt.Errorf("error creating user: %s", err)
	}

	if err := auth.UpdateUsers(*client, *parsedUser); err != nil {
		return fmt.Errorf("error creating user: %s", err)
	}

	d.SetId(policyTypeUsers)
	return readUsers(d, meta)
}

func readUsers(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	retrievedUser, err := auth.GetUsers(*client)
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

func updateUsers(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	parsedUser, err := parseUsers(d)
	if err != nil {
		return fmt.Errorf("error updating user: %s", err)
	}

	if err := auth.UpdateUsers(*client, *parsedUser); err != nil {
		return fmt.Errorf("error updating user: %s", err)
	}

	return readUsers(d, meta)
}

func deleteUsers(d *schema.ResourceData, meta interface{}) error {
	// TODO: reset to default user
	return nil
}

func parseUsers(d *schema.ResourceData) (*user.User, error) {
	parsedUser := user.User{}
	
	if d.Get("authType") != nil {
		parsedUser.AuthType = d.Get("authType").(string)
	}
	if d.Get("password") != nil {
		parsedUser.Password = d.Get("password").(string)
	}
	if d.Get("permissions") != nil && len(d.Get("permissions").([]interface{})) > 0 {
		parsedUser.Permissions = flattenUserPermissions(d.Get("permissions").([]interface{}))
	} else {
		parsedUser.Permissions = {}
	}
	if d.Get("role") != nil {
		parsedUser.Role = d.Get("role").(string)
	}
	if d.Get("username") != nil {
		parsedUser.Username = d.Get("username").(string)
	}

	return parsedUser
}

func flattenUserPermissions(in []user.UserPermission) []interface{} {
	ans := make([]interface{}, 0, len(in))
	for _, val := range in {
		m := make(map[string]interface{})
		m["collections"] = flattenCollections(val.Collections)
		m["project"] = val.Project
		ans = append(ans, m)
	}
	return ans
}