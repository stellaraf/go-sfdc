// Methods that handle methods through the Bulk API.
package sfdc

import "fmt"

func (client *Client) SendFeedItem(item *FeedItemOptions) (*BulkJobCompleteResponse, error) {
	bulk := client.Bulk()
	job := bulk.NewInsertJob("FeedItem")
	res, err := bulk.Insert(job, item)
	if err != nil {
		return nil, err
	}
	if res.State != "UploadComplete" {
		return nil, fmt.Errorf("failed to send job %s; state=%s operation=%s", res.ID, res.State, res.Operation)
	}
	return res, nil
}

func (client *Client) SendCaseUpdate(id string, data *Case, cf ...map[string]any) (*BulkJobCompleteResponse, error) {
	data.ID = id
	bulk := client.Bulk()
	job := bulk.NewUpsertJob("Case", "Id")
	res, err := bulk.Upsert(job, data, cf...)
	if err != nil {
		return nil, err
	}
	if res.State != "UploadComplete" {
		return nil, fmt.Errorf("failed to send job %s; state=%s operation=%s", res.ID, res.State, res.Operation)
	}
	return res, nil
}

func (client *Client) SendCloseCase(id string) error {
	data := &Case{Status: "Closed"}
	data.ID = id
	bulk := client.Bulk()
	job := bulk.NewUpsertJob("Case", "Id")
	res, err := bulk.Upsert(job, data)
	if err != nil {
		return err
	}
	if res.State != "UploadComplete" {
		return fmt.Errorf("failed to send job %s; state=%s operation=%s", res.ID, res.State, res.Operation)
	}
	return nil
}
