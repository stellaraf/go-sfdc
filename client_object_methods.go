// Methods that handle objects.
package sfdc

import (
	"fmt"

	"github.com/stellaraf/go-sfdc/types"
)

func (client *Client) Account(id string) (account types.Account, err error) {
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

func (client *Client) Case(id string) (_case types.Case, err error) {
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

func (client *Client) Contact(id string) (contact types.Contact, err error) {
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
