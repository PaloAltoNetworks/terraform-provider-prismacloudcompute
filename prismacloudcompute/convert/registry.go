package convert

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/settings"
)

func SchemaToRegistrySpecification(d *schema.ResourceData) []settings.RegistrySpecification {
	parsedRegistrySpecifications := make([]settings.RegistrySpecification, 0)
	if specifications, ok := d.GetOk("specification"); ok {
		presentRegistrySpecifications := specifications.([]interface{})
		for _, val := range presentRegistrySpecifications {
			presentRegistrySpecification := val.(map[string]interface{})
			parsedRegistrySpecifications = append(parsedRegistrySpecifications, settings.RegistrySpecification{
				Cap:                      presentRegistrySpecification["cap"].(int),
				Collections:              SchemaToStringSlice(presentRegistrySpecification["collections"].([]interface{})),
				Credential:               presentRegistrySpecification["credential"].(string),
				ExcludedRepositories:     SchemaToStringSlice(presentRegistrySpecification["excluded_repositories"].([]interface{})),
				ExcludedTags:             SchemaToStringSlice(presentRegistrySpecification["excluded_tags"].([]interface{})),
				HarborDeploymentSecurity: presentRegistrySpecification["harbor_deployment_security"].(bool),
				JfrogRepoTypes:           SchemaToStringSlice(presentRegistrySpecification["jfrog_repo_types"].([]interface{})),
				Namespace:                presentRegistrySpecification["namespace"].(string),
				Os:                       presentRegistrySpecification["os"].(string),
				Tag:                      presentRegistrySpecification["tag"].(string),
				Registry:                 presentRegistrySpecification["registry"].(string),
				Repository:               presentRegistrySpecification["repository"].(string),
				Scanners:                 presentRegistrySpecification["scanners"].(int),
				Version:                  presentRegistrySpecification["type"].(string),
				VersionPattern:           presentRegistrySpecification["version_pattern"].(string),
			})
		}
	}

	return parsedRegistrySpecifications
}

func RegistrySpecificationToSchema(s []settings.RegistrySpecification) []interface{} {
	ans := make([]interface{}, 0, len(s))
	for _, v := range s {
		m := make(map[string]interface{})
		m["cap"] = v.Cap
		m["collections"] = v.Collections
		m["credential"] = v.Credential
		m["excluded_repositories"] = v.ExcludedRepositories
		m["excluded_tags"] = v.ExcludedTags
		m["harbor_deployment_security"] = v.HarborDeploymentSecurity
		m["jfrog_repo_types"] = v.JfrogRepoTypes
		m["namespace"] = v.Namespace
		m["os"] = v.Os
		m["tag"] = v.Tag
		m["registry"] = v.Registry
		m["repository"] = v.Repository
		m["scanners"] = v.Scanners
		m["type"] = v.Version
		m["version_pattern"] = v.VersionPattern
		ans = append(ans, m)
	}
	return ans
}
