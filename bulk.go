package sfdc

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"github.com/stellaraf/go-sfdc/internal/util"
)

type BulkClient struct {
	main *Client
}

func NewBulkClient(client *Client) *BulkClient {
	return &BulkClient{
		main: client,
	}
}

func (bc *BulkClient) NewJob(object string) *BulkJob {
	return &BulkJob{
		Object:      object,
		ContentType: "CSV",
		Operation:   "insert",
		LineEnding:  "LF",
	}
}

func (bc *BulkClient) jobReq(job *BulkJob) (*BulkJobResponse, error) {
	err := bc.main.prepare()
	if err != nil {
		return nil, err
	}
	req := bc.main.httpClient.R().SetHeader("accept", "application/json").SetHeader("content-type", "application/json")
	req.SetBody(job)
	req.SetResult(&BulkJobResponse{})
	res, err := req.Post(fmt.Sprintf(PATH_BULK_INGEST, API_VERSION))
	if err != nil {
		return nil, err
	}
	if !res.IsSuccess() {
		return nil, fmt.Errorf("failed to create job: %s", string(res.Body()))
	}
	jobData, ok := res.Result().(*BulkJobResponse)
	if !ok {
		return nil, fmt.Errorf("failed to parse job response")
	}
	return jobData, nil
}

func (bc *BulkClient) uploadData(id string, data any) error {
	err := bc.main.prepare()
	if err != nil {
		return err
	}
	path := fmt.Sprintf(PATH_BULK_INGEST, API_VERSION) + fmt.Sprintf("/%s", id)
	req := bc.main.httpClient.R().SetHeader("accept", "application/json").SetHeader("content-type", "text/csv")
	req.SetBody(data)
	res, err := req.Put(path + "/batches")
	if err != nil {
		return err
	}
	if !res.IsSuccess() {
		return fmt.Errorf("failed to upload data to job %s: %s", id, string(res.Body()))
	}
	return nil
}

func (bc *BulkClient) complete(id string) (*BulkJobCompleteResponse, error) {
	err := bc.main.prepare()
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf(PATH_BULK_INGEST, API_VERSION) + fmt.Sprintf("/%s", id)
	req := bc.main.httpClient.R().SetHeader("accept", "application/json").SetHeader("content-type", "application/json")
	req.SetBody(&BulkJobCompleteData)
	req.SetResult(&BulkJobCompleteResponse{})
	res, err := req.Patch(path)
	if err != nil {
		return nil, err
	}
	if !res.IsSuccess() {
		return nil, fmt.Errorf("failed to mark job %s as complete: %s", id, string(res.Body()))
	}
	data, ok := res.Result().(*BulkJobCompleteResponse)
	if !ok {
		return nil, fmt.Errorf("failed to parse job completion response")
	}
	return data, nil
}

func (bc *BulkClient) Insert(bulkJob *BulkJob, data any) (*BulkJobCompleteResponse, error) {
	bulkData, err := bulkJob.CSV(data)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert data to CSV")
	}
	job, err := bc.jobReq(bulkJob)
	if err != nil {
		return nil, err
	}
	err = bc.uploadData(job.ID, bulkData)
	if err != nil {
		return nil, err
	}
	complete, err := bc.complete(job.ID)
	if err != nil {
		return nil, err
	}
	return complete, nil

}

func MarshalCSV(base any, customFields ...CustomFields) (string, error) {
	b, err := json.Marshal(&base)
	if err != nil {
		return "", errors.Wrap(err, "failed initial json encoding")
	}
	bm := make(map[string]any)
	err = json.Unmarshal(b, &bm)
	if err != nil {
		return "", errors.Wrap(err, "failed initial json decoding")
	}
	for _, cf := range customFields {
		for k, v := range cf {
			bm[k] = v
		}
	}
	bm = util.SortMap(bm)
	headers := make([]string, 0, len(bm))
	values := make([]string, 0, len(bm))
	for k, v := range bm {
		headers = append(headers, k)
		values = append(values, fmt.Sprint(v))
	}
	buffer := new(bytes.Buffer)
	writer := csv.NewWriter(buffer)
	err = writer.Write(headers)
	if err != nil {
		return "", err
	}
	err = writer.Write(values)
	if err != nil {
		return "", err
	}
	writer.Flush()
	return buffer.String(), nil
}
