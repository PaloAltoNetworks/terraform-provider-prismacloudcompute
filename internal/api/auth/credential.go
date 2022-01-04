package auth

import (
	"fmt"
	"net/http"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
)

const CredentialsEndpoint = "api/v1/credentials"

// type Credential struct {
// 	AccountGuid string           `json:"accountGUID,omitempty"`
// 	AccountId   string           `json:"accountID,omitempty"`
// 	Description string           `json:"description,omitempty"`
// 	Name        string           `json:"_id,omitempty"`
// 	Secret      CredentialSecret `json:"secret,omitempty"`
// 	Type        string           `json:"type,omitempty"`
// 	UseAwsRole  bool             `json:"useAWSRole,omitempty"`
// }

// type CredentialSecret struct {
// 	Plain string `json:"plain,omitempty"`
// }

type Credential struct {
	Id           string           `json:"_id,omitempty"`
	AccountGUID  string           `json:"accountGUID,omitempty"`
	AccountID    string           `json:"accountID,omitempty"`
	ApiToken     Secret           `json:"apiToken,omitempty"`
	CaCert       string           `json:"caCert,omitempty"`
	Created      string           `json:"created,omitempty"`
	Description  string           `json:"description,omitempty"`
	External     bool             `json:"external,omitempty"`
	LastModified string           `json:"lastModified,omitempty"`
	Owner        string           `json:"owner,omitempty"`
	RoleArn      string           `json:"roleArn,omitempty"`
	Secret       Secret           `json:"secret,omitempty"`
	SkipVerify   bool             `json:"skipVerify,omitempty"`
	Tokens       []TemporaryToken `json:"tokens,omitempty"`
	Type         string           `json:"type,omitempty"`
	Url          string           `json:"url,omitempty"`
	UseAWSRole   bool             `json:"useAWSRole,omitempty"`
}

type Secret struct {
	Encrypted string `json:"encrypted,omitempty"`
	Plain     string `json:"plain,omitempty"`
}

type TemporaryToken struct {
	AwsAccessKeyId     string `json:"awsAccessKeyId,omitempty"`
	AwsSecretAccessKey Secret `json:"awsSecretAccessKey,omitempty"`
	Duration           int    `json:"duration,omitempty"`
	ExpirationTime     string `json:"expirationTime,omitempty"`
	Token              Secret `json:"token,omitempty"`
}

// Get all credentials.
func ListCredentials(c api.Client) ([]Credential, error) {
	var ans []Credential
	if err := c.Request(http.MethodGet, CredentialsEndpoint, nil, nil, &ans); err != nil {
		return nil, fmt.Errorf("error listing credentials: %s", err)
	}
	return ans, nil
}

// Get a specific credential.
func GetCredential(c api.Client, name string) (*Credential, error) {
	credentials, err := ListCredentials(c)
	if err != nil {
		return nil, fmt.Errorf("error getting credential '%s': %s", name, err)
	}
	for _, val := range credentials {
		if val.Id == name {
			return &val, nil
		}
	}
	return nil, fmt.Errorf("credential '%s' not found", name)
}

// Create a new or update an existing credential.
func UpdateCredential(c api.Client, credential Credential) error {
	return c.Request(http.MethodPost, CredentialsEndpoint, nil, credential, nil)
}

// Delete an existing credential.
func DeleteCredential(c api.Client, name string) error {
	return c.Request(http.MethodDelete, fmt.Sprintf("%s/%s", CredentialsEndpoint, name), nil, nil, nil)
}
