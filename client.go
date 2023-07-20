package sfdc

import (
	"github.com/go-resty/resty/v2"
	_auth "github.com/stellaraf/go-sfdc/auth"
	"github.com/stellaraf/go-sfdc/types"
	"github.com/stellaraf/go-sfdc/util"
)

// Salesforce Client
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

func (client *Client) mergeCustomFields(obj any, fields []types.CustomFields) (merged map[string]any, err error) {
	size := 0
	for _, m := range fields {
		size += len(m)
	}
	allFields := make(map[string]any, size)
	for _, m := range fields {
		for k, v := range m {
			allFields[k] = v
		}
	}
	merged, err = util.MergeStructToMap(obj, allFields)
	return
}

func (client *Client) handleObjectError(res *resty.Response) (err error) {
	if res.IsError() {
		e := res.Error().(*types.SalesforceErrorResponse)
		return util.GetSFDCError(e)
	}
	return
}

// Create a go-sfdc client and performs initial authentication.
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
