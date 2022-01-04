package provider

// import (
// 	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/client"
// 	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/auth"

// 	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
// )

// func dataSourceRbacRoles() *schema.Resource {
// 	return &schema.Resource{
// 		Read: dataSourceRbacRolesRead,

// 		Schema: map[string]*schema.Schema{

// 			// Output.
// 			"description": {
// 				Type:        schema.TypeString,
// 				Optional:    true,
// 				Description: "Role description.",
// 			},
// 			"name": {
// 				Type:        schema.TypeString,
// 				Optional:    true,
// 				Description: "Role name.",
// 			},
// 			"permissions": {
// 				Type:        schema.TypeList,
// 				Optional:    true,
// 				Description: "List of permissions.",
// 				Elem: &schema.Resource{
// 					Schema: map[string]*schema.Schema{
// 						"name": {
// 							Type:        schema.TypeString,
// 							Optional:    true,
// 							Description: "Names roles for the user.",
// 						},
// 						"readwrite": {
// 							Type:        schema.TypeBool,
// 							Optional:    true,
// 							Description: "Indicates the type of permission.",
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}
// }

// func dataSourceRbacRolesRead(d *schema.ResourceData, meta interface{}) error {
// 	client := meta.(*api.Client)

// 	i, err := auth.ListRoles(*client)

// 	if err != nil {
// 		return err
// 	}

// 	list := make([]interface{}, 0, 1)
// 	for _, val := range i {
// 		list = append(list, map[string]interface{}{
// 			"description": val.Description,
// 			"name": val.Name,
// 			"perms": val.Permissions,
// 		})
// 	}

// 	return nil
// }
