package sfdc

import (
	"time"

	"github.com/muesli/cache2go"
	"github.com/stellaraf/go-sfdc/types"
	"github.com/stellaraf/go-sfdc/util"
)

func setup() (
	getAccessTokenCallback types.CachedTokenCallback,
	setAccessTokenCallback types.SetTokenCallback,
	getRefreshTokenCallback types.CachedTokenCallback,
	setRefreshTokenCallback types.SetTokenCallback,
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

func initClient() (client *Client, env types.Environment, err error) {
	env, err = util.LoadEnv()
	if err != nil {
		return
	}
	getAccessToken, setAccessToken, getRefreshToken, setRefreshToken, err := setup()
	if err != nil {
		return
	}
	var encryptionPassphrase *string
	if env.EncryptionPassphrase != "" {
		encryptionPassphrase = &env.EncryptionPassphrase
	}
	client, err = NewClient(
		env.ClientID, env.PrivateKey, env.AuthUsername, env.AuthURL,
		encryptionPassphrase,
		getAccessToken, setAccessToken, getRefreshToken, setRefreshToken,
	)
	return
}
