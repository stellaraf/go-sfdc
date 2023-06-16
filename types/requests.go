package types

type CachedTokenCallback func() (string, error)

type SetTokenCallback func(token string, expiresIn float64) error

type RecordCreatedResponse struct {
	Errors  []string `json:"errors"`
	ID      string   `json:"id"`
	Success bool     `json:"success"`
}

type RecordResponse[R any] struct {
	Done      bool `json:"done"`
	TotalSize int  `json:"totalSize"`
	Records   []R  `json:"records"`
}

type GenericResponse struct {
	Message string         `json:"message"`
	Data    map[string]any `json:"data,omitempty"`
}
