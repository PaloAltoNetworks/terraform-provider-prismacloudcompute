package prismacloudcompute

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/paloaltonetworks/prisma-cloud-compute-go/pcc"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider returns a terraform.ResourceProvider.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"console_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Prisma Cloud Compute Console URL",
				DefaultFunc: schema.EnvDefaultFunc("PRISMACLOUDCOMPUTE_CONSOLE_URL", nil),
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
			"prismacloudcompute_collection":                    resourceCollection(),
			"prismacloudcompute_ci_image_compliance_policy":    resourcePoliciesComplianceCiImage(),
			"prismacloudcompute_container_compliance_policy":   resourcePoliciesComplianceContainer(),
			"prismacloudcompute_host_compliance_policy":        resourcePoliciesComplianceHost(),
			"prismacloudcompute_container_runtime_policy":      resourcePoliciesRuntimeContainer(),
			"prismacloudcompute_host_runtime_policy":           resourcePoliciesRuntimeHost(),
			"prismacloudcompute_ci_image_vulnerability_policy": resourcePoliciesVulnerabilityCiImage(),
			"prismacloudcompute_host_vulnerability_policy":     resourcePoliciesVulnerabilityHost(),
			"prismacloudcompute_image_vulnerability_policy":    resourcePoliciesVulnerabilityImage(),
			"prismacloudcompute_settings_registry":             resourceRegistry(),
			"prismacloudcompute_users":                         resourceUsers(),
			"prismacloudcompute_groups":                        resourceGroups(),
			"prismacloudcompute_rbac_roles":                    resourceRbacRoles(),
			"prismacloudcompute_credentials":                   resourceCredentials(),
		},

		DataSourcesMap: map[string]*schema.Resource{},

		ConfigureFunc: configure,
	}
}

func configure(d *schema.ResourceData) (interface{}, error) {
	var config pcc.Credentials
	if val, ok := d.GetOk("config_file"); ok {
		configFile, err := os.Open(val.(string))
		if err != nil {
			fmt.Printf("error opening config file: %v", err)
		}
		defer configFile.Close()

		fileContent, err := ioutil.ReadAll(configFile)
		if err != nil {
			fmt.Printf("error reading config file: %v", err)
			return nil, err
		}
		if err := json.Unmarshal(fileContent, &config); err != nil {
			fmt.Printf("error unmarshalling config file: %v", err)
			return nil, err
		}
	}

	if val, ok := d.GetOk("console_url"); ok {
		config.ConsoleURL = val.(string)
	}
	if val, ok := d.GetOk("username"); ok {
		config.Username = val.(string)
	}
	if val, ok := d.GetOk("password"); ok {
		config.ConsoleURL = val.(string)
	}
	if val, ok := d.GetOk("skip_cert_verification"); ok {
		config.SkipCertVerification = val.(bool)
	}

	client, err := pcc.APIClient(config.ConsoleURL, config.Username, config.Password, config.SkipCertVerification)
	if err != nil {
		fmt.Printf("failed creating API client")
		return nil, err
	}

	return client, err
}
