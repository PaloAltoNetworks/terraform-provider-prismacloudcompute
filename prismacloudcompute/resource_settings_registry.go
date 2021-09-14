package prismacloudcompute

import (
	"fmt"
	"time"

	pcc "github.com/paloaltonetworks/prisma-cloud-compute-go"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/settings/registry"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceRegistry() *schema.Resource {
	return &schema.Resource{
		Create: createRegistrySettings,
		Read:   readRegistrySettings,
		Update: updateRegistrySettings,
		Delete: deleteRegistrySettings,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"specification": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of specifications.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cap": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Specifies the maximum number of images from each repo to fetch and scan, sorted by most recently modified.",
						},
						"collections": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Specifies the set of Defenders in-scope for working on a scan job.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"credential": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "ID of the credentials in the credentials store to use for authenticating with the registry.",
						},
						"excluded_repositories": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Repositories to exclude from scanning.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"excluded_tags": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Tags to exclude from scanning.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"harbor_deployment_security": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Indicates whether the Prisma Cloud plugin uses temporary tokens provided by Harbor to scan images in projects where Harbor's deployment security setting is enabled.",
						},
						"jfrog_repo_types": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "JFrog Artifactory repository types to scan.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"namespace": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IBM Bluemix namespace.",
						},
						"os": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The registry images base OS type.",
						},
						"registry": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Registry address (e.g., https://gcr.io)..",
						},
						"repository": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Repositories to scan.",
						},
						"scanners": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Number of Defenders that can be utilized for each scan job.",
						},
						"tag": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Tags to scan.",
						},
						"version": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Registry type. Determines the protocol Prisma Cloud uses to communicate with the registry.",
						},
						"version_pattern": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Pattern heuristic for quickly filtering images by tags without having to query all images for modification dates.",
						},
					},
				},
			},
		},
	}
}

func parseRegistry(d *schema.ResourceData) registry.Registry {
	return registry.Registry{
		Specifications: parseRegistrySpecification(d.Get("specification").([]interface{})),
	}
}

func parseRegistrySpecification(specifications []interface{}) []registry.Specification {
	ans := make([]registry.Specification, 0, len(specifications))
	for _, v := range specifications {
		m := v.(map[string]interface{})
		ans = append(ans, registry.Specification{
			Cap:                      m["cap"].(int),
			Collections:              parseStringArray(m["collections"].([]interface{})),
			CredentialId:             m["credential"].(string),
			ExcludedRepositories:     parseStringArray(m["excluded_repositories"].([]interface{})),
			ExcludedTags:             parseStringArray(m["excluded_tags"].([]interface{})),
			HarborDeploymentSecurity: m["harbor_deployment_security"].(bool),
			JfrogRepoTypes:           parseStringArray(m["jfrog_repo_types"].([]interface{})),
			Namespace:                m["namespace"].(string),
			Os:                       m["os"].(string),
			Tag:                      m["tag"].(string),
			Registry:                 m["registry"].(string),
			Repository:               m["repository"].(string),
			Scanners:                 m["scanners"].(int),
			Version:                  m["version"].(string),
			VersionPattern:           m["version_pattern"].(string),
		})
	}
	return ans
}

func flattenRegistrySpecification(s []registry.Specification) []interface{} {
	ans := make([]interface{}, 0, len(s))
	for _, v := range s {
		m := make(map[string]interface{})
		m["cap"] = v.Cap
		m["collections"] = v.Collections
		m["credential"] = v.CredentialId
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
		m["version"] = v.Version
		m["version_pattern"] = v.VersionPattern
		ans = append(ans, m)
	}
	return ans
}

func createRegistrySettings(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	obj := parseRegistry(d)

	if err := registry.Update(*client, obj); err != nil {
		return fmt.Errorf("failed to create registry: %s", err)
	}

	d.SetId("registrySettings")
	return readRegistrySettings(d, meta)
}

func readRegistrySettings(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)

	obj, err := registry.Get(*client)
	if err != nil {
		return fmt.Errorf("failed to read registry: %s", err)
	}

	if err := d.Set("specification", flattenRegistrySpecification(obj.Specifications)); err != nil {
		return fmt.Errorf("failed setting 'specification' for %s: %s", d.Id(), err)
	}

	return nil
}

func updateRegistrySettings(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	obj := parseRegistry(d)

	if err := registry.Update(*client, obj); err != nil {
		return fmt.Errorf("failed to update registry: %s", err)
	}

	return readRegistrySettings(d, meta)
}

func deleteRegistrySettings(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	obj := registry.Registry{
		Specifications: make([]registry.Specification, 0),
	}
	if err := registry.Update(*client, obj); err != nil {
		return fmt.Errorf("failed to delete registry: %s", err)
	}
	d.SetId("")
	return nil
}
