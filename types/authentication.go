package types

import (
	"encoding/json"
	"time"
)

type Token struct {
	ID          string    `json:"id"`
	Scope       string    `json:"scope"`
	TokenType   string    `json:"token_type"`
	AccessToken string    `json:"access_token"`
	InstanceURL string    `json:"instance_url"`
	ExpiresAt   time.Time `json:"expiresAt"`
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
