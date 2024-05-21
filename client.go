package sfdc

import (
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/go-resty/resty/v2"
)

const DefaultRetryDuration time.Duration = time.Second * 10

// Salesforce Client
type Client struct {
	httpClient *resty.Client
	auth       *Auth
	timeout    time.Duration
	backoff    backoff.BackOff
}

func (client *Client) prepare() error {
	token, err := client.auth.GetAccessToken()
	if err != nil {
		return err
	}
	client.httpClient.SetAuthToken(token)
	return nil
}

// do executes a given resty request method such as Get/Post. If a timeout/backoff is specified,
// the request will be executed and retried within that timeout period.
func (client *Client) Do(doer func(u string) (*resty.Response, error), url string) (*resty.Response, error) {
	op := func() (*resty.Response, error) {
		return doer(url)
	}
	if client.timeout == 0 {
		return op()
	}
	return backoff.RetryWithData(op, client.backoff)
}

// WithRetry specifies a time period in which to retry all requests if a errors are returned.
func (client *Client) WithRetry(timeout time.Duration) *Client {
	client.timeout = timeout
	client.backoff = backoff.NewExponentialBackOff(backoff.WithMaxElapsedTime(timeout))
	return client
}

// Create a go-sfdc client and performs initial authentication.
func New(
	clientID, clientSecret, authURL string,
	encryption *string,
	getToken CachedTokenCallback,
	setToken SetTokenCallback,
) (*Client, error) {

	auth, err := NewAuth(
		clientID, clientSecret, authURL,
		encryption,
		getToken,
		setToken,
	)
	if err != nil {
		return nil, err
	}
	httpClient := resty.New()
	httpClient.SetBaseURL(auth.InstanceURL.String())
	client := &Client{
		httpClient: httpClient,
		auth:       auth,
		timeout:    DefaultRetryDuration,
		backoff:    backoff.NewExponentialBackOff(backoff.WithMaxElapsedTime(DefaultRetryDuration)),
	}
	return client, nil
}
