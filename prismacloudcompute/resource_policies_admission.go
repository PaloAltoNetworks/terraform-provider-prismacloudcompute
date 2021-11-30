package prismacloudcompute

import (
	"fmt"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/prismacloudcompute/convert"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/pcc"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/policy"
)

func resourcePoliciesAdmission() *schema.Resource {
	return &schema.Resource{
		Create: createPolicyAdmission,
		Read:   readPolicyAdmission,
		Update: updatePolicyAdmission,
		Delete: deletePolicyAdmission,

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
							Description: "Uhe effect to be used. Can be set to 'allow', 'block' or 'alert'.",
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

func createPolicyAdmission(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	parsedRules, err := convert.SchemaToAdmissionRules(d)

	if err != nil {
		return fmt.Errorf("error creating %s policy: %s", policyTypeAdmission, err)
	}

	parsedPolicy := policy.AdmissionPolicy{
		Id:    policyTypeAdmission,
		Rules: parsedRules,
	}

	if err := policy.UpdateAdmission(*client, parsedPolicy); err != nil {
		return fmt.Errorf("error creating %s policy: %s", policyTypeAdmission, err)
	}

	d.SetId(policyTypeAdmission)
	return readPolicyAdmission(d, meta)
}

func readPolicyAdmission(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	retrievedPolicy, err := policy.GetAdmission(*client)
	if err != nil {
		return fmt.Errorf("error reading %s policy: %s", policyTypeAdmission, err)
	}

	if err := d.Set("rule", convert.AdmissionRulesToSchema(retrievedPolicy.Rules)); err != nil {
		return fmt.Errorf("error reading %s policy: %s", policyTypeAdmission, err)
	}

	return nil
}

func updatePolicyAdmission(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	parsedRules, err := convert.SchemaToAdmissionRules(d)
	if err != nil {
		return fmt.Errorf("error updating %s policy: %s", policyTypeAdmission, err)
	}

	parsedPolicy := policy.AdmissionPolicy{
		Id:    policyTypeAdmission,
		Rules: parsedRules,
	}

	if err := policy.UpdateAdmission(*client, parsedPolicy); err != nil {
		return fmt.Errorf("error updating %s policy: %s", policyTypeAdmission, err)
	}

	return readPolicyAdmission(d, meta)
}

func deletePolicyAdmission(d *schema.ResourceData, meta interface{}) error {
	// TODO: reset to default policy
	return nil
}
