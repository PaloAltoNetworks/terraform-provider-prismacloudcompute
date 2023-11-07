package provider

import (
	"context"
	"log"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/alertprofile"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/convert"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAlertprofile() *schema.Resource {
	return &schema.Resource{
		CreateContext: createAlertprofile,
		ReadContext:   readAlertprofile,
		UpdateContext: updateAlertprofile,
		DeleteContext: deleteAlertprofile,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Alert Profile ID",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Alert Profile name",
			},
			"enable_immediate_vulnerabilities_alerts": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable immediate vulnerabilities alerts",
			},
			"owner": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Owner",
			},
			"slack": {
				Type:        schema.TypeList,
				Optional:    true,
				MinItems:    1,
				MaxItems:    1,
				Description: "Alert Profile configuration, the values depend on the alert profile type",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"webhook_url": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Webhook URL",
						},
					},
				},
			},
			"webhook": {
				Type:        schema.TypeList,
				Optional:    true,
				MinItems:    1,
				MaxItems:    1,
				Description: "Alert Profile configuration, the values depend on the alert profile type",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"url": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Webhook URL",
						},
						"credential_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Credential ID",
						},
						"custom_ca": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Custom CA Cert",
						},
						"custom_json": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Custom JSON payload",
						},
					},
				},
			},
			"policy": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Policy configuration. Configure triggers for alerts.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"admission": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Admission audits",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"all_rules": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"rules": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"agentless_app_firewall": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "WAAS Firewall (Agentless)",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"all_rules": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"rules": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"app_embedded_app_firewall": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "WAAS Firewall (App-Embedded Defender)",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"all_rules": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"rules": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"app_embedded_runtime": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "App-Embedded Defender runtime",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"all_rules": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"rules": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"cloud_discovery": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Cloud discovery",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"all_rules": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"rules": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"code_repo_vulnerability": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Code repository vulnerabilities",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"all_rules": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"rules": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"container_app_firewall": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "WAAS Firewall (container)",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"all_rules": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"rules": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						// TODO: what is the UI element for this?
						"container_compliance": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"all_rules": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"rules": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"container_compliance_scan": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Container and image compliance",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"all_rules": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"rules": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"container_runtime": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Container runtime",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"all_rules": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"rules": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"container_vulnerability": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Deployed image vulnerabilities",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"all_rules": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"rules": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"defender": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Defender health",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"all_rules": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"rules": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"docker": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Access",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"all_rules": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"rules": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"host_app_firewall": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "WAAS Firewall (host)",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"all_rules": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"rules": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						// TODO: what UI element is this?
						"host_compliance": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"all_rules": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"rules": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"host_compliance_scan": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Host compliance",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"all_rules": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"rules": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"host_runtime": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Host runtime",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"all_rules": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"rules": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"host_vulnerability": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Host vulnerabilities",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"all_rules": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"rules": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"incident": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Incidents",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"all_rules": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"rules": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"kubernetes_audit": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Kubernetes audits",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"all_rules": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"rules": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"network_firewall": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Cloud Native Network Segmentation (CNNS)",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"all_rules": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"rules": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"registry_vulnerability": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Registry image vulnerabilities",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"all_rules": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"rules": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"serverless_app_firewall": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "WAAS Firewall (serverless)",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"all_rules": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"rules": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"serverless_runtime": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Serverless runtime",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"all_rules": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"rules": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"vm_compliance": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "VM images compliance",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"all_rules": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"rules": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"vm_vulnerability": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "VM images vulnerabilities",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"all_rules": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"rules": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"waas_health": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "WAAS health",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"all_rules": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"rules": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func createAlertprofile(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)
	parsedAlertprofile, err := convert.SchemaToAlertprofile(d)
	if err != nil {
		return diag.Errorf("failed to create Alert Profile '%+v': %s", parsedAlertprofile, err)
	}
	if err := alertprofile.CreateAlertprofile(*client, parsedAlertprofile); err != nil {
		return diag.Errorf("error creating alertprofile '%+v': %s", parsedAlertprofile, err)
	}

	d.SetId(parsedAlertprofile.Name)

	return readAlertprofile(ctx, d, meta)
}

func readAlertprofile(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)

	var diags diag.Diagnostics

	retrievedAlertProfile, err := alertprofile.GetAlertprofile(*client, d.Id())
	if err != nil {
		return diag.Errorf("error reading alertprofile: %s", err)
	}

	d.Set("name", retrievedAlertProfile.Name)
	d.Set("enable_immediate_vulnerabilities_alerts", retrievedAlertProfile.VulnerabilityImmediateAlertsEnabled)

	config := make(map[string]interface{})

	if retrievedAlertProfile.Webhook.Enabled {
		config["url"] = retrievedAlertProfile.Webhook.Url
		config["credential_id"] = retrievedAlertProfile.Webhook.CredentialId
		config["custom_ca"] = retrievedAlertProfile.Webhook.CaCert
		config["custom_json"] = retrievedAlertProfile.Webhook.Json
		if err = d.Set("webhook", []interface{}{config}); err != nil {
			log.Printf("[WARN] Error setting 'webhook' for %s: %s", d.Id(), err)
		}
	}

	if retrievedAlertProfile.Slack.Enabled {
		config["webhook_url"] = retrievedAlertProfile.Slack.WebhookUrl
		if err = d.Set("slack", []interface{}{config}); err != nil {
			log.Printf("[WARN] Error setting 'slack' for %s: %s", d.Id(), err)
		}
	}

	alertTriggerPolicies := convert.AlertProfilePoliciesToSchema(&retrievedAlertProfile.Policy)

	if err = d.Set("policy", []interface{}{alertTriggerPolicies}); err != nil {
		log.Printf("[WARN] Error setting 'policy' for %s: %s", d.Id(), err)
	}

	return diags
}

func updateAlertprofile(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)

	parsedAlertprofile, err := convert.SchemaToAlertprofile(d)
	if err != nil {
		return diag.Errorf("failed to update Alert Profile '%+v': %s", parsedAlertprofile, err)
	}

	if err := alertprofile.UpdateAlertprofile(*client, parsedAlertprofile); err != nil {
		return diag.Errorf("error updating alertprofile '%s': %s", d.Id(), err)
	}

	return readAlertprofile(ctx, d, meta)
}

func deleteAlertprofile(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)

	var diags diag.Diagnostics

	if err := alertprofile.DeleteAlertprofile(*client, d.Id()); err != nil {
		return diag.Errorf("error deleting alertprofile '%s': %s", d.Id(), err)
	}

	d.SetId("")

	return diags
}

func stringInSlice(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
