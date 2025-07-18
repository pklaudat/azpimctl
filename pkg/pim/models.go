package pim

type ElegibleRolesResponse struct {
}

type RoleActivationResponse struct {
}

type ExpirationProperties struct {
	Type        string
	Duration    string
	EndDateTime string
}

type ScheduledInfoProperties struct {
	StartDateTime string
	Expiration    ExpirationProperties
}

type RoleActivationProperties struct {
	RoleDefinitionId string
	PrincipalId      string
	RequestType      string
	ScheduledInfo    ScheduledInfoProperties
}

type RoleActivationRequest struct {
	Properties struct {
	}
}
