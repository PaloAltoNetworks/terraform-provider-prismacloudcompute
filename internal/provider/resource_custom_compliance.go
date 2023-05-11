package provider

import (
	"fmt"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/policy"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/convert"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCustomCompliance() *schema.Resource {
	return &schema.Resource{
		Create: createCustomCompliance,
		Read:   readCustomCompliance,
		Update: updateCustomCompliance,
		Delete: deleteCustomCompliance,

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

func createCustomCompliance(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*api.Client)
	parsedCustomCompliance := convert.SchemaToCustomCompliance(d)
	id, err := policy.CreateCustomCompliance(*client, parsedCustomCompliance)

	if err != nil {
		return fmt.Errorf("error creating custom Compliance '%+v': %s", parsedCustomCompliance, err)
	}
	if err := d.Set("prisma_id", id); err != nil {
		return fmt.Errorf("error creating custom Compliance '%+v': %s", parsedCustomCompliance, err)
	}
	d.SetId(parsedCustomCompliance.Name)
	return readCustomCompliance(d, meta)
}

func readCustomCompliance(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*api.Client)
	retrievedCustomCompliance, err := policy.GetCustomComplianceById(*client, d.Get("prisma_id").(int))
	if err != nil {
		return fmt.Errorf("error reading custom Compliance: %s", err)
	}

	if err := d.Set("name", retrievedCustomCompliance.Name); err != nil {
		return fmt.Errorf("error reading custom Compliance: %s", err)
	}
	if err := d.Set("prisma_id", retrievedCustomCompliance.Id); err != nil {
		return fmt.Errorf("error reading custom rule: %s", err)
	}
	if err := d.Set("title", retrievedCustomCompliance.Title); err != nil {
		return fmt.Errorf("error reading custom Compliance: %s", err)
	}
	if err := d.Set("severity", retrievedCustomCompliance.Severity); err != nil {
		return fmt.Errorf("error reading custom Compliance: %s", err)
	}
	if err := d.Set("script", retrievedCustomCompliance.Script); err != nil {
		return fmt.Errorf("error reading custom Compliance: %s", err)
	}

	return nil
}

func updateCustomCompliance(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*api.Client)
	parsedCustomCompliance := convert.SchemaToCustomCompliance(d)

	if _, err := policy.UpdateCustomCompliance(*client, parsedCustomCompliance); err != nil {
		return fmt.Errorf("error updating custom Compliance: %s", err)
	}

	return readCustomCompliance(d, meta)
}

func deleteCustomCompliance(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*api.Client)
	if err := policy.DeleteCustomCompliance(*client, d.Get("prisma_id").(int)); err != nil {
		return fmt.Errorf("error updating custom Compliance '%s': %s", d.Id(), err)
	}

	d.SetId("")
	return nil
}
