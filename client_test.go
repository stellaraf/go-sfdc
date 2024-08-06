package sfdc_test

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/muesli/cache2go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.stellar.af/go-sfdc"
	"go.stellar.af/go-sfdc/internal/env"
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
		e.ClientID, e.ClientSecret, e.AuthURL,
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

func mockSuccessFn(_ string) (*resty.Response, error) {
	cache := cache2go.Cache("go-sfdc-test-client-backoff-1")
	iter, err := cache.Value("iter")
	if err != nil {
		return nil, err
	}
	c, ok := iter.Data().(int)
	if !ok {
		return nil, fmt.Errorf("failed to retrieve current iteration from 'go-sfdc-test-client-backoff-1' cache")
	}
	if c == 3 {
		return &resty.Response{}, nil
	}
	cache.Add("iter", time.Hour, c+1)
	return nil, fmt.Errorf("failure %d", c)
}

func mockFailureFn(_ string) (*resty.Response, error) {
	cache := cache2go.Cache("go-sfdc-test-client-backoff-2")
	iter, err := cache.Value("iter")
	if err != nil {
		return nil, err
	}
	c, ok := iter.Data().(int)
	if !ok {
		return nil, fmt.Errorf("failed to retrieve current iteration from 'go-sfdc-test-client-backoff-2' cache")
	}
	cache.Add("iter", time.Hour, c+1)
	return nil, fmt.Errorf("failure %d", c)
}

func Test_Client(t *testing.T) {
	cache1 := cache2go.Cache("go-sfdc-test-client-backoff-1")
	cache1.Add("iter", time.Hour, 0)
	cache2 := cache2go.Cache("go-sfdc-test-client-backoff-2")
	cache2.Add("iter", time.Hour, 0)
	t.Run("backoff success", func(t *testing.T) {
		t.Parallel()
		Client.WithRetry(time.Second * 5)
		res, err := Client.Do(mockSuccessFn, "")
		require.NoError(t, err)
		assert.IsType(t, &resty.Response{}, res)
	})
	t.Run("backoff failure", func(t *testing.T) {
		t.Parallel()
		Client.WithRetry(time.Second * 5)
		res, err := Client.Do(mockFailureFn, "")
		require.Error(t, err)
		require.Nil(t, res)
		assert.Regexp(t, regexp.MustCompile("failure [3-9]"), err.Error())
	})
	t.Cleanup(func() {
		cache1.Flush()
		cache2.Flush()
	})
}
