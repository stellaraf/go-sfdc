package sfdc

import (
	"github.com/go-resty/resty/v2"
	_auth "github.com/stellaraf/go-sfdc/auth"
	"github.com/stellaraf/go-sfdc/types"
)

type Client struct {
	httpClient *resty.Client
	auth       *_auth.Auth
}

func (client *Client) prepare() (err error) {
	token, err := client.auth.GetAccessToken()
	if err != nil {
		return
	}
	client.httpClient.SetAuthToken(token)
	return
}

func NewClient(
	clientID, privateKey, username, authURL string,
	encryption *string,
	getAccessTokenCallback types.CachedTokenCallback,
	setAccessTokenCallback types.SetTokenCallback,
	getRefreshTokenCallback types.CachedTokenCallback,
	setRefreshTokenCallback types.SetTokenCallback,
) (client *Client, err error) {

	auth, err := _auth.NewAuth(
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
