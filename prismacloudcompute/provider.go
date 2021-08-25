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
			//			"prismacloudcompute_settings_registry":                resourceSettingsRegistry(),
			"prismacloudcompute_policiesruntimecontainer":      resourcePoliciesRuntimeContainer(),
			"prismacloudcompute_policiesvulnerabilityimages":   resourcePoliciesVulnerabilityImages(),
			"prismacloudcompute_policiesvulnerabilityciimages": resourcePoliciesVulnerabilityCiImages(),
			"prismacloudcompute_policiescompliancecontainer":   resourcePoliciesComplianceContainer(),
			"prismacloudcompute_policiescomplianceciimages":    resourcePoliciesComplianceCiImages(),
			"prismacloudcompute_policiesruntimehost":      resourcePoliciesRuntimeHost(),
			"prismacloudcompute_policiesvulnerabilityhost":      resourcePoliciesVulnerabilityHost(),
			"prismacloudcompute_policiescompliancehost":         resourcePoliciesComplianceHost(),
/*															"prismacloudcompute_users":                            resourceUsers(),
															"prismacloudcompute_usersid":                         resourceUsersId(),
															"prismacloudcompute_groups":                           resourceGroups(),
															"prismacloudcompute_groupsid":                        resourceGroupsId(),
															"prismacloudcompute_settingslogon":                   resourceSettingsLogon(),*/
		},

		DataSourcesMap: map[string]*schema.Resource{
			"prismacloudcompute_collections": dataSourceCollections(),
			//			"prismacloudcompute_settingsregistry":                dataSourceSettingsRegistry(),
			"prismacloudcompute_policiesruntimecontainer":      dataSourcePoliciesRuntimeContainer(),
			"prismacloudcompute_policiesvulnerabilityimages":   dataSourcePoliciesVulnerabilityImages(),
			"prismacloudcompute_policiesvulnerabilityciimages": dataSourcePoliciesVulnerabilityCiImages(),
			"prismacloudcompute_policiescompliancecontainer":   dataSourcePoliciesComplianceContainer(),
			"prismacloudcompute_policiescomplianceciimages":    dataSourcePoliciesComplianceCiImages(),
			"prismacloudcompute_policiesruntimehost":            dataSourcePoliciesRuntimeHost(),
			"prismacloudcompute_policiesvulnerabilityhost":      dataSourcePoliciesVulnerabilityHost(),
			"prismacloudcompute_policiescompliancehost":         dataSourcePoliciesComplianceHost(),
/*			"prismacloudcompute_users":                            dataSourceUsers(),
			"prismacloudcompute_usersid":                         dataSourceUsersId(),
			"prismacloudcompute_groups":                           dataSourceGroups(),
			"prismacloudcompute_groupsid":                        dataSourceGroupsId(),
			"prismacloudcompute_settingslogon":                   dataSourceSettingsLogon(),*/
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
