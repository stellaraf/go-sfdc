// Methods that query SOQL.
package sfdc

import (
	"fmt"

	"github.com/stellaraf/go-sfdc/types"
)

type soqlClient[T any] struct {
	Client
}

func (client *soqlClient[T]) Query(soqlQuery *soql) (results types.RecordResponse[T], err error) {
	query, err := soqlQuery.String()
	if err != nil {
		return
	}
	err = client.prepare()
	if err != nil {
		return
	}
	path, err := getPath("soql")
	if err != nil {
		return
	}
	req := client.httpClient.R()
	req.SetQueryParam("q", query)
	res, err := req.Get(path)
	if err != nil {
		return
	}
	err = handleResponse(res, &results)
	return
}

func initSOQL[T any](client *Client) (sc *soqlClient[T]) {
	sc = &soqlClient[T]{*client}
	return
}

func (client *Client) OpenCases() (cases types.RecordResponse[types.Case], err error) {
	sc := initSOQL[types.Case](client)
	q := SOQL().
		Select("Id", "Subject", "OwnerId").
		From("Case").
		Where("IsClosed", "=", false).
		Limit(200)
	cases, err = sc.Query(q)
	return
}

func (client *Client) UserName(id string) (name string, err error) {
	sc := initSOQL[types.User](client)
	q := SOQL().
		Select("Id", "Name").
		From("User").
		Where("Id", "=", id).
		Limit(1)
	user, err := sc.Query(q)
	if err != nil {
		return
	}
	if user.TotalSize != 1 {
		err = fmt.Errorf("no user found with ID '%s'", id)
		return
	}
	name = user.Records[0].Name
	return
}

func (client *Client) GroupName(id string) (name string, err error) {
	sc := initSOQL[types.Group](client)
	q := SOQL().
		Select("Id", "Name").
		From("Group").
		Where("Id", "=", id).
		Limit(1)
	group, err := sc.Query(q)
	if err != nil {
		return
	}
	if group.TotalSize != 1 {
		err = fmt.Errorf("no group found with ID '%s'", id)
		return
	}
	name = group.Records[0].Name
	return
}

func (client *Client) AccountIDFromName(name string) (id string, err error) {
	sc := initSOQL[types.Account](client)
	q := SOQL().
		Select("Id", "Name").
		From("Account").
		Where("Name", "=", name).
		Limit(1)
	account, err := sc.Query(q)
	if err != nil {
		return
	}
	if account.TotalSize != 1 {
		err = fmt.Errorf("no account found with name '%s'", name)
		return
	}
	id = account.Records[0].ID
	return
}

func (client *Client) Customers() (accounts []types.Customer, err error) {
	sc := initSOQL[types.Customer](client)
	q := SOQL().
		Select("Id", "Name", "Type", "Service_Identifier__c").
		From("Account").
		Where("Type", "=", "Customer")
	res, err := sc.Query(q)
	if err != nil {
		return
	}
	accounts = res.Records
	return
}
