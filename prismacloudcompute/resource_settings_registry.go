package prismacloudcompute

import (
	"log"
	"time"

	pcc "github.com/paloaltonetworks/prisma-cloud-compute-go"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/settings/registry"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceRegistry() *schema.Resource {
	return &schema.Resource{
		Create: createRegistry,
		Read:   readRegistry,
		Update: updateRegistry,
		Delete: deleteRegistry,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"harborscannerurlsuffix": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "URL suffix for the harbor scanner.",
			},
			"webhookurlsuffix": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "URL suffix for the webhook.",
			},
			"specifications": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of specifications.",
				Elem: &schema.Resource {
					Schema: map[string]*schema.Schema{
						"cap": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Specifies the maximum number of images from each repo to fetch and scan, sorted by most recently modified.",
						},
						"collections": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Specifies the set of Defenders in-scope for working on a scan job.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"credential": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Credential contains external provider authentication data",
							Elem: &schema.Resource {
								Schema: map[string]*schema.Schema{
									"_id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique ID for the credential.",
									},
									"accountguid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Account identifier (e.g., username, access key, account GUID, etc.).",
									},
									"accountid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Account identifier (e.g., username, access key, account GUID, etc.).",
									},
									"apitoken": {
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "Secret contains the plain and encrypted version of a value (the plain version is never stored in the DB)",
										Elem: &schema.Resource {
											Schema: map[string]*schema.Schema{
												"encrypted": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Encrypted value for the secret.",
												},
												"plain": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Plain text value for the secret. Note: marshalling to JSON will convert to an encrypted value.",
												},
											},
										},
									},
									"cacert": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "CA certificate for certificate-based authentication.",
									},
									"created": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Created when the credential was created (or the account ID was changed for AWS).",
									},
									"description": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Description of the credential.",
									},
									"external": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Indicates if the credential is external (true) or not (false).",
									},
									"lastmodified": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Datetime when the credential was last modified.",
									},
									"owner": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "User who created or modified the credential.",
									},
									"rolearn": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Amazon Resource Name (ARN) of the role to assume.",
									},
									"secret": {
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "Secret contains the plain and encrypted version of a value (the plain version is never stored in the DB).",
										Elem: &schema.Resource {
											Schema: map[string]*schema.Schema{
												"encrypted": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Encrypted value for the secret.",
												},
												"plain": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Plain text value for the secret. Note: marshalling to JSON will convert to an encrypted value.",
												},
											},
										},
									},
									"skipverify": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "SkipVerify if should skip certificate verification in tls communication.",
									},
									"type": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The credential type.",
									},
									"url": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The server base URL.",
									},
									"useawsrole": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Indicates if authentication should be done with the instance's attached credentials (EC2 IAM Role).",
									},
								},
							},
						},
						"credentialid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "ID of the credentials in the credentials store to use for authenticating with the registry.",
						},
						"excludedrepositories": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Repositories to exclude from scanning.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"excludedtags": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Tags to exclude from scanning.",
						},
						"harbordeploymentsecurity": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Indicates whether the Prisma Cloud plugin uses temporary tokens provided by Harbor to scan images in projects where Harbor's deployment security setting is enabled.",
						},
						"jfrogrepotypes": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "JFrog Artifactory repository types to scan.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"namespace": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IBM Bluemix namespace https://console.bluemix.net/docs/services/Regis.",
						},
						"os": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "RegistryOSType specifies the registry images base OS type.",
						},
						"registry": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Registry address (e.g., https://gcr.io)..",
						},
						"repository": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Repositories to scan.",
						},
						"scanners": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Number of Defenders that can be utilized for each scan job.",
						},
						"tag": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Tags to scan.",
						},
						"version": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Registry type. Determines the protocol Prisma Cloud uses to communicate with the registry.",
						},
						"versionpattern": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Pattern heuristic for quickly filtering images by tags without having to query all images for modification dates.",
						},						
					},
				},
			},
		},
	}
}

func parseRegistry(d *schema.ResourceData, id string) registry.Registry {
	ans := registry.Registry{}
	if d.Get("harborscannerurlsuffix") != nil {
		ans.HarborScannerUrlSuffix = d.Get("harborscannerurlsuffix").(string)
	}
	if d.Get("specifications") != nil && len(d.Get("specifications").([]interface{})) > 0 {
		ans.Specifications = getSpecifications(d.Get("specifications").([]interface{}))
	}
	if d.Get("webhookurlsuffix") != nil {
		ans.WebhookUrlSuffix = d.Get("webhookurlsuffix").(string)
	}

	return ans
}

func saveRegistry(d *schema.ResourceData, obj registry.Registry) {
	d.Set("harborScannerUrlSuffix", obj.HarborScannerUrlSuffix)
	if err := d.Set("specifications", obj.Specifications); err != nil {
		log.Printf("[WARN] Error setting 'specifications' for %q: %s", d.Id(), err)
	}
	d.Set("webhookUrlSuffix", obj.WebhookUrlSuffix)
}

func createRegistry(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	obj := parseRegistry(d, "")

	if err := registry.Update(*client, obj); err != nil {
		log.Printf("Failed to create Registry: %s\n", err)
		return err
	}

	reg, err := registry.Get(*client)
	if err != nil {
		return err
	}

	d.SetId(reg.WebhookUrlSuffix)
	return readRegistry(d, meta)	
}

func readRegistry(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)

	obj, err := registry.Get(*client)
	if err != nil {
		return err
	}

	saveRegistry(d, obj)

	return nil
}

func updateRegistry(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	id := d.Id()
	obj := parseRegistry(d, id)

	if err := registry.Update(*client, obj); err != nil {
		return err
	}

	return readRegistry(d, meta)
}

func deleteRegistry(d *schema.ResourceData, meta interface{}) error {
/*	client := meta.(*pcc.Client)
	id := d.Id()

	err := registry.Delete(*client, id)
	if err != nil {
		return err
	}
*/
	d.SetId("")
	return nil
}

func getSpecifications(specs []interface{}) []registry.Specification {
	specsArray := []registry.Specification{}
		
	for i := 0; i < len(specs); i++ {
		specItem := specs[i].(map[string]interface{})
		spec := registry.Specification{}

		if specItem["cap"] != nil {
			spec.Cap = specItem["cap"].(int)
		}

		if specItem["collections"] != nil {
			spec.Collections = parseStringArray(specItem["collections"].([]interface{}))
		}
		if specItem["credential"] != nil {
			credItem := specItem["credential"].(map[string]interface{})
			cred := registry.Credential{}
				
			if credItem["accountguid"] != nil {
				cred.AccountGUID = credItem["accountguid"].(string)
			}
			if credItem["accountid"] != nil {
				cred.AccountID = credItem["accountid"].(string)
			}
			if credItem["apitoken"] != nil {
				apiTokenItem := credItem["apitoken"]
				cred.ApiToken = getStringResult(apiTokenItem)
			}
			if credItem["cacert"] != nil {
				cred.CaCert = credItem["cacert"].(string)
			}
			if credItem["created"] != nil {
				cred.Created = credItem["created"].(string)
			}
			if credItem["description"] != nil {
				cred.Description = credItem["description"].(string)
			}
			if credItem["external"] != nil {
				cred.External = credItem["external"].(bool)
			}
			if credItem["_id"] != nil {
				cred.Id = credItem["_id"].(string)
			}
			if credItem["lastmodified"] != nil {
				cred.LastModified = credItem["lastmodified"].(string)
			}
			if credItem["owner"] != nil {
				cred.Owner = credItem["owner"].(string)
			}
			if credItem["rolearn"] != nil {
				cred.RoleArn = credItem["rolearn"].(string)
			}
			if credItem["secret"] != nil {
				secretItem := credItem["secret"].(string)
				cred.Secret = getStringResult(secretItem )
			}
			if credItem["skipverify"] != nil {
				cred.RoleArn = credItem["skipverify"].(string)
			}
			if credItem["type"] != nil {
				cred.SkipVerify = credItem["type"].(bool)
			}
			if credItem["url"] != nil {
				cred.Url = credItem["url"].(string)
			}
			if credItem["useawsrole"] != nil {
				cred.UseAWSRole = credItem["useawsrole"].(bool)
			}
			
			spec.Credential = cred
		}
		if specItem["credentialid"] != nil {
			spec.CredentialID = specItem["credentialid"].(string)
		}
		if specItem["excludedrepositories"] != nil {
			spec.ExcludedRepositories = parseStringArray(specItem["excludedrepositories"].([]interface{}))
		}
		if specItem["excludedtags"] != nil {
			spec.ExcludedTags = specItem["excludedtags"].(string)
		}
		if specItem["harbordeploymentsecurity"] != nil {
			spec.HarborDeploymentSecurity = specItem["harbordeploymentsecurity"].(bool)
		}
		if specItem["jfrogrepotypes"] != nil {
			spec.JfrogRepoTypes = parseStringArray(specItem["jfrogrepotypes"].([]interface{}))
		}
		if specItem["namespace"] != nil {
			spec.Namespace = specItem["namespace"].(string)
		}
		if specItem["os"] != nil {
			spec.Os = specItem["os"].(string)
		}
		if specItem["tag"] != nil {
			spec.Tag = specItem["tag"].(string)
		}
		if specItem["registry"] != nil {
			spec.Registry = specItem["registry"].(string)
		}
		if specItem["repository"] != nil {
			spec.Repository = specItem["repository"].(string)
		}
		if specItem["scanners"] != nil {
			spec.Scanners = specItem["scanners"].(int)
		}
		if specItem["version"] != nil {
			spec.Version = specItem["version"].(string)
		}
		if specItem["versionpattern"] != nil {
			spec.VersionPattern = specItem["versionpattern"].(string)
		}
		
		specsArray = append(specsArray, spec)
	}
	
	return specsArray
}

func getStringResult(stringResultItem interface{}) registry.StringResult {
	item := stringResultItem.(map[string]interface{})
	stringResult := registry.StringResult{}
	
	if item["encrypted"] != nil {
		stringResult.Encrypted = item["encrypted"].(string)
	}
	if item["plain"] != nil {
		stringResult.Plain = item["plain"].(string)
	}
	return stringResult
}				
