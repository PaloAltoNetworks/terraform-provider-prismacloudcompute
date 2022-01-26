package provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"console_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Prisma Cloud Compute Console URL",
				DefaultFunc: schema.EnvDefaultFunc("PRISMACLOUDCOMPUTE_CONSOLE_URL", nil),
			},
			"project": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Prisma Cloud Compute project",
				DefaultFunc: schema.EnvDefaultFunc("PRISMACLOUDCOMPUTE_PROJECT", nil),
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Prisma Cloud Compute username",
				DefaultFunc: schema.EnvDefaultFunc("PRISMACLOUDCOMPUTE_USERNAME", nil),
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Prisma Cloud Compute password",
				DefaultFunc: schema.EnvDefaultFunc("PRISMACLOUDCOMPUTE_PASSWORD", nil),
				Sensitive:   true,
			},
			"skip_cert_verification": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether or not to skip certificate verification",
				Default:     true,
			},
			"config_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Configuration file in JSON format. See examples/creds.json",
				DefaultFunc: schema.EnvDefaultFunc("PRISMACLOUDCOMPUTE_CONFIG_FILE", nil),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"prismacloudcompute_collection":                       resourceCollection(),
			"prismacloudcompute_custom_rule":                      resourceCustomRule(),
			"prismacloudcompute_admission_policy":                 resourcePoliciesAdmission(),
			"prismacloudcompute_ci_image_compliance_policy":       resourcePoliciesComplianceCiImage(),
			"prismacloudcompute_container_compliance_policy":      resourcePoliciesComplianceContainer(),
			"prismacloudcompute_host_compliance_policy":           resourcePoliciesComplianceHost(),
			"prismacloudcompute_container_runtime_policy":         resourcePoliciesRuntimeContainer(),
			"prismacloudcompute_host_runtime_policy":              resourcePoliciesRuntimeHost(),
			"prismacloudcompute_ci_coderepo_vulnerability_policy": resourcePoliciesVulnerabilityCiCoderepo(),
			"prismacloudcompute_ci_image_vulnerability_policy":    resourcePoliciesVulnerabilityCiImage(),
			"prismacloudcompute_coderepo_vulnerability_policy":    resourcePoliciesVulnerabilityCoderepo(),
			"prismacloudcompute_host_vulnerability_policy":        resourcePoliciesVulnerabilityHost(),
			"prismacloudcompute_image_vulnerability_policy":       resourcePoliciesVulnerabilityImage(),
			"prismacloudcompute_registry_settings":                resourceRegistry(),
			"prismacloudcompute_user":                             resourceUsers(),
			"prismacloudcompute_group":                            resourceGroups(),
			"prismacloudcompute_role":                             resourceRbacRoles(),
			"prismacloudcompute_credential":                       resourceCredentials(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			"prismacloudcompute_custom_rule": dataSourceCustomRule(),
		},

		ConfigureFunc: configure,
	}
}

func configure(d *schema.ResourceData) (interface{}, error) {
	var config api.APIClientConfig
	if val, ok := d.GetOk("config_file"); ok {
		configFile, err := os.Open(val.(string))
		if err != nil {
			return nil, fmt.Errorf("error opening config file: %s", err)
		}
		defer configFile.Close()

		fileContent, err := ioutil.ReadAll(configFile)
		if err != nil {
			return nil, fmt.Errorf("error reading config file: %s", err)
		}
		if err := json.Unmarshal(fileContent, &config); err != nil {
			return nil, fmt.Errorf("error unmarshalling config file: %s", err)
		}
	}

	if val, ok := d.GetOk("console_url"); ok {
		config.ConsoleURL = val.(string)
	}
	if val, ok := d.GetOk("project"); ok {
		config.Project = val.(string)
	}
	if val, ok := d.GetOk("username"); ok {
		config.Username = val.(string)
	}
	if val, ok := d.GetOk("password"); ok {
		config.Password = val.(string)
	}
	if val, ok := d.GetOk("skip_cert_verification"); ok {
		config.SkipCertVerification = val.(bool)
	}

	client, err := api.APIClient(config)
	if err != nil {
		return nil, fmt.Errorf("error creating API client: %s", err)
	}

	return client, nil
}
