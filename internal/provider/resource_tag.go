package provider

import (
	"context"
	"fmt"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/tag"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/convert"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTag() *schema.Resource {
	return &schema.Resource{
		CreateContext: createTag,
		ReadContext:   readTag,
		UpdateContext: updateTag,
		DeleteContext: deleteTag,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The ID of the tag.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"color": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A hex color code for the tag to display in the Console.",
				Default:     "#A020F0",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A free-form text description of the tag.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A unique tag name.",
			},
			"assignment": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specify how vulnerabilities are tagged, based on CVE ID, package, and resources.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"check_base_layer": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether or not to check the base layer.",
						},
						"comment": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Free-form text field.",
						},
						"id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Common Vulnerability and Exposures (CVE) ID.",
						},
						"package_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Source or binary package name where the vulnerability is found.",
							Default:     "*",
						},
						"resource_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "*",
							Description: "Specifies the resource type for tagging where the vulnerability is found.",
							ValidateDiagFunc: func(v interface{}, p cty.Path) diag.Diagnostics {
								validValues := []string{"image", "host", "function", "codeRepo"}
								value := v.(string)

								for _, item := range validValues {
									if item == value {
										return diag.Diagnostics{}
									}
								}
								return diag.Diagnostics{
									diag.Diagnostic{
										Severity: diag.Error,
										Summary:  "Invalid resource_type",
										Detail:   fmt.Sprintf("Valid values are %v", validValues),
									},
								}
							},
						},
						"resources": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Resource names separated by a comma or use the wildcard * to apply the tag to all the resources where the vulnerability is found.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func createTag(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)
	parsedTag := convert.SchemaToTag(d)
	if err := tag.CreateTag(*client, parsedTag); err != nil {
		return diag.Errorf("error creating tag'%+v': %s", parsedTag, err)
	}

	d.SetId(parsedTag.Name)

	return readTag(ctx, d, meta)
}

func readTag(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)

	var diags diag.Diagnostics

	retrievedTag, err := tag.GetTag(*client, d.Id())
	if err != nil {
		return diag.Errorf("error reading tag: %s", err)
	}

	if err := d.Set("assignment", convert.TagVulnsToSchema(retrievedTag.Vulns)); err != nil {
		return diag.Errorf("error reading tag: %s, %v", err, retrievedTag.Vulns)
	}
	d.Set("color", retrievedTag.Color)
	d.Set("description", retrievedTag.Description)
	d.Set("name", retrievedTag.Name)

	return diags
}

func updateTag(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)

	parsedTag := convert.SchemaToTag(d)

	if err := tag.UpdateTag(*client, parsedTag); err != nil {
		return diag.Errorf("error updating tag: %s", err)
	}

	return readTag(ctx, d, meta)
}

func deleteTag(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)

	var diags diag.Diagnostics

	if err := tag.DeleteTag(*client, d.Id()); err != nil {
		return diag.Errorf("error updating tag '%s': %s", d.Id(), err)
	}

	d.SetId("")

	return diags
}
