package sfdc

const (
	DEFAULT_RECORD_TYPE_ID string = "012000000000000AAA"
)

type CustomFields map[string]any

type Attributes struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type BaseObject struct {
	Attributes       *Attributes `json:"attributes,omitempty"`
	CreatedByID      string      `json:"CreatedById,omitempty"`
	CreatedDate      string      `json:"CreatedDate,omitempty"`
	ID               string      `json:"Id,omitempty"`
	LastModifiedByID string      `json:"LastModifiedById,omitempty"`
	LastModifiedDate string      `json:"LastModifiedDate,omitempty"`
	OwnerID          string      `json:"OwnerId,omitempty"`
	SystemModStamp   string      `json:"SystemModstamp,omitempty"`
}

type Contact struct {
	BaseObject
	AccountID              string       `json:"AccountId,omitempty"`
	IsDeleted              bool         `json:"IsDeleted,omitempty"`
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
	MailingAddress         *Address     `json:"MailingAddress,omitempty"`
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
	Address           *Address     `json:"Address,omitempty"`
	Email             string       `json:"Email,omitempty"`
	Phone             string       `json:"Phone,omitempty"`
	Alias             string       `json:"Alias,omitempty"`
	CommunityNickname string       `json:"CommunityNickname,omitempty"`
	IsActive          bool         `json:"IsActive,omitempty"`
	TimeZoneSidKey    string       `json:"TimeZoneSidKey,omitempty"`
	UserRoleID        string       `json:"UserRoleId,omitempty"`
	ManagerID         string       `json:"ManagerId,omitempty"`
	UserType          string       `json:"UserType,omitempty"`
	ProfileID         string       `json:"ProfileId,omitempty"`
	ContactID         string       `json:"ContactId,omitempty"`
	AccountID         string       `json:"AccountId,omitempty"`
	CustomFields      CustomFields `json:"-"`
}

type Group struct {
	BaseObject
	Name                   string       `json:"Name"`
	Developer_Name         string       `json:"Developer_Name,omitempty"`
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
	BillingAddress          *Address     `json:"BillingAddress,omitempty"`
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
	ShippingAddress         *Address     `json:"ShippingAddress,omitempty"`
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
	AccountId          string       `json:"AccountId,omitempty"`
	CaseNumber         string       `json:"CaseNumber,omitempty"`
	ClosedDate         string       `json:"ClosedDate,omitempty"`
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
	LastReferencedDate string       `json:"LastReferencedDate,omitempty"`
	LastViewedDate     string       `json:"LastViewedDate,omitempty"`
	MasterRecordID     string       `json:"MasterRecordId,omitempty"`
	MilestoneStatus    string       `json:"MilestoneStatus,omitempty"`
	Origin             string       `json:"Origin,omitempty"`
	OwnerID            string       `json:"OwnerId,omitempty"`
	ParentID           string       `json:"ParentId,omitempty"`
	Reason             string       `json:"Reason,omitempty"`
	Status             string       `json:"Status,omitempty"`
	StopStartDate      string       `json:"StopStartDate,omitempty"`
	Subject            string       `json:"Subject,omitempty"`
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
	Attributes         *Attributes `json:"attributes,omitempty"`
	ID                 string      `json:"Id,omitempty"`
	ParentID           string      `json:"ParentId,omitempty"`
	Type               string      `json:"Type,omitempty"`
	CreatedByID        string      `json:"CreatedById,omitempty"`
	CreatedDate        string      `json:"CreatedDate,omitempty"`
	IsDeleted          bool        `json:"IsDeleted,omitempty"`
	LastModifiedDate   string      `json:"LastModifiedDate,omitempty"`
	SystemModstamp     string      `json:"SystemModstamp,omitempty"`
	Revision           int         `json:"Revision,omitempty"`
	LastEditByID       string      `json:"LastEditById,omitempty"`
	LastEditDate       string      `json:"LastEditDate,omitempty"`
	ConnectionID       string      `json:"ConnectionId,omitempty"`
	CommentCount       int         `json:"CommentCount,omitempty"`
	LikeCount          int         `json:"LikeCount,omitempty"`
	Title              string      `json:"Title,omitempty"`
	Body               string      `json:"Body,omitempty"`
	LinkURL            string      `json:"LinkUrl,omitempty"`
	IsRichText         bool        `json:"IsRichText,omitempty"`
	RelatedRecordID    string      `json:"RelatedRecordId,omitempty"`
	InsertedByID       string      `json:"InsertedById,omitempty"`
	NetworkScope       string      `json:"NetworkScope,omitempty"`
	Visibility         string      `json:"Visibility,omitempty"`
	BestCommentID      string      `json:"BestCommentId,omitempty"`
	HasContent         bool        `json:"HasContent,omitempty"`
	HasLink            bool        `json:"HasLink,omitempty"`
	HasFeedEntity      bool        `json:"HasFeedEntity,omitempty"`
	HasVerifiedComment bool        `json:"HasVerifiedComment,omitempty"`
	IsClosed           bool        `json:"IsClosed,omitempty"`
	Status             string      `json:"Status,omitempty"`
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

type Lead struct {
	BaseObject
	Address                *Address     `json:"Address,omitempty"`
	AnnualRevenue          float64      `json:"AnnualRevenue,omitempty"`
	FirstName              string       `json:"FirstName,omitempty"`
	LastName               string       `json:"LastName,omitempty"`
	Description            string       `json:"Description,omitempty"`
	Company                string       `json:"Company,omitempty"`
	Name                   string       `json:"Name,omitempty"`
	Title                  string       `json:"Title,omitempty"`
	Email                  string       `json:"Email,omitempty"`
	Latitude               float64      `json:"latitude,omitempty"`
	Longitude              float64      `json:"longitude,omitempty"`
	Website                string       `json:"Website,omitempty"`
	PhotoURL               string       `json:"PhotoUrl,omitempty"`
	LeadSource             string       `json:"LeadSource,omitempty"`
	ActivityMetricId       string       `json:"ActivityMetricID,omitempty"`
	ActivityMetricRollupID string       `json:"ActivityMetricRollupId,omitempty"`
	Status                 string       `json:"Status,omitempty"`
	Phone                  string       `json:"Phone,omitempty"`
	MobilePhone            string       `json:"MobilePhone,omitempty"`
	Industry               string       `json:"Industry,omitempty"`
	Rating                 string       `json:"Rating,omitempty"`
	NumberOfEmployees      uint         `json:"NumberOfEmployees,omitempty"`
	HasOptedOutOfEmail     bool         `json:"HasOptedOutOfEmail,omitempty"`
	IsConverted            bool         `json:"IsConverted,omitempty"`
	ConvertedDate          *Time        `json:"ConvertedDate,omitempty"`
	ConvertedAccountID     string       `json:"ConvertedAccountId,omitempty"`
	ConvertedContactID     string       `json:"ConvertedContactId,omitempty"`
	ConvertedOpportunityID string       `json:"ConvertedOpportunityId,omitempty"`
	IsUnreadByOwner        bool         `json:"IsUnreadByOwner,omitempty"`
	LastActivityDate       *Time        `json:"LastActivityDate,omitempty"`
	LastViewedDate         *Time        `json:"LastViewedDate,omitempty"`
	LastReferencedDate     *Time        `json:"LastReferencedDate,omitempty"`
	Jigsaw                 any          `json:"Jigsaw,omitempty"`
	JigsawContactID        string       `json:"JigsawContactId,omitempty"`
	ConnectionReceivedID   string       `json:"ConnectionReceivedId,omitempty"`
	ConnectionSentID       string       `json:"ConnectionSentId,omitempty"`
	EmailBouncedReason     string       `json:"EmailBouncedReason,omitempty"`
	EmailBouncedDate       *Time        `json:"EmailBouncedDate,omitempty"`
	IndividualID           string       `json:"IndividualId,omitempty"`
	CustomFields           CustomFields `json:"-"`
}
