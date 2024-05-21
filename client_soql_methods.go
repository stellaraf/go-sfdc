// Methods that query SOQL.
package sfdc

import (
	"fmt"
)

type SOQLClient[T any] struct {
	Client
}

// Query the Salesforce SOQL API with the created SOQL query.
func (soqlClient *SOQLClient[T]) Query(soqlQuery *soql) (results RecordResponse[T], err error) {
	query, err := soqlQuery.String()
	if err != nil {
		return
	}
	err = soqlClient.prepare()
	if err != nil {
		return
	}
	path, err := getPath("soql")
	if err != nil {
		return
	}
	req := soqlClient.httpClient.R().
		SetQueryParam("q", query).
		SetResult(RecordResponse[T]{}).
		SetError(SalesforceErrorResponse{})

	res, err := soqlClient.Do(req.Get, path)
	if err != nil {
		return
	}
	err = soqlClient.handleResponse(res, &results)
	return
}

func NewSOQL[T any](client *Client) (sc *SOQLClient[T]) {
	sc = &SOQLClient[T]{*client}
	return
}

// Retrieve a summary of all open cases.
func (client *Client) OpenCases() (cases []OpenCase, err error) {
	sc := NewSOQL[OpenCase](client)
	q := SOQL().
		Select("Id", "Subject", "OwnerId").
		From("Case").
		Where("IsClosed", EQUALS, false).
		Limit(200)
	res, err := sc.Query(q)
	if err != nil {
		return
	}
	cases = res.Records
	return
}

// Retrieve the name of a user based on that user's ID.
func (client *Client) UserName(id string) (name string, err error) {
	sc := NewSOQL[User](client)
	q := SOQL().
		Select("Id", "Name").
		From("User").
		Where("Id", EQUALS, id).
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

// Retrieve the name of a group based on the group's ID.
func (client *Client) GroupName(id string) (name string, err error) {
	sc := NewSOQL[ObjectSummary](client)
	q := SOQL().
		Select("Id", "Name").
		From("Group").
		Where("Id", EQUALS, id).
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

// Find an account's ID based on the account's name.
func (client *Client) AccountIDFromName(name string) (id string, err error) {
	sc := NewSOQL[Account](client)
	q := SOQL().
		Select("Id", "Name").
		From("Account").
		Where("Name", EQUALS, name).
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

// Retrieve a summary of all accounts where the 'Type' field is 'Customer'.
func (client *Client) Customers() (accounts []Customer, err error) {
	sc := NewSOQL[Customer](client)
	q := SOQL().
		Select("Id", "Name", "Type").
		From("Account").
		Where("Type", "=", "Customer")
	res, err := sc.Query(q)
	if err != nil {
		return
	}
	accounts = res.Records
	return
}

// Retrieve a case by its case number.
func (client *Client) CaseByNumber(caseNumber string) (result *Case, err error) {
	err = client.prepare()
	if err != nil {
		return
	}
	query := SOQL().Select("Id").From("Case").Where("CaseNumber", EQUALS, caseNumber)
	sc := NewSOQL[ObjectID](client)
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

// Retrieve a user by its email address.
func (client *Client) UserByEmail(email string) (user *User, err error) {
	err = client.prepare()
	if err != nil {
		return
	}
	query := SOQL().Select("Id").From("User").Where("Email", EQUALS, email)
	sc := NewSOQL[ObjectID](client)
	res, err := sc.Query(query)
	if err != nil {
		return
	}
	if len(res.Records) == 0 {
		err = fmt.Errorf("user with email '%s' not found", email)
		return
	}
	user, err = client.User(res.Records[0].ID)
	return
}
