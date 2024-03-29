package provider

import (
	"context"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/policy"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/convert"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePoliciesComplianceCiCoderepo() *schema.Resource {
	return &schema.Resource{
		CreateContext: createPolicyComplianceCiCoderepo,
		ReadContext:   readPolicyComplianceCiCoderepo,
		UpdateContext: updatePolicyComplianceCiCoderepo,
		DeleteContext: deletePolicyComplianceCiCoderepo,

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
									"block_threshold": {
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

func createPolicyComplianceCiCoderepo(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)
	parsedRules, err := convert.SchemaToComplianceCiCoderepoRules(d)
	if err != nil {
		return diag.Errorf("error creating %s policy: %s", policyTypeComplianceCiCoderepo, err)
	}

	parsedPolicy := policy.ComplianceCoderepoPolicy{
		Type:  policyTypeComplianceCiCoderepo,
		Rules: parsedRules,
	}

	if err := policy.UpdateComplianceCiCoderepo(*client, parsedPolicy); err != nil {
		return diag.Errorf("error creating %s policy: %s", policyTypeComplianceCiCoderepo, err)
	}

	d.SetId(policyTypeComplianceCiCoderepo)
	return readPolicyComplianceCiCoderepo(ctx, d, meta)
}

func readPolicyComplianceCiCoderepo(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)

	var diags diag.Diagnostics

	retrievedPolicy, err := policy.GetComplianceCiCoderepo(*client)
	if err != nil {
		return diag.Errorf("error reading %s policy: %s", policyTypeComplianceCiCoderepo, err)
	}

	if err := d.Set("rule", convert.ComplianceCoderepoCiRulesToSchema(retrievedPolicy.Rules)); err != nil {
		return diag.Errorf("error reading %s policy: %s", policyTypeComplianceCiCoderepo, err)
	}

	return diags
}

func updatePolicyComplianceCiCoderepo(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)
	parsedRules, err := convert.SchemaToComplianceCiCoderepoRules(d)
	if err != nil {
		return diag.Errorf("error updating %s policy: %s", policyTypeComplianceCiCoderepo, err)
	}

	parsedPolicy := policy.ComplianceCoderepoPolicy{
		Type:  policyTypeComplianceCiCoderepo,
		Rules: parsedRules,
	}

	if err := policy.UpdateComplianceCiCoderepo(*client, parsedPolicy); err != nil {
		return diag.Errorf("error updating %s policy: %s", policyTypeComplianceCiCoderepo, err)
	}

	return readPolicyComplianceCiCoderepo(ctx, d, meta)
}

func deletePolicyComplianceCiCoderepo(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// TODO: reset to default policy
	var diags diag.Diagnostics
	return diags
}
