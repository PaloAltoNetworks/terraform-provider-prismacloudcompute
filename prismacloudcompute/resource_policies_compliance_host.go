package prismacloudcompute

import (
	"fmt"
	"time"

	pcc "github.com/paloaltonetworks/prisma-cloud-compute-go"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/policy"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourcePoliciesComplianceHost() *schema.Resource {
	return &schema.Resource{
		Create: createPolicyComplianceHost,
		Read:   readPolicyComplianceHost,
		Update: updatePolicyComplianceHost,
		Delete: deletePolicyComplianceHost,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"_id": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  policyTypeComplianceHost,
			},
			"policy_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  policyTypeComplianceHost,
			},
			"rule": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of policy rules.",
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
							Description: "List of collections used to scope the rule.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"conditions": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "The set of compliance checks.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"compliance_check": {
										Type:        schema.TypeSet,
										Optional:    true,
										Description: "A compliance check. Omitted compliance checks are ignored.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"block": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Whether or not to block if this check is failed. Setting to 'false' will only alert on failure.",
												},
												"id": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "Compliance check ID.",
												},
											},
										},
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

func parsePolicyComplianceHost(d *schema.ResourceData, policyId string) (*policy.Policy, error) {
	parsedPolicy, err := parsePolicy(d, policyId, d.Get("policy_type").(string))
	if err != nil {
		return nil, fmt.Errorf("error parsing %s policy: %s", policyId, err)
	}
	for _, v := range parsedPolicy.Rules {
		v.Action = []string{""}
		v.Group = []string{""}
		v.License = policy.License{}
		v.Principal = []string{""}
	}
	return parsedPolicy, nil
}

func flattenPolicyComplianceHostRules(in []policy.Rule) []interface{} {
	ans := make([]interface{}, 0, len(in))
	for _, val := range in {
		m := make(map[string]interface{})
		m["block_message"] = val.BlockMsg
		m["collections"] = flattenCollections(val.Collections)
		m["conditions"] = flattenConditions(val.Condition)
		m["disabled"] = val.Disabled
		m["effect"] = val.Effect
		m["name"] = val.Name
		m["notes"] = val.Notes
		m["show_passed_checks"] = val.AllCompliance
		m["verbose"] = val.Verbose
		ans = append(ans, m)
	}
	return ans
}

func createPolicyComplianceHost(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	parsedPolicy, err := parsePolicyComplianceHost(d, "")
	if err != nil {
		return fmt.Errorf("error creating %s policy: %s", policyTypeComplianceHost, err)
	}

	if err := policy.Update(*client, policy.ComplianceHostEndpoint, *parsedPolicy); err != nil {
		return err
	}

	pol, err := policy.Get(*client, policy.ComplianceHostEndpoint)
	if err != nil {
		return err
	}

	d.SetId(pol.PolicyId)
	return readPolicyComplianceHost(d, meta)
}

func readPolicyComplianceHost(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)

	retrievedPolicy, err := policy.Get(*client, policy.ComplianceHostEndpoint)
	if err != nil {
		return err
	}

	d.Set("_id", policyTypeComplianceHost)
	d.Set("policy_type", policyTypeComplianceHost)
	if err := d.Set("rule", flattenPolicyComplianceHostRules(retrievedPolicy.Rules)); err != nil {
		return fmt.Errorf("error setting rule for resource %s: %s", d.Id(), err)
	}

	return nil
}

func updatePolicyComplianceHost(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	id := d.Id()
	parsedPolicy, err := parsePolicyComplianceHost(d, id)
	if err != nil {
		return fmt.Errorf("error updating %s policy: %s", policyTypeComplianceHost, err)
	}

	if err := policy.Update(*client, policy.ComplianceHostEndpoint, *parsedPolicy); err != nil {
		return err
	}

	return readPolicyComplianceHost(d, meta)
}

func deletePolicyComplianceHost(d *schema.ResourceData, meta interface{}) error {
	return nil
}
