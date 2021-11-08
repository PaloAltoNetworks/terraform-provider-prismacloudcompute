package prismacloudcompute

import (
	"fmt"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/prismacloudcompute/convert"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/pcc"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/policy"
)

func resourcePoliciesComplianceCiImage() *schema.Resource {
	return &schema.Resource{
		Create: createPolicyComplianceCiImage,
		Read:   readPolicyComplianceCiImage,
		Update: updatePolicyComplianceCiImage,
		Delete: deletePolicyComplianceCiImage,

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

func createPolicyComplianceCiImage(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	parsedRules, err := convert.SchemaToComplianceCiRules(d)
	if err != nil {
		return fmt.Errorf("error creating %s policy: %s", policyTypeComplianceCiImage, err)
	}

	parsedPolicy := policy.CompliancePolicy{
		Type:  policyTypeComplianceCiImage,
		Rules: parsedRules,
	}

	if err := policy.UpdateComplianceCiImage(*client, parsedPolicy); err != nil {
		return fmt.Errorf("error creating %s policy: %s", policyTypeComplianceCiImage, err)
	}

	d.SetId(policyTypeComplianceCiImage)
	return readPolicyComplianceCiImage(d, meta)
}

func readPolicyComplianceCiImage(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	retrievedPolicy, err := policy.GetComplianceCiImage(*client)
	if err != nil {
		return fmt.Errorf("error reading %s policy: %s", policyTypeComplianceCiImage, err)
	}

	if err := d.Set("rule", convert.ComplianceCiRulesToSchema(retrievedPolicy.Rules)); err != nil {
		return fmt.Errorf("error reading %s policy: %s", policyTypeComplianceCiImage, err)
	}
	return nil
}

func updatePolicyComplianceCiImage(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	parsedRules, err := convert.SchemaToComplianceCiRules(d)
	if err != nil {
		return fmt.Errorf("error updating %s policy: %s", policyTypeComplianceCiImage, err)
	}

	parsedPolicy := policy.CompliancePolicy{
		Type:  policyTypeComplianceCiImage,
		Rules: parsedRules,
	}

	if err := policy.UpdateComplianceCiImage(*client, parsedPolicy); err != nil {
		return fmt.Errorf("error updating %s policy: %s", policyTypeComplianceCiImage, err)
	}

	return readPolicyComplianceCiImage(d, meta)
}

func deletePolicyComplianceCiImage(d *schema.ResourceData, meta interface{}) error {
	// TODO: reset to default policy
	return nil
}
