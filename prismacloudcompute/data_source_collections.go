package prismacloudcompute

import (
	"bytes"
	"encoding/base64"
	"log"

	pc "github.com/paloaltonetworks/prisma-cloud-compute-go"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/collection"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceCollections() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCollectionsRead,

		Schema: map[string]*schema.Schema{

			// Output.
			"accountids": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "Account IDs",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"appids": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "App IDs",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"clusters": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "Clusters",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"coderepos": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "Code repositories",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"color": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Color",
			},
			"containers": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "Containers",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Description",
			},
			"functions": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "Serverless functions",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"hosts": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "Hosts",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"images": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "Images",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"labels": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "Labels",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"modified": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Last modified date",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name",
			},
			"namespaces": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "Namespaces",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"owner": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Owner",
			},
			"prisma": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Prisma",
			},
			"system": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "System",
			},
		},
	}
}

func dataSourceCollectionsRead(d *schema.ResourceData, meta interface{}) error {
	var buf bytes.Buffer
	client := meta.(*pc.Client)

	items, err := collection.List(client)
	if err != nil {
		return err
	}

	if buf.Len() == 0 {
		d.SetId("all")
	} else {
		d.SetId(base64.StdEncoding.EncodeToString(buf.Bytes()))
	}
	d.Set("total", len(items))

	list := make([]interface{}, 0, len(items))
	for _, i := range items {
		list = append(list, map[string]interface{}{
			"accountIDs":  StringSliceToSet(i.AccountIDs),
			"appIDs":      StringSliceToSet(i.AppIDs),
			"clusters":    StringSliceToSet(i.Clusters),
			"codeRepos":   StringSliceToSet(i.CodeRepos),
			"color":       i.Color,
			"containers":  StringSliceToSet(i.Containers),
			"description": i.Description,
			"functions":   StringSliceToSet(i.Functions),
			"hosts":       StringSliceToSet(i.Hosts),
			"images":      StringSliceToSet(i.Images),
			"labels":      StringSliceToSet(i.Labels),
			"modified":    i.Modified,
			"name":        i.Name,
			"namespaces":  StringSliceToSet(i.Namespaces),
			"owner":       i.Owner,
			"prisma":      i.Prisma,
			"system":      i.System,
		})
	}

	if err := d.Set("listing", list); err != nil {
		log.Printf("[WARN] Error setting 'listing' field for %q: %s", d.Id(), err)
	}

	return nil
}
