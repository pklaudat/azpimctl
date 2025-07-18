package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

type AzureContext struct {
	SubscriptionID   string `json:"id"`
	SubscriptionName string `json:"name"`
	IsDefault        bool   `json:"isDefault"`
	TenantId         string `json:"tenantId"`
}

type AzureProfile struct {
	// InstallationId string         `json:"installationId"`
	Subscriptions []AzureContext
}

func GetToken(scope string) string {
	ctx := context.Background()

	cred, err := azidentity.NewAzureCLICredential(nil)

	if err != nil {
		fmt.Printf("Failed to fetch Azure CLI Credentials. Make sure you have Azure CLI Installed")
	}

	token, err := cred.GetToken(
		ctx, policy.TokenRequestOptions{
			Scopes: []string{scope},
		},
	)

	if err != nil {
		fmt.Printf("Failed to fetch Azure Token.")
	}

	return token.Token

}

func GetContext() (*AzureContext, error) {
	var configDir string
	if runtime.GOOS == "windows" {
		configDir = filepath.Join(os.Getenv("USERPROFILE"), ".azure")
	} else {
		configDir = filepath.Join(os.Getenv("HOME"), ".azure")
	}

	contextFile := filepath.Join(configDir, "azureProfile.json")

	file, err := os.Open(contextFile)

	if err != nil {
		return nil, fmt.Errorf("Failed to open the azure profile context file.")
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	// trim bom bytes to avoid error opening in windows machines
	data = bytes.TrimPrefix(data, []byte{0xEF, 0xBB, 0xBF})

	var profile struct {
		Subscriptions []AzureContext `json:"subscriptions"`
	}

	if err := json.Unmarshal(data, &profile); err != nil {
		return nil, fmt.Errorf("Failed to decode the azure profile to json.")
	}

	if len(profile.Subscriptions) == 0 {
		return nil, nil
	}

	return &profile.Subscriptions[0], nil
}
