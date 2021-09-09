package prismacloudcompute

import (
	"log"
	"time"

	pcc "github.com/paloaltonetworks/prisma-cloud-compute-go"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/collections"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceCollection() *schema.Resource {
	return &schema.Resource{
		Create: createCollection,
		Read:   readCollection,
		Update: updateCollection,
		Delete: deleteCollection,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"accountids": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of account IDs.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"appids": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of application IDs.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"clusters": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of cluster names.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"coderepos": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of code repositories.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"color": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A hex color code for the collection",
			},
			"containers": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of containers.",
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
				Description: "List of functions.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"hosts": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of hosts.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"images": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of images.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"labels": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of labels.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique collection name.",
			},
			"namespaces": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of cluster namespaces.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func stringSliceToSet(list []string) *schema.Set {
	items := make([]interface{}, len(list))
	for i := range list {
		items[i] = list[i]
	}

	return schema.NewSet(schema.HashString, items)
}

func parseCollection(d *schema.ResourceData, id string) collections.Collection {
	ans := collections.Collection{
		Name: d.Get("name").(string),
	}
	if d.Get("accountids") != nil && len(d.Get("accountids").([]interface{})) > 0 {
		ans.AccountIDs = parseStringArray(d.Get("accountids").([]interface{}))
	} else {
		ans.AccountIDs = []string{"*"}
	}
	if d.Get("appids") != nil && len(d.Get("appids").([]interface{})) > 0 {
		ans.AppIDs = parseStringArray(d.Get("appids").([]interface{}))
	} else {
		ans.AppIDs = []string{"*"}
	}
	if d.Get("clusters") != nil && len(d.Get("clusters").([]interface{})) > 0 {
		ans.Clusters = parseStringArray(d.Get("clusters").([]interface{}))
	} else {
		ans.Clusters = []string{"*"}
	}
	if d.Get("coderepos") != nil && len(d.Get("coderepos").([]interface{})) > 0 {
		ans.CodeRepos = parseStringArray(d.Get("coderepos").([]interface{}))
	} else {
		ans.CodeRepos = []string{"*"}
	}
	if d.Get("color") != nil {
		ans.Color = d.Get("color").(string)
	}
	if d.Get("containers") != nil && len(d.Get("containers").([]interface{})) > 0 {
		ans.Containers = parseStringArray(d.Get("containers").([]interface{}))
	} else {
		ans.Containers = []string{"*"}
	}
	if d.Get("description") != nil {
		ans.Description = d.Get("description").(string)
	}
	if d.Get("functions") != nil && len(d.Get("functions").([]interface{})) > 0 {
		ans.Functions = parseStringArray(d.Get("functions").([]interface{}))
	} else {
		ans.Functions = []string{"*"}
	}
	if d.Get("hosts") != nil && len(d.Get("hosts").([]interface{})) > 0 {
		ans.Hosts = parseStringArray(d.Get("hosts").([]interface{}))
	} else {
		ans.Hosts = []string{"*"}
	}
	if d.Get("images") != nil && len(d.Get("images").([]interface{})) > 0 {
		ans.Images = parseStringArray(d.Get("images").([]interface{}))
	} else {
		ans.Images = []string{"*"}
	}
	if d.Get("labels") != nil && len(d.Get("labels").([]interface{})) > 0 {
		ans.Labels = parseStringArray(d.Get("labels").([]interface{}))
	} else {
		ans.Labels = []string{"*"}
	}
	if d.Get("namespaces") != nil && len(d.Get("namespaces").([]interface{})) > 0 {
		ans.Namespaces = parseStringArray(d.Get("namespaces").([]interface{}))
	} else {
		ans.Namespaces = []string{"*"}
	}

	return ans
}

func saveCollection(d *schema.ResourceData, obj collections.Collection) {
	if err := d.Set("accountids", stringSliceToSet(obj.AccountIDs)); err != nil {
		log.Printf("[WARN] Error setting 'accountIDs' for %q: %s", d.Id(), err)
	}
	if err := d.Set("appids", stringSliceToSet(obj.AppIDs)); err != nil {
		log.Printf("[WARN] Error setting 'appIDs' for %q: %s", d.Id(), err)
	}
	if err := d.Set("clusters", stringSliceToSet(obj.Clusters)); err != nil {
		log.Printf("[WARN] Error setting 'clusters' for %q: %s", d.Id(), err)
	}
	if err := d.Set("coderepos", stringSliceToSet(obj.CodeRepos)); err != nil {
		log.Printf("[WARN] Error setting 'codeRepos' for %q: %s", d.Id(), err)
	}
	d.Set("color", obj.Color)
	if err := d.Set("containers", stringSliceToSet(obj.Containers)); err != nil {
		log.Printf("[WARN] Error setting 'containers' for %q: %s", d.Id(), err)
	}
	d.Set("description", obj.Description)
	if err := d.Set("functions", stringSliceToSet(obj.Functions)); err != nil {
		log.Printf("[WARN] Error setting 'functions' for %q: %s", d.Id(), err)
	}
	if err := d.Set("hosts", stringSliceToSet(obj.Hosts)); err != nil {
		log.Printf("[WARN] Error setting 'hosts' for %q: %s", d.Id(), err)
	}
	if err := d.Set("images", stringSliceToSet(obj.Images)); err != nil {
		log.Printf("[WARN] Error setting 'images' for %q: %s", d.Id(), err)
	}
	if err := d.Set("labels", stringSliceToSet(obj.Labels)); err != nil {
		log.Printf("[WARN] Error setting 'labels' for %q: %s", d.Id(), err)
	}
	d.Set("name", obj.Name)
	if err := d.Set("namespaces", stringSliceToSet(obj.Namespaces)); err != nil {
		log.Printf("[WARN] Error setting 'namespaces' for %q: %s", d.Id(), err)
	}
}

func createCollection(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	obj := parseCollection(d, "")

	if err := collections.Create(*client, obj); err != nil {
		log.Printf("Failed to create collection: %s\n", err)
		return err
	}

	d.SetId(obj.Name)
	return readCollection(d, meta)
}

func readCollection(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	id := d.Id()

	obj, err := collections.Get(*client, id)
	if err != nil {
		return err
	}

	saveCollection(d, *obj)

	return nil
}

func updateCollection(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	id := d.Id()
	obj := parseCollection(d, id)

	if err := collections.Update(*client, obj); err != nil {
		return err
	}

	return readCollection(d, meta)
}

func deleteCollection(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	id := d.Id()

	err := collections.Delete(*client, id)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func parseStringArray(itemList []interface{}) []string {
	listArray := make([]string, 0, len(itemList))
	if len(itemList) > 0 {
		for i := 0; i < len(itemList); i++ {
			item := itemList[i].(string)
			listArray = append(listArray, item)
		}
	}
	return listArray
}
