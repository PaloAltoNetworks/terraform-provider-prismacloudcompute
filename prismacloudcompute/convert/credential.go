package convert

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/auth"
)

func SchemaToCredential(d *schema.ResourceData) (auth.Credential, error) {
	parsedCredential := auth.Credential{}

	if val, ok := d.GetOk("name"); ok {
		parsedCredential.Id = val.(string)
	}
	if val, ok := d.GetOk("ibm_account_guid"); ok {
		parsedCredential.AccountGUID = val.(string)
	}
	if val, ok := d.GetOk("account_id"); ok {
		parsedCredential.AccountID = val.(string)
	}
	if val, ok := d.GetOk("api_token"); ok {
		parsedCredential.ApiToken = schemaToCredentialSecret(val.([]interface{}))
	}
	if val, ok := d.GetOk("ca_cert"); ok {
		parsedCredential.CaCert = val.(string)
	}
	if val, ok := d.GetOk("description"); ok {
		parsedCredential.Description = val.(string)
	}
	if val, ok := d.GetOk("role_arn"); ok {
		parsedCredential.RoleArn = val.(string)
	}
	if val, ok := d.GetOk("secret"); ok {
		parsedCredential.Secret = schemaToCredentialSecret(val.([]interface{}))
	}
	if val, ok := d.GetOk("skip_cert_verification"); ok {
		parsedCredential.SkipVerify = val.(bool)
	}
	if val, ok := d.GetOk("type"); ok {
		parsedCredential.Type = val.(string)
	}
	if val, ok := d.GetOk("url"); ok {
		parsedCredential.Url = val.(string)
	}
	if val, ok := d.GetOk("use_aws_role"); ok {
		parsedCredential.UseAWSRole = val.(bool)
	}

	return parsedCredential, nil
}

func schemaToCredentialSecret(in []interface{}) auth.Secret {
	ans := auth.Secret{}
	if in[0] != nil {
		presentSecret := in[0].(map[string]interface{})
		ans.Encrypted = presentSecret["encrypted"].(string)
		ans.Plain = presentSecret["plain"].(string)
	}
	return ans
}

func CredentialSecretToSchema(in auth.Secret) map[string]interface{} {
	ans := make(map[string]interface{})
	ans["encrypted"] = in.Encrypted
	ans["plain"] = in.Plain
	return ans
}
