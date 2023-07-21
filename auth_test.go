package sfdc_test

import (
	"testing"
	"time"

	"github.com/muesli/cache2go"
	"github.com/stellaraf/go-sfdc"
	"github.com/stellaraf/go-sfdc/internal/env"
	"github.com/stretchr/testify/assert"
)

func setupAuth() (
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

func initAuth() (auth *sfdc.Auth, err error) {
	env, err := env.LoadEnv()
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
	auth, err = sfdc.NewAuth(
		env.ClientID,
		env.PrivateKey,
		env.AuthUsername,
		env.AuthURL,
		encryptionPassphrase,
		getAccessToken,
		setAccessToken,
		getRefreshToken,
		setRefreshToken,
	)
	return
}

func Test_Auth(t *testing.T) {
	t.Run("get an access token", func(t *testing.T) {
		auth, err := initAuth()
		assert.NoError(t, err)
		token, err := auth.GetNewToken()
		assert.NoError(t, err)
		assert.NotEqual(t, "", token.AccessToken)
	})

	t.Run("test auth errors", func(t *testing.T) {
		env, err := env.LoadEnv()
		assert.NoError(t, err)
		getAccessToken, setAccessToken, getRefreshToken, setRefreshToken, err := setupAuth()
		assert.NoError(t, err)
		_, err = sfdc.NewAuth(
			"invalid-client-key",
			env.PrivateKey,
			env.AuthUsername,
			env.AuthURL,
			nil,
			getAccessToken,
			setAccessToken,
			getRefreshToken,
			setRefreshToken,
		)
		assert.ErrorContains(t, err, "client identifier invalid")
	})
}
