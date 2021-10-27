package prismacloudcompute

import (
	"fmt"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/prismacloudcompute/convert"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/auth"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/pcc"
)

func resourceGroups() *schema.Resource {
	return &schema.Resource{
		Create: createGroup,
		Read:   readGroup,
		Update: updateGroup,
		Delete: deleteGroup,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Group ID.",
			},
			"ldap_group": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether or not the group is an LDAP group.",
			},
			"group_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Group name.",
			},
			"oauth_group": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether or not the group is an OAuth group.",
			},
			"oidc_group": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether or not the group is an OpenID Connect group.",
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
			"saml_group": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether or not the group is a SAML group.",
			},
			"users": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Users in the group.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func createGroup(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	parsedGroup, err := convert.SchemaToGroup(d)
	if err != nil {
		return fmt.Errorf("error creating group '%+v': %s", parsedGroup, err)
	}

	if err := auth.CreateGroup(*client, parsedGroup); err != nil {
		return fmt.Errorf("error creating group '%+v': %s", parsedGroup, err)
	}

	d.SetId(parsedGroup.Name)
	return readGroup(d, meta)
}

func readGroup(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	retrievedGroup, err := auth.GetGroup(*client, d.Id())
	if err != nil {
		return fmt.Errorf("error reading group: %s", err)
	}

	d.Set("groupid", retrievedGroup.Id)
	d.Set("ldapgroup", retrievedGroup.LdapGroup)
	d.Set("groupname", retrievedGroup.Name)
	d.Set("oauthgroup", retrievedGroup.OauthGroup)
	d.Set("oidcgroup", retrievedGroup.OidcGroup)
	if err := d.Set("permissions", convert.GroupPermissionsToSchema(retrievedGroup.Permissions)); err != nil {
		return fmt.Errorf("error reading group: %s", err)
	}
	d.Set("role", retrievedGroup.Role)
	d.Set("samlgroup", retrievedGroup.SamlGroup)
	if err := d.Set("users", convert.GroupUsersToSchema(retrievedGroup.Users)); err != nil {
		return fmt.Errorf("error reading group: %s", err)
	}
	return nil
}

func updateGroup(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	parsedGroup, err := convert.SchemaToGroup(d)
	if err != nil {
		return fmt.Errorf("error updating group: %s", err)
	}

	if err := auth.UpdateGroup(*client, parsedGroup); err != nil {
		return fmt.Errorf("error updating group: %s", err)
	}

	return readGroup(d, meta)
}

func deleteGroup(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	if err := auth.DeleteGroup(*client, d.Id()); err != nil {
		return fmt.Errorf("error deleting group '%s': %s", d.Id(), err)
	}

	d.SetId("")
	return nil
}
