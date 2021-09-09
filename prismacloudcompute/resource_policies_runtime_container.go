package prismacloudcompute

import (
	"fmt"
	"time"

	pcc "github.com/paloaltonetworks/prisma-cloud-compute-go"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/policies"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourcePoliciesRuntimeContainer() *schema.Resource {
	return &schema.Resource{
		Create: createPolicyRuntimeContainer,
		Read:   readPolicyRuntimeContainer,
		Update: updatePolicyRuntimeContainer,
		Delete: deletePolicyRuntimeContainer,

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
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the policy set.",
				Default:     policyTypeRuntimeContainer,
			},
			"learningdisabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If set to 'true', automatic behavioural learning is enabled.",
				Default:     false,
			},
			"rule": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Rules in the policies.",
				MinItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"advancedprotection": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Prisma Cloud advanced threat protection",
						},
						"cloudmetadataenforcement": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Suspicious queries to cloud provider APIs",
						},
						"collections": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of collections used to scope the rule.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"customrules": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of custom rules.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"_id": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Custom rule ID.",
									},
									"action": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "The action to perform if the custom rule applies. Can be set to 'audit', 'incident'.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The effect to be used for the custom rule. Can be set to 'block', 'prevent', 'alert', 'allow', 'ban', or 'disable'.",
									},
								},
							},
						},
						"disabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "If set to 'true', the rule is currently disabled.",
						},
						"dns": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "The DNS runtime rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"blacklist": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Deny-listed domain names (e.g., www.bad-url.com, *.bad-url.com).",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The effect to be used in the runtime rule. Can be set to 'block', 'prevent', 'alert', 'disable'.",
									},
									"whitelist": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Allow-listed domain names (e.g., *.gmail.com, *.s3.amazon.com).",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"filesystem": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Represents restrictions or suppression for filesystem changes.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"backdoorfiles": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', monitors files that can create or persist backdoors (SSH or admin account config files).",
									},
									"blacklist": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of denied file system paths.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"checknewfiles": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', Detects changes to binaries and certificates.",
									},
									"effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The effect that will be used in the runtime rule. Can be set to 'block', 'prevent', 'alert', or 'disable'.",
									},
									"skipencryptedbinaries": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', the encrypted binaries check will be skipped.",
									},
									"suspiciouselfheaders": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', enables malware detection based on suspicious ELF headers.",
									},
									"whitelist": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of allowed filesystem paths.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"kubernetesenforcement": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Detects containers that attempt to compromise the orchestrator.",
						},
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Name of the rule.",
						},
						"network": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Represents the restrictions and suppression for networking.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"blacklistips": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Deny-listed IP addresses.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"blacklistlisteningports": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Deny-listed listening ports.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"deny": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "If set to 'true', the connection is denied.",
												},
												"end": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "end",
												},
												"start": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "start",
												},
											},
										},
									},
									"blacklistoutboundports": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Deny-listed outbound ports.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"deny": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "If set to 'true', the connection is denied.",
												},
												"end": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "end",
												},
												"start": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "start",
												},
											},
										},
									},
									"detectportscan": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', port scanning detection is enabled.",
									},
									"effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Effect used in the runtime rule. Can be set to: 'block', 'prevent', 'alert', or 'disable'.",
									},
									"skipmodifiedproc": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', Prisma Cloud can detect malicious networking activity from modified processes.",
									},
									"skiprawsockets": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', raw socket detection will be skipped.",
									},
									"whitelistips": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Allow-listed IP addresses.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"whitelistlisteningports": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Allow-listed listening ports.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"deny": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "If set to 'true', the connection is denied.",
												},
												"end": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "end",
												},
												"start": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "start",
												},
											},
										},
									},
									"whitelistoutboundports": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Allow-listed outbound ports.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"deny": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "If set to 'true', the connection is denied.",
												},
												"end": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "end",
												},
												"start": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "start",
												},
											},
										},
									},
								},
							},
						},
						"notes": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A free-form text description of the collection.",
						},
						"processes": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Represents restrictions or suppression for running processes.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"blacklist": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of processes to deny.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"checkcryptominers": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', detect crypto miners.",
									},
									"checklateralmovement": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', enables detection of processes that can be used for lateral movement exploits.",
									},
									"checkparentchild": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', enables check for parent-child relationship when comparing spawned processes in the model.",
									},
									"checksuidbinaries": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', enables check for process elevanting privileges (SUID bit).",
									},
									"effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The effect to be used in the runtime rule. Can be set to 'block', 'prevent', 'alert', 'disable'.",
									},
									"skipmodified": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', trigger audits/incidents when a modified proc is spawned.",
									},
									"skipreverseshell": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', reverse shell detection is disabled.",
									},
									"whitelist": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Allow-list of processes.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"wildfireanalysis": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The effect that will be used in the runtime rule. Can be set to 'block', 'prevent', 'alert', or 'disable'.",
						},
					},
				},
			},
		},
	}
}

func parsePolicyRuntimeContainer(d *schema.ResourceData, policyID string) policies.Policy {
	return parsePolicy(d, policyID, "")
}

func createPolicyRuntimeContainer(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	obj := parsePolicyRuntimeContainer(d, "")

	if err := policies.Update(*client, policies.RuntimeContainerEndpoint, obj); err != nil {
		return err
	}

	pol, err := policies.Get(*client, policies.RuntimeContainerEndpoint)
	if err != nil {
		return err
	}

	d.SetId(pol.PolicyId)
	return readPolicyRuntimeContainer(d, meta)
}

func readPolicyRuntimeContainer(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)

	obj, err := policies.Get(*client, policies.RuntimeContainerEndpoint)
	if err != nil {
		return err
	}

	d.Set("_id", policyTypeRuntimeContainer)
	d.Set("learningdisabled", obj.LearningDisabled)
	if err := d.Set("rule", obj.Rules); err != nil {
		return fmt.Errorf("error setting rule for resource %s: %s", d.Id(), err)
	}

	return nil
}

func updatePolicyRuntimeContainer(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	id := d.Id()
	obj := parsePolicyRuntimeContainer(d, id)

	if err := policies.Update(*client, policies.RuntimeContainerEndpoint, obj); err != nil {
		return err
	}

	return readPolicyRuntimeContainer(d, meta)
}

func deletePolicyRuntimeContainer(d *schema.ResourceData, meta interface{}) error {
	/*	client := meta.(*pcc.Client)
		id := d.Id()

		err := policies.Delete(client, id)
		if err != nil {
			if err != pcc.ObjectNotFoundError {
				return err
			}
		}*/

	d.SetId("")
	return nil
}
