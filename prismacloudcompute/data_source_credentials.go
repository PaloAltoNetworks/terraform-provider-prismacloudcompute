package prismacloudcompute

import (
	"github.com/paloaltonetworks/prisma-cloud-compute-go/pcc"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/auth"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCredentials() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCredentialsRead,

		Schema: map[string]*schema.Schema{

			// Output.
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

func dataSourceCredentialsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)

	i, err := auth.ListCredentials(*client)

	if err != nil {
		return err
	}

	list := make([]interface{}, 0, 1)
	for _, val := range i {
		list = append(list, map[string]interface{}{
			"_id": val.Id,
			"accountGUID": val.AccountGUID,
			"accountID": val.AccountID,
			"apiToken": flattenSecret(val.ApiToken),
			"caCert": val.CaCert,
			"created": val.Created,
			"description": val.Description,
			"external": val.External,
			"lastModified": val.LastModified,
			"owner": val.Owner,
			"roleArn": val.RoleArn,
			"secret": flattenSecret(val.Secret),
			"skipVerify": val.SkipVerify,
			"tokens": flattenTokens(val.Tokens),
			"type": val.Type,
			"url": val.Url,
			"useAWSRole": val.UseAWSRole,
		})
	}

	return nil
}
