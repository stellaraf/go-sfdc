// Methods that handle objects.
package sfdc

import (
	"fmt"

	"github.com/stellaraf/go-sfdc/types"
)

// Retrieve an Account.
func (client *Client) Account(id string) (account *types.Account, err error) {
	err = client.prepare()
	if err != nil {
		return
	}
	basePath, err := getPath("account")
	if err != nil {
		return
	}
	path := fmt.Sprintf("%s/%s", basePath, id)
	req := client.httpClient.R()
	res, err := req.Get(path)
	if err != nil {
		return
	}
	err = client.handleObjectError(res)
	if err != nil {
		return
	}
	account = &types.Account{}
	err = handleResponse(res, account)
	return
}

// Retrieve a User.
func (client *Client) User(id string) (user *types.User, err error) {
	err = client.prepare()
	if err != nil {
		return
	}
	basePath, err := getPath("user")
	if err != nil {
		return
	}
	path := fmt.Sprintf("%s/%s", basePath, id)
	req := client.httpClient.R()
	res, err := req.Get(path)
	if err != nil {
		return
	}
	err = client.handleObjectError(res)
	if err != nil {
		return
	}
	user = &types.User{}
	err = handleResponse(res, user)
	return
}

// Retrieve a Group.
func (client *Client) Group(id string) (group *types.Group, err error) {
	err = client.prepare()
	if err != nil {
		return
	}
	basePath, err := getPath("group")
	if err != nil {
		return
	}
	path := fmt.Sprintf("%s/%s", basePath, id)
	req := client.httpClient.R()
	res, err := req.Get(path)
	if err != nil {
		return
	}
	err = client.handleObjectError(res)
	if err != nil {
		return
	}
	group = &types.Group{}
	err = handleResponse(res, group)
	return
}

// Retrieve a Case.
func (client *Client) Case(id string) (_case *types.Case, err error) {
	err = client.prepare()
	if err != nil {
		return
	}
	basePath, err := getPath("case")
	if err != nil {
		return
	}
	path := fmt.Sprintf("%s/%s", basePath, id)
	req := client.httpClient.R()
	res, err := req.Get(path)
	if err != nil {
		return
	}
	err = client.handleObjectError(res)
	if err != nil {
		return
	}
	_case = &types.Case{}
	err = handleResponse(res, _case)
	return
}

// Retrieve a Contact.
func (client *Client) Contact(id string) (contact *types.Contact, err error) {
	err = client.prepare()
	if err != nil {
		return
	}
	basePath, err := getPath("contact")
	if err != nil {
		return
	}
	path := fmt.Sprintf("%s/%s", basePath, id)
	req := client.httpClient.R()
	res, err := req.Get(path)
	if err != nil {
		return
	}
	err = client.handleObjectError(res)
	if err != nil {
		return
	}
	contact = &types.Contact{}
	err = handleResponse(res, contact)
	return
}

// Update an account's fields.
func (client *Client) UpdateAccount(id string, data any, customFields ...types.CustomFields) (err error) {
	err = client.prepare()
	if err != nil {
		return
	}
	basePath, err := getPath("account")
	if err != nil {
		return
	}
	body, err := client.mergeCustomFields(data, customFields)
	if err != nil {
		return
	}
	path := fmt.Sprintf("%s/%s", basePath, id)
	req := client.httpClient.R().SetBody(body)
	res, err := req.Patch(path)
	if err != nil {
		return
	}
	return client.handleObjectError(res)
}

// Update a case's fields.
func (client *Client) UpdateCase(id string, data *types.CaseUpdate, customFields ...types.CustomFields) (err error) {
	_, err = client.Case(id)
	if err != nil {
		return
	}
	err = client.prepare()
	if err != nil {
		return
	}
	basePath, err := getPath("case")
	if err != nil {
		return
	}
	body, err := client.mergeCustomFields(data, customFields)
	if err != nil {
		return
	}
	path := fmt.Sprintf("%s/%s", basePath, id)
	req := client.httpClient.R().
		SetBody(body).
		SetResult(&types.RecordCreatedResponse{}).
		SetError(types.SalesforceErrorResponse{})
	res, err := req.Patch(path)
	if err != nil {
		return
	}
	return client.handleObjectError(res)
}

// Create a new case.
func (client *Client) CreateCase(data *types.CaseCreate, customFields ...types.CustomFields) (result *types.RecordCreatedResponse, err error) {
	err = client.prepare()
	if err != nil {
		return
	}

	if data.ContactID == "" {
		accountContact, err := client.AccountContact(data.AccountID)
		if err != nil {
			return nil, err
		}
		data.ContactID = accountContact.ID
	}

	basePath, err := getPath("case")
	if err != nil {
		return
	}
	body, err := client.mergeCustomFields(data, customFields)
	if err != nil {
		return
	}
	req := client.httpClient.R().
		SetBody(body).
		SetResult(&types.RecordCreatedResponse{}).
		SetError(types.SalesforceErrorResponse{})
	res, err := req.Post(basePath)
	if err != nil {
		return
	}
	err = client.handleObjectError(res)
	if err != nil {
		return
	}
	result = res.Result().(*types.RecordCreatedResponse)
	return
}

func (client *Client) FeedItem(id string) (result *types.FeedItem, err error) {
	err = client.prepare()
	if err != nil {
		return
	}
	basePath, err := getPath("feed_item")
	if err != nil {
		return
	}
	req := client.httpClient.R().SetResult(&types.FeedItem{})
	res, err := req.Get(fmt.Sprintf("%s/%s", basePath, id))
	if err != nil {
		return
	}
	err = client.handleObjectError(res)
	if err != nil {
		return
	}
	result = res.Result().(*types.FeedItem)
	return
}
