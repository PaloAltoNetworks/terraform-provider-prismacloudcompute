package convert

import (
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/policy"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Converts a custom rule schema to a custom rule object for SDK compatibility.
func SchemaToCustomCompliance(d *schema.ResourceData) policy.CustomCompliance {
	parsedCompliance := policy.CustomCompliance{}

	if val, ok := d.GetOk("name"); ok {
		parsedCompliance.Name = val.(string)
	}
	if val, ok := d.GetOk("title"); ok {
		parsedCompliance.Title = val.(string)
	}
	if val, ok := d.GetOk("severity"); ok {
		parsedCompliance.Severity = val.(string)
	}
	if val, ok := d.GetOk("script"); ok {
		parsedCompliance.Script = val.(string)
	}
	if val, ok := d.GetOk("prisma_id"); ok {
		parsedCompliance.Id = val.(int)
	}
	return parsedCompliance
}
