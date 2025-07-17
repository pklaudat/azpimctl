package auth

import (
	"context"
	"encoding/json"
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
}

func GetToken(scope string) string {
	ctx := context.Background()

	cred, err := azidentity.NewAzureCLICredential(nil)

	if err != nil {
		panic("Failed to fetch Azure CLI Credentials. Make sure you have Azure CLI Installed")
	}

	token, err := cred.GetToken(
		ctx, policy.TokenRequestOptions{
			Scopes: []string{scope},
		},
	)

	if err != nil {
		panic("Failed to fetch Azure Token.")
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
		return nil, err
	}
	defer file.Close()

	var profile struct {
		Subscriptions []AzureContext `json:"subscriptions"`
	}
	if err := json.NewDecoder(file).Decode(&profile); err != nil {
		return nil, err
	}
	if len(profile.Subscriptions) == 0 {
		return nil, nil
	}

	return &profile.Subscriptions[0], nil
}
