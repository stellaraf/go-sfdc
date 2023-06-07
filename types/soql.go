package types

type Customer struct {
	ID                string `json:"Id"`
	Name              string `json:"Name"`
	Type              string `json:"Type"`
	ServiceIdentifier string `json:"Service_Identifier__c"`
}
