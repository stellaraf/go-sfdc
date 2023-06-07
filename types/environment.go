package types

type TestData struct {
	AccountID   string `json:"accountId"`
	UserID      string `json:"userId"`
	GroupID     string `json:"groupId"`
	CaseID      string `json:"caseId"`
	AccountName string `json:"accountName"`
	ContactID   string `json:"contactId"`
}

type Environment struct {
	ClientID             string   `json:"clientId"`
	PrivateKey           string   `json:"privateKey"`
	AuthURL              string   `json:"authUrl"`
	AuthUsername         string   `json:"authUsername"`
	EncryptionPassphrase string   `json:"encryptionPassphrase"`
	TestData             TestData `json:"testData"`
}
