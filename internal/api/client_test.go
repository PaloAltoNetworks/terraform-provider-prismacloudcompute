package api

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestAPIClient(t *testing.T) {
	credsFile, err := os.Open("examples/creds.json")
	if err != nil {
		fmt.Printf("error opening creds file: %v", err)
	}
	defer credsFile.Close()

	fileContent, err := io.ReadAll(credsFile)
	if err != nil {
		fmt.Printf("error reading creds file: %v", err)
		return
	}
	var config APIClientConfig
	if err := json.Unmarshal(fileContent, &config); err != nil {
		fmt.Printf("error unmarshalling creds file: %v", err)
		return
	}

	client, _ := APIClient(config)
	if client.JWT == "" {
		t.Errorf("JWT is empty. Authenticate did not work.")
	}
}
