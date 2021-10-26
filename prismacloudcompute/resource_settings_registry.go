package prismacloudcompute

import (
	"fmt"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/prismacloudcompute/convert"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/pcc"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/settings"
)

func resourceRegistry() *schema.Resource {
	return &schema.Resource{
		Create: createRegistrySettings,
		Read:   readRegistrySettings,
		Update: updateRegistrySettings,
		Delete: deleteRegistrySettings,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"specification": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Registry scanning specifications.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cap": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The maximum number of images to scan from each repository, sorted by most recently modified.",
						},
						"collections": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "The set of Defenders available for scanning.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"credential": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the credential from the credentials store to use for authenticating with the registry.",
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
							Description: "Use temporary tokens provided by Harbor to scan images in projects with the deployment security setting enabled.",
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
							Description: "IBM Cloud namespace.",
						},
						"os": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The base OS of the registry images.",
						},
						"registry": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Registry address.",
						},
						"repository": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Repositories to scan. Pattern matching is supported.",
						},
						"scanners": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Number of Defenders that can be utilized for each scan job.",
						},
						"tag": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Tags to scan. Pattern matching is supported.",
						},
						"type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Registry type.",
						},
						"version_pattern": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Pattern used by the scanner to identify the latest tags without querying the registry for additional metadata. If a pattern specifies both date and version, date takes precedence over version.",
						},
					},
				},
			},
		},
	}
}

func createRegistrySettings(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	parsedRegistry := settings.RegistrySettings{
		Specifications: convert.SchemaToRegistrySpecification(d),
	}

	if err := settings.UpdateRegistrySettings(*client, parsedRegistry); err != nil {
		return fmt.Errorf("failed to create registry: %s", err)
	}

	d.SetId("registrySettings")
	return readRegistrySettings(d, meta)
}

func readRegistrySettings(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	retrievedRegistry, err := settings.GetRegistrySettings(*client)
	if err != nil {
		return fmt.Errorf("failed to read registry: %s", err)
	}

	if err := d.Set("specification", convert.RegistrySpecificationToSchema(retrievedRegistry.Specifications)); err != nil {
		return fmt.Errorf("failed to read registry: %s", err)
	}

	return nil
}

func updateRegistrySettings(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	parsedRegistry := settings.RegistrySettings{
		Specifications: convert.SchemaToRegistrySpecification(d),
	}

	if err := settings.UpdateRegistrySettings(*client, parsedRegistry); err != nil {
		return fmt.Errorf("failed to update registry: %s", err)
	}

	return readRegistrySettings(d, meta)
}

func deleteRegistrySettings(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	defaults := settings.RegistrySettings{
		Specifications: make([]settings.RegistrySpecification, 0),
	}
	if err := settings.UpdateRegistrySettings(*client, defaults); err != nil {
		return fmt.Errorf("failed to delete registry: %s", err)
	}
	d.SetId("")
	return nil
}
