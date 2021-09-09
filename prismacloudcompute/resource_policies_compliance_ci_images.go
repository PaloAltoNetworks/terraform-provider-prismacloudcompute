package prismacloudcompute

import (
	"fmt"
	"time"

	pcc "github.com/paloaltonetworks/prisma-cloud-compute-go"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/policies"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourcePoliciesComplianceCiImages() *schema.Resource {
	return &schema.Resource{
		Create: createPolicyComplianceCiImages,
		Read:   readPolicyComplianceCiImages,
		Update: updatePolicyComplianceCiImages,
		Delete: deletePolicyComplianceCiImages,

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
				Default:  policyTypeComplianceCIImages,
			},
			"policytype": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  policyTypeComplianceCIImages,
			},
			"rule": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of policy rules.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allcompliance": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether or not to report both failed and passed compliance checks.",
						},
						"blockmsg": {
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

func parsePolicyComplianceCiImages(d *schema.ResourceData, policyID string) policies.Policy {
	policy := parsePolicy(d, policyID, d.Get("policytype").(string))
	for _, v := range policy.Rules {
		v.Action = []string{""}
		v.Group = []string{""}
		v.License = policies.License{}
		v.Principal = []string{""}
	}
	return policy
}

func createPolicyComplianceCiImages(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	obj := parsePolicyComplianceCiImages(d, "")

	if err := policies.Update(*client, policies.ComplianceCiImagesEndpoint, obj); err != nil {
		return err
	}

	pol, err := policies.Get(*client, policies.ComplianceCiImagesEndpoint)
	if err != nil {
		return err
	}

	d.SetId(pol.PolicyId)
	return readPolicyComplianceCiImages(d, meta)
}

func readPolicyComplianceCiImages(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)

	obj, err := policies.Get(*client, policies.ComplianceCiImagesEndpoint)
	if err != nil {
		return err
	}

	d.Set("_id", policyTypeComplianceCIImages)
	d.Set("policytype", policyTypeComplianceCIImages)
	if err := d.Set("rule", obj.Rules); err != nil {
		return fmt.Errorf("error setting rule for resource %s: %s", d.Id(), err)
	}

	return nil
}

func updatePolicyComplianceCiImages(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	id := d.Id()
	obj := parsePolicyComplianceCiImages(d, id)

	if err := policies.Update(*client, policies.ComplianceCiImagesEndpoint, obj); err != nil {
		return err
	}

	return readPolicyComplianceCiImages(d, meta)
}

func deletePolicyComplianceCiImages(d *schema.ResourceData, meta interface{}) error {
	/*	client := meta.(*pcc.Client)
		id := d.Id()

		err := policy.Delete(client, id)
		if err != nil {
			if err != pcc.ObjectNotFoundError {
				return err
			}
		}*/

	d.SetId("")
	return nil
}
