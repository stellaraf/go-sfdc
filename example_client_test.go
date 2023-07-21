package sfdc_test

import (
	"time"

	"github.com/muesli/cache2go"
	"github.com/stellaraf/go-sfdc"
)

// In this example, the cache2go caching backend is used, but any caching backend can be used.
func ExampleNew() {
	cache := cache2go.Cache("go-sfdc")
	cache.Flush()

	// Define a callback function to retrieve the Salesforce OAuth2 access token from the cache.
	getAccessToken := func() (string, error) {
		res, err := cache.Value("access-token")
		if err != nil {
			return "", nil
		}
		token := res.Data().(string)
		return token, nil
	}

	// Define a callback function to add the Salesforce OAuth2 access token from the cache.
	setAccessToken := func(token string, expiresIn float64) error {
		cache.Add("access-token", time.Duration(expiresIn), token)
		return nil
	}

	// Define a callback function to retrieve the Salesforce OAuth2 refresh token from the cache.
	getRefreshToken := func() (string, error) {
		res, err := cache.Value("refresh-token")
		if err != nil {
			return "", nil
		}
		token := res.Data().(string)
		return token, nil
	}
	// Define a callback function to add the Salesforce OAuth2 refresh token from the cache.
	setRefreshToken := func(token string, expiresIn float64) error {
		cache.Add("refresh-token", time.Duration(expiresIn), token)
		return nil
	}

	// Salesforce Connected App OAuth2 Client ID.
	clientID := "abcdef1234567890"

	// Provide the private key used for the Salesforce Connected App.
	privateKey := "<private key>"

	// If set, the encryption passphrase is used to encrypt all values written to the cache
	// using AES-256-GCM encryption.
	var encryptionPassphrase *string
	passphrase := "xY2jK9a8s6d5fE7H"
	encryptionPassphrase = &passphrase

	// If you do not wish to encrypt cached values, pass nil.
	encryptionPassphrase = nil

	// Salesforce authentication URL.
	// See: https://help.salesforce.com/s/articleView?id=sf.remoteaccess_oauth_jwt_flow.htm
	authURL := "https://login.salesforce.com"

	// Username with which go-sfdc will authenticate to the Salesforce API.
	// See: https://help.salesforce.com/s/articleView?id=sf.remoteaccess_oauth_jwt_flow.htm
	username := "user@example.com"

	client, err := sfdc.New(
		clientID,
		privateKey,
		username,
		authURL,
		encryptionPassphrase,
		getAccessToken,
		setAccessToken,
		getRefreshToken,
		setRefreshToken,
	)
	if err != nil {
		// handle error
	}
	client.Account("001G00000789QPONML")
}
