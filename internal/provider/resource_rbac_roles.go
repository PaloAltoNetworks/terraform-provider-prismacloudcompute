package provider

import (
	"context"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/auth"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/convert"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRbacRoles() *schema.Resource {
	return &schema.Resource{
		CreateContext: createRbacRole,
		ReadContext:   readRbacRole,
		UpdateContext: updateRbacRole,
		DeleteContext: deleteRbacRole,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The ID of the role.",
				Type:        schema.TypeString,
				Computed:    true,
			},
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
			"permission": {
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
						"read_write": {
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

func createRbacRole(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)
	parsedRole, err := convert.SchemaToRbacRole(d)
	if err != nil {
		return diag.Errorf("error creating role '%+v': %s", parsedRole, err)
	}

	if err := auth.CreateRole(*client, parsedRole); err != nil {
		return diag.Errorf("error creating role '%+v': %s", parsedRole, err)
	}

	d.SetId(parsedRole.Name)
	return readRbacRole(ctx, d, meta)
}

func readRbacRole(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)

	var diags diag.Diagnostics

	retrievedRole, err := auth.GetRole(*client, d.Id())
	if err != nil {
		return diag.Errorf("error reading role: %s", err)
	}

	d.Set("description", retrievedRole.Description)
	d.Set("name", retrievedRole.Name)
	if err := d.Set("permission", convert.RbacRolePermissionsToSchema(retrievedRole.Permissions)); err != nil {
		return diag.Errorf("error reading role: %s", err)
	}

	return diags
}

func updateRbacRole(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)
	parsedRole, err := convert.SchemaToRbacRole(d)
	if err != nil {
		return diag.Errorf("error updating role: %s", err)
	}

	if err := auth.UpdateRole(*client, parsedRole); err != nil {
		return diag.Errorf("error updating role: %s", err)
	}

	return readRbacRole(ctx, d, meta)
}

func deleteRbacRole(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)

	var diags diag.Diagnostics

	if err := auth.DeleteRole(*client, d.Id()); err != nil {
		return diag.Errorf("error deleting role '%s': %s", d.Id(), err)
	}

	d.SetId("")

	return diags
}
