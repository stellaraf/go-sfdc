package sfdc_test

import (
	"testing"
	"time"

	"github.com/muesli/cache2go"
	"github.com/stellaraf/go-sfdc"
	"github.com/stellaraf/go-sfdc/internal/env"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupAuth() (
	getAccessTokenCallback sfdc.CachedTokenCallback,
	setAccessTokenCallback sfdc.SetTokenCallback,
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

	setAccessToken := func(token string, expiresIn time.Duration) error {
		cache.Add("access-token", expiresIn, token)
		return nil
	}
	return getAccessToken, setAccessToken, nil
}

func initAuth() (auth *sfdc.Auth, err error) {
	env, err := env.LoadEnv()
	if err != nil {
		return
	}
	getAccessToken, setAccessToken, err := setup()
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
	)
	return
}

func Test_Auth(t *testing.T) {
	t.Run("get an access token", func(t *testing.T) {
		auth, err := initAuth()
		require.NoError(t, err)
		token, err := auth.GetNewToken()
		require.NoError(t, err)
		assert.NotEqual(t, "", token.AccessToken)
	})

	t.Run("test auth errors", func(t *testing.T) {
		env, err := env.LoadEnv()
		require.NoError(t, err)
		getAccessToken, setAccessToken, err := setupAuth()
		require.NoError(t, err)
		_, err = sfdc.NewAuth(
			"invalid-client-key",
			env.PrivateKey,
			env.AuthUsername,
			env.AuthURL,
			nil,
			getAccessToken,
			setAccessToken,
		)
		assert.ErrorContains(t, err, "client identifier invalid")
	})
}
