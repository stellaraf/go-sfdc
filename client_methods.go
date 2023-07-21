// Methods that consume object _and_ SOQL methods, or that handle data processing.
package sfdc

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
	res, err := req.Post(path)
	if err != nil {
		return
	}
	result = res.Result().(*RecordCreatedResponse)
	return
}
