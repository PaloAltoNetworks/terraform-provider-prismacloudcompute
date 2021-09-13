package prismacloudcompute

import (
	"log"

	pcc "github.com/paloaltonetworks/prisma-cloud-compute-go"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/settings/registry"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceRegistry() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRegistryRead,

		Schema: map[string]*schema.Schema{

			// Output.
			"harborscannerurlsuffix": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "URL suffix for the harbor scanner.",
			},
			"webhookurlsuffix": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "URL suffix for the webhook.",
			},
			"specification": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of specifications.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cap": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Specifies the maximum number of images from each repo to fetch and scan, sorted by most recently modified.",
						},
						"collections": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Specifies the set of Defenders in-scope for working on a scan job.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"credential": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Credential contains external provider authentication data",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"_id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique ID for the credential.",
									},
									"account_guid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Account identifier (e.g., username, access key, account GUID, etc.).",
									},
									"account_id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Account identifier (e.g., username, access key, account GUID, etc.).",
									},
									"api_token": {
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "Secret contains the plain and encrypted version of a value (the plain version is never stored in the DB)",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"encrypted": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Encrypted value for the secret.",
												},
												"plain": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Plain text value for the secret. Note: marshalling to JSON will convert to an encrypted value.",
												},
											},
										},
									},
									"ca_cert": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "CA certificate for certificate-based authentication.",
									},
									"created": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Created when the credential was created (or the account ID was changed for AWS).",
									},
									"description": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Description of the credential.",
									},
									"external": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Indicates if the credential is external (true) or not (false).",
									},
									"lastmodified": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Datetime when the credential was last modified.",
									},
									"owner": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "User who created or modified the credential.",
									},
									"role_arn": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Amazon Resource Name (ARN) of the role to assume.",
									},
									"secret": {
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "Secret contains the plain and encrypted version of a value (the plain version is never stored in the DB).",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"encrypted": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Encrypted value for the secret.",
												},
												"plain": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Plain text value for the secret. Note: marshalling to JSON will convert to an encrypted value.",
												},
											},
										},
									},
									"skip_verify": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "SkipVerify if should skip certificate verification in tls communication.",
									},
									"type": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The credential type.",
									},
									"url": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The server base URL.",
									},
									"use_aws_role": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Indicates if authentication should be done with the instance's attached credentials (EC2 IAM Role).",
									},
								},
							},
						},
						"credential_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "ID of the credentials in the credentials store to use for authenticating with the registry.",
						},
						"excluded_repositories": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Repositories to exclude from scanning.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"excluded_tags": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Tags to exclude from scanning.",
						},
						"harbor_deployment_security": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Indicates whether the Prisma Cloud plugin uses temporary tokens provided by Harbor to scan images in projects where Harbor's deployment security setting is enabled.",
						},
						"jfrog_repo_types": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "JFrog Artifactory repository types to scan.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"namespace": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IBM Bluemix namespace https://console.bluemix.net/docs/services/Regis.",
						},
						"os": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "RegistryOSType specifies the registry images base OS type.",
						},
						"registry": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Registry address (e.g., https://gcr.io)..",
						},
						"repository": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Repositories to scan.",
						},
						"scanners": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Number of Defenders that can be utilized for each scan job.",
						},
						"tag": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Tags to scan.",
						},
						"version": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Registry type. Determines the protocol Prisma Cloud uses to communicate with the registry.",
						},
						"version_pattern": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Pattern heuristic for quickly filtering images by tags without having to query all images for modification dates.",
						},
					},
				},
			},
		},
	}
}

func dataSourceRegistryRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)

	i, err := registry.Get(*client)

	if err != nil {
		return err
	}

	list := make([]interface{}, 0, 1)
	list = append(list, map[string]interface{}{
		"specification": i.Specifications,
	})

	if err := d.Set("listing", list); err != nil {
		log.Printf("[WARN] Error setting 'listing' field for %q: %s", d.Id(), err)
	}

	return nil
}
