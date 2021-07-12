package prismacloudcompute

import (
	"log"
	"time"

	pc "github.com/paloaltonetworks/prisma-cloud-compute-go"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/collection"

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
				Description: "List of Kubernetes cluster names.",
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
				Description: "A hex color code for a collection",
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
			"modified": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Date/time when the collection was last modified.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique collection name.",
			},
			"namespaces": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of Kubernetes namespaces.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"owner": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User who created or last modified the collection.",
			},
			"prisma": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If set to 'true', this collection originates from Prisma Cloud.",
			},
			"system": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If set to 'true', this collection was created by the system (i.e., a non-user). Otherwise it was created by a real user.",
			},
		},
	}
}

func parseCollection(d *schema.ResourceData, id string) collection.Collection {
	ans := collection.Collection{
		Name: d.Get("name").(string),
	}
	if d.Get("accountIDs") != nil && len(d.Get("accountIDs").([]interface{})) > 0 {
		ans.AccountIDs = parseStringArray(d.Get("accountIDs").([]interface{}))
    }
	if d.Get("appIDs") != nil && len(d.Get("appIDs").([]interface{})) > 0 {
		ans.AppIDs = parseStringArray(d.Get("appIDs").([]interface{}))
	}
	if d.Get("clusters") != nil && len(d.Get("clusters").([]interface{})) > 0 {
		ans.Clusters = parseStringArray(d.Get("clusters").([]interface{}))
	}
	if d.Get("codeRepos") != nil && len(d.Get("codeRepos").([]interface{})) > 0 {
		ans.CodeRepos = parseStringArray(d.Get("codeRepos").([]interface{}))
	}
	if d.Get("color") != nil {
		ans.Color = d.Get("color").(string)
	}
	if d.Get("containers") != nil && len(d.Get("containers").([]interface{})) > 0 {
		ans.Containers = parseStringArray(d.Get("containers").([]interface{}))
	}
	if d.Get("description") != nil {
		ans.Description = d.Get("description").(string)
	}
	if d.Get("functions") != nil && len(d.Get("functions").([]interface{})) > 0 {
		ans.Functions = parseStringArray(d.Get("functions").([]interface{}))
	}
	if d.Get("hosts") != nil && len(d.Get("hosts").([]interface{})) > 0 {
		ans.Hosts = parseStringArray(d.Get("hosts").([]interface{}))
	}
	if d.Get("images") != nil && len(d.Get("images").([]interface{})) > 0 {
		ans.Images = parseStringArray(d.Get("images").([]interface{}))
	}
	if d.Get("labels") != nil && len(d.Get("labels").([]interface{})) > 0 {
		ans.Labels = parseStringArray(d.Get("labels").([]interface{}))
	}
	if d.Get("modified") != nil {
		ans.Modified = d.Get("modified").(string)
	}
	if d.Get("namespaces") != nil && len(d.Get("namespaces").([]interface{})) > 0 {
		ans.Namespaces = parseStringArray(d.Get("namespaces").([]interface{}))
	}
	if d.Get("owner") != nil {
		ans.Owner = d.Get("owner").(string)
	}
	if d.Get("prisma") != nil {
		ans.Prisma = d.Get("prisma").(interface{}).(bool)
	}
	if d.Get("system") != nil {
		ans.System = d.Get("system").(interface{}).(bool)
	}

	return ans
}

func saveCollection(d *schema.ResourceData, obj collection.Collection) {
	if err := d.Set("accountIDs", StringSliceToSet(obj.AccountIDs)); err != nil {
		log.Printf("[WARN] Error setting 'accountIDs' for %q: %s", d.Id(), err)
	}
	if err := d.Set("appIDs", StringSliceToSet(obj.AppIDs)); err != nil {
		log.Printf("[WARN] Error setting 'appIDs' for %q: %s", d.Id(), err)
	}
	if err := d.Set("clusters", StringSliceToSet(obj.Clusters)); err != nil {
		log.Printf("[WARN] Error setting 'clusters' for %q: %s", d.Id(), err)
	}
	if err := d.Set("codeRepos", StringSliceToSet(obj.CodeRepos)); err != nil {
		log.Printf("[WARN] Error setting 'codeRepos' for %q: %s", d.Id(), err)
	}
	d.Set("color", obj.Color)
	if err := d.Set("containers", StringSliceToSet(obj.Containers)); err != nil {
		log.Printf("[WARN] Error setting 'containers' for %q: %s", d.Id(), err)
	}
	d.Set("description", obj.Description)
	if err := d.Set("functions", StringSliceToSet(obj.Functions)); err != nil {
		log.Printf("[WARN] Error setting 'functions' for %q: %s", d.Id(), err)
	}
	if err := d.Set("hosts", StringSliceToSet(obj.Hosts)); err != nil {
		log.Printf("[WARN] Error setting 'hosts' for %q: %s", d.Id(), err)
	}
	if err := d.Set("images", StringSliceToSet(obj.Images)); err != nil {
		log.Printf("[WARN] Error setting 'images' for %q: %s", d.Id(), err)
	}
	if err := d.Set("labels", StringSliceToSet(obj.Labels)); err != nil {
		log.Printf("[WARN] Error setting 'labels' for %q: %s", d.Id(), err)
	}
	d.Set("modified", obj.Modified)
	d.Set("name", obj.Name)
	if err := d.Set("namespaces", StringSliceToSet(obj.Namespaces)); err != nil {
		log.Printf("[WARN] Error setting 'namespaces' for %q: %s", d.Id(), err)
	}
	d.Set("owner", obj.Owner)
	d.Set("prisma", obj.Prisma)
	d.Set("system", obj.System)
}

func createCollection(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	obj := parseCollection(d, "")

	if err := collection.Create(client, obj); err != nil {
		log.Printf("Failed to create collection: %s\n", err)
		return err
	}

	PollApiUntilSuccess(func() error {
		_, err := collection.Get(client, obj.Name)
		log.Printf("Failed to get collection %s: %s\n", obj.Name, err)
		return err
	})

	d.SetId(obj.Name)
	return readCollection(d, meta)
}

func readCollection(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	obj := parseCollection(d, "")
	id := d.Id()

	obj, err := collection.Get(client, id)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return err
	}

	saveCollection(d, obj)

	return nil
}

func updateCollection(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	id := d.Id()
	obj := parseCollection(d, id)

	if err := collection.Update(client, obj); err != nil {
		return err
	}

	return readCollection(d, meta)
}

func deleteCollection(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	id := d.Id()

	err := collection.Delete(client, id)
	if err != nil {
		if err != pc.ObjectNotFoundError {
			return err
		}
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