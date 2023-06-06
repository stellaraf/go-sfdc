package types

type Contact struct {
	BaseObject
	IsDeleted              bool
	MasterRecordID         string `json:"MasterRecordId,omitempty"`
	LastName               string `json:"LastName,omitempty"`
	FirstName              string `json:"FirstName,omitempty"`
	MiddleName             string `json:"MiddleName,omitempty"`
	Suffix                 string `json:"Suffix,omitempty"`
	Name                   string `json:"Name,omitempty"`
	Salutation             string `json:"Salutation,omitempty"`
	MailingStreet          string `json:"MailingStreet,omitempty"`
	MailingCity            string `json:"MailingCity,omitempty"`
	MailingState           string `json:"MailingState,omitempty"`
	MailingPostalCode      string `json:"MailingPostalCode,omitempty"`
	MailingCountry         string `json:"MailingCountry,omitempty"`
	MailingLatitude        int    `json:"MailingLatitude,omitempty"`
	MailingLongitude       int    `json:"MailingLongitude,omitempty"`
	MailingGeocodeAccuracy string `json:"MailingGeocodeAccuracy,omitempty"`
	MailingAddress         string `json:"MailingAddress,omitempty"`
	Phone                  string `json:"Phone,omitempty"`
	Fax                    string `json:"Fax,omitempty"`
	MobilePhone            string `json:"MobilePhone,omitempty"`
	ReportsToId            string `json:"ReportsToId,omitempty"`
	Email                  string `json:"Email,omitempty"`
	Title                  string `json:"Title,omitempty"`
	Department             string `json:"Department,omitempty"`
	LastActivityDate       string `json:"LastActivityDate,omitempty"`
}

type Address struct {
	City            string `json:"city,omitempty"`
	Country         string `json:"country,omitempty"`
	GeocodeAccuracy string `json:"geocodeAccuracy,omitempty"`
	Latitude        int    `json:"latitude,omitempty"`
	Longitude       int    `json:"longitude,omitempty"`
	State           string `json:"state,omitempty"`
	Street          string `json:"street,omitempty"`
}

type User struct {
	BaseObject
	Username          string  `json:"Username"`
	LastName          string  `json:"LastName,omitempty"`
	FirstName         string  `json:"FirstName,omitempty"`
	MiddleName        string  `json:"MiddleName,omitempty"`
	Suffix            string  `json:"Suffix,omitempty"`
	Name              string  `json:"Name,omitempty"`
	CompanyName       string  `json:"CompanyName,omitempty"`
	Division          string  `json:"Division,omitempty"`
	Department        string  `json:"Department,omitempty"`
	Title             string  `json:"Title,omitempty"`
	Street            string  `json:"Street,omitempty"`
	City              string  `json:"City,omitempty"`
	State             string  `json:"State,omitempty"`
	PostalCode        string  `json:"PostalCode,omitempty"`
	Country           string  `json:"Country,omitempty"`
	Latitude          int     `json:"Latitude,omitempty"`
	Longitude         int     `json:"Longitude,omitempty"`
	GeocodeAccuracy   string  `json:"GeocodeAccuracy,omitempty"`
	Address           Address `json:"Address,omitempty"`
	Email             string  `json:"Email,omitempty"`
	Phone             string  `json:"Phone,omitempty"`
	Alias             string  `json:"Alias"`
	CommunityNickname string  `json:"CommunityNickname,omitempty"`
	IsActive          bool    `json:"IsActive,omitempty"`
	TimeZoneSidKey    string  `json:"TimeZoneSidKey,omitempty"`
	UserRoleID        string  `json:"UserRoleId,omitempty"`
	ManagerID         string  `json:"ManagerId,omitempty"`
	UserType          string  `json:"UserType,omitempty"`
	ProfileID         string  `json:"ProfileId"`
	ContactID         string  `json:"ContactId,omitempty"`
	AccountID         string  `json:"AccountId,omitempty"`
}

type Group struct {
	BaseObject
	Developer_Name         string `json:"Developer_Name"`
	RelatedID              string `json:"RelatedId,omitempty"`
	QueueRoutingConfigID   string `json:"QueueRoutingConfigId,omitempty"`
	DoesSendEmailToMembers bool   `json:"DoesSendEmailToMembers,omitempty"`
	DoesIncludeBosses      bool   `json:"DoesIncludeBosses,omitempty"`
	Email                  string `json:"Email,omitempty"`
	Type                   string `json:"Type,omitempty"`
}

type Account struct {
	BaseObject
	AccountSource            string  `json:"AccountSource,omitempty"`
	Additional_Domains__c    string  `json:"Additional_Domains__c,omitempty"`
	BillingAddress           Address `json:"BillingAddress,omitempty"`
	BillingCity              string  `json:"BillingCity,omitempty"`
	BillingCountry           string  `json:"BillingCountry,omitempty"`
	BillingGeocodeAccuracy   string  `json:"BillingGeocodeAccuracy,omitempty"`
	BillingLatitude          int     `json:"BillingLatitude,omitempty"`
	BillingLongitude         int     `json:"BillingLongitude,omitempty"`
	BillingPostalCode        string  `json:"BillingPostalCode,omitempty"`
	BillingState             string  `json:"BillingState,omitempty"`
	BillingStreet            string  `json:"BillingStreet,omitempty"`
	Case_BCC__c              string  `json:"Case_BCC__c,omitempty"`
	ConnectionReceivedId     string  `json:"ConnectionReceivedId,omitempty"`
	Description              string  `json:"Description,omitempty"`
	Email_Domains__c         string  `json:"Email_Domains__c,omitempty"`
	Has_RMM__c               bool    `json:"Has_RMM__c,omitempty"`
	Industry                 string  `json:"Industry,omitempty"`
	IsDeleted                bool    `json:"IsDeleted,omitempty"`
	IT_Glue_Organizations__c string  `json:"IT_Glue_Organizations__c,omitempty"`
	Jigsaw                   string  `json:"Jigsaw,omitempty"`
	JigsawCompanyId          string  `json:"JigsawCompanyId,omitempty"`
	LastReferencedDate       string  `json:"LastReferencedDate"`
	LastViewedDate           string  `json:"LastViewedDate,omitempty"`
	MasterRecordId           string  `json:"MasterRecordId,omitempty"`
	Name                     string  `json:"Name"`
	NumberOfEmployees        string  `json:"NumberOfEmployees,omitempty"`
	ParentId                 string  `json:"ParentId,omitempty"`
	Phone                    string  `json:"Phone,omitempty"`
	PhotoUrl                 string  `json:"PhotoUrl,omitempty"`
	RMM_Organizations__c     string  `json:"RMM_Organizations__c,omitempty"`
	Service_Identifier__c    string  `json:"Service_Identifier__c"`
	ShippingAddress          Address `json:"ShippingAddress,omitempty"`
	ShippingCity             string  `json:"ShippingCity,omitempty"`
	ShippingCountry          string  `json:"ShippingCountry,omitempty"`
	ShippingGeocodeAccuracy  string  `json:"ShippingGeocodeAccuracy,omitempty"`
	ShippingLatitude         int     `json:"ShippingLatitude,omitempty"`
	ShippingLongitude        int     `json:"ShippingLongitude,omitempty"`
	ShippingPostalCode       string  `json:"ShippingPostalCode,omitempty"`
	ShippingState            string  `json:"ShippingState,omitempty"`
	ShippingStreet           string  `json:"ShippingStreet,omitempty"`
	SicDesc                  string  `json:"SicDesc,omitempty"`
	Type                     string  `json:"Type,omitempty"`
	Website                  string  `json:"Website,omitempty"`
}

type CaseSummary struct {
	ID           string `json:"Id"`
	OwnerID      string `json:"OwnerId"`
	RMMSeriesUID string `json:"rmmSeriesUid__c"`
	Status       string
}

type Case struct {
	BaseObject
	AccountId   string `json:"AccountId"`
	CaseNumber  string `json:"CaseNumber"`
	CaseComment string `json:"Case_Comment__c,omitempty"`
	// LEGACY
	CaseSubReason      string `json:"Case_Sub_Reason__c,omitempty"`
	ClosedDate         string `json:"ClosedDate"`
	Comments           string `json:"Comments,omitempty"`
	ContactEmail       string `json:"ContactEmail,omitempty"`
	ContactFax         string `json:"ContactFax,omitempty"`
	ContactID          string `json:"ContactId,omitempty"`
	ContactMobile      string `json:"ContactMobile,omitempty"`
	ContactPhone       string `json:"ContactPhone,omitempty"`
	Description        string `json:"Description,omitempty"`
	DiscoveredSolution bool   `json:"Discovered_Solution__c,omitempty"`
	EmailDomain        string `json:"Email_Domain__c,omitempty"`
	EntitlementID      string `json:"EntitlementId,omitempty"`
	ExternalTicketID   string `json:"External_Ticket_Id__c,omitempty"`
	IsClosed           bool   `json:"IsClosed,omitempty"`
	IsDeleted          bool   `json:"IsDeleted,omitempty"`
	IsEscalated        bool   `json:"IsEscalated,omitempty"`
	IsStopped          bool   `json:"IsStopped,omitempty"`
	Language           string `json:"Language,omitempty"`
	LastReferencedDate string `json:"LastReferencedDate"`
	LastViewedDate     string `json:"LastViewedDate"`
	MasterRecordID     string `json:"MasterRecordId,omitempty"`
	MilestoneStatus    string `json:"MilestoneStatus"`
	Origin             string `json:"Origin,omitempty"`
	ParentId           string `json:"ParentId,omitempty"`
	// LEGACY
	Reason                string `json:"Reason,omitempty"`
	ResolutionImplemented bool   `json:"Resolution_Implemented__c,omitempty"`
	ServiceContract       string `json:"Service_Contract__c,omitempty"`
	ServiceReason         string `json:"Service_Reason__c"`
	ServiceSubReason      string `json:"Service_SubReason__c"`
	ServiceType           string `json:"Service_Type__c"`
	Status                string `json:"Status,omitempty"`
	StopStartDate         string `json:"StopStartDate,omitempty"`
	Subject               string `json:"Subject"`
	SuppliedCompany       string `json:"SuppliedCompany,omitempty"`
	SuppliedEmail         string `json:"SuppliedEmail,omitempty"`
	SuppliedName          string `json:"SuppliedName,omitempty"`
	SuppliedPhone         string `json:"SuppliedPhone,omitempty"`
	Type                  string `json:"Type,omitempty"`
	RMMSeriesUID          string `json:"rmmSeriesUid__c,omitempty"`
}

type FeedItem struct {
	// ID of whatever this is attached to (most likely, the case ID).
	ParentID string `json:"ParentId"`
	// Body - required if Type is TextPost.
	Body string `json:"Body"`
	// Temporarily only supporting "TextPost".
	Type string `json:"Type"`
	/* If true, these HTML tags are supported:
	<p/>
	<a/>
	<b/>
	<code/>
	<i/>
	<u/>
	<s/>
	<ul/>
	<ol/>
	<li/>
	<img/>
	*/
	IsRichText bool `json:"IsRichText,omitempty"`
	/*
	   Specifies whether this feed item is available to all users or internal users only.
	   "AllUsers": The feed item is available to all users who have permission to see the feed item.
	   "InternalUsers": The feed item is available to internal users only.

	   For record posts, Visibility is set to InternalUsers for all internal users by default.
	   External users can set Visibility only to AllUsers.
	   Visibility can be updated on record posts.
	   The Update property is supported only for feed items posted on records.
	*/
	Visibility string `json:"Visibility,omitempty"`
}
