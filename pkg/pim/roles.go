package pim

import (
	"encoding/json"
	"fmt"

	"github.com/pklaudat/azpimctl/pkg/utils"
)

const ROLE_ARM_API_VERSION = "2022-04-01"

type RoleDefinition struct {
	Properties struct {
		RoleName string `json:"roleName"`
	} `json:"properties"`
}

func GetRoleDisplayName(roleDefinitionID string, token string) (string, error) {
	var url string

	if len(roleDefinitionID) == 36 {
		url = fmt.Sprintf("https://management.azure.com/providers/Microsoft.Authorization/roleDefinitions/%s?api-version=%s", roleDefinitionID, ROLE_ARM_API_VERSION)
	} else {
		url = fmt.Sprintf("https://management.azure.com%s?api-version=%s", roleDefinitionID, ROLE_ARM_API_VERSION)
	}

	body := utils.Request("GET", url, token, nil)

	var result RoleDefinition
	if err := json.NewDecoder(body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	return result.Properties.RoleName, nil
}
