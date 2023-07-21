package sfdc

import (
	"github.com/go-resty/resty/v2"
)

// Salesforce Client
type Client struct {
	httpClient *resty.Client
	auth       *Auth
}

func (client *Client) prepare() (err error) {
	token, err := client.auth.GetAccessToken()
	if err != nil {
		return
	}
	client.httpClient.SetAuthToken(token)
	return
}

// Create a go-sfdc client and performs initial authentication.
func New(
	clientID, privateKey, username, authURL string,
	encryption *string,
	getAccessTokenCallback CachedTokenCallback,
	setAccessTokenCallback SetTokenCallback,
	getRefreshTokenCallback CachedTokenCallback,
	setRefreshTokenCallback SetTokenCallback,
) (client *Client, err error) {

	auth, err := NewAuth(
		clientID, privateKey, username, authURL,
		encryption,
		getAccessTokenCallback,
		setAccessTokenCallback,
		getRefreshTokenCallback,
		setRefreshTokenCallback,
	)
	if err != nil {
		return
	}
	httpClient := resty.New()
	httpClient.SetBaseURL(auth.InstanceURL.String())
	client = &Client{
		httpClient: httpClient,
		auth:       auth,
	}
	return
}
