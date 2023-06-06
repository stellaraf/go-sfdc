package types

// type Options struct {
// 	User string `json:"string"`
// 	Scope string `json:"scope"`
// 	BaseURL string `json:"scope"`
// }

type StringOrInt interface {
	string | int
}

type Attributes struct {
	Type string `json:"type"`
	URL  string `json:"url"`
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

type RecordCreatedResponse struct {
	Errors  []string `json:"errors"`
	ID      string   `json:"id"`
	Success bool     `json:"success"`
}

type BaseObject struct {
	Attributes       Attributes `json:"attributes"`
	CreatedByID      string     `json:"CreatedById"`
	CreatedDate      string     `json:"CreatedDate"`
	ID               string     `json:"Id"`
	LastModifiedByID string     `json:"LastModifiedById"`
	LastModifiedDate string     `json:"LastModifiedDate"`
	OwnerID          string     `json:"OwnerId"`
	SystemModStamp   string     `json:"SystemModstamp"`
}
