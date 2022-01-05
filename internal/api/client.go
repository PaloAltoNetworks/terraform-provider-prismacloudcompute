package api

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
)

type APIClientConfig struct {
	ConsoleURL           string `json:"console_url"`
	Project              string `json:"project"`
	Username             string `json:"username"`
	Password             string `json:"password"`
	SkipCertVerification bool   `json:"skip_cert_verification"`
}

// A connection to Prisma Cloud Compute.
type Client struct {
	Config     APIClientConfig
	HTTPClient *http.Client
	JWT        string
}

// Communicate with the Prisma Cloud Compute API.
func (c *Client) Request(method, endpoint string, query, data, response interface{}) (err error) {
	parsedURL, err := url.Parse(c.Config.ConsoleURL)
	if err != nil {
		return err
	}
	if parsedURL.Scheme == "" {
		parsedURL.Scheme = "https"
	}
	parsedURL.Path = path.Join(parsedURL.Path, endpoint)

	var buf bytes.Buffer

	if data != nil {
		data_json, err := json.Marshal(data)
		if err != nil {
			return err
		}
		buf = *bytes.NewBuffer(data_json)
	}

	req, err := http.NewRequest(method, parsedURL.String(), &buf)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.JWT)
	req.Header.Set("Content-Type", "application/json")
	if c.Config.Project != "" {
		queryParams := req.URL.Query()
		queryParams.Set("project", c.Config.Project)
		req.URL.RawQuery = queryParams.Encode()
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Non-OK status: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if len(body) > 0 {
		if err = json.Unmarshal(body, response); err != nil {
			return err
		}
	}
	return nil
}

// Authenticate with the Prisma Cloud Compute Console.
func (c *Client) authenticate() (err error) {

	type AuthRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	type AuthResponse struct {
		Token string `json:"token"`
	}

	res := AuthResponse{}
	if err := c.Request(http.MethodPost, "/api/v1/authenticate", nil, AuthRequest{c.Config.Username, c.Config.Password}, &res); err != nil {
		return fmt.Errorf("error POSTing to authenticate endpoint: %v", err)
	}
	c.JWT = res.Token
	return nil
}

// Create Client and authenticate.
func APIClient(config APIClientConfig) (*Client, error) {
	apiClient := &Client{
		Config: config,
	}

	if config.SkipCertVerification {
		apiClient.HTTPClient = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}
	} else {
		apiClient.HTTPClient = &http.Client{}
	}

	if err := apiClient.authenticate(); err != nil {
		return nil, fmt.Errorf("error authenticating: %v", err)
	}

	return apiClient, nil
}
