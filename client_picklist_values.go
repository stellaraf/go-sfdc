package sfdc

import "fmt"

func (client *Client) PicklistValues(soAPIName, recordTypeID, fieldAPIName string) (*PicklistValues, error) {
	err := client.prepare()
	if err != nil {
		return nil, err
	}
	base, err := getPath("picklist")
	if err != nil {
		return nil, err
	}
	base += "/ui-api/object-info/%s/picklist-values/%s/%s"
	path := fmt.Sprintf(base, soAPIName, recordTypeID, fieldAPIName)
	req := client.httpClient.R().
		SetResult(&PicklistValues{}).
		SetError(SalesforceErrorResponse{})
	res, err := client.Do(req.Get, path)
	if err != nil {
		return nil, err
	}
	err = client.handleObjectError(res)
	if err != nil {
		return nil, err
	}
	result, ok := res.Result().(*PicklistValues)
	if !ok {
		return nil, fmt.Errorf("failed to parse response '%s'", string(res.Body()))
	}
	return result, nil
}

func (client *Client) StandardValueSet(name string) (*StandardValueSet, error) {

	q, err := SOQL().
		Select("Id", "DurableId", "Metadata", "FullName", "MasterLabel").
		From("StandardValueSet").
		Where("MasterLabel", EQUALS, name).
		String()
	if err != nil {
		return nil, err
	}
	path, err := getPath("tooling")
	if err != nil {
		return nil, err
	}
	err = client.prepare()
	if err != nil {
		return nil, err
	}
	req := client.httpClient.R().
		SetResult(&StandardValueSet{}).
		SetError(SalesforceErrorResponse{}).
		SetQueryParam("q", q)
	res, err := client.Do(req.Get, path)
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, getSFDCError(res.Error())
	}
	result, ok := res.Result().(*StandardValueSet)
	if !ok {
		return nil, fmt.Errorf("failed to parse response '%s'", string(res.Body()))
	}
	return result, nil
}
