package provider

import (
	"fmt"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/auth"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/convert"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCredentials() *schema.Resource {
	return &schema.Resource{
		Create: createCredentials,
		Read:   readCredentials,
		Update: updateCredentials,
		Delete: deleteCredentials,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The ID of the credential.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"account_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Account identifier (username, access key, etc.).",
			},
			"api_token": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "The plain and encrypted version of the API token (the plain version is never stored in the database)",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"encrypted": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Encrypted value for the secret",
						},
						"plain": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Plain text value for the secret. Note: marshalling to JSON will convert to an encrypted value",
						},
					},
				},
			},
			"ca_cert": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "CA certificate for certificate-based authentication.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the credential.",
			},
			"ibm_account_guid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IBM Cloud account GUID.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Unique name for the credential.",
			},
			"role_arn": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Amazon Resource Name (ARN) of the role to assume.",
			},
			"secret": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Plain and encrypted version of the credential (the plain version is never stored in the database)",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"encrypted": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Encrypted value for the secret",
						},
						"plain": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Plain text value for the secret. Note: marshalling to JSON will convert to an encrypted value",
						},
					},
				},
			},
			"skip_cert_verification": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "SkipVerify if should skip certificate verification in tls communication.",
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Credential type.",
			},
			"url": { // GitHub Enterprise
				Type:        schema.TypeString,
				Optional:    true,
				Description: "URL is the server base url.",
			},
			"use_aws_role": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates if authentication should be done with the instance's attached credentials (EC2 IAM Role).",
			},
		},
	}
}

func createCredentials(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*api.Client)
	parsedCredential, err := convert.SchemaToCredential(d)
	if err != nil {
		return fmt.Errorf("error converting schema to credential: %s", err)
	}

	if err := auth.UpdateCredential(*client, parsedCredential); err != nil {
		return fmt.Errorf("error creating credential '%v': %s", parsedCredential.Id, err)
	}
	d.SetId(parsedCredential.Id)
	return readCredentials(d, meta)
}

func readCredentials(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*api.Client)
	retrievedCredential, err := auth.GetCredential(*client, d.Id())
	if err != nil {
		return fmt.Errorf("error getting credential '%s' from Console: %s", d.Id(), err)
	}

	d.Set("account_id", retrievedCredential.AccountID)
	if err := d.Set("api_token", convert.CredentialSecretToSchema(retrievedCredential.ApiToken)); err != nil {
		return fmt.Errorf("error converting credential secret to schema: %s", err)
	}
	d.Set("ca_cert", retrievedCredential.CaCert)
	d.Set("description", retrievedCredential.Description)
	d.Set("ibm_account_guid", retrievedCredential.AccountGUID)
	d.Set("name", retrievedCredential.Id)
	d.Set("role_arn", retrievedCredential.RoleArn)
	if err := d.Set("secret", convert.CredentialSecretToSchema(retrievedCredential.Secret)); err != nil {
		return fmt.Errorf("error converting credential secret to schema: %s", err)
	}
	d.Set("skip_cert_verification", retrievedCredential.SkipVerify)
	d.Set("type", retrievedCredential.Type)
	d.Set("url", retrievedCredential.Url)
	d.Set("use_aws_role", retrievedCredential.UseAWSRole)

	return nil
}

func updateCredentials(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*api.Client)
	parsedCredential, err := convert.SchemaToCredential(d)
	if err != nil {
		return fmt.Errorf("error parsing schema to credential: %s", err)
	}

	if err := auth.UpdateCredential(*client, parsedCredential); err != nil {
		return fmt.Errorf("error updating credential: %s", err)
	}
	return readCredentials(d, meta)
}

func deleteCredentials(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*api.Client)
	if err := auth.DeleteCredential(*client, d.Id()); err != nil {
		return fmt.Errorf("error deleting credential: %s", err)
	}
	d.SetId("")
	return nil
}
