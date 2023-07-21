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

func setup() (
	getAccessTokenCallback sfdc.CachedTokenCallback,
	setAccessTokenCallback sfdc.SetTokenCallback,
	getRefreshTokenCallback sfdc.CachedTokenCallback,
	setRefreshTokenCallback sfdc.SetTokenCallback,
	err error) {
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

	setAccessToken := func(token string, expiresIn float64) error {
		cache.Add("access-token", time.Duration(expiresIn), token)
		return nil
	}
	getRefreshToken := func() (string, error) {
		res, err := cache.Value("refresh-token")
		if err != nil {
			return "", nil
		}
		token := res.Data().(string)
		return token, nil
	}
	setRefreshToken := func(token string, expiresIn float64) error {
		cache.Add("refresh-token", time.Duration(expiresIn), token)
		return nil
	}
	return getAccessToken, setAccessToken, getRefreshToken, setRefreshToken, nil
}

func initClient() (client *sfdc.Client, e env.Environment, err error) {
	e, err = env.LoadEnv()
	if err != nil {
		return
	}
	getAccessToken, setAccessToken, getRefreshToken, setRefreshToken, err := setup()
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
		getAccessToken, setAccessToken, getRefreshToken, setRefreshToken,
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
