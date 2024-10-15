package response

type UserpassLoginResponse struct {
	RequestId     string       `json:"request_id"`
	LeaseId       string       `json:"lease_id"`
	Renewable     bool         `json:"renewable"`
	LeaseDuration int          `json:"lease_duration"`
	Data          interface{}  `json:"data"`
	WrapInfo      interface{}  `json:"wrap_info"`
	Warnings      interface{}  `json:"warnings"`
	UserpassAuth  UserpassAuth `json:"auth"`
}

type UserpassAuth struct {
	ClientToken    string      `json:"client_token"`
	Accessor       string      `json:"accessor"`
	Policies       []string    `json:"policies"`
	TokenPolicies  []string    `json:"token_policies"`
	Metadata       Metadata    `json:"metadata"`
	LeaseDuration  int         `json:"lease_duration"`
	Renewable      bool        `json:"renewable"`
	EntityId       string      `json:"entity_id"`
	TokenType      string      `json:"token_type"`
	Orphan         bool        `json:"orphan"`
	MfaRequirement interface{} `json:"mfa_requirement"`
	NumUses        int         `json:"num_uses"`
}

type Metadata struct {
	Username string `json:"username"`
}

type AppRoleLoginResponse struct {
	AppRoleAuth   AppRoleAuth `json:"auth"`
	Warnings      interface{} `json:"warnings"`
	WrapInfo      interface{} `json:"wrap_info"`
	Data          interface{} `json:"data"`
	LeaseDuration int         `json:"lease_duration"`
	Renewable     bool        `json:"renewable"`
	LeaseId       string      `json:"lease_id"`
}

type AppRoleAuth struct {
	Renewable     bool        `json:"renewable"`
	LeaseDuration int         `json:"lease_duration"`
	Metadata      interface{} `json:"metadata"`
	TokenPolicies []string    `json:"token_policies"`
	Accessor      string      `json:"accessor"`
	ClientToken   string      `json:"client_token"`
}
