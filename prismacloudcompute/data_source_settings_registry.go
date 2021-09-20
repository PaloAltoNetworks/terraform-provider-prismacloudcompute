package prismacloudcompute

import (
	"log"

	"github.com/paloaltonetworks/prisma-cloud-compute-go/pcc"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/settings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRegistry() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRegistryRead,

		Schema: map[string]*schema.Schema{

			// Output.
			"harborscannerurlsuffix": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "URL suffix for the harbor scanner.",
			},
			"webhookurlsuffix": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "URL suffix for the webhook.",
			},
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
						"credential_id": {
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
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Tags to exclude from scanning.",
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
							Description: "IBM Bluemix namespace https://console.bluemix.net/docs/services/Regis.",
						},
						"os": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "RegistryOSType specifies the registry images base OS type.",
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

func dataSourceRegistryRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)

	i, err := settings.GetRegistrySettings(*client)

	if err != nil {
		return err
	}

	list := make([]interface{}, 0, 1)
	list = append(list, map[string]interface{}{
		"specification": i.Specifications,
	})

	if err := d.Set("listing", list); err != nil {
		log.Printf("[WARN] Error setting 'listing' field for %q: %s", d.Id(), err)
	}

	return nil
}
