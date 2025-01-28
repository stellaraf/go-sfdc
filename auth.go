package sfdc

import (
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/go-resty/resty/v2"
	"go.stellar.af/go-utils/encryption"
)

type Auth struct {
	InstanceURL            *url.URL
	clientSecret           string
	clientID               string
	httpClient             *resty.Client
	authURL                *url.URL
	encryption             bool
	encryptionPassphrase   string
	getAccessTokenCallback CachedTokenCallback
	setAccessTokenCallback SetTokenCallback
}

func (auth *Auth) IntrospectToken(token string) (*TokenIntrospection, error) {
	data := map[string]string{
		"client_id":       auth.clientID,
		"client_secret":   auth.clientSecret,
		"token_type_hint": "access_token",
		"token":           token,
	}
	req := auth.httpClient.R().
		SetFormData(data).
		SetResult(&TokenIntrospection{}).
		SetError(&AuthErrorResponse{})
	res, err := req.Post("/services/oauth2/introspect")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		err = getSFDCError(res.Error())
		return nil, err
	}

	intro, ok := res.Result().(*TokenIntrospection)
	if !ok {
		detail := string(res.Body())
		m := "failed to introspect access token"
		if detail != "" {
			m += fmt.Sprintf(" due to error: %s", detail)
		}
		err = errors.New(m)
		return nil, err
	}
	return intro, nil
}

func (auth *Auth) GetNewToken() (*Token, error) {
	req := auth.httpClient.R().
		SetQueryParam("grant_type", "client_credentials").
		SetQueryParam("client_id", auth.clientID).
		SetQueryParam("client_secret", auth.clientSecret).
		SetResult(&Token{}).
		SetError(&AuthErrorResponse{})

	res, err := req.Post("/services/oauth2/token")
	if err != nil {
		return nil, err
	}

	if res.IsError() {
		err = getSFDCError(res.Error())
		return nil, err
	}

	token, ok := res.Result().(*Token)
	if !ok {
		detail := string(res.Body())
		m := "failed to retrieve Salesforce access token"
		if detail != "" {
			m += fmt.Sprintf(" due to error: %s", detail)
		}
		err = errors.New(m)
		return nil, err
	}

	intro, err := auth.IntrospectToken(token.AccessToken)
	if err != nil {
		return nil, err
	}

	token.SetExpiry(intro.Exp)

	return token, nil
}

func (auth *Auth) GetAccessToken() (token string, err error) {
	cachedToken, err := auth.getAccessTokenCallback()
	if err != nil {
		return
	}
	if cachedToken == "" {
		newToken, err := auth.GetNewToken()
		if err != nil {
			return "", err
		}
		err = auth.CacheNewToken(newToken)
		if err != nil {
			return "", err
		}
		return newToken.AccessToken, nil
	}
	if auth.encryption {
		decrypted, err := encryption.Decrypt(auth.encryptionPassphrase, cachedToken)
		if err != nil {
			return "", err
		}
		return decrypted, nil
	}
	return cachedToken, nil
}

func (auth *Auth) SetAccessToken(token *Token) (err error) {
	exp := time.Until(token.expiresAt)
	if auth.encryption {
		var encrypted string
		encrypted, err = encryption.Encrypt(token.AccessToken, auth.encryptionPassphrase)
		if err != nil {
			return
		}
		auth.setAccessTokenCallback(encrypted, exp)
		return
	}
	auth.setAccessTokenCallback(token.AccessToken, exp)
	return
}

func (auth *Auth) CacheNewToken(token *Token) (err error) {
	err = auth.SetAccessToken(token)
	if err != nil {
		return
	}
	return
}

func NewAuth(
	clientID, clientSecret, authURL string,
	encryption *string,
	getAccessTokenCallback CachedTokenCallback,
	setAccessTokenCallback SetTokenCallback,
) (auth *Auth, err error) {
	// parse auth base URL and set base URL of http client
	var doEncrypt bool
	passphrase := ""
	if encryption == nil {
		doEncrypt = false
	} else {
		doEncrypt = true
		passphrase = *encryption
	}
	httpClient := resty.New()
	parsedAuthURL, err := url.Parse(authURL)
	if err != nil {
		return
	}
	httpClient.SetHeader("user-agent", "go-sfdc")
	httpClient.SetBaseURL(fmt.Sprintf("%s://%s", parsedAuthURL.Scheme, parsedAuthURL.Host))
	auth = &Auth{
		InstanceURL:            nil,
		authURL:                parsedAuthURL,
		clientID:               clientID,
		clientSecret:           clientSecret,
		encryption:             doEncrypt,
		encryptionPassphrase:   passphrase,
		getAccessTokenCallback: getAccessTokenCallback,
		setAccessTokenCallback: setAccessTokenCallback,
		httpClient:             httpClient,
	}
	token, err := auth.GetNewToken()
	if err != nil {
		return
	}
	instanceURL, err := url.Parse(token.InstanceURL)
	if err != nil {
		return
	}
	auth.InstanceURL = instanceURL
	err = auth.CacheNewToken(token)
	return
}
