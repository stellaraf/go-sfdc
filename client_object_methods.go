// Methods that handle objects.
package sfdc

import (
	"fmt"

	"github.com/stellaraf/go-sfdc/types"
	"github.com/stellaraf/go-sfdc/util"
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
	err = handleResponse(res, &account)
	return
}

// Retrieve a User.
func (client *Client) User(id string) (user types.User, err error) {
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
	err = handleResponse(res, &user)
	return
}

// Retrieve a Group.
func (client *Client) Group(id string) (group types.Group, err error) {
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
	err = handleResponse(res, &group)
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
	err = handleResponse(res, &_case)
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
	err = handleResponse(res, &contact)
	return
}

// Update an account's fields.
func (client *Client) UpdateAccount(id string, data interface{}) (err error) {
	err = client.prepare()
	if err != nil {
		return
	}
	basePath, err := getPath("account")
	if err != nil {
		return
	}
	path := fmt.Sprintf("%s/%s", basePath, id)
	req := client.httpClient.R().SetBody(data)
	_, err = req.Patch(path)
	return
}

// Update a case's fields.
func (client *Client) UpdateCase(id string, data *types.CaseUpdate) (err error) {
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
	path := fmt.Sprintf("%s/%s", basePath, id)
	req := client.httpClient.R().SetBody(data)
	_, err = req.Patch(path)
	return
}

// Create a new case.
func (client *Client) CreateCase(data *types.CaseCreate, extra ...map[string]any) (result *types.RecordCreatedResponse, err error) {
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
	firstExtra := map[string]any{}
	if len(extra) != 0 {
		firstExtra = extra[0]
	}
	body, err := util.MergeStructToMap(data, firstExtra)
	if err != nil {
		return
	}
	req := client.httpClient.R().SetBody(body).SetResult(&types.RecordCreatedResponse{})
	res, err := req.Post(basePath)
	if err != nil {
		return
	}
	result = res.Result().(*types.RecordCreatedResponse)
	return
}
