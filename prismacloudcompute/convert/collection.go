package convert

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/collection"
)

// Converts a list of collection objects to a list of strings (collection names)
// for policy resource schema compatibility.
func CollectionsToPolicySchema(in []collection.Collection) []interface{} {
	ans := make([]interface{}, 0, len(in))
	for _, val := range in {
		ans = append(ans, val.Name)
	}
	return ans
}

// Converts a list of strings (collection names) provided by the 'collections' key
// in all policy resources to a list of collection objects for SDK compatibility.
func PolicySchemaToCollections(in []interface{}) []collection.Collection {
	ans := make([]collection.Collection, 0, len(in))
	for _, val := range in {
		parsedCollection := collection.Collection{
			Name: val.(string),
		}
		ans = append(ans, parsedCollection)
	}
	return ans
}

// Converts a collection schema to a collection object for SDK compatibility.
func SchemaToCollection(d *schema.ResourceData) collection.Collection {
	ans := collection.Collection{
		AccountIds: []string{"*"},
		AppIds:     []string{"*"},
		Clusters:   []string{"*"},
		CodeRepos:  []string{"*"},
		Color:      d.Get("color").(string),
		Containers: []string{"*"},
		Functions:  []string{"*"},
		Hosts:      []string{"*"},
		Images:     []string{"*"},
		Labels:     []string{"*"},
		Name:       d.Get("name").(string),
		Namespaces: []string{"*"},
	}
	if val, ok := d.GetOk("account_ids"); ok {
		ans.AccountIds = SchemaToStringSlice(val.([]interface{}))
	}
	if val, ok := d.GetOk("application_ids"); ok {
		ans.AppIds = SchemaToStringSlice(val.([]interface{}))
	}
	if val, ok := d.GetOk("clusters"); ok {
		ans.Clusters = SchemaToStringSlice(val.([]interface{}))
	}
	if val, ok := d.GetOk("code_repositories"); ok {
		ans.CodeRepos = SchemaToStringSlice(val.([]interface{}))
	}
	if val, ok := d.GetOk("containers"); ok {
		ans.Containers = SchemaToStringSlice(val.([]interface{}))
	}
	if val, ok := d.GetOk("description"); ok {
		ans.Description = val.(string)
	}
	if val, ok := d.GetOk("functions"); ok {
		ans.Functions = SchemaToStringSlice(val.([]interface{}))
	}
	if val, ok := d.GetOk("hosts"); ok {
		ans.Hosts = SchemaToStringSlice(val.([]interface{}))
	}
	if val, ok := d.GetOk("images"); ok {
		ans.Images = SchemaToStringSlice(val.([]interface{}))
	}
	if val, ok := d.GetOk("labels"); ok {
		ans.Labels = SchemaToStringSlice(val.([]interface{}))
	}
	if val, ok := d.GetOk("namespaces"); ok {
		ans.Namespaces = SchemaToStringSlice(val.([]interface{}))
	}
	return ans
}
