package sfdc

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/url"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stellaraf/go-sfdc/internal/util"
	"github.com/stellaraf/go-utils/encryption"
)

const GRANT_TYPE_JWT_BEARER string = "urn:ietf:params:oauth:grant-type:jwt-bearer"

type Auth struct {
	InstanceURL            *url.URL
	privateKey             string
	clientID               string
	username               string
	httpClient             *resty.Client
	authURL                *url.URL
	encryption             bool
	encryptionPassphrase   string
	getAccessTokenCallback CachedTokenCallback
	setAccessTokenCallback SetTokenCallback
}

func parsePrivateKey(key []byte) (parsed any, err error) {
	parsed, err = x509.ParsePKCS8PrivateKey(key)
	if err != nil {
		parsed, err = x509.ParsePKCS1PrivateKey(key)
		if err != nil {
			parsed, err = x509.ParseECPrivateKey(key)
			if err != nil {
				return
			}
		}
	}
	if parsed == nil {
		err = fmt.Errorf("failed to parse private key")
		return
	}
	return
}

func (auth *Auth) GetNewToken() (token *Token, err error) {
	expiresAt := time.Now()
	expiresAt = expiresAt.Add(time.Second * 300)
	// SFDC requires that the audience be a single string, not an array.
	jwt.MarshalSingleStringAsArray = false
	claims := &jwt.RegisteredClaims{
		Issuer:    auth.clientID,
		Subject:   auth.username,
		Audience:  jwt.ClaimStrings{auth.authURL.String()},
		ExpiresAt: jwt.NewNumericDate(expiresAt),
	}
	initialToken := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	block, _ := pem.Decode([]byte(auth.privateKey))
	if block == nil {
		err = fmt.Errorf("failed to decode private key")
		return
	}
	rsaKey, err := parsePrivateKey(block.Bytes)
	if err != nil {
		return
	}
	assertion, err := initialToken.SignedString(rsaKey)
	if err != nil {
		return
	}

	req := auth.httpClient.R().
		SetHeader("content-type", "application/x-www-form-urlencoded").
		SetQueryParam("grant_type", GRANT_TYPE_JWT_BEARER).
		SetQueryParam("assertion", assertion).
		SetResult(&Token{}).
		SetError(&AuthErrorResponse{})

	res, err := req.Post("/services/oauth2/token")
	if err != nil {
		return
	}
	if res.IsError() {
		err = getSFDCError(res.Error())
		return
	}
	token, ok := res.Result().(*Token)
	if !ok {
		detail := string(res.Body())
		m := "failed to retrieve Salesforce access token"
		if detail != "" {
			m += fmt.Sprintf(" due to error: %s", detail)
		}
		err = fmt.Errorf(m)
		return
	}
	token.ExpiresAt = expiresAt
	return
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
	exp := time.Until(token.ExpiresAt)
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
	clientID, privateKey, username, authURL string,
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
	key := util.FormatPrivateKey(privateKey)
	auth = &Auth{
		InstanceURL:            nil,
		authURL:                parsedAuthURL,
		username:               username,
		clientID:               clientID,
		privateKey:             key,
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
