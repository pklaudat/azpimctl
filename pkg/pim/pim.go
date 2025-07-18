package pim

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/pklaudat/azpimctl/pkg/auth"
	"github.com/pklaudat/azpimctl/pkg/utils"
)

const (
	PIM_BASE_URL        = "https://management.azure.com/providers/Microsoft.Authorization/roleEligibilityScheduleInstances"
	PIM_DEFAULT_SCOPE   = "https://management.azure.com"
	PIM_ARM_API_VERSION = "2020-10-01"
)

type pimClient struct {
	scope string
	token string
}

func NewPimClient(scope string) *pimClient {
	token := auth.GetToken(scope)
	return &pimClient{scope, token}
}

func (p *pimClient) ListElegibleRoles() {

	url := fmt.Sprint("%s?api-version=%s&$filter=asTarget()", PIM_BASE_URL, PIM_ARM_API_VERSION)
	roles := utils.Request("GET", url, p.token, nil)

	var response ElegibleRolesResponse

	fmt.Printf("%v", roles)

	err := json.NewDecoder(roles).Decode(&response)

	if err != nil {
		panic("Failed to decode elegible role response.")
	}

}

func (p *pimClient) ActivateElegibleRole(scope, roleId, duration string) {

	payload := RoleActivationRequest{
		Properties: RoleActivationProperties{
			RoleDefinitionId: roleId,
			PrincipalId:      "",
			RequestType:      "SelfActivate",
			ScheduledInfo: ScheduledInfoProperties{
				StartDateTime: time.Now().Local().String(),
				Expiration: ExpirationProperties{
					Duration: "P8H",
					Type:     "AfterDuration",
				},
			},
		},
	}

	payloadBytes, err := json.Marshal(payload)

	if err != nil {
		panic("Failed to marshal payload")
	}

	body := io.NopCloser(bytes.NewReader(payloadBytes))

	utils.Request("PUT", PIM_BASE_URL, p.token, body)

}
