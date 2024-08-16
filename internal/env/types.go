package env

type TestData struct {
	AccountID             string `json:"accountId"`
	UserID                string `json:"userId"`
	GroupID               string `json:"groupId"`
	CaseID                string `json:"caseId"`
	AccountName           string `json:"accountName"`
	ContactID             string `json:"contactId"`
	AccountCustomFieldKey string `json:"accountCustomFieldKey"`
	CaseCustomFieldKey    string `json:"caseCustomFieldKey"`
	UserEmail             string `json:"userEmail"`
	ServiceContractID     string `json:"serviceContractId"`
	LeadID                string `json:"leadId"`
	PicklistObject        string `json:"picklistObject"`
	PicklistField         string `json:"picklistField"`
}

type Environment struct {
	ClientID             string   `json:"clientId"`
	ClientSecret         string   `json:"clientSecret"`
	AuthURL              string   `json:"authUrl"`
	EncryptionPassphrase string   `json:"encryptionPassphrase"`
	TestData             TestData `json:"testData"`
}
