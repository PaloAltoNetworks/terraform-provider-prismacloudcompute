package convert

import (
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/account"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/auth"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func SchemaToCloudAccountCredential(d *schema.ResourceData) (auth.Credential, error) {
	var parsedCredential auth.Credential

	if val, ok := d.GetOk("credential"); ok {
		for _, val := range val.([]interface{}) {
			parsedCredential.Id = val.(map[string]interface{})["name"].(string)
			parsedCredential.AccountID = val.(map[string]interface{})["account_id"].(string)
			if len(val.(map[string]interface{})["api_token"].([]interface{})) > 0 {
				parsedCredential.ApiToken = schemaToCredentialSecret(val.(map[string]interface{})["api_token"].([]interface{}))
			}
			parsedCredential.CaCert = val.(map[string]interface{})["ca_cert"].(string)
			parsedCredential.Description = val.(map[string]interface{})["description"].(string)
			parsedCredential.External = val.(map[string]interface{})["external"].(bool)
			//parsedCredential.AccountGUID = val.(map[string]interface{})["ibm_account_guide"].(string)
			parsedCredential.RoleArn = val.(map[string]interface{})["role_arn"].(string)
			parsedCredential.Id = val.(map[string]interface{})["id"].(string)
			if len(val.(map[string]interface{})["secret"].([]interface{})) > 0 {
				parsedCredential.Secret = schemaToCredentialSecret(val.(map[string]interface{})["secret"].([]interface{}))
			}
			parsedCredential.SkipVerify = val.(map[string]interface{})["skip_cert_verification"].(bool)
			parsedCredential.Type = val.(map[string]interface{})["type"].(string)
			parsedCredential.Url = val.(map[string]interface{})["url"].(string)
			parsedCredential.UseAWSRole = val.(map[string]interface{})["use_aws_role"].(bool)
			parsedCredential.UseSTSRegionalEndpoint = val.(map[string]interface{})["use_sts_regional_endpoint"].(bool)
		}
	}

	return parsedCredential, nil
}

func SchemaToCloudScanRule(d *schema.ResourceData) (account.CloudScanRule, error) {
	var parsedCloudScanRule account.CloudScanRule

	if val, ok := d.GetOk("credential"); ok {
		parsedCloudScanRule.CredentialId = val.([]interface{})[0].(map[string]interface{})["name"].(string)
	}

	if val, ok := d.GetOk("aws_region_type"); ok {
		parsedCloudScanRule.AwsRegionType = val.(string)
	}

	if val, ok := d.GetOk("discovery_enabled"); ok {
		parsedCloudScanRule.DiscoveryEnabled = val.(bool)
	}

	if val, ok := d.GetOk("serverless_radar_enabled"); ok {
		parsedCloudScanRule.ServerlessRadarEnabled = val.(bool)
	}

	if val, ok := d.GetOk("vm_tags_enabled"); ok {
		parsedCloudScanRule.VmTagsEnabled = val.(bool)
	}

	if val, ok := d.GetOk("discover_all_function_versions"); ok {
		parsedCloudScanRule.DiscoverAllFunctionVersions = val.(bool)
	}

	if val, ok := d.GetOk("serverless_radar_cap"); ok {
		parsedCloudScanRule.ServerlessRadarCap = val.(int)
	}

	if val, ok := d.GetOk("agentless_scan_spec"); ok {
		specs := val.(map[string]interface{})
		parsedCloudScanRule.AgentlessScanSpec.Enabled = specs["enabled"].(bool)
		parsedCloudScanRule.AgentlessScanSpec.HubAccount = specs["hub_account"].(bool)
		parsedCloudScanRule.AgentlessScanSpec.ConsoleAddr = specs["console_addr"].(string)
		parsedCloudScanRule.AgentlessScanSpec.ScanNonRunning = specs["scan_non_running"].(bool)
		parsedCloudScanRule.AgentlessScanSpec.ProxyAddress = specs["proxy_address"].(string)
		parsedCloudScanRule.AgentlessScanSpec.ProxyCA = specs["proxy_ca"].(string)
		parsedCloudScanRule.AgentlessScanSpec.SkipPermissionsCheck = specs["skip_permissions_check"].(bool)
		parsedCloudScanRule.AgentlessScanSpec.AutoScale = specs["auto_scale"].(bool)
		parsedCloudScanRule.AgentlessScanSpec.Scanners = specs["scanners"].(int)
		parsedCloudScanRule.AgentlessScanSpec.SecurityGroup = specs["security_group"].(string)
		parsedCloudScanRule.AgentlessScanSpec.SubNet = specs["subnet"].(string)
		parsedCloudScanRule.AgentlessScanSpec.Regions = SchemaToStringSlice(specs["regions"].([]interface{}))

		presentCustomTags := specs["custom_tags"].([]interface{})
		parsedCustomTags := make([]account.Tag, 0, len(presentCustomTags))
		for _, val := range presentCustomTags {
			presentCustomTag := val.(map[string]interface{})
			parsedCustomTags = append(parsedCustomTags, account.Tag{
				Key:   presentCustomTag["key"].(string),
				Value: presentCustomTag["value"].(string),
			})
		}
		parsedCloudScanRule.AgentlessScanSpec.CustomTags = parsedCustomTags

		presentIncludedTags := specs["included_tags"].([]interface{})
		parsedIncludedTags := make([]account.Tag, 0, len(presentIncludedTags))
		for _, val := range presentIncludedTags {
			presentIncludedTag := val.(map[string]interface{})
			parsedIncludedTags = append(parsedIncludedTags, account.Tag{
				Key:   presentIncludedTag["key"].(string),
				Value: presentIncludedTag["value"].(string),
			})
		}
		parsedCloudScanRule.AgentlessScanSpec.IncludedTags = parsedIncludedTags
	}

	if val, ok := d.GetOk("aws_region_type"); ok {
		parsedCloudScanRule.AwsRegionType = val.(string)
	}

	return parsedCloudScanRule, nil
}

func ServerlessScanSpecToSchema(d *account.ServerLessScanSpec) []interface{} {
	ans := make([]interface{}, 0, 1)
	serverlessScanSpec := make(map[string]interface{})
	serverlessScanSpec["enabled"] = d.Enabled
	serverlessScanSpec["cap"] = d.Cap
	serverlessScanSpec["scan_all_versions"] = d.ScanAllVersions
	serverlessScanSpec["scan_layers"] = d.ScanLayers
	ans = append(ans, serverlessScanSpec)
	return ans
}

func AgentlessScanSpecToSchema(d *account.AgentlessScanSpec) []interface{} {
	ans := make([]interface{}, 0, 1)
	agentlessScanSpec := make(map[string]interface{})
	agentlessScanSpec["enabled"] = d.Enabled
	agentlessScanSpec["hub_account"] = d.HubAccount
	agentlessScanSpec["console_addr"] = d.ConsoleAddr
	agentlessScanSpec["scan_non_running"] = d.ScanNonRunning
	agentlessScanSpec["proxy_address"] = d.ProxyAddress
	agentlessScanSpec["proxy_ca"] = d.ProxyCA
	agentlessScanSpec["skip_permissions_check"] = d.SkipPermissionsCheck
	agentlessScanSpec["auto_scale"] = d.AutoScale
	agentlessScanSpec["scanners"] = d.Scanners
	agentlessScanSpec["security_group"] = d.SecurityGroup
	agentlessScanSpec["subnet"] = d.SubNet
	agentlessScanSpec["regions"] = d.Regions
	ans = append(ans, agentlessScanSpec)
	return ans
}

func CloudAccountCredentialToSchema(d auth.Credential) []interface{} {
	ans := make([]interface{}, 0, 1)
	credential := make(map[string]interface{})
	credential["id"] = d.Id
	credential["type"] = d.Type
	credential["account_id"] = d.AccountID
	credential["account_guid"] = d.AccountGUID
	credential["secret"] = CredentialSecretToSchema(d.Secret)
	credential["api_token"] = CredentialSecretToSchema(d.ApiToken)
	credential["use_aws_role"] = d.UseAWSRole
	ans = append(ans, credential)
	return ans
}
