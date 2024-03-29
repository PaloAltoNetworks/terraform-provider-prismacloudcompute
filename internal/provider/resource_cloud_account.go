package provider

import (
	"context"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/account"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/auth"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/convert"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudAccount() *schema.Resource {
	return &schema.Resource{
		CreateContext: createCloudAccount,
		ReadContext:   readCloudAccount,
		UpdateContext: updateCloudAccount,
		DeleteContext: deleteCloudAccount,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"credential_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Credential ID",
			},
			"credential": {
				Type:        schema.TypeList,
				Optional:    true,
				MinItems:    1,
				MaxItems:    1,
				Description: "Serverless Scan Configuration",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Description: "The ID of the credential.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"account_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Account identifier (username, access key, etc.).",
						},
						"account_guid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
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
						"external": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Indicates if the credential is external (true) or not (false).",
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
						"use_sts_regional_endpoint": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Indicates whether to use the regional STS endpoint for an STS session.",
						},
					},
				},
			},
			"aws_region_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "AWS Region Type",
			},
			"discovery_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enables cloud discovery, which will discover all workloads in the account and their scan status.",
			},
			"serverless_radar_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enables the discovery of serverless functions.",
			},
			"vm_tags_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enables the discovery of tags on VMs in AWS accounts.",
			},
			"discover_all_function_versions": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Cloud Discovery Enabled",
			},
			"serverless_radar_cap": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Serverless Radar Cap",
			},
			"agentless_scan_spec": {
				Type:        schema.TypeList,
				Optional:    true,
				MinItems:    1,
				MaxItems:    1,
				Description: "Serverless Scan Configuration",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "",
						},
						"hub_account": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Indicates whether the Prisma Cloud scanner will be centralized in the hub account and scan the target accounts from there (enabled) or the actual scanning will occur within each account that is being scanned (false).",
						},
						"console_addr": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Console URL.",
						},
						"scan_non_running": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Scan non running hosts.",
						},
						"proxy_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Example: http://proxyserver.company.com:8081",
						},
						"proxy_ca": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Proxy CA certificate. Required when using TLS intercept proxies.",
						},
						"skip_permissions_check": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "When enabled, Prisma Cloud will scan this account even if there are missing permissions.",
						},
						"auto_scale": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "When enabled, Prisma Cloud automatically spins up multiple scanners in the environment to parallel scan for faster results.",
						},
						"scanners": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Limit on the number of scanners that Prisma Cloud can spin up at any given time.",
						},
						"security_group": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Security group name. Should be identical and unique across all regions.",
						},
						"subnet": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Subnet name. Should be identical and unique across all regions. Note: if the subnet allows auto-assignment of public IPs, a public IP will be attached to the scanner instance.",
						},
						"regions": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"custom_tags": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "These tags will be applied to resources created by Prisma Cloud in the Agentless scan process.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"value": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"included_tags": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"value": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
					},
				},
			},
			"serverless_scan_spec": {
				Type:        schema.TypeList,
				Optional:    true,
				MinItems:    1,
				MaxItems:    1,
				Description: "Serverless Scan Configuration",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "",
						},
						"cap": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The number of most recently modified functions to scan, on a per-scope basis. For example, if there are 100 functions in scope, and you set this value to 50, Prisma Cloud will only scan the fifty most recently modified functions. To scan all functions in scope, set this to 0.",
						},
						"scan_all_versions": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Indicates whether Prisma Cloud will scan all versions (enabled) or only the latest versions (false) of serverless functions.",
						},
						"scan_layers": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Indicates whether or not Prisma Cloud will scan Lambda layers.",
						},
					},
				},
			},
		},
	}
}

func createCloudAccount(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)

	parsedCredential, err := convert.SchemaToCloudAccountCredential(d)
	if err != nil {
		return diag.Errorf("failed to create cloud account credential '%+v': %s", parsedCredential, err)
	}
	if err := auth.UpdateCredential(*client, parsedCredential); err != nil {
		return diag.Errorf("error creating cloud account credential '%+v': %s", parsedCredential, err)
	}

	var scanRules []account.CloudScanRule
	parsedCloudScanRule, err := convert.SchemaToCloudScanRule(d)
	if err != nil {
		return diag.Errorf("failed to create cloud scan rule '%+v': %s", parsedCloudScanRule, err)
	}
	scanRules = append(scanRules, parsedCloudScanRule)
	if err := account.UpdateCloudScanRule(*client, scanRules); err != nil {
		return diag.Errorf("error creating cloud account '%+v': %s", parsedCloudScanRule, err)
	}

	d.SetId(parsedCredential.Id)

	return readCloudAccount(ctx, d, meta)
}

func readCloudAccount(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)

	var diags diag.Diagnostics

	retrievedCloudScanRule, err := account.GetCloudScanRule(*client, d.Id())
	if err != nil {
		return diag.Errorf("error reading cloud account: %s", err)
	}

	d.Set("credential_id", retrievedCloudScanRule.CredentialId)

	if err := d.Set("credential", convert.CloudAccountCredentialToSchema(retrievedCloudScanRule.Credential)); err != nil {
		return diag.Errorf("error reading cloud account credential: %s", err)
	}

	d.Set("discovery_enabled", retrievedCloudScanRule.DiscoveryEnabled)
	d.Set("serverless_radar_enabled", retrievedCloudScanRule.ServerlessRadarEnabled)
	d.Set("vm_tags_enabled", retrievedCloudScanRule.VmTagsEnabled)
	d.Set("discover_all_function_versions", retrievedCloudScanRule.DiscoverAllFunctionVersions)
	d.Set("serverless_radar_cap", retrievedCloudScanRule.ServerlessRadarCap)
	d.Set("aws_region_type", retrievedCloudScanRule.AwsRegionType)

	if err := d.Set("serverless_scan_spec", convert.ServerlessScanSpecToSchema(&retrievedCloudScanRule.ServerlessScanSpec)); err != nil {
		return diag.Errorf("error reading serverless scan spec: %s", err)
	}

	if err := d.Set("agentless_scan_spec", convert.AgentlessScanSpecToSchema(&retrievedCloudScanRule.AgentlessScanSpec)); err != nil {
		return diag.Errorf("error reading agentless scan spec: %s", err)
	}

	return diags
}

func updateCloudAccount(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)

	parsedCloudAccountCredential, err := convert.SchemaToCloudAccountCredential(d)
	if err != nil {
		return diag.Errorf("failed to parse cloud account credential '%+v': %s", parsedCloudAccountCredential, err)
	}
	if err := auth.UpdateCredential(*client, parsedCloudAccountCredential); err != nil {
		return diag.Errorf("error updating cloud account credential '%s': %s", d.Id(), err)
	}

	var scanRules []account.CloudScanRule
	parsedCloudScanRule, err := convert.SchemaToCloudScanRule(d)
	if err != nil {
		return diag.Errorf("failed to parse cloud scan rule '%+v': %s", parsedCloudScanRule, err)
	}
	scanRules = append(scanRules, parsedCloudScanRule)

	if err := account.UpdateCloudScanRule(*client, scanRules); err != nil {
		return diag.Errorf("error updating cloud scan rule '%s': %s", d.Id(), err)
	}

	return readCloudAccount(ctx, d, meta)
}

func deleteCloudAccount(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)

	var diags diag.Diagnostics

	if err := account.DeleteCloudScanRule(*client, d.Id()); err != nil {
		return diag.Errorf("error deleting credential: %s", err)
	}

	if err := auth.DeleteCredential(*client, d.Id()); err != nil {
		return diag.Errorf("error deleting credential: %s", err)
	}

	d.SetId("")

	return diags
}
