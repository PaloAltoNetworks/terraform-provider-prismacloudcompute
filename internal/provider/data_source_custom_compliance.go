package provider

import (
	"fmt"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/policy"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCustomCompliance() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to retrieve ID of a custom compliance.",
		Read:        dataSourceCustomComplianceRead,

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
				Description: "Name of the custom compliance.",
			},
			"title": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the custom compliance.",
			},
			"severity": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Severity of the custom compliance",
			},
			"script": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Script of the custom compliance",
			},
		},
	}
}

func dataSourceCustomComplianceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*api.Client)

	if name := d.Get("name").(string); name != "" {
		retrievedCustomCompliance, err := policy.GetCustomComplianceByName(*client, name)
		if err != nil {
			return fmt.Errorf("error reading custom Compliance: %s", err)
		}
		if err := d.Set("name", retrievedCustomCompliance.Name); err != nil {
			return fmt.Errorf("error reading custom Compliance: %s", err)
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
		if err := d.Set("prisma_id", retrievedCustomCompliance.Id); err != nil {
			return fmt.Errorf("error reading custom Compliance: %s", err)
		}
		d.SetId(retrievedCustomCompliance.Name)

		return nil
	}

	return fmt.Errorf("missing name parameter")
}
