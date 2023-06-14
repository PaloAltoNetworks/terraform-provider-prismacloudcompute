package convert

import (
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/settings"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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

func SchemaToRegistry(d *schema.ResourceData) settings.RegistrySpecification {
	parsedRegistrySpecification := settings.RegistrySpecification{}

	if val, ok := d.GetOk("cap"); ok {
		parsedRegistrySpecification.Cap = val.(int)
	}
	if val, ok := d.GetOk("collections"); ok {
		parsedRegistrySpecification.Collections = SchemaToStringSlice(val.([]interface{}))
	}
	if val, ok := d.GetOk("credential"); ok {
		parsedRegistrySpecification.Credential = val.(string)
	}
	if val, ok := d.GetOk("excluded_repositories"); ok {
		parsedRegistrySpecification.ExcludedRepositories = SchemaToStringSlice(val.([]interface{}))
	}
	if val, ok := d.GetOk("excluded_tags"); ok {
		parsedRegistrySpecification.ExcludedTags = SchemaToStringSlice(val.([]interface{}))
	}
	if val, ok := d.GetOk("harbor_deployment_security"); ok {
		parsedRegistrySpecification.HarborDeploymentSecurity = val.(bool)
	}
	if val, ok := d.GetOk("jfrog_repo_types"); ok {
		parsedRegistrySpecification.JfrogRepoTypes = SchemaToStringSlice(val.([]interface{}))
	}
	if val, ok := d.GetOk("namespace"); ok {
		parsedRegistrySpecification.Namespace = val.(string)
	}
	if val, ok := d.GetOk("os"); ok {
		parsedRegistrySpecification.Os = val.(string)
	}
	if val, ok := d.GetOk("tag"); ok {
		parsedRegistrySpecification.Tag = val.(string)
	}
	if val, ok := d.GetOk("registry"); ok {
		parsedRegistrySpecification.Registry = val.(string)
	}
	if val, ok := d.GetOk("repository"); ok {
		parsedRegistrySpecification.Repository = val.(string)
	}
	if val, ok := d.GetOk("scanners"); ok {
		parsedRegistrySpecification.Scanners = val.(int)
	}
	if val, ok := d.GetOk("type"); ok {
		parsedRegistrySpecification.Version = val.(string)
	}
	if val, ok := d.GetOk("version_pattern"); ok {
		parsedRegistrySpecification.VersionPattern = val.(string)
	}

	return parsedRegistrySpecification
}
