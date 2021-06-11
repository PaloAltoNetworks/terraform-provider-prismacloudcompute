package prismacloudcompute

import (
	"fmt"

	pc "github.com/paloaltonetworks/prisma-cloud-compute-go"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The API URL without the leading protocol",
				DefaultFunc: schema.EnvDefaultFunc("PRISMACLOUDCOMPUTE_URL", nil),
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User's registered email address",
				DefaultFunc: schema.EnvDefaultFunc("PRISMACLOUDCOMPUTE_USERNAME", nil),
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Secret key",
				DefaultFunc: schema.EnvDefaultFunc("PRISMACLOUDCOMPUTE_PASSWORD", nil),
				Sensitive:   true,
			},
			"port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "If the port is non-standard for the protocol, the port number to use",
				DefaultFunc: schema.EnvDefaultFunc("PRISMACLOUD_PORT", nil),
			},
			"endpoint_prefix": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Prefix for the API endpoints",
				Default:     "/api/v1/",
			},
			"protocol": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "The protocol (https or http)",
				DefaultFunc:  schema.EnvDefaultFunc("PRISMACLOUDCOMPUTE_PROTOCOL", nil),
				ValidateFunc: validation.StringInSlice([]string{"", "https", "http"}, false),
			},
			"timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The timeout in seconds for all communications with Prisma Cloud",
				Default:     90,
			},
			"skip_ssl_cert_verification": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Skip SSL certificate verification",
				Default:     true,
			},
			"logging": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeBool,
				},
				Optional:    true,
				Description: "Logging options for the API connection",
			},
			"disable_reconnect": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Disable reconnecting on JWT expiration",
			},
			"json_web_token": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "JSON web token to use",
				DefaultFunc: schema.EnvDefaultFunc("PRISMACLOUDCOMPUTE_JSON_WEB_TOKEN", nil),
				Sensitive:   true,
			},
			"json_config_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Retrieve the provider configuration from this JSON file",
				DefaultFunc: schema.EnvDefaultFunc("PRISMACLOUDCOMPUTE_JSON_CONFIG_FILE", nil),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"prismacloudcompute_collection": resourceCollection(),
			//			"prismacloudcompute_collections_id":                   resourceCollectionsId(),
			/*			"prismacloudcompute_settings_registry":                resourceSettingsRegistry(),
						"prismacloudcompute_policies_runtime_container":       resourcePoliciesRuntimeContainer(),
						"prismacloudcompute_policies_vulnerability_images":    resourcePoliciesVulnerabilityImages(),
						"prismacloudcompute_policies_vulnerability_ci_images": resourcePoliciesVulnerabilityCiImages(),
						"prismacloudcompute_policies_compliance_container":    resourcePoliciesComplianceContainer(),
						"prismacloudcompute_policies_compliance_ci_images":    resourcePoliciesComplianceCiImages(),
						"prismacloudcompute_policies_runtime_host":            resourcePoliciesRuntimeHost(),
						"prismacloudcompute_policies_vulnerability_host":      resourcePoliciesVulnerabilityHost(),
						"prismacloudcompute_policies_compliance_host":         resourcePoliciesComplianceHost(),
						"prismacloudcompute_users":                            resourceUsers(),
						"prismacloudcompute_users_id":                         resourceUsersId(),
						"prismacloudcompute_groups":                           resourceGroups(),
						"prismacloudcompute_groups_id":                        resourceGroupsId(),
						"prismacloudcompute_settings_logon":                   resourceSettingsLogon(),*/
		},

		DataSourcesMap: map[string]*schema.Resource{
			"prismacloudcompute_collections": dataSourceCollections(),
			//			"prismacloudcompute_collections_id":                   dataSourceCollectionsId(),
			/*			"prismacloudcompute_settings_registry":                dataSourceSettingsRegistry(),
						"prismacloudcompute_policies_runtime_container":       dataSourcePoliciesRuntimeContainer(),
						"prismacloudcompute_policies_vulnerability_images":    dataSourcePoliciesVulnerabilityImages(),
						"prismacloudcompute_policies_vulnerability_ci_images": dataSourcePoliciesVulnerabilityCiImages(),
						"prismacloudcompute_policies_compliance_container":    dataSourcePoliciesComplianceContainer(),
						"prismacloudcompute_policies_compliance_ci_images":    dataSourcePoliciesComplianceCiImages(),
						"prismacloudcompute_policies_runtime_host":            dataSourcePoliciesRuntimeHost(),
						"prismacloudcompute_policies_vulnerability_host":      dataSourcePoliciesVulnerabilityHost(),
						"prismacloudcompute_policies_compliance_host":         dataSourcePoliciesComplianceHost(),
						"prismacloudcompute_users":                            dataSourceUsers(),
						"prismacloudcompute_users_id":                         dataSourceUsersId(),
						"prismacloudcompute_groups":                           dataSourceGroups(),
						"prismacloudcompute_groups_id":                        dataSourceGroupsId(),
						"prismacloudcompute_settings_logon":                   dataSourceSettingsLogon(),*/
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	/*
	   An int in Terraform is a Go "int", which can be either 32 or 64bit
	   depending on what the underlying OS is.  A Terraform "schema.TypeInt" is
	   also a Go "int".

	   Timestamps returned Prisma Cloud are 64bit ints.  In addition to this,
	   a Go time.Duration is an int64.

	   Thus, require that the OS is 64bit or bail.

	   If this becomes a problem in the future, then the fix is to go through all
	   resources and anywhere where a timestamp is returned, that needs to be either
	   a float64 or a string, either of which will require lots of casting.
	*/
	is64Bit := uint64(^uintptr(0)) == ^uint64(0)
	if !is64Bit {
		return nil, fmt.Errorf("This provider requires a 64bit OS")
	}

	logSetting := make(map[string]bool)
	logConfig := d.Get("logging").(map[string]interface{})
	for key := range logConfig {
		logSetting[key] = logConfig[key].(bool)
	}

	con := &pc.Client{
		Url:                     d.Get("url").(string),
		Username:                d.Get("username").(string),
		Password:                d.Get("password").(string),
		Port:                    d.Get("port").(int),
		Protocol:                d.Get("protocol").(string),
		Timeout:                 d.Get("timeout").(int),
		SkipSslCertVerification: d.Get("skip_ssl_cert_verification").(bool),
		DisableReconnect:        d.Get("disable_reconnect").(bool),
		JsonWebToken:            d.Get("json_web_token").(string),
		Logging:                 logSetting,
	}

	err := con.Initialize(d.Get("json_config_file").(string))
	return con, err
}
