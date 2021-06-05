package prismacloudcompute

import (
	//	"encoding/json"
	"log"
	"time"

	pc "github.com/paloaltonetworks/prisma-cloud-compute-go"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/collection"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	//	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
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
			/*			"collections": {
						Type:        schema.TypeList,
						Required:    true,
						Description: "Collections",
						MinItems:    1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{*/
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
				Computed:    true,
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
			//					},
			//				},
			//			},
		},
	}
}

func parseCollection(d *schema.ResourceData, id string) collection.Collection {
	//	rspec := d.Get("collections").([]interface{})[0].(map[string]interface{})
	ans := collection.Collection{
		Name:        d.Get("name").(string),
		AccountIDs:  d.Get("accountIDs").([]string),
		AppIDs:      d.Get("appIDs").([]string),
		Clusters:    d.Get("clusters").([]string),
		CodeRepos:   d.Get("codeRepos").([]string),
		Color:       d.Get("color").(string),
		Containers:  d.Get("containers").([]string),
		Description: d.Get("description").(string),
		Functions:   d.Get("functions").([]string),
		Hosts:       d.Get("hosts").([]string),
		Images:      d.Get("images").([]string),
		Labels:      d.Get("labels").([]string),
		Modified:    d.Get("modified").(string),
		Namespaces:  d.Get("namespaces").([]string),
		Owner:       d.Get("owner").(string),
		Prisma:      d.Get("prisma").(bool),
		System:      d.Get("system").(bool),
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
		return err
	}

	/*	PollApiUntilSuccess(func() error {
			_, err := collection.Identify(client, obj.Name)
			return err
		})

		id, err := collection.Identify(client, obj.Name)
		if err != nil {
			return err
		}

		PollApiUntilSuccess(func() error {
			_, err := collection.Get(client, obj.Name)
			return err
		})

		d.SetId(id)
	*/
	return readCollection(d, meta)
}

func readCollection(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	obj := parseCollection(d, "")
	//	id := d.Id()

	obj, err := collection.Get(client, obj.Name)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			//			d.SetName("")
			return nil
		}
		return err
	}

	//	saveCollection(d, obj)

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
