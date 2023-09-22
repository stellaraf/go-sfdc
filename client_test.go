package sfdc_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/muesli/cache2go"
	"github.com/stellaraf/go-sfdc"
	"github.com/stellaraf/go-sfdc/internal/env"
)

var Client *sfdc.Client
var Env env.Environment

func setup() (getAccessTokenCallback sfdc.CachedTokenCallback, setAccessTokenCallback sfdc.SetTokenCallback, err error) {
	cache := cache2go.Cache("go-sfdc-test")
	cache.Flush()
	getAccessToken := func() (string, error) {
		res, err := cache.Value("access-token")
		if err != nil {
			return "", nil
		}
		token := res.Data().(string)
		return token, nil
	}

	setAccessToken := func(token string, expiresIn time.Duration) error {
		cache.Add("access-token", expiresIn, token)
		return nil
	}
	return getAccessToken, setAccessToken, nil
}

func initClient() (client *sfdc.Client, e env.Environment, err error) {
	e, err = env.LoadEnv()
	if err != nil {
		return
	}
	getAccessToken, setAccessToken, err := setup()
	if err != nil {
		return
	}
	var encryptionPassphrase *string
	if e.EncryptionPassphrase != "" {
		encryptionPassphrase = &e.EncryptionPassphrase
	}
	client, err = sfdc.New(
		e.ClientID, e.PrivateKey, e.AuthUsername, e.AuthURL,
		encryptionPassphrase,
		getAccessToken, setAccessToken,
	)
	return
}

func init() {
	client, env, err := initClient()
	if err != nil {
		panic(err)
	}
	Client = client
	Env = env
}

func createCaseSubject(t *testing.T) string {
	now := time.Now()
	return fmt.Sprintf("go-sfdc %s at %s", t.Name(), now.Format(time.RFC3339Nano))
}
