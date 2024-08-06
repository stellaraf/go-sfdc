package sfdc

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"go.stellar.af/go-sfdc/internal/util"
	"go.stellar.af/go-utils/slice"
)

type BulkClient struct {
	main *Client
}

func NewBulkClient(client *Client) *BulkClient {
	return &BulkClient{
		main: client,
	}
}

func (bc *BulkClient) NewInsertJob(object string) *BulkJob {
	return &BulkJob{
		Object:      object,
		ContentType: "CSV",
		Operation:   "insert",
		LineEnding:  "LF",
	}
}

func (bc *BulkClient) NewUpsertJob(object string, extIDField string) *BulkJob {
	return &BulkJob{
		Object:              object,
		ContentType:         "CSV",
		Operation:           "upsert",
		LineEnding:          "LF",
		ExternalIDFieldName: extIDField,
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

func (bc *BulkClient) JobStatus(id string) (*BulkJobStatus, error) {
	err := bc.main.prepare()
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf(PATH_BULK_INGEST_COMPLETE, API_VERSION, id)
	req := bc.main.httpClient.R().SetHeader("accept", "application/json").SetHeader("content-type", "application/json")
	req.SetResult(&BulkJobStatus{})
	res, err := req.Get(path)
	if err != nil {
		return nil, err
	}
	if !res.IsSuccess() {
		return nil, fmt.Errorf("failed to get job %s status: %s", id, string(res.Body()))
	}
	data, ok := res.Result().(*BulkJobStatus)
	if !ok {
		return nil, fmt.Errorf("failed to parse job status response")
	}
	return data, nil
}

func (bc *BulkClient) jobIsState(id, state string) (bool, error) {
	status, err := bc.JobStatus(id)
	if err != nil {
		return false, err
	}
	return status.State == state, nil
}

func (bc *BulkClient) JobIsComplete(id string) (bool, error) {
	return bc.jobIsState(id, "JobComplete")
}

func (bc *BulkClient) JobIsFailed(id string) (bool, error) {
	return bc.jobIsState(id, "Failed")
}

func (bc *BulkClient) JobIsInProgress(id string) (bool, error) {
	return bc.jobIsState(id, "InProgress")
}

func (bc *BulkClient) JobIsOpen(id string) (bool, error) {
	return bc.jobIsState(id, "Open")
}

func (bc *BulkClient) JobIsAborted(id string) (bool, error) {
	return bc.jobIsState(id, "Aborted")
}

func (bc *BulkClient) Insert(bulkJob *BulkJob, data any, cf ...map[string]any) (*BulkJobCompleteResponse, error) {
	bulkData, err := MarshalCSV(data, cf...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert data to CSV")
	}
	return bc.InsertRaw(bulkJob, bulkData)
}

func (bc *BulkClient) InsertRaw(bulkJob *BulkJob, data string) (*BulkJobCompleteResponse, error) {
	job, err := bc.jobReq(bulkJob)
	if err != nil {
		return nil, err
	}
	err = bc.uploadData(job.ID, data)
	if err != nil {
		return nil, err
	}
	complete, err := bc.complete(job.ID)
	if err != nil {
		return nil, err
	}
	return complete, nil
}

func (bc *BulkClient) Upsert(bulkJob *BulkJob, data any, cf ...map[string]any) (*BulkJobCompleteResponse, error) {
	return bc.Insert(bulkJob, data, cf...)
}

func (bc *BulkClient) UpsertRaw(bulkJob *BulkJob, data string) (*BulkJobCompleteResponse, error) {
	return bc.InsertRaw(bulkJob, data)
}

func (bc *BulkClient) InsertMultiple(bulkJob *BulkJob, data any) (*BulkJobCompleteResponse, error) {
	csv, err := MarshalCSVSlice(data)
	if err != nil {
		return nil, err
	}
	return bc.InsertRaw(bulkJob, csv)
}

func (bc *BulkClient) UpsertMultiple(bulkJob *BulkJob, data any) (*BulkJobCompleteResponse, error) {
	return bc.InsertMultiple(bulkJob, data)
}

func mapToCSV(m map[string]any) (string, error) {
	bm := util.SortMap(m)
	headers := make([]string, 0, len(bm))
	values := make([]string, 0, len(bm))
	for k, v := range bm {
		headers = append(headers, k)
		values = append(values, fmt.Sprint(v))
	}
	buffer := new(bytes.Buffer)
	writer := csv.NewWriter(buffer)
	err := writer.Write(headers)
	if err != nil {
		return "", err
	}
	err = writer.Write(values)
	if err != nil {
		return "", err
	}
	writer.Flush()
	out := buffer.String()
	return strings.TrimLeft(out, "\n"), nil
}

func MarshalCSV(base any, customFields ...map[string]any) (string, error) {
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
	return mapToCSV(bm)
}

func MarshalCSVSlice(base any) (string, error) {
	b, err := json.Marshal(&base)
	if err != nil {
		return "", errors.Wrap(err, "failed initial json encoding")
	}
	sl := make([]map[string]any, 0)
	err = json.Unmarshal(b, &sl)
	if err != nil {
		return "", errors.Wrap(err, "failed initial json decoding")
	}

	headers := make([]string, 0)
	rows := make([][]string, 0)
	for _, s := range sl {
		for k := range s {
			headers = append(headers, k)
		}
	}
	headers = slice.Dedup(headers)
	for _, m := range sl {
		values := make([]string, 0, len(headers))
		for _, h := range headers {
			values = append(values, fmt.Sprint(m[h]))
		}
		rows = append(rows, values)
	}
	buffer := new(bytes.Buffer)
	writer := csv.NewWriter(buffer)
	err = writer.Write(headers)
	if err != nil {
		return "", err
	}
	for _, row := range rows {
		err = writer.Write(row)
		if err != nil {
			return "", err
		}
	}
	writer.Flush()
	out := buffer.String()
	return strings.TrimLeft(out, "\n"), nil
}
