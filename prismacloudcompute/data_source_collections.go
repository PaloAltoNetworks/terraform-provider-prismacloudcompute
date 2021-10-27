package prismacloudcompute

// import (
// 	"bytes"
// 	"encoding/base64"
// 	"log"

// 	"github.com/paloaltonetworks/prisma-cloud-compute-go/collection"
// 	"github.com/paloaltonetworks/prisma-cloud-compute-go/pcc"

// 	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
// )

// func dataSourceCollections() *schema.Resource {
// 	return &schema.Resource{
// 		Read: dataSourceCollectionsRead,

// 		Schema: map[string]*schema.Schema{

// 			// Output.
// 			"account_ids": {
// 				Type:        schema.TypeList,
// 				Optional:    true,
// 				Description: "List of account IDs.",
// 				Elem: &schema.Schema{
// 					Type: schema.TypeString,
// 				},
// 			},
// 			"application_ids": {
// 				Type:        schema.TypeList,
// 				Optional:    true,
// 				Description: "List of application IDs.",
// 				Elem: &schema.Schema{
// 					Type: schema.TypeString,
// 				},
// 			},
// 			"clusters": {
// 				Type:        schema.TypeList,
// 				Optional:    true,
// 				Description: "List of Kubernetes cluster names.",
// 				Elem: &schema.Schema{
// 					Type: schema.TypeString,
// 				},
// 			},
// 			"code_repositories": {
// 				Type:        schema.TypeList,
// 				Optional:    true,
// 				Description: "List of code repositories.",
// 				Elem: &schema.Schema{
// 					Type: schema.TypeString,
// 				},
// 			},
// 			"color": {
// 				Type:        schema.TypeString,
// 				Optional:    true,
// 				Description: "A hex color code for a collection.",
// 			},
// 			"containers": {
// 				Type:        schema.TypeList,
// 				Optional:    true,
// 				Description: "List of containers.",
// 				Elem: &schema.Schema{
// 					Type: schema.TypeString,
// 				},
// 			},
// 			"description": {
// 				Type:        schema.TypeString,
// 				Optional:    true,
// 				Description: "A free-form text description of the collection.",
// 			},
// 			"functions": {
// 				Type:        schema.TypeList,
// 				Optional:    true,
// 				Description: "List of functions.",
// 				Elem: &schema.Schema{
// 					Type: schema.TypeString,
// 				},
// 			},
// 			"hosts": {
// 				Type:        schema.TypeList,
// 				Optional:    true,
// 				Description: "List of hosts.",
// 				Elem: &schema.Schema{
// 					Type: schema.TypeString,
// 				},
// 			},
// 			"images": {
// 				Type:        schema.TypeList,
// 				Optional:    true,
// 				Description: "List of images.",
// 				Elem: &schema.Schema{
// 					Type: schema.TypeString,
// 				},
// 			},
// 			"labels": {
// 				Type:        schema.TypeList,
// 				Optional:    true,
// 				Description: "List of labels.",
// 				Elem: &schema.Schema{
// 					Type: schema.TypeString,
// 				},
// 			},
// 			"name": {
// 				Type:        schema.TypeString,
// 				Required:    true,
// 				Description: "Unique collection name.",
// 			},
// 			"namespaces": {
// 				Type:        schema.TypeList,
// 				Optional:    true,
// 				Description: "List of Kubernetes namespaces.",
// 				Elem: &schema.Schema{
// 					Type: schema.TypeString,
// 				},
// 			},
// 		},
// 	}
// }

// func dataSourceCollectionsRead(d *schema.ResourceData, meta interface{}) error {
// 	var buf bytes.Buffer
// 	client := meta.(*pcc.Client)

// 	items, err := collection.List(*client)
// 	if err != nil {
// 		return err
// 	}

// 	if buf.Len() == 0 {
// 		d.SetId("all")
// 	} else {
// 		d.SetId(base64.StdEncoding.EncodeToString(buf.Bytes()))
// 	}
// 	d.Set("total", len(items))

// 	list := make([]interface{}, 0, len(items))
// 	for _, i := range items {
// 		list = append(list, map[string]interface{}{
// 			"accountIDs":  stringSliceToSet(i.AccountIds),
// 			"appIDs":      stringSliceToSet(i.AppIds),
// 			"clusters":    stringSliceToSet(i.Clusters),
// 			"codeRepos":   stringSliceToSet(i.CodeRepos),
// 			"color":       i.Color,
// 			"containers":  stringSliceToSet(i.Containers),
// 			"description": i.Description,
// 			"functions":   stringSliceToSet(i.Functions),
// 			"hosts":       stringSliceToSet(i.Hosts),
// 			"images":      stringSliceToSet(i.Images),
// 			"labels":      stringSliceToSet(i.Labels),
// 			"name":        i.Name,
// 			"namespaces":  stringSliceToSet(i.Namespaces),
// 		})
// 	}

// 	if err := d.Set("listing", list); err != nil {
// 		log.Printf("[WARN] Error setting 'listing' field for %q: %s", d.Id(), err)
// 	}

// 	return nil
// }
