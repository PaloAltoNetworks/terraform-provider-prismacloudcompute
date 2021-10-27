package prismacloudcompute

import (
	"fmt"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/prismacloudcompute/convert"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/collection"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/pcc"
)

func resourceCollection() *schema.Resource {
	return &schema.Resource{
		Create: createCollection,
		Read:   readCollection,
		Update: updateCollection,
		Delete: deleteCollection,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The ID of the collection.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"account_ids": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "Targeted cloud account IDs.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"application_ids": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Targeted application IDs (for app-embedded).",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"clusters": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Targeted cluster names.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"code_repositories": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Targeted code repositories.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"color": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A hex color code for the collection to display in the Console.",
				Default:     "#A020F0",
			},
			"containers": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Targeted containers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A free-form text description of the collection.",
			},
			"functions": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Targeted functions.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"hosts": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Targeted hosts.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"images": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Targeted images.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"labels": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Targeted labels.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A unique collection name.",
			},
			"namespaces": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Targeted cluster namespaces.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func createCollection(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	parsedCollection := convert.SchemaToCollection(d)
	if err := collection.CreateCollection(*client, parsedCollection); err != nil {
		return fmt.Errorf("error creating collection '%+v': %s", parsedCollection, err)
	}

	d.SetId(parsedCollection.Name)
	return readCollection(d, meta)
}

func readCollection(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	retrievedCollection, err := collection.GetCollection(*client, d.Id())
	if err != nil {
		return fmt.Errorf("error reading collection: %s", err)
	}

	if err := d.Set("account_ids", retrievedCollection.AccountIds); err != nil {
		return fmt.Errorf("error reading collection: %s", err)
	}
	if err := d.Set("application_ids", retrievedCollection.AppIds); err != nil {
		return fmt.Errorf("error reading collection: %s", err)
	}
	if err := d.Set("clusters", retrievedCollection.Clusters); err != nil {
		return fmt.Errorf("error reading collection: %s", err)
	}
	if err := d.Set("code_repositories", retrievedCollection.CodeRepos); err != nil {
		return fmt.Errorf("error reading collection: %s", err)
	}
	d.Set("color", retrievedCollection.Color)
	if err := d.Set("containers", retrievedCollection.Containers); err != nil {
		return fmt.Errorf("error reading collection: %s", err)
	}
	d.Set("description", retrievedCollection.Description)
	if err := d.Set("functions", retrievedCollection.Functions); err != nil {
		return fmt.Errorf("error reading collection: %s", err)
	}
	if err := d.Set("hosts", retrievedCollection.Hosts); err != nil {
		return fmt.Errorf("error reading collection: %s", err)
	}
	if err := d.Set("images", retrievedCollection.Images); err != nil {
		return fmt.Errorf("error reading collection: %s", err)
	}
	if err := d.Set("labels", retrievedCollection.Labels); err != nil {
		return fmt.Errorf("error reading collection: %s", err)
	}
	d.Set("name", retrievedCollection.Name)
	if err := d.Set("namespaces", retrievedCollection.Namespaces); err != nil {
		return fmt.Errorf("error reading collection: %s", err)
	}

	return nil
}

func updateCollection(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	parsedCollection := convert.SchemaToCollection(d)

	if err := collection.UpdateCollection(*client, parsedCollection); err != nil {
		return fmt.Errorf("error updating collection: %s", err)
	}

	return readCollection(d, meta)
}

func deleteCollection(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	if err := collection.DeleteCollection(*client, d.Id()); err != nil {
		return fmt.Errorf("error updating collection '%s': %s", d.Id(), err)
	}

	d.SetId("")
	return nil
}
