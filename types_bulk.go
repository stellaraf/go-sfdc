package sfdc

type BulkJob struct {
	Object      string `json:"object"`
	ContentType string `json:"contentType"`
	Operation   string `json:"operation"`
	LineEnding  string `json:"lineEnding"`
}

func (job *BulkJob) CSV(data any, cf ...CustomFields) (string, error) {
	return MarshalCSV(data, cf...)
}

type BulkJobComplete struct {
	State string `json:"state"`
}

var BulkJobCompleteData = BulkJobComplete{State: "UploadComplete"}

type BulkJobResponse struct {
	ID              string  `json:"id"`
	Operation       string  `json:"operation"`
	Object          string  `json:"object"`
	CreatedByID     string  `json:"createdById"`
	CreatedDate     string  `json:"createdDate"`
	SystemModstamp  string  `json:"systemModstamp"`
	State           string  `json:"state"`
	ConcurrencyMode string  `json:"concurrencyMode"`
	ContentType     string  `json:"contentType"`
	APIVersion      float64 `json:"apiVersion"`
	ContentURL      string  `json:"contentUrl"`
	LineEnding      string  `json:"lineEnding"`
	ColumnDelimiter string  `json:"columnDelimiter"`
}

type BulkJobCompleteResponse struct {
	ID              string  `json:"id"`
	Operation       string  `json:"operation"`
	Object          string  `json:"object"`
	CreatedByID     string  `json:"createdById"`
	CreatedDate     string  `json:"createdDate"`
	SystemModstamp  string  `json:"systemModstamp"`
	State           string  `json:"state"`
	ConcurrencyMode string  `json:"concurrencyMode"`
	ContentType     string  `json:"contentType"`
	APIVersion      float64 `json:"apiVersion"`
}

type BulkJobStatus struct {
	ID                      string  `json:"id"`
	Operation               string  `json:"operation"`
	Object                  string  `json:"object"`
	CreatedByID             string  `json:"createdById"`
	CreatedDate             string  `json:"createdDate"`
	SystemModstamp          string  `json:"systemModstamp"`
	State                   string  `json:"state"`
	ConcurrencyMode         string  `json:"concurrencyMode"`
	ContentType             string  `json:"contentType"`
	APIVersion              float64 `json:"apiVersion"`
	JobType                 string  `json:"jobType"`
	LineEnding              string  `json:"lineEnding"`
	ColumnDelimiter         string  `json:"columnDelimiter"`
	NumberRecordsProcessed  int     `json:"numberRecordsProcessed"`
	NumberRecordsFailed     int     `json:"numberRecordsFailed"`
	Retries                 int     `json:"retries"`
	TotalProcessingTime     int     `json:"totalProcessingTime"`
	APIActiveProcessingTime int     `json:"apiActiveProcessingTime"`
	ApexProcessingTime      int     `json:"apexProcessingTime"`
}
