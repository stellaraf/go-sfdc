// Methods that consume object _and_ SOQL methods, or that handle data processing.
package sfdc

import (
	"fmt"

	"github.com/stellaraf/go-sfdc/types"
)

// Retrieve the primary contact for an account.
func (client *Client) AccountContact(accountId string) (contact *types.Contact, err error) {
	err = client.prepare()
	if err != nil {
		return
	}
	account, err := client.Account(accountId)
	if err != nil {
		return
	}
	if account.MainPointOfContact == "" {
		err = fmt.Errorf("account '%s' (%s) has no Main Point of Contact associated", accountId, account.Name)
		return
	}
	contact, err = client.Contact(account.MainPointOfContact)
	return
}

// Close an open case.
func (client *Client) CloseCase(caseID string) (err error) {
	err = client.prepare()
	if err != nil {
		return
	}
	err = client.UpdateCase(caseID, &types.CaseUpdate{Status: "Closed"})
	return
}

// Post a comment to a case.
func (client *Client) PostToCase(caseID string, content string, feedOptions *types.FeedItemOptions) (result *types.RecordCreatedResponse, err error) {
	err = client.prepare()
	if err != nil {
		return
	}
	path, err := getPath("feed_item")
	if err != nil {
		return
	}
	if feedOptions == nil {
		feedOptions = &types.FeedItemOptions{}
	}
	feedOptions.ParentID = caseID
	feedOptions.Body = content
	feedOptions.Type = "TextPost"
	req := client.httpClient.R().SetBody(feedOptions).SetResult(&types.RecordCreatedResponse{})
	res, err := req.Post(path)
	if err != nil {
		return
	}
	result = res.Result().(*types.RecordCreatedResponse)
	return
}

// Retrieve a case by its case number.
func (client *Client) CaseByNumber(caseNumber string) (result *types.Case, err error) {
	err = client.prepare()
	if err != nil {
		return
	}
	query := SOQL().Select("Id").From("Case").Where("CaseNumber", "=", caseNumber)
	if err != nil {
		return
	}
	sc := NewSOQLClient[types.ObjectID](client)
	queryResult, err := sc.Query(query)
	if err != nil {
		return
	}
	if len(queryResult.Records) == 0 {
		err = fmt.Errorf("case with case number '%s' not found", caseNumber)
		return
	}
	caseID := queryResult.Records[0].ID
	result, err = client.Case(caseID)
	return
}
