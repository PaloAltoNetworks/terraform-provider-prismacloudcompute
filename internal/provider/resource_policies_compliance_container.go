package provider

import (
	"context"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/policy"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/convert"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePoliciesComplianceContainer() *schema.Resource {
	return &schema.Resource{
		CreateContext: createPolicyComplianceContainer,
		ReadContext:   readPolicyComplianceContainer,
		UpdateContext: updatePolicyComplianceContainer,
		DeleteContext: deletePolicyComplianceContainer,

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
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"block_message": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Message to display for blocked requests.",
						},
						"collections": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Collections used to scope the rule.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"compliance_check": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Compliance checks. Omitted checks are ignored.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"block": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether or not to block if this check is failed. Setting to 'false' will only alert if the check is failed.",
									},
									"id": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Compliance check number.",
									},
								},
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
							Description: "The effect of the rule. Can be set to 'ignore', 'alert', 'block', or 'alert, block'.",
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
						"show_passed_checks": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether or not to report both failed and passed compliance checks.",
						},
						"verbose": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether or not to provide verbose output for blocked requests.",
						},
					},
				},
			},
		},
	}
}

func createPolicyComplianceContainer(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)
	parsedRules, err := convert.SchemaToComplianceDeployedRules(d)
	if err != nil {
		return diag.Errorf("error creating %s policy: %s", policyTypeComplianceContainer, err)
	}

	parsedPolicy := policy.CompliancePolicy{
		Type:  policyTypeComplianceContainer,
		Rules: parsedRules,
	}

	if err := policy.UpdateComplianceContainer(*client, parsedPolicy); err != nil {
		return diag.Errorf("error creating %s policy: %s", policyTypeComplianceContainer, err)
	}

	d.SetId(policyTypeComplianceContainer)
	return readPolicyComplianceContainer(ctx, d, meta)
}

func readPolicyComplianceContainer(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)

	var diags diag.Diagnostics

	retrievedPolicy, err := policy.GetComplianceContainer(*client)
	if err != nil {
		return diag.Errorf("error reading %s policy: %s", policyTypeComplianceContainer, err)
	}

	if err := d.Set("rule", convert.ComplianceDeployedRulesToSchema(retrievedPolicy.Rules)); err != nil {
		return diag.Errorf("error reading %s policy: %s", policyTypeComplianceContainer, err)
	}
	return diags
}

func updatePolicyComplianceContainer(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)
	parsedRules, err := convert.SchemaToComplianceDeployedRules(d)
	if err != nil {
		return diag.Errorf("error updating %s policy: %s", policyTypeComplianceContainer, err)
	}

	parsedPolicy := policy.CompliancePolicy{
		Type:  policyTypeComplianceContainer,
		Rules: parsedRules,
	}

	if err := policy.UpdateComplianceContainer(*client, parsedPolicy); err != nil {
		return diag.Errorf("error updating %s policy: %s", policyTypeComplianceContainer, err)
	}

	return readPolicyComplianceContainer(ctx, d, meta)
}

func deletePolicyComplianceContainer(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// TODO: reset to default policy
	var diags diag.Diagnostics
	return diags
}
