package convert

import (
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/tag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Converts a tag schema to a tag object for SDK compatibility.
func SchemaToTag(d *schema.ResourceData) tag.Tag {
	ans := tag.Tag{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Color:       d.Get("color").(string),
		Vulns:       []tag.Vuln{},
	}
	if assignments, ok := d.GetOk("assignment"); ok {
		presentAssignments := assignments.([]interface{})
		for _, val := range presentAssignments {
			presentAssignment := val.(map[string]interface{})
			parsedAssignment := tag.Vuln{
				CheckBaseLayer: presentAssignment["check_base_layer"].(bool),
				Comment:        presentAssignment["comment"].(string),
				Id:             presentAssignment["id"].(string),
				PackageName:    presentAssignment["package_name"].(string),
				ResourceType:   presentAssignment["resource_type"].(string),
				Resources:      SchemaToStringSlice(presentAssignment["resources"].([]interface{})),
			}
			if len(parsedAssignment.Resources) == 0 {
				parsedAssignment.Resources = []string{"*"}
			}
			ans.Vulns = append(ans.Vulns, parsedAssignment)
		}
	}
	return ans
}

func TagVulnsToSchema(in []tag.Vuln) []interface{} {
	ans := make([]interface{}, 0, len(in))
	for _, val := range in {
		m := make(map[string]interface{})
		m["resources"] = val.Resources
		m["comment"] = val.Comment
		m["id"] = val.Id
		m["resource_type"] = val.ResourceType
		m["package_name"] = val.PackageName
		ans = append(ans, m)
	}
	return ans
}
