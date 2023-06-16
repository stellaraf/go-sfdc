package types

type Customer struct {
	ID                string `json:"Id"`
	Name              string `json:"Name"`
	Type              string `json:"Type"`
	ServiceIdentifier string `json:"Service_Identifier__c"`
}

type OpenCase struct {
	ID      string `json:"Id"`
	Subject string `json:"Subject"`
	OwnerID string `json:"OwnerId"`
}

type ObjectSummary struct {
	ID   string `json:"Id"`
	Name string `json:"Name"`
}

type ObjectID struct {
	ID string `json:"Id"`
}
