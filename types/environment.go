package types

type Environment struct {
	ClientID             string   `json:"clientId"`
	PrivateKey           string   `json:"privateKey"`
	AuthURL              string   `json:"authUrl"`
	AuthUsername         string   `json:"authUsername"`
	EncryptionPassphrase string   `json:"encryptionPassphrase"`
	TestData             TestData `json:"testData"`
}
