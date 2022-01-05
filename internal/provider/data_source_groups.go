package provider

// import (
// 	"log"

// 	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/auth"
// 	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/client"

// 	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
// )

// func dataSourceGroups() *schema.Resource {
// 	return &schema.Resource{
// 		Read: dataSourceGroupsRead,

// 		Schema: map[string]*schema.Schema{

// 			// Output.
// 			"groupid": {
// 				Type:        schema.TypeString,
// 				Optional:    true,
// 				Description: "Group ID.",
// 			},
// 			"ldapgroup": {
// 				Type:        schema.TypeBool,
// 				Optional:    true,
// 				Description: "Indicates if the group is an LDAP group (true) or not (false).",
// 			},
// 			"groupname": {
// 				Type:        schema.TypeString,
// 				Optional:    true,
// 				Description: "Group name.",
// 			},
// 			"oauthgroup": {
// 				Type:        schema.TypeBool,
// 				Optional:    true,
// 				Description: "Indicates if the group is an OAuth group (true) or not (false).",
// 			},
// 			"oidcgroup": {
// 				Type:        schema.TypeBool,
// 				Optional:    true,
// 				Description: "Indicates if the group is an OpenID Connect group (true) or not (false).",
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
// 				Description: "Role of the group.",
// 			},
// 			"samlgroup": {
// 				Type:        schema.TypeBool,
// 				Optional:    true,
// 				Description: "Indicates if the group is a SAML group (true) or not (false).",
// 			},
// 			"users": {
// 				Type:        schema.TypeList,
// 				Optional:    true,
// 				Description: "Users in the group.",
// 				Elem: &schema.Resource{
// 					Schema: map[string]*schema.Schema{
// 						"username": {
// 							Type:        schema.TypeString,
// 							Optional:    true,
// 							Description: "Name of a user.",
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}
// }

// func dataSourceGroupsRead(d *schema.ResourceData, meta interface{}) error {
// 	client := meta.(*api.Client)

// 	i, err := auth.ListGroups(*client)

// 	if err != nil {
// 		return err
// 	}

// 	list := make([]interface{}, 0, 1)
// 	for _, val := range i {
// 		list = append(list, map[string]interface{}{
// 			"groupId":     val.Id,
// 			"ldapGroup":   val.LdapGroup,
// 			"groupName":   val.Name,
// 			"oauthGroup":  val.OauthGroup,
// 			"oidcGroup":   val.OidcGroup,
// 			"permissions": val.Permissions,
// 			"role":        val.Role,
// 			"samlGroup":   val.SamlGroup,
// 			"users":       val.Users,
// 		})
// 	}

// 	if err := d.Set("listing", list); err != nil {
// 		log.Printf("[WARN] Error setting 'listing' field for %q: %s", d.Id(), err)
// 	}

// 	return nil
// }
