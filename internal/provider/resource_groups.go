package provider

import (
	"context"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/auth"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/convert"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGroups() *schema.Resource {
	return &schema.Resource{
		CreateContext: createGroup,
		ReadContext:   readGroup,
		UpdateContext: updateGroup,
		DeleteContext: deleteGroup,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The ID of the group.",
				Type:        schema.TypeString,
				Computed:    true,
			},
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
			"name": {
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

func createGroup(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)
	parsedGroup, err := convert.SchemaToGroup(d)
	if err != nil {
		return diag.Errorf("error creating group '%+v': %s", parsedGroup, err)
	}

	if err := auth.CreateGroup(*client, parsedGroup); err != nil {
		return diag.Errorf("error creating group '%+v': %s", parsedGroup, err)
	}

	d.SetId(parsedGroup.Name)
	return readGroup(ctx, d, meta)
}

func readGroup(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)

	var diags diag.Diagnostics

	retrievedGroup, err := auth.GetGroup(*client, d.Id())
	if err != nil {
		return diag.Errorf("error reading group: %s", err)
	}

	d.Set("group_id", retrievedGroup.Id)
	d.Set("ldap_group", retrievedGroup.LdapGroup)
	d.Set("name", retrievedGroup.Name)
	d.Set("oauth_group", retrievedGroup.OauthGroup)
	d.Set("oidc_group", retrievedGroup.OidcGroup)
	if err := d.Set("permissions", convert.GroupPermissionsToSchema(retrievedGroup.Permissions)); err != nil {
		return diag.Errorf("error reading group: %s", err)
	}
	d.Set("role", retrievedGroup.Role)
	d.Set("saml_group", retrievedGroup.SamlGroup)
	if err := d.Set("users", convert.GroupUsersToSchema(retrievedGroup.Users)); err != nil {
		return diag.Errorf("error reading group: %s", err)
	}
	return diags
}

func updateGroup(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)
	parsedGroup, err := convert.SchemaToGroup(d)
	if err != nil {
		return diag.Errorf("error updating group: %s", err)
	}

	if err := auth.UpdateGroup(*client, parsedGroup); err != nil {
		return diag.Errorf("error updating group: %s", err)
	}

	return readGroup(ctx, d, meta)
}

func deleteGroup(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)

	var diags diag.Diagnostics

	if err := auth.DeleteGroup(*client, d.Id()); err != nil {
		return diag.Errorf("error deleting group '%s': %s", d.Id(), err)
	}

	d.SetId("")
	return diags
}
