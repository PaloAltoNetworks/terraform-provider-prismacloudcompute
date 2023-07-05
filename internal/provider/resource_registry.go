package provider

import (
	"context"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/settings"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/convert"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
 * Note: This resource is not fully implemented. Registries can be added but must be manually updated/deleted
   using the UI. Any deleted registries must also be manually removed from the Terraform state file.
*/

func resourceRegistry() *schema.Resource {
	return &schema.Resource{
		CreateContext: createRegistry,
		ReadContext:   readRegistry,
		UpdateContext: updateRegistry,
		DeleteContext: deleteRegistry,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

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
	}
}

func createRegistry(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)
	parsedRegistry := convert.SchemaToRegistry(d)

	if err := settings.AddRegistrySetting(*client, parsedRegistry); err != nil {
		return diag.Errorf("error creating registry: %s", err)
	}

	d.SetId("registrySettings")
	return readRegistry(ctx, d, meta)
}

func readRegistry(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// client := meta.(*api.Client)

	// var diags diag.Diagnostics

	// retrievedRegistry, err := settings.GetRegistrySettings(*client)
	// if err != nil {
	// 	return diag.Errorf("error reading registry: %s", err)
	// }

	// if err := d.Set("specification", convert.RegistrySpecificationToSchema(retrievedRegistry.Specifications)); err != nil {
	// 	return diag.Errorf("error reading registry: %s", err)
	// }

	// return diags
	return diag.Diagnostics{}
}

func updateRegistry(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// client := meta.(*api.Client)

	// parsedRegistry := convert.SchemaToRegistry(d)

	// retrievedRegistries, err := settings.GetRegistrySettings(*client)
	// if err != nil {
	// 	return diag.Errorf("error reading registry: %s", err)
	// }

	// registries := retrievedRegistries.Specifications

	// if err := settings.UpdateRegistrySettings(*client, parsedRegistry); err != nil {
	// 	return diag.Errorf("error updating registry: %s", err)
	// }

	// return readRegistrySettings(ctx, d, meta)
	return diag.Diagnostics{}
}

func deleteRegistry(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// client := meta.(*api.Client)

	// var diags diag.Diagnostics

	// defaults := settings.RegistrySettings{
	// 	Specifications: make([]settings.RegistrySpecification, 0),
	// }
	// if err := settings.UpdateRegistrySettings(*client, defaults); err != nil {
	// 	return diag.Errorf("error deleting registry: %s", err)
	// }

	// d.SetId("")

	// return diags
	return diag.Diagnostics{}
}
