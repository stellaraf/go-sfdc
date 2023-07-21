package sfdc

type CustomFields map[string]any

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

type Contact struct {
	BaseObject
	IsDeleted              bool
	MasterRecordID         string       `json:"MasterRecordId,omitempty"`
	LastName               string       `json:"LastName,omitempty"`
	FirstName              string       `json:"FirstName,omitempty"`
	MiddleName             string       `json:"MiddleName,omitempty"`
	Suffix                 string       `json:"Suffix,omitempty"`
	Name                   string       `json:"Name,omitempty"`
	Salutation             string       `json:"Salutation,omitempty"`
	MailingStreet          string       `json:"MailingStreet,omitempty"`
	MailingCity            string       `json:"MailingCity,omitempty"`
	MailingState           string       `json:"MailingState,omitempty"`
	MailingPostalCode      string       `json:"MailingPostalCode,omitempty"`
	MailingCountry         string       `json:"MailingCountry,omitempty"`
	MailingLatitude        int          `json:"MailingLatitude,omitempty"`
	MailingLongitude       int          `json:"MailingLongitude,omitempty"`
	MailingGeocodeAccuracy string       `json:"MailingGeocodeAccuracy,omitempty"`
	MailingAddress         Address      `json:"MailingAddress,omitempty"`
	Phone                  string       `json:"Phone,omitempty"`
	Fax                    string       `json:"Fax,omitempty"`
	MobilePhone            string       `json:"MobilePhone,omitempty"`
	ReportsToId            string       `json:"ReportsToId,omitempty"`
	Email                  string       `json:"Email,omitempty"`
	Title                  string       `json:"Title,omitempty"`
	Department             string       `json:"Department,omitempty"`
	LastActivityDate       string       `json:"LastActivityDate,omitempty"`
	CustomFields           CustomFields `json:"-"`
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
	Username          string       `json:"Username"`
	LastName          string       `json:"LastName,omitempty"`
	FirstName         string       `json:"FirstName,omitempty"`
	MiddleName        string       `json:"MiddleName,omitempty"`
	Suffix            string       `json:"Suffix,omitempty"`
	Name              string       `json:"Name,omitempty"`
	CompanyName       string       `json:"CompanyName,omitempty"`
	Division          string       `json:"Division,omitempty"`
	Department        string       `json:"Department,omitempty"`
	Title             string       `json:"Title,omitempty"`
	Street            string       `json:"Street,omitempty"`
	City              string       `json:"City,omitempty"`
	State             string       `json:"State,omitempty"`
	PostalCode        string       `json:"PostalCode,omitempty"`
	Country           string       `json:"Country,omitempty"`
	Latitude          int          `json:"Latitude,omitempty"`
	Longitude         int          `json:"Longitude,omitempty"`
	GeocodeAccuracy   string       `json:"GeocodeAccuracy,omitempty"`
	Address           Address      `json:"Address,omitempty"`
	Email             string       `json:"Email,omitempty"`
	Phone             string       `json:"Phone,omitempty"`
	Alias             string       `json:"Alias"`
	CommunityNickname string       `json:"CommunityNickname,omitempty"`
	IsActive          bool         `json:"IsActive,omitempty"`
	TimeZoneSidKey    string       `json:"TimeZoneSidKey,omitempty"`
	UserRoleID        string       `json:"UserRoleId,omitempty"`
	ManagerID         string       `json:"ManagerId,omitempty"`
	UserType          string       `json:"UserType,omitempty"`
	ProfileID         string       `json:"ProfileId"`
	ContactID         string       `json:"ContactId,omitempty"`
	AccountID         string       `json:"AccountId,omitempty"`
	CustomFields      CustomFields `json:"-"`
}

type Group struct {
	BaseObject
	Name                   string       `json:"Name"`
	Developer_Name         string       `json:"Developer_Name"`
	RelatedID              string       `json:"RelatedId,omitempty"`
	QueueRoutingConfigID   string       `json:"QueueRoutingConfigId,omitempty"`
	DoesSendEmailToMembers bool         `json:"DoesSendEmailToMembers,omitempty"`
	DoesIncludeBosses      bool         `json:"DoesIncludeBosses,omitempty"`
	Email                  string       `json:"Email,omitempty"`
	Type                   string       `json:"Type,omitempty"`
	CustomFields           CustomFields `json:"-"`
}

type Account struct {
	BaseObject
	AccountSource           string       `json:"AccountSource,omitempty"`
	BillingAddress          Address      `json:"BillingAddress,omitempty"`
	BillingCity             string       `json:"BillingCity,omitempty"`
	BillingCountry          string       `json:"BillingCountry,omitempty"`
	BillingGeocodeAccuracy  string       `json:"BillingGeocodeAccuracy,omitempty"`
	BillingLatitude         int          `json:"BillingLatitude,omitempty"`
	BillingLongitude        int          `json:"BillingLongitude,omitempty"`
	BillingPostalCode       string       `json:"BillingPostalCode,omitempty"`
	BillingState            string       `json:"BillingState,omitempty"`
	BillingStreet           string       `json:"BillingStreet,omitempty"`
	ConnectionReceivedID    string       `json:"ConnectionReceivedId,omitempty"`
	Description             string       `json:"Description,omitempty"`
	Industry                string       `json:"Industry,omitempty"`
	IsCustomerPortal        bool         `json:"IsCustomerPortal"`
	IsDeleted               bool         `json:"IsDeleted,omitempty"`
	Jigsaw                  string       `json:"Jigsaw,omitempty"`
	JigsawCompanyID         string       `json:"JigsawCompanyId,omitempty"`
	LastReferencedDate      string       `json:"LastReferencedDate,omitempty"`
	LastViewedDate          string       `json:"LastViewedDate,omitempty"`
	MasterRecordID          string       `json:"MasterRecordId,omitempty"`
	Name                    string       `json:"Name"`
	NumberOfEmployees       int          `json:"NumberOfEmployees,omitempty"`
	ParentID                string       `json:"ParentId,omitempty"`
	Phone                   string       `json:"Phone,omitempty"`
	PhotoUrl                string       `json:"PhotoUrl,omitempty"`
	ShippingAddress         Address      `json:"ShippingAddress,omitempty"`
	ShippingCity            string       `json:"ShippingCity,omitempty"`
	ShippingCountry         string       `json:"ShippingCountry,omitempty"`
	ShippingGeocodeAccuracy string       `json:"ShippingGeocodeAccuracy,omitempty"`
	ShippingLatitude        int          `json:"ShippingLatitude,omitempty"`
	ShippingLongitude       int          `json:"ShippingLongitude,omitempty"`
	ShippingPostalCode      string       `json:"ShippingPostalCode,omitempty"`
	ShippingState           string       `json:"ShippingState,omitempty"`
	ShippingStreet          string       `json:"ShippingStreet,omitempty"`
	SicDesc                 string       `json:"SicDesc,omitempty"`
	Type                    string       `json:"Type,omitempty"`
	Website                 string       `json:"Website,omitempty"`
	CustomFields            CustomFields `json:"-"`
}

type CaseUpdate struct {
	// Options
	SkipAutoAssign bool `json:"-"`
	// Fields
	Comments        string `json:"Comments,omitempty"`
	ContactEmail    string `json:"ContactEmail,omitempty"`
	ContactFax      string `json:"ContactFax,omitempty"`
	ContactID       string `json:"ContactId,omitempty"`
	ContactMobile   string `json:"ContactMobile,omitempty"`
	ContactPhone    string `json:"ContactPhone,omitempty"`
	Description     string `json:"Description,omitempty"`
	EntitlementID   string `json:"EntitlementId,omitempty"`
	IsEscalated     bool   `json:"IsEscalated,omitempty"`
	IsStopped       bool   `json:"IsStopped,omitempty"`
	Language        string `json:"Language,omitempty"`
	MasterRecordID  string `json:"MasterRecordId,omitempty"`
	MilestoneStatus string `json:"MilestoneStatus,omitempty"`
	Origin          string `json:"Origin,omitempty"`
	OwnerID         string `json:"OwnerId,omitempty"`
	ParentId        string `json:"ParentId,omitempty"`
	Reason          string `json:"Reason,omitempty"`
	Status          string `json:"Status,omitempty"`
	StopStartDate   string `json:"StopStartDate,omitempty"`
	Subject         string `json:"Subject,omitempty"`
	SuppliedCompany string `json:"SuppliedCompany,omitempty"`
	SuppliedEmail   string `json:"SuppliedEmail,omitempty"`
	SuppliedName    string `json:"SuppliedName,omitempty"`
	SuppliedPhone   string `json:"SuppliedPhone,omitempty"`
	Type            string `json:"Type,omitempty"`
}

type CaseCreate struct {
	AccountID    string       `json:"AccountId,omitempty"`
	Comments     string       `json:"Comments,omitempty"`
	ContactID    string       `json:"ContactId,omitempty"`
	Description  string       `json:"Description,omitempty"`
	Origin       string       `json:"Origin,omitempty"`
	Status       string       `json:"Status,omitempty"`
	Subject      string       `json:"Subject,omitempty"`
	CustomFields CustomFields `json:"-"`
}

type Case struct {
	BaseObject
	AccountId          string       `json:"AccountId"`
	CaseNumber         string       `json:"CaseNumber"`
	ClosedDate         string       `json:"ClosedDate"`
	Comments           string       `json:"Comments,omitempty"`
	ContactEmail       string       `json:"ContactEmail,omitempty"`
	ContactFax         string       `json:"ContactFax,omitempty"`
	ContactID          string       `json:"ContactId,omitempty"`
	ContactMobile      string       `json:"ContactMobile,omitempty"`
	ContactPhone       string       `json:"ContactPhone,omitempty"`
	Description        string       `json:"Description,omitempty"`
	EntitlementID      string       `json:"EntitlementId,omitempty"`
	IsClosed           bool         `json:"IsClosed,omitempty"`
	IsDeleted          bool         `json:"IsDeleted,omitempty"`
	IsEscalated        bool         `json:"IsEscalated,omitempty"`
	IsStopped          bool         `json:"IsStopped,omitempty"`
	Language           string       `json:"Language,omitempty"`
	LastReferencedDate string       `json:"LastReferencedDate"`
	LastViewedDate     string       `json:"LastViewedDate"`
	MasterRecordID     string       `json:"MasterRecordId,omitempty"`
	MilestoneStatus    string       `json:"MilestoneStatus"`
	Origin             string       `json:"Origin,omitempty"`
	OwnerID            string       `json:"OwnerId,omitempty"`
	ParentID           string       `json:"ParentId,omitempty"`
	Reason             string       `json:"Reason,omitempty"`
	Status             string       `json:"Status,omitempty"`
	StopStartDate      string       `json:"StopStartDate,omitempty"`
	Subject            string       `json:"Subject"`
	SuppliedCompany    string       `json:"SuppliedCompany,omitempty"`
	SuppliedEmail      string       `json:"SuppliedEmail,omitempty"`
	SuppliedName       string       `json:"SuppliedName,omitempty"`
	SuppliedPhone      string       `json:"SuppliedPhone,omitempty"`
	Type               string       `json:"Type,omitempty"`
	CustomFields       CustomFields `json:"-"`
}

type FeedItemOptions struct {
	// ID of the user that posts the feed item.
	CreatedByID string `json:"CreatedById,omitempty"`
	// ID of whatever this is attached to (most likely, the case ID).
	ParentID string `json:"ParentId,omitempty"`
	// Title of the feed item.
	Title string `json:"Title,omitempty"`
	// Body - required if Type is TextPost.
	Body string `json:"Body,omitempty"`
	// Temporarily only supporting "TextPost".
	Type string `json:"Type,omitempty"`
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

type FeedItem struct {
	Attributes         Attributes `json:"attributes"`
	ID                 string     `json:"Id"`
	ParentID           string     `json:"ParentId"`
	Type               string     `json:"Type"`
	CreatedByID        string     `json:"CreatedById"`
	CreatedDate        string     `json:"CreatedDate"`
	IsDeleted          bool       `json:"IsDeleted"`
	LastModifiedDate   string     `json:"LastModifiedDate"`
	SystemModstamp     string     `json:"SystemModstamp"`
	Revision           int        `json:"Revision"`
	LastEditByID       string     `json:"LastEditById,omitempty"`
	LastEditDate       string     `json:"LastEditDate,omitempty"`
	ConnectionID       string     `json:"ConnectionId,omitempty"`
	CommentCount       int        `json:"CommentCount"`
	LikeCount          int        `json:"LikeCount"`
	Title              string     `json:"Title,omitempty"`
	Body               string     `json:"Body"`
	LinkURL            string     `json:"LinkUrl,omitempty"`
	IsRichText         bool       `json:"IsRichText"`
	RelatedRecordID    string     `json:"RelatedRecordId,omitempty"`
	InsertedByID       string     `json:"InsertedById"`
	NetworkScope       string     `json:"NetworkScope"`
	Visibility         string     `json:"Visibility"`
	BestCommentID      string     `json:"BestCommentId,omitempty"`
	HasContent         bool       `json:"HasContent"`
	HasLink            bool       `json:"HasLink"`
	HasFeedEntity      bool       `json:"HasFeedEntity"`
	HasVerifiedComment bool       `json:"HasVerifiedComment"`
	IsClosed           bool       `json:"IsClosed"`
	Status             string     `json:"Status"`
}

type ServiceContract struct {
	BaseObject
	AccountID               string   `json:"AccountId,omitempty"`
	ActivationDate          string   `json:"ActivationDate,omitempty"`
	ApprovalStatus          string   `json:"ApprovalStatus,omitempty"`
	BillingAddress          *Address `json:"BillingAddress,omitempty"`
	BillingCity             string   `json:"BillingCity,omitempty"`
	BillingCountry          string   `json:"BillingCountry,omitempty"`
	BillingGeocodeAccuracy  string   `json:"BillingGeocodeAccuracy,omitempty"`
	BillingLatitude         int      `json:"BillingLatitude,omitempty"`
	BillingLongitude        int      `json:"BillingLongitude,omitempty"`
	BillingPostalCode       string   `json:"BillingPostalCode,omitempty"`
	BillingState            string   `json:"BillingState,omitempty"`
	BillingStreet           string   `json:"BillingStreet,omitempty"`
	ContactID               string   `json:"ContactId,omitempty"`
	ContractNumber          string   `json:"ContractNumber,omitempty"`
	Description             string   `json:"Description,omitempty"`
	Discount                float64  `json:"Discount,omitempty"`
	EndDate                 string   `json:"EndDate,omitempty"`
	GrandTotal              float64  `json:"GrandTotal,omitempty"`
	IsDeleted               bool     `json:"IsDeleted,omitempty"`
	LastReferencedDate      string   `json:"LastReferencedDate,omitempty"`
	LastViewedDate          string   `json:"LastViewedDate,omitempty"`
	LineItemCount           int      `json:"LineItemCount,omitempty"`
	Name                    string   `json:"Name,omitempty"`
	ParentServiceContractID string   `json:"ParentServiceContractId,omitempty"`
	Pricebook2ID            string   `json:"Pricebook2Id,omitempty"`
	RecordTypeID            string   `json:"RecordTypeId,omitempty"`
	RootServiceContractID   string   `json:"RootServiceContractId,omitempty"`
	ShippingAddress         *Address `json:"ShippingAddress,omitempty"`
	ShippingCity            string   `json:"ShippingCity,omitempty"`
	ShippingCountry         string   `json:"ShippingCountry,omitempty"`
	ShippingGeocodeAccuracy string   `json:"ShippingGeocodeAccuracy,omitempty"`
	ShippingLatitude        int      `json:"ShippingLatitude,omitempty"`
	ShippingLongitude       int      `json:"ShippingLongitude,omitempty"`
	ShippingPostalCode      string   `json:"ShippingPostalCode,omitempty"`
	ShippingState           string   `json:"ShippingState,omitempty"`
	ShippingStreet          string   `json:"ShippingStreet,omitempty"`
	ShippingHandling        float64  `json:"ShippingHandling,omitempty"`
	SpecialTerms            string   `json:"SpecialTerms,omitempty"`
	StartDate               string   `json:"StartDate,omitempty"`
	Status                  string   `json:"Status,omitempty"`
	Subtotal                float64  `json:"Subtotal,omitempty"`
	SystemModstamp          string   `json:"SystemModstamp,omitempty"`
	Tax                     float64  `json:"Tax,omitempty"`
	Term                    int      `json:"Term,omitempty"`
	TotalPrice              float64  `json:"TotalPrice,omitempty"`
}
