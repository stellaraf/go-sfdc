package types

type CachedTokenCallback func() (string, error)

type SetTokenCallback func(token string, expiresIn float64) error
