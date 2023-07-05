package provider

import (
	"context"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/policy"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/convert"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePoliciesComplianceCoderepo() *schema.Resource {
	return &schema.Resource{
		CreateContext: createPolicyComplianceCoderepo,
		ReadContext:   readPolicyComplianceCoderepo,
		UpdateContext: updatePolicyComplianceCoderepo,
		DeleteContext: deletePolicyComplianceCoderepo,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The ID of the policy.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"rule": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Rules that make up the policy.",
				MinItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"license": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "License compliance section.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"alert_threshold": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Threshold for generating license alerts.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enabled": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Whether or not to disable compliance alerts.",
												},
												"value": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "Minimum compliance severity to generate an alert. Can be set to 0=off, 1=low, 4=medium, 7=high, and 9=critical.",
												},
											},
										},
									},
									"critical": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of licenses with critical level of violation.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"high": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of licenses with high level of violation.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"medium": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of licenses with medium level of violation.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"low": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of licenses with low level of violation.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"collections": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Collections used to scope the rule.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"disabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether or not to disable the rule.",
						},
						"effect": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The effect of the rule. Can be set to 'ignore' or 'alert'.",
						},
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Unique name of the rule.",
						},
						"notes": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Free-form text field.",
						},
					},
				},
			},
		},
	}
}

func createPolicyComplianceCoderepo(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)
	parsedRules, err := convert.SchemaToComplianceCoderepoRules(d)
	if err != nil {
		return diag.Errorf("error creating %s policy: %s", policyTypeComplianceCoderepo, err)
	}

	parsedPolicy := policy.ComplianceCoderepoPolicy{
		Type:  policyTypeComplianceCoderepo,
		Rules: parsedRules,
	}

	if err := policy.UpdateComplianceCoderepo(*client, parsedPolicy); err != nil {
		return diag.Errorf("error creating %s policy: %s", policyTypeComplianceCoderepo, err)
	}

	d.SetId(policyTypeComplianceCoderepo)
	return readPolicyComplianceCoderepo(ctx, d, meta)
}

func readPolicyComplianceCoderepo(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)

	var diags diag.Diagnostics

	retrievedPolicy, err := policy.GetComplianceCoderepo(*client)
	if err != nil {
		return diag.Errorf("error reading %s policy: %s", policyTypeComplianceCoderepo, err)
	}

	if err := d.Set("rule", convert.ComplianceCoderepoRulesToSchema(retrievedPolicy.Rules)); err != nil {
		return diag.Errorf("error reading %s policy: %s", policyTypeComplianceCoderepo, err)
	}

	return diags
}

func updatePolicyComplianceCoderepo(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)
	parsedRules, err := convert.SchemaToComplianceCoderepoRules(d)
	if err != nil {
		return diag.Errorf("error updating %s policy: %s", policyTypeComplianceCoderepo, err)
	}

	parsedPolicy := policy.ComplianceCoderepoPolicy{
		Type:  policyTypeComplianceCoderepo,
		Rules: parsedRules,
	}

	if err := policy.UpdateComplianceCoderepo(*client, parsedPolicy); err != nil {
		return diag.Errorf("error updating %s policy: %s", policyTypeComplianceCoderepo, err)
	}

	return readPolicyComplianceCoderepo(ctx, d, meta)
}

func deletePolicyComplianceCoderepo(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// TODO: reset to default policy
	var diags diag.Diagnostics
	return diags
}
