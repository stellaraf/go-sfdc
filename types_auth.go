package sfdc

import (
	"encoding/json"
	"time"
)

type CachedTokenCallback func() (string, error)

type SetTokenCallback func(token string, expiresIn time.Duration) error

type Token struct {
	ID          string `json:"id"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
	InstanceURL string `json:"instance_url"`
	IssuedAt    string `json:"issued_at"`
	Signature   string `json:"signature"`
	expiresAt   time.Time
}

func (token *Token) SetExpiry(exp int) *Token {
	token.expiresAt = time.Unix(0, int64(exp)*int64(time.Millisecond))
	return token
}

func (token *Token) IsExpired() bool {
	return time.Now().After(token.expiresAt)
}

type TokenIntrospection struct {
	Active    bool   `json:"active"`
	Scope     string `json:"scope"`
	ClientID  string `json:"client_id"`
	Username  string `json:"username"`
	Sub       string `json:"sub"`
	TokenType string `json:"token_type"`
	Exp       int    `json:"exp"`
	Iat       int    `json:"iat"`
	Nbf       int    `json:"nbf"`
}

type JWTClaim struct {
	Iss string `json:"iss"`
	Aud string `json:"aud"`
	Sub string `json:"sub"`
	Exp string `json:"exp"`
}

type JWTHeader struct {
	Alg string `json:"alg"`
}

func (token *Token) JSON() string {
	b, err := json.Marshal(token)
	if err != nil {
		return ""
	}
	return string(b)
}
