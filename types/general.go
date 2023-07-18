package types

type StringOrInt interface {
	string | int
}

type Attributes struct {
	Type string `json:"type"`
	URL  string `json:"url"`
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
