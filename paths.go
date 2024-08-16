package sfdc

import (
	"fmt"
)

const API_VERSION string = "v58.0"

const PATH_ACCOUNT string = "/services/data/%s/sobjects/Account"
const PATH_CASE string = "/services/data/%s/sobjects/Case"
const PATH_CONTACT string = "/services/data/%s/sobjects/Contact"
const PATH_FEED_ITEM string = "/services/data/%s/sobjects/FeedItem"
const PATH_GROUP string = "/services/data/%s/sobjects/Group"
const PATH_PRICEBOOK_ENTRY string = "/services/data/%s/sobjects/PricebookEntry"
const PATH_SERVICE_CONTRACT string = "/services/data/%s/sobjects/ServiceContract"
const PATH_SOQL string = "/services/data/%s/query"
const PATH_USER string = "/services/data/%s/sobjects/User"
const PATH_LEAD string = "/services/data/%s/sobjects/Lead"
const PATH_PICKLIST string = "/services/data/%s"
const PATH_TOOLING string = "/services/data/%s/tooling/query"

const PATH_BULK_INGEST string = "/services/data/%s/jobs/ingest"
const PATH_BULK_INGEST_COMPLETE string = "/services/data/%s/jobs/ingest/%s"

var pathMap map[string]string = map[string]string{
	"account":          PATH_ACCOUNT,
	"case":             PATH_CASE,
	"contact":          PATH_CONTACT,
	"feed_item":        PATH_FEED_ITEM,
	"group":            PATH_GROUP,
	"lead":             PATH_LEAD,
	"picklist":         PATH_PICKLIST,
	"pricebook_entry":  PATH_PRICEBOOK_ENTRY,
	"service_contract": PATH_SERVICE_CONTRACT,
	"soql":             PATH_SOQL,
	"tooling":          PATH_TOOLING,
	"user":             PATH_USER,
}

func getPath(pathName string) (path string, err error) {
	pathT, has := pathMap[pathName]
	if !has {
		err = fmt.Errorf("no path is defined for '%s'", pathName)
		return
	}
	path = fmt.Sprintf(pathT, API_VERSION)
	return
}
