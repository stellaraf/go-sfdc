package sfdc

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/stellaraf/go-sfdc/util"
)

const API_VERSION string = "v56.0"

const PATH_SOQL string = "/services/data/%s/query"
const PATH_ACCOUNT string = "/services/data/%s/sobjects/Account"
const PATH_CASE string = "/services/data/%s/sobjects/Case"
const PATH_FEED_ITEM string = "/services/data/%s/sobjects/FeedItem"
const PATH_GROUP string = "/services/data/%s/sobjects/Group"
const PATH_USER string = "/services/data/%s/sobjects/User"
const PATH_PRICEBOOK_ENTRY string = "/services/data/%s/sobjects/PricebookEntry"
const PATH_SERVICE_CONTRACT string = "/services/data/%s/sobjects/ServiceContract"

var pathMap map[string]string = map[string]string{
	"soql":             PATH_SOQL,
	"account":          PATH_ACCOUNT,
	"case":             PATH_CASE,
	"feed_item":        PATH_FEED_ITEM,
	"group":            PATH_GROUP,
	"user":             PATH_USER,
	"pricebook_entry":  PATH_PRICEBOOK_ENTRY,
	"service_contract": PATH_SERVICE_CONTRACT,
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

func handleResponse(res *resty.Response, data any) (err error) {
	err = util.CheckForError(res)
	if err != nil {
		return
	}
	body := res.Body()
	err = json.Unmarshal(body, data)
	return
}
