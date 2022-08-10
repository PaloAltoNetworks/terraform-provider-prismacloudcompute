package provider

import (
	"fmt"
	"log"
	"strings"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/alertprofile"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/convert"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAlertprofile() *schema.Resource {
	return &schema.Resource{
		Create: createAlertprofile,
		Read:   readAlertprofile,
		Update: updateAlertprofile,
		Delete: deleteAlertprofile,

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
			"alert_profile_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Alert Profile type",
				ValidateDiagFunc: func(v interface{}, p cty.Path) diag.Diagnostics {
					value := v.(string)
					expected := []string{"webhook"}

					var diags diag.Diagnostics
					if !stringInSlice(expected, value) {
						diag := diag.Diagnostic{
							Severity: diag.Error,
							Summary:  "invalid/unsupported alert_profile_type",
							Detail:   fmt.Sprintf("alert_profile_type %q is invalid or unsupported. Supported types: %q", value, strings.Join(expected, ",")),
						}
						diags = append(diags, diag)
					}
					return diags
				},
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enabled",
				Default:     true,
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
			"alert_profile_config": {
				Type:        schema.TypeList,
				Required:    true,
				MinItems:    1,
				MaxItems:    1,
				Description: "Alert Profile configuration, the values depend on the alert profile type",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"prisma_cloud_integration_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "ID of the Prisma Cloud Integration",
						},
						"webhook_url": {
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
			"alert_triggers": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Policy configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Access (Docker)",
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
						"app_embedded_defender_runtime": {
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
						"cloud_native_network_firewall": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Cloud Native Network Firewall (CNNF)",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Required: true,
									},
								},
							},
						},
						"container_and_image_compliance": {
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
						"defender_health": {
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
								},
							},
						},
						"host_compliance": {
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
						"host_vulnerabilities": {
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
						"image_vulnerabilities": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Image vulnerabilities (registry and deployed)",
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
						"incidents": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "incidents",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Required: true,
									},
								},
							},
						},
						"kubernetes_audits": {
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
						"waas_firewall_app_embedded_defender": {
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
						"waas_firewall_container": {
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
						"waas_firewall_host": {
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
						"waas_firewall_serverless": {
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
								},
							},
						},
					},
				},
			},
		},
	}
}

func createAlertprofile(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*api.Client)
	parsedAlertprofile, err := convert.SchemaToAlertprofile(d)
	if err != nil {
		return fmt.Errorf("failed to create Alert Profile '%+v': %s", parsedAlertprofile, err)
	}
	if err := alertprofile.CreateAlertprofile(*client, parsedAlertprofile); err != nil {
		return fmt.Errorf("error creating alertprofile '%+v': %s", parsedAlertprofile, err)
	}

	d.SetId(parsedAlertprofile.Name)
	return readAlertprofile(d, meta)
}

func readAlertprofile(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*api.Client)
	retrievedAlertprofile, err := alertprofile.GetAlertprofile(*client, d.Id())
	if err != nil {
		return fmt.Errorf("error reading alertprofile: %s", err)
	}
	d.Set("name", retrievedAlertprofile.Name)
	d.Set("owner", retrievedAlertprofile.Owner)

	config := make(map[string]interface{})
	config["prisma_cloud_integration_id"] = retrievedAlertprofile.IntegrationID
	config["enable_immediate_vulnerabilities_alerts"] = retrievedAlertprofile.VulnerabilityImmediateAlertsEnabled

	if retrievedAlertprofile.Webhook.Enabled {
		d.Set("enabled", true)
		d.Set("alert_profile_type", "webhook")
		config["webhook_url"] = retrievedAlertprofile.Webhook.Url
		config["credential_id"] = retrievedAlertprofile.Webhook.CredentialId
		config["custom_ca"] = retrievedAlertprofile.Webhook.CaCert
		config["custom_json"] = retrievedAlertprofile.Webhook.Json
	}

	if err = d.Set("alert_profile_config", []interface{}{config}); err != nil {
		log.Printf("[WARN] Error setting 'alert_profile_config' for %s: %s", d.Id(), err)
	}

	alertTriggerPolicies := convert.AlertProfilePoliciesToSchema(&retrievedAlertprofile.Policy)

	if err = d.Set("alert_triggers", []interface{}{alertTriggerPolicies}); err != nil {
		log.Printf("[WARN] Error setting 'alert_triggers' for %s: %s", d.Id(), err)
	}

	return nil
}

func updateAlertprofile(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*api.Client)

	parsedAlertprofile, err := convert.SchemaToAlertprofile(d)
	if err != nil {
		return fmt.Errorf("failed to update Alert Profile '%+v': %s", parsedAlertprofile, err)
	}

	if err := alertprofile.UpdateAlertprofile(*client, parsedAlertprofile); err != nil {
		return fmt.Errorf("error updating alertprofile '%s': %s", d.Id(), err)
	}

	return readAlertprofile(d, meta)
}

func deleteAlertprofile(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*api.Client)
	if err := alertprofile.DeleteAlertprofile(*client, d.Id()); err != nil {
		return fmt.Errorf("error deleting alertprofile '%s': %s", d.Id(), err)
	}

	d.SetId("")
	return nil
}

func stringInSlice(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
