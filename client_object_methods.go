// Methods that handle objects.
package sfdc

import (
	"fmt"

	"github.com/stellaraf/go-sfdc/internal/util"
)

// Retrieve an Account.
func (client *Client) Account(id string) (account *Account, err error) {
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
	account = &Account{}
	err = client.handleResponse(res, account)
	return
}

// Retrieve a User.
func (client *Client) User(id string) (user *User, err error) {
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
	user = &User{}
	err = client.handleResponse(res, user)
	return
}

// Retrieve a Group.
func (client *Client) Group(id string) (group *Group, err error) {
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
	group = &Group{}
	err = client.handleResponse(res, group)
	return
}

// Retrieve a Case.
func (client *Client) Case(id string) (_case *Case, err error) {
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
	_case = &Case{}
	err = client.handleResponse(res, _case)
	return
}

// Retrieve a Service Contract.
func (client *Client) ServiceContract(id string) (contract *ServiceContract, err error) {
	err = client.prepare()
	if err != nil {
		return
	}
	basePath, err := getPath("service_contract")
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
	contract = &ServiceContract{}
	err = client.handleResponse(res, contract)
	return
}

// Retrieve a Contact.
func (client *Client) Contact(id string) (contact *Contact, err error) {
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
	contact = &Contact{}
	err = client.handleResponse(res, contact)
	return
}

// Update an account's fields.
func (client *Client) UpdateAccount(id string, data any, customFields ...map[string]any) (err error) {
	err = client.prepare()
	if err != nil {
		return
	}
	basePath, err := getPath("account")
	if err != nil {
		return
	}
	body, err := util.MergeCustomFields(data, customFields)
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
func (client *Client) UpdateCase(id string, data *CaseUpdate, customFields ...map[string]any) (err error) {
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
	body, err := util.MergeCustomFields(data, customFields)
	if err != nil {
		return
	}
	path := fmt.Sprintf("%s/%s", basePath, id)
	req := client.httpClient.R().
		SetBody(body).
		SetResult(&RecordCreatedResponse{}).
		SetError(SalesforceErrorResponse{})

	if data.SkipAutoAssign {
		req.SetHeader("Sforce-Auto-Assign", "FALSE")
	}
	res, err := req.Patch(path)
	if err != nil {
		return
	}
	return client.handleObjectError(res)
}

// Create a new case.
func (client *Client) CreateCase(data *CaseCreate, customFields ...map[string]any) (result *RecordCreatedResponse, err error) {
	err = client.prepare()
	if err != nil {
		return
	}

	basePath, err := getPath("case")
	if err != nil {
		return
	}
	body, err := util.MergeCustomFields(data, customFields)
	if err != nil {
		return
	}
	req := client.httpClient.R().
		SetBody(body).
		SetResult(&RecordCreatedResponse{}).
		SetError(SalesforceErrorResponse{})
	res, err := req.Post(basePath)
	if err != nil {
		return
	}
	err = client.handleObjectError(res)
	if err != nil {
		return
	}
	result = res.Result().(*RecordCreatedResponse)
	return
}

func (client *Client) FeedItem(id string) (result *FeedItem, err error) {
	err = client.prepare()
	if err != nil {
		return
	}
	basePath, err := getPath("feed_item")
	if err != nil {
		return
	}
	req := client.httpClient.R().SetResult(&FeedItem{})
	res, err := req.Get(fmt.Sprintf("%s/%s", basePath, id))
	if err != nil {
		return
	}
	err = client.handleObjectError(res)
	if err != nil {
		return
	}
	result = res.Result().(*FeedItem)
	return
}

func (client *Client) CreateFeedItem(data *FeedItemOptions) (*RecordCreatedResponse, error) {
	err := client.prepare()
	if err != nil {
		return nil, err
	}
	basePath, err := getPath("feed_item")
	if err != nil {
		return nil, err
	}
	req := client.httpClient.R().SetResult(&RecordCreatedResponse{}).SetError(SalesforceErrorResponse{}).SetBody(data)
	res, err := req.Post(basePath)
	if err != nil {
		return nil, err
	}
	err = client.handleObjectError(res)
	if err != nil {
		return nil, err
	}
	result := res.Result().(*RecordCreatedResponse)
	return result, nil
}
