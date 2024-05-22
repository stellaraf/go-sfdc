// Methods that consume object _and_ SOQL methods, or that handle data processing.
package sfdc

import "fmt"

// Close an open case.
func (client *Client) CloseCase(caseID string) (err error) {
	err = client.prepare()
	if err != nil {
		return
	}
	err = client.UpdateCase(caseID, &CaseUpdate{Status: "Closed", SkipAutoAssign: true})
	return
}

// Post a comment to a case.
func (client *Client) PostToCase(caseID string, content string, feedOptions *FeedItemOptions) (result *RecordCreatedResponse, err error) {
	err = client.prepare()
	if err != nil {
		return
	}
	path, err := getPath("feed_item")
	if err != nil {
		return
	}
	if feedOptions == nil {
		feedOptions = &FeedItemOptions{}
	}
	feedOptions.ParentID = caseID
	feedOptions.Body = content
	feedOptions.Type = "TextPost"
	req := client.httpClient.R().SetBody(feedOptions).SetResult(&RecordCreatedResponse{})
	res, err := client.Do(req.Post, path)
	if err != nil {
		return
	}
	result = res.Result().(*RecordCreatedResponse)
	return
}

// GetObject retrieves an object and provides an interface to access its data.
func (client *Client) GetObject(path string) (*ObjectResponse, error) {
	err := client.prepare()
	if err != nil {
		return nil, err
	}
	req := client.httpClient.R()
	res, err := client.Do(req.Get, path)
	if err != nil {
		return nil, err
	}
	obj, err := NewObjectResponse(res.Body())
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// PostObject creates a new object.
func (client *Client) PostObject(path string, data map[string]any) (*RecordCreatedResponse, error) {
	err := client.prepare()
	if err != nil {
		return nil, err
	}
	req := client.httpClient.R().SetBody(data).SetResult(&RecordCreatedResponse{})
	res, err := client.Do(req.Post, path)
	if err != nil {
		return nil, err
	}
	result, ok := res.Result().(*RecordCreatedResponse)
	if !ok {
		return nil, fmt.Errorf("failed to marshal response")
	}
	return result, nil
}

// PatchObject updates an existing object.
func (client *Client) PatchObject(path string, data map[string]any) error {
	err := client.prepare()
	if err != nil {
		return err
	}
	req := client.httpClient.R().SetBody(data)
	res, err := client.Do(req.Patch, path)
	if err != nil {
		return err
	}
	if res.IsError() {
		return fmt.Errorf("failed to update object at path '%s'", path)
	}
	return nil
}

// DeleteObject deletes an existing object.
func (client *Client) DeleteObject(path string) error {
	err := client.prepare()
	if err != nil {
		return err
	}
	req := client.httpClient.R()
	res, err := client.Do(req.Delete, path)
	if err != nil {
		return err
	}
	if res.IsError() {
		return fmt.Errorf("failed to delete object at path '%s'", path)
	}
	return nil
}
