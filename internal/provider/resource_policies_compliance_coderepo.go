package provider

import (
	"fmt"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/policy"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/convert"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePoliciesComplianceCoderepo() *schema.Resource {
	return &schema.Resource{
		Create: createPolicyComplianceCoderepo,
		Read:   readPolicyComplianceCoderepo,
		Update: updatePolicyComplianceCoderepo,
		Delete: deletePolicyComplianceCoderepo,

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

func createPolicyComplianceCoderepo(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*api.Client)
	parsedRules, err := convert.SchemaToComplianceCoderepoRules(d)
	if err != nil {
		return fmt.Errorf("error creating %s policy: %s", policyTypeComplianceCoderepo, err)
	}

	parsedPolicy := policy.ComplianceCoderepoPolicy{
		Type:  policyTypeComplianceCoderepo,
		Rules: parsedRules,
	}

	if err := policy.UpdateComplianceCoderepo(*client, parsedPolicy); err != nil {
		return fmt.Errorf("error creating %s policy: %s", policyTypeComplianceCoderepo, err)
	}

	d.SetId(policyTypeComplianceCoderepo)
	return readPolicyComplianceCoderepo(d, meta)
}

func readPolicyComplianceCoderepo(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*api.Client)
	retrievedPolicy, err := policy.GetComplianceCoderepo(*client)
	if err != nil {
		return fmt.Errorf("error reading %s policy: %s", policyTypeComplianceCoderepo, err)
	}

	if err := d.Set("rule", convert.ComplianceCoderepoRulesToSchema(retrievedPolicy.Rules)); err != nil {
		return fmt.Errorf("error reading %s policy: %s", policyTypeComplianceCoderepo, err)
	}

	return nil
}

func updatePolicyComplianceCoderepo(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*api.Client)
	parsedRules, err := convert.SchemaToComplianceCoderepoRules(d)
	if err != nil {
		return fmt.Errorf("error updating %s policy: %s", policyTypeComplianceCoderepo, err)
	}

	parsedPolicy := policy.ComplianceCoderepoPolicy{
		Type:  policyTypeComplianceCoderepo,
		Rules: parsedRules,
	}

	if err := policy.UpdateComplianceCoderepo(*client, parsedPolicy); err != nil {
		return fmt.Errorf("error updating %s policy: %s", policyTypeComplianceCoderepo, err)
	}

	return readPolicyComplianceCoderepo(d, meta)
}

func deletePolicyComplianceCoderepo(d *schema.ResourceData, meta interface{}) error {
	// TODO: reset to default policy
	return nil
}
