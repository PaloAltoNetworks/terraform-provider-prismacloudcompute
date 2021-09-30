package prismacloudcompute

import (
	"fmt"
	"time"

	"github.com/paloaltonetworks/prisma-cloud-compute-go/pcc"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/auth"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCredentials() *schema.Resource {
	return &schema.Resource{
		Create: createCredentials,
		Read:   readCredentials,
		Update: updateCredentials,
		Delete: deleteCredentials,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Unique ID for the credential.",
			},
			"accountguid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Unique ID for an IBM Cloud account.",
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
				Elem: &schema.Schema{
					Type: schema.TypeString,
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
				Description: "Secret contains the plain and encrypted version of a value (the plain version is never stored in the DB)",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"skipverify": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "SkipVerify if should skip certificate verification in tls communication.",
			},
			"tokens": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "TemporaryToken is a temporary session token for cloud provider APIs AWS - https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp.html GCP - https://cloud.google.com/iam/docs/creating-short-lived-service-account-credentials Azure - https://docs.microsoft.com/en-us/azure/active-directory/manage-apps/what-is-single-sign-on",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"awsaccesskeyid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Temporary access key.",
						},
						"awssecretaccesskey": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Secret contains the plain and encrypted version of a value (the plain version is never stored in the DB)",
						},
						"duration": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Duration of the token.",
						},
						"expirationtime": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Expiration time for the token.",
						},
						"token": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Secret contains the plain and encrypted version of a value (the plain version is never stored in the DB)",
							Elem: &schema.Resource{
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
					},
				},
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Credential type.",
			},
			"url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "URL is the server base url.",
			},
			"useawsrole": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates if authentication should be done with the instance's attached credentials (EC2 IAM Role).",
			},
		},
	}
}

func createCredentials(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	parsedCredential, err := parseCredentials(d)
	if err != nil {
		return fmt.Errorf("error creating credential: %s", err)
	}

	if err := auth.UpdateCredential(*client, parsedCredential); err != nil {
		return fmt.Errorf("error creating credential: %s %s", err, parsedCredential.Id)
	}

	d.SetId(parsedCredential.Id)
	return readCredentials(d, meta)
}

func readCredentials(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	retrievedCredentials, err := auth.ListCredentials(*client)
	if err != nil {
		return fmt.Errorf("error reading credential: %s", err)
	}

	retrievedCredential := retrievedCredentials[0]

	if err := d.Set("id", retrievedCredential.Id); err != nil {
		return fmt.Errorf("error reading %s _id: %s", retrievedCredential.Id, err)
	}
	if err := d.Set("accountguid", retrievedCredential.AccountGUID); err != nil {
		return fmt.Errorf("error reading %s accountGUID: %s", retrievedCredential.AccountGUID, err)
	}
	if err := d.Set("accountid", retrievedCredential.AccountID); err != nil {
		return fmt.Errorf("error reading %s accountID: %s", retrievedCredential.AccountID, err)
	}
	if err := d.Set("apitoken", flattenSecret(retrievedCredential.ApiToken)); err != nil {
		return fmt.Errorf("error reading %s apiToken: %s", flattenSecret(retrievedCredential.ApiToken), err)
	}
	if err := d.Set("cacert", retrievedCredential.CaCert); err != nil {
		return fmt.Errorf("error reading %s caCert: %s", retrievedCredential.CaCert, err)
	}
	if err := d.Set("created", retrievedCredential.Created); err != nil {
		return fmt.Errorf("error reading %s created: %s", retrievedCredential.Created, err)
	}
	if err := d.Set("description", retrievedCredential.Description); err != nil {
		return fmt.Errorf("error reading %s description: %s", retrievedCredential.Description, err)
	}
	if err := d.Set("external", retrievedCredential.External); err != nil {
		return fmt.Errorf("error reading %s external: %s", retrievedCredential.External, err)
	}
	if err := d.Set("lastmodified", retrievedCredential.LastModified); err != nil {
		return fmt.Errorf("error reading %s lastModified: %s", retrievedCredential.LastModified, err)
	}
	if err := d.Set("owner", retrievedCredential.Owner); err != nil {
		return fmt.Errorf("error reading %s owner: %s", retrievedCredential.Owner, err)
	}
	if err := d.Set("rolearn", retrievedCredential.RoleArn); err != nil {
		return fmt.Errorf("error reading %s roleArn: %s", retrievedCredential.RoleArn, err)
	}
	if err := d.Set("secret", flattenSecret(retrievedCredential.Secret)); err != nil {
		return fmt.Errorf("error reading %s secret: %s", flattenSecret(retrievedCredential.Secret), err)
	}
	if err := d.Set("skipverify", retrievedCredential.SkipVerify); err != nil {
		return fmt.Errorf("error reading %s skipVerify: %s", retrievedCredential.SkipVerify, err)
	}
	if err := d.Set("tokens", retrievedCredential.Tokens); err != nil {
		return fmt.Errorf("error reading %s tokens: %s", flattenTokens(retrievedCredential.Tokens), err)
	}
	if err := d.Set("type", retrievedCredential.Type); err != nil {
		return fmt.Errorf("error reading %s type: %s", retrievedCredential.Type, err)
	}
	if err := d.Set("url", retrievedCredential.Url); err != nil {
		return fmt.Errorf("error reading %s url: %s", retrievedCredential.Url, err)
	}
	if err := d.Set("useawsrole", retrievedCredential.UseAWSRole); err != nil {
		return fmt.Errorf("error reading %s useAWSRole: %s", retrievedCredential.UseAWSRole, err)
	}

	return nil
}

func updateCredentials(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	parsedCredential, err := parseCredentials(d)
	if err != nil {
		return fmt.Errorf("error updating credential: %s", err)
	}

	if err := auth.UpdateCredential(*client, parsedCredential); err != nil {
		return fmt.Errorf("error updating credential: %s", err)
	}

	return readCredentials(d, meta)
}

func deleteCredentials(d *schema.ResourceData, meta interface{}) error {
	// TODO: reset to default credential
	client := meta.(*pcc.Client)
	id := d.Id()

	if err := auth.DeleteCredential(*client, id); err != nil {
		return fmt.Errorf("failed to update credential: %s", err)
	}

	d.SetId("")
	return nil
}

func parseCredentials(d *schema.ResourceData) (auth.Credential, error) {
	parsedCredential := auth.Credential{}
	
	if d.Get("id") != nil {
		parsedCredential.Id = d.Get("id").(string)
	}
	if d.Get("accountguid") != nil {
		parsedCredential.AccountGUID = d.Get("accountguid").(string)
	}
	if d.Get("accountid") != nil {
		parsedCredential.AccountID = d.Get("accountid").(string)
	}
	if d.Get("apitoken") != nil {
		parsedCredential.ApiToken = convertSecret(d.Get("apitoken").(map[string]interface{}))
	}
	if d.Get("cacert") != nil {
		parsedCredential.CaCert = d.Get("cacert").(string)
	}
	if d.Get("created") != nil {
		parsedCredential.Created = d.Get("created").(string)
	}
	if d.Get("description") != nil {
		parsedCredential.Description = d.Get("description").(string)
	}
	if d.Get("external") != nil {
		parsedCredential.External = d.Get("external").(bool)
	}
	if d.Get("lastmodified") != nil {
		parsedCredential.LastModified = d.Get("lastmodified").(string)
	}
	if d.Get("owner") != nil {
		parsedCredential.Owner = d.Get("owner").(string)
	}
	if d.Get("rolearn") != nil {
		parsedCredential.RoleArn = d.Get("rolearn").(string)
	}
	if d.Get("secret") != nil {
		parsedCredential.Secret = convertSecret(d.Get("secret").(map[string]interface{}))
	}
	if d.Get("skipverify") != nil {
		parsedCredential.SkipVerify = d.Get("skipverify").(bool)
	}
	if d.Get("tokens") != nil && len(d.Get("tokens").([]interface{})) > 0 {
		parsedCredential.Tokens = convertTokens(d.Get("tokens").([]interface{}))
	} else {
		parsedCredential.Tokens = []auth.TemporaryToken{}
	}
	if d.Get("type") != nil {
		parsedCredential.Type = d.Get("type").(string)
	}
	if d.Get("url") != nil {
		parsedCredential.Url = d.Get("url").(string)
	}
	if d.Get("useawsrole") != nil {
		parsedCredential.UseAWSRole = d.Get("useawsrole").(bool)
	}

	return parsedCredential, nil
}

func convertSecret(valMap map[string]interface{}) auth.Secret {
	ans := auth.Secret{}
	
	if valMap["encrypted"] != nil {
		ans.Encrypted = valMap["encrypted"].(string)
	}
	if valMap["plain"] != nil {
		ans.Plain = valMap["plain"].(string)
	}
	return ans
}

func convertTokens(in []interface{}) []auth.TemporaryToken {
	ans := []auth.TemporaryToken{}
	for _, val := range in {
		m := auth.TemporaryToken{}
		valMap := val.(map[string]interface{})
		if valMap["awsaccesskeyid"] != nil {
			m.AwsAccessKeyId = valMap["awsaccesskeyid"].(string)
		}
		if valMap["awssecretaccessKey"] != nil {
			m.AwsSecretAccessKey = convertSecret(valMap["awssecretaccessKey"].(map[string]interface{}))
		}
		if valMap["duration"] != nil {
			m.Duration = valMap["duration"].(int)
		}
		if valMap["expirationtime"] != nil {
			m.ExpirationTime = valMap["expirationtime"].(string)
		}
		if valMap["token"] != nil {
			m.Token = convertSecret(valMap["token"].(map[string]interface{}))
		}
		ans = append(ans, m)
	}
	
	return ans
}

func flattenSecret(val auth.Secret) map[string]interface{} {
	ans := make(map[string]interface{})
	ans["collections"] = val.Encrypted
	ans["project"] = val.Plain
	
	return ans
}


func flattenTokens(in []auth.TemporaryToken) []interface{} {
	ans := make([]interface{}, 0, len(in))
	for _, val := range in {
		m := make(map[string]interface{})
		m["awsAccessKeyId"] = val.AwsAccessKeyId
		m["awsSecretAccessKey"] = flattenSecret(val.AwsSecretAccessKey)
		m["duration"] = val.Duration
		m["expirationTime"] = val.ExpirationTime
		m["token"] = flattenSecret(val.Token)
		
		ans = append(ans, m)
	}
	return ans
}
