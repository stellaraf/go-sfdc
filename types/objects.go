package types

type Contact struct {
	BaseObject
	IsDeleted              bool
	MasterRecordID         string         `json:"MasterRecordId,omitempty"`
	LastName               string         `json:"LastName,omitempty"`
	FirstName              string         `json:"FirstName,omitempty"`
	MiddleName             string         `json:"MiddleName,omitempty"`
	Suffix                 string         `json:"Suffix,omitempty"`
	Name                   string         `json:"Name,omitempty"`
	Salutation             string         `json:"Salutation,omitempty"`
	MailingStreet          string         `json:"MailingStreet,omitempty"`
	MailingCity            string         `json:"MailingCity,omitempty"`
	MailingState           string         `json:"MailingState,omitempty"`
	MailingPostalCode      string         `json:"MailingPostalCode,omitempty"`
	MailingCountry         string         `json:"MailingCountry,omitempty"`
	MailingLatitude        int            `json:"MailingLatitude,omitempty"`
	MailingLongitude       int            `json:"MailingLongitude,omitempty"`
	MailingGeocodeAccuracy string         `json:"MailingGeocodeAccuracy,omitempty"`
	MailingAddress         Address        `json:"MailingAddress,omitempty"`
	Phone                  string         `json:"Phone,omitempty"`
	Fax                    string         `json:"Fax,omitempty"`
	MobilePhone            string         `json:"MobilePhone,omitempty"`
	ReportsToId            string         `json:"ReportsToId,omitempty"`
	Email                  string         `json:"Email,omitempty"`
	Title                  string         `json:"Title,omitempty"`
	Department             string         `json:"Department,omitempty"`
	LastActivityDate       string         `json:"LastActivityDate,omitempty"`
	CustomFields           map[string]any `json:"-"`
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
	Username          string         `json:"Username"`
	LastName          string         `json:"LastName,omitempty"`
	FirstName         string         `json:"FirstName,omitempty"`
	MiddleName        string         `json:"MiddleName,omitempty"`
	Suffix            string         `json:"Suffix,omitempty"`
	Name              string         `json:"Name,omitempty"`
	CompanyName       string         `json:"CompanyName,omitempty"`
	Division          string         `json:"Division,omitempty"`
	Department        string         `json:"Department,omitempty"`
	Title             string         `json:"Title,omitempty"`
	Street            string         `json:"Street,omitempty"`
	City              string         `json:"City,omitempty"`
	State             string         `json:"State,omitempty"`
	PostalCode        string         `json:"PostalCode,omitempty"`
	Country           string         `json:"Country,omitempty"`
	Latitude          int            `json:"Latitude,omitempty"`
	Longitude         int            `json:"Longitude,omitempty"`
	GeocodeAccuracy   string         `json:"GeocodeAccuracy,omitempty"`
	Address           Address        `json:"Address,omitempty"`
	Email             string         `json:"Email,omitempty"`
	Phone             string         `json:"Phone,omitempty"`
	Alias             string         `json:"Alias"`
	CommunityNickname string         `json:"CommunityNickname,omitempty"`
	IsActive          bool           `json:"IsActive,omitempty"`
	TimeZoneSidKey    string         `json:"TimeZoneSidKey,omitempty"`
	UserRoleID        string         `json:"UserRoleId,omitempty"`
	ManagerID         string         `json:"ManagerId,omitempty"`
	UserType          string         `json:"UserType,omitempty"`
	ProfileID         string         `json:"ProfileId"`
	ContactID         string         `json:"ContactId,omitempty"`
	AccountID         string         `json:"AccountId,omitempty"`
	CustomFields      map[string]any `json:"-"`
}

type Group struct {
	BaseObject
	Name                   string `json:"Name"`
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
	AccountSource           string         `json:"AccountSource,omitempty"`
	AdditionalDomains       string         `json:"Additional_Domains__c,omitempty"`
	BillingAddress          Address        `json:"BillingAddress,omitempty"`
	BillingCity             string         `json:"BillingCity,omitempty"`
	BillingCountry          string         `json:"BillingCountry,omitempty"`
	BillingGeocodeAccuracy  string         `json:"BillingGeocodeAccuracy,omitempty"`
	BillingLatitude         int            `json:"BillingLatitude,omitempty"`
	BillingLongitude        int            `json:"BillingLongitude,omitempty"`
	BillingPostalCode       string         `json:"BillingPostalCode,omitempty"`
	BillingState            string         `json:"BillingState,omitempty"`
	BillingStreet           string         `json:"BillingStreet,omitempty"`
	CaseBCC                 string         `json:"Case_BCC__c,omitempty"`
	ConnectionReceivedID    string         `json:"ConnectionReceivedId,omitempty"`
	Description             string         `json:"Description,omitempty"`
	EmailDomains            string         `json:"Email_Domains__c,omitempty"`
	HasRMM                  bool           `json:"Has_RMM__c,omitempty"`
	Industry                string         `json:"Industry,omitempty"`
	IsDeleted               bool           `json:"IsDeleted,omitempty"`
	ITGlueOrganizations     string         `json:"IT_Glue_Organizations__c,omitempty"`
	Jigsaw                  string         `json:"Jigsaw,omitempty"`
	JigsawCompanyID         string         `json:"JigsawCompanyId,omitempty"`
	LastReferencedDate      string         `json:"LastReferencedDate"`
	LastViewedDate          string         `json:"LastViewedDate,omitempty"`
	MainPointOfContact      string         `json:"Main_Point_of_Contact__c,omitempty"`
	MasterRecordID          string         `json:"MasterRecordId,omitempty"`
	Name                    string         `json:"Name"`
	NumberOfEmployees       string         `json:"NumberOfEmployees,omitempty"`
	ParentID                string         `json:"ParentId,omitempty"`
	Phone                   string         `json:"Phone,omitempty"`
	PhotoUrl                string         `json:"PhotoUrl,omitempty"`
	RMMOrganizations        string         `json:"RMM_Organizations__c,omitempty"`
	ServiceIdentifier       string         `json:"Service_Identifier__c"`
	ShippingAddress         Address        `json:"ShippingAddress,omitempty"`
	ShippingCity            string         `json:"ShippingCity,omitempty"`
	ShippingCountry         string         `json:"ShippingCountry,omitempty"`
	ShippingGeocodeAccuracy string         `json:"ShippingGeocodeAccuracy,omitempty"`
	ShippingLatitude        int            `json:"ShippingLatitude,omitempty"`
	ShippingLongitude       int            `json:"ShippingLongitude,omitempty"`
	ShippingPostalCode      string         `json:"ShippingPostalCode,omitempty"`
	ShippingState           string         `json:"ShippingState,omitempty"`
	ShippingStreet          string         `json:"ShippingStreet,omitempty"`
	SicDesc                 string         `json:"SicDesc,omitempty"`
	Type                    string         `json:"Type,omitempty"`
	Website                 string         `json:"Website,omitempty"`
	CustomFields            map[string]any `json:"-"`
}

type CaseSummary struct {
	ID      string `json:"Id"`
	OwnerID string `json:"OwnerId"`
	// RMMSeriesUID string `json:"rmmSeriesUid__c"`
	Status string
}

type CaseUpdate struct {
	Comments         string `json:"Comments,omitempty"`
	Description      string `json:"Description,omitempty"`
	ServiceReason    string `json:"Service_Reason__c,omitempty"`
	ServiceSubReason string `json:"Service_SubReason__c,omitempty"`
	ServiceType      string `json:"Service_Type__c,omitempty"`
	Status           string `json:"Status,omitempty"`
	Subject          string `json:"Subject,omitempty"`
	Type             string `json:"Type,omitempty"`
}

type CaseCreate struct {
	AccountID        string `json:"AccountId,omitempty"`
	Comments         string `json:"Comments,omitempty"`
	ContactID        string `json:"ContactId,omitempty"`
	Description      string `json:"Description,omitempty"`
	Origin           string `json:"Origin,omitempty"`
	ServiceReason    string `json:"Service_Reason__c,omitempty"`
	ServiceSubReason string `json:"Service_SubReason__c,omitempty"`
	ServiceType      string `json:"Service_Type__c,omitempty"`
	Status           string `json:"Status,omitempty"`
	Subject          string `json:"Subject,omitempty"`
}

type Case struct {
	BaseObject
	AccountId   string `json:"AccountId"`
	CaseNumber  string `json:"CaseNumber"`
	CaseComment string `json:"Case_Comment__c,omitempty"`
	// LEGACY
	CaseSubReason      string         `json:"Case_Sub_Reason__c,omitempty"`
	ClosedDate         string         `json:"ClosedDate"`
	Comments           string         `json:"Comments,omitempty"`
	ContactEmail       string         `json:"ContactEmail,omitempty"`
	ContactFax         string         `json:"ContactFax,omitempty"`
	ContactID          string         `json:"ContactId,omitempty"`
	ContactMobile      string         `json:"ContactMobile,omitempty"`
	ContactPhone       string         `json:"ContactPhone,omitempty"`
	Description        string         `json:"Description,omitempty"`
	DiscoveredSolution bool           `json:"Discovered_Solution__c,omitempty"`
	EmailDomain        string         `json:"Email_Domain__c,omitempty"`
	EntitlementID      string         `json:"EntitlementId,omitempty"`
	ExternalTicketID   string         `json:"External_Ticket_Id__c,omitempty"`
	IsClosed           bool           `json:"IsClosed,omitempty"`
	IsDeleted          bool           `json:"IsDeleted,omitempty"`
	IsEscalated        bool           `json:"IsEscalated,omitempty"`
	IsStopped          bool           `json:"IsStopped,omitempty"`
	Language           string         `json:"Language,omitempty"`
	LastReferencedDate string         `json:"LastReferencedDate"`
	LastViewedDate     string         `json:"LastViewedDate"`
	MasterRecordID     string         `json:"MasterRecordId,omitempty"`
	MilestoneStatus    string         `json:"MilestoneStatus"`
	Origin             string         `json:"Origin,omitempty"`
	ParentId           string         `json:"ParentId,omitempty"`
	CustomFields       map[string]any `json:"-"`
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
	AlertID               string `json:"Alert_ID__c"`
}

type FeedItemOptions struct {
	// ID of whatever this is attached to (most likely, the case ID).
	ParentID string `json:"ParentId,omitempty"`
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
