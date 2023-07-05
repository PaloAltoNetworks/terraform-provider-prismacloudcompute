package provider

import (
	"context"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/policy"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/convert"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCustomCompliance() *schema.Resource {
	return &schema.Resource{
		CreateContext: createCustomCompliance,
		ReadContext:   readCustomCompliance,
		UpdateContext: updateCustomCompliance,
		DeleteContext: deleteCustomCompliance,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "ID of the custom Compliance.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"prisma_id": {
				Description: "Prisma Cloud Compute ID of the custom rule.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Free-form text description of the custom Compliance.",
			},
			"title": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Description of the custom compliance",
			},
			"severity": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Severity of this custom compliance",
			},
			"script": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Script of this custom compliance",
			},
		},
	}
}

func createCustomCompliance(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)
	parsedCustomCompliance := convert.SchemaToCustomCompliance(d)
	err := policy.CreateCustomCompliance(*client, parsedCustomCompliance)

	if err != nil {
		return diag.Errorf("error creating custom Compliance '%+v': %s", parsedCustomCompliance, err)
	}

	d.SetId(parsedCustomCompliance.Name)
	return readCustomCompliance(ctx, d, meta)
}

func readCustomCompliance(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)
	retrievedCustomCompliance, err := policy.GetCustomComplianceByName(*client, d.Id())
	if err != nil {
		return diag.Errorf("error reading custom Compliance: %s", err)
	}

	if err := d.Set("name", retrievedCustomCompliance.Name); err != nil {
		return diag.Errorf("error reading custom Compliance: %s", err)
	}
	if err := d.Set("prisma_id", retrievedCustomCompliance.Id); err != nil {
		return diag.Errorf("error reading custom Compliance: %s", err)
	}
	if err := d.Set("title", retrievedCustomCompliance.Title); err != nil {
		return diag.Errorf("error reading custom Compliance: %s", err)
	}
	if err := d.Set("severity", retrievedCustomCompliance.Severity); err != nil {
		return diag.Errorf("error reading custom Compliance: %s", err)
	}
	if err := d.Set("script", retrievedCustomCompliance.Script); err != nil {
		return diag.Errorf("error reading custom Compliance: %s", err)
	}

	return nil
}

func updateCustomCompliance(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)
	parsedCustomCompliance := convert.SchemaToCustomCompliance(d)

	if err := policy.UpdateCustomCompliance(*client, parsedCustomCompliance); err != nil {
		return diag.Errorf("error updating custom Compliance: %s", err)
	}

	return readCustomCompliance(ctx, d, meta)
}

func deleteCustomCompliance(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)

	var diags diag.Diagnostics

	if err := policy.DeleteCustomCompliance(*client, d.Id()); err != nil {
		return diag.Errorf("error deleting custom Compliance '%s': %s", d.Id(), err)
	}

	d.SetId("")

	return diags
}
