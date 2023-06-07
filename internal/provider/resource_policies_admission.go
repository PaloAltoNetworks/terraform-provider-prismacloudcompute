package provider

import (
	"context"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/policy"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/convert"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePoliciesAdmission() *schema.Resource {
	return &schema.Resource{
		CreateContext: createPolicyAdmission,
		ReadContext:   readPolicyAdmission,
		UpdateContext: updatePolicyAdmission,
		DeleteContext: deletePolicyAdmission,

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
						"description": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Free-form text field.",
						},
						"disabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether or not to disable the rule.",
						},
						"effect": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The effect to be used. Can be set to 'allow', 'block' or 'alert'.",
						},
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Unique name of the rule.",
						},
						"script": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Policy script in Rego syntax.",
						},
					},
				},
			},
		},
	}
}

func createPolicyAdmission(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)
	parsedRules, err := convert.SchemaToAdmissionRules(d)

	if err != nil {
		return diag.Errorf("error creating %s policy: %s", policyTypeAdmission, err)
	}

	parsedPolicy := policy.AdmissionPolicy{
		Id:    policyTypeAdmission,
		Rules: parsedRules,
	}

	if err := policy.UpdateAdmission(*client, parsedPolicy); err != nil {
		return diag.Errorf("error creating %s policy: %s", policyTypeAdmission, err)
	}

	d.SetId(policyTypeAdmission)
	return readPolicyAdmission(ctx, d, meta)
}

func readPolicyAdmission(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)

	var diags diag.Diagnostics

	retrievedPolicy, err := policy.GetAdmission(*client)
	if err != nil {
		return diag.Errorf("error reading %s policy: %s", policyTypeAdmission, err)
	}

	if err := d.Set("rule", convert.AdmissionRulesToSchema(retrievedPolicy.Rules)); err != nil {
		return diag.Errorf("error reading %s policy: %s", policyTypeAdmission, err)
	}

	return diags
}

func updatePolicyAdmission(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)
	parsedRules, err := convert.SchemaToAdmissionRules(d)
	if err != nil {
		return diag.Errorf("error updating %s policy: %s", policyTypeAdmission, err)
	}

	parsedPolicy := policy.AdmissionPolicy{
		Id:    policyTypeAdmission,
		Rules: parsedRules,
	}

	if err := policy.UpdateAdmission(*client, parsedPolicy); err != nil {
		return diag.Errorf("error updating %s policy: %s", policyTypeAdmission, err)
	}

	return readPolicyAdmission(ctx, d, meta)
}

func deletePolicyAdmission(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// TODO: reset to default policy
	var diags diag.Diagnostics
	return diags
}
