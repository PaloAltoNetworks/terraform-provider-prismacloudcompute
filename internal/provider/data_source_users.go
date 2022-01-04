package provider

// import (
// 	"log"

// 	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/auth"
// 	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/client"

// 	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
// )

// func dataSourceUsers() *schema.Resource {
// 	return &schema.Resource{
// 		Read: dataSourceUsersRead,

// 		Schema: map[string]*schema.Schema{

// 			// Output.
// 			"authtype": {
// 				Type:        schema.TypeString,
// 				Optional:    true,
// 				Description: "The user authentication type.",
// 			},
// 			"password": {
// 				Type:        schema.TypeString,
// 				Optional:    true,
// 				Description: "Password for authentication.",
// 			},
// 			"permissions": {
// 				Type:        schema.TypeList,
// 				Optional:    true,
// 				Description: "List of permissions.",
// 				Elem: &schema.Resource{
// 					Schema: map[string]*schema.Schema{
// 						"collections": {
// 							Type:        schema.TypeList,
// 							Optional:    true,
// 							Description: "Specifies the set of Defenders in-scope for working on a scan job.",
// 							Elem: &schema.Schema{
// 								Type: schema.TypeString,
// 							},
// 						},
// 						"project": {
// 							Type:        schema.TypeString,
// 							Optional:    true,
// 							Description: "Names of projects which the user can access.",
// 						},
// 					},
// 				},
// 			},
// 			"role": {
// 				Type:        schema.TypeString,
// 				Optional:    true,
// 				Description: "User role.",
// 			},
// 			"username": {
// 				Type:        schema.TypeString,
// 				Optional:    true,
// 				Description: "Username for authentication.",
// 			},
// 		},
// 	}
// }

// func dataSourceUsersRead(d *schema.ResourceData, meta interface{}) error {
// 	client := meta.(*api.Client)

// 	i, err := auth.ListUsers(*client)

// 	if err != nil {
// 		return err
// 	}

// 	list := make([]interface{}, 0, 1)
// 	for _, val := range i {
// 		list = append(list, map[string]interface{}{
// 			"authType":    val.AuthType,
// 			"password":    val.Password,
// 			"permissions": val.Permissions,
// 			"role":        val.Role,
// 			"username":    val.Username,
// 		})
// 	}

// 	if err := d.Set("listing", list); err != nil {
// 		log.Printf("[WARN] Error setting 'listing' field for %q: %s", d.Id(), err)
// 	}

// 	return nil
// }
