package sfdc_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.stellar.af/go-sfdc"
)

func BulkJobStatus(client *sfdc.BulkClient, jobID string, maxTime time.Duration) (*sfdc.BulkJobStatus, error) {
	bo := backoff.NewExponentialBackOff(backoff.WithMaxElapsedTime(maxTime))
	backoffFn := func() (*sfdc.BulkJobStatus, error) {
		status, err := client.JobStatus(jobID)
		if err != nil {
			return nil, err
		}
		if status.State != "JobComplete" {
			return nil, fmt.Errorf("not complete yet")
		}
		return status, nil
	}
	return backoff.RetryWithData(backoffFn, bo)
}

func TestBulkClient_Insert(t *testing.T) {
	t.Parallel()
	client := sfdc.NewBulkClient(Client)
	contact := sfdc.Contact{LastName: t.Name(), AccountID: Env.TestData.AccountID}
	job := client.NewInsertJob("Contact")
	res, err := client.Insert(job, &contact)
	require.NoError(t, err)
	assert.NotEmpty(t, res.ID)
	status, err := BulkJobStatus(client, res.ID, 60*time.Second)
	require.NoError(t, err)
	t.Cleanup(func() {
		if status.State == "JobComplete" {
			soql := sfdc.NewSOQL[sfdc.ObjectID](Client)
			q := sfdc.SOQL().Select("Id").From("Contact").Where("LastName", sfdc.EQUALS, t.Name())
			res, err := soql.Query(q)
			require.NoError(t, err)
			for _, record := range res.Records {
				path := fmt.Sprintf(sfdc.PATH_CONTACT, sfdc.API_VERSION) + fmt.Sprintf("/%s", record.ID)
				err := Client.DeleteObject(path)
				require.NoError(t, err)
			}
		}
	})
}

func TestBulkClient_InsertMultiple(t *testing.T) {
	t.Parallel()
	client := sfdc.NewBulkClient(Client)
	names := make([]string, 10)
	now := time.Now()
	lastName := now.Format("20060106_150405_00")
	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("%s--%d", lastName, i)
		names[i] = name
	}
	contacts := make([]sfdc.Contact, len(names))
	for i := 0; i < len(names); i++ {
		contact := sfdc.Contact{FirstName: t.Name(), LastName: names[i], AccountID: Env.TestData.AccountID}
		contacts[i] = contact
	}
	job := client.NewInsertJob("Contact")
	res, err := client.InsertMultiple(job, contacts)
	require.NoError(t, err)
	require.NotEmpty(t, res.ID)
	status, err := BulkJobStatus(client, res.ID, 60*time.Second)
	require.NoError(t, err)
	if status.State == "Failed" {
		t.Logf("state=%s failed=%d processed=%d", status.State, status.NumberRecordsFailed, status.NumberRecordsProcessed)
	}
	require.Equal(t, "JobComplete", status.State)
	assert.Equal(t, 0, status.NumberRecordsFailed)
	t.Cleanup(func() {
		if status.State == "JobComplete" {
			soql := sfdc.NewSOQL[sfdc.ObjectID](Client)
			q := sfdc.SOQL().Select("Id").From("Contact").Where("FirstName", sfdc.EQUALS, t.Name()).Where("LastName", sfdc.CONTAINS, lastName)
			res, err := soql.Query(q)
			require.NoError(t, err)
			for _, record := range res.Records {
				path := fmt.Sprintf(sfdc.PATH_CONTACT, sfdc.API_VERSION) + fmt.Sprintf("/%s", record.ID)
				err := Client.DeleteObject(path)
				require.NoError(t, err)
			}
			assert.Len(t, res.Records, len(contacts))
		}
	})
}

func Test_MarshalCSV(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		t.Parallel()
		contact := sfdc.Contact{Name: t.Name()}
		csv, err := sfdc.MarshalCSV(&contact)
		contains := fmt.Sprintf("Name\n%s\n", t.Name())
		require.NoError(t, err)
		assert.Contains(t, csv, contains)
	})
	t.Run("custom fields", func(t *testing.T) {
		t.Parallel()
		contact := sfdc.Contact{Name: "John", CustomFields: map[string]any{}}
		contact.CustomFields["Field__c"] = "Value"
		csv, err := sfdc.MarshalCSV(&contact, contact.CustomFields)
		require.NoError(t, err)
		if !strings.Contains(csv, "Field__c,Name\nValue,John\n") && !strings.Contains(csv, "Name,Field__c\nJohn,Value\n") {
			assert.FailNow(t, "mismatching CSV")
		}
	})
}

func Test_MarshalCSVSlice(t *testing.T) {
	t.Parallel()
	names := make([]string, 10)
	now := time.Now()
	lastName := now.Format("20060106_150405_00")
	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("%s--%d", lastName, i)
		names[i] = name
	}
	contacts := make([]sfdc.Contact, 0, len(names))
	for _, ln := range names {
		contact := sfdc.Contact{FirstName: t.Name(), LastName: ln, AccountID: Env.TestData.AccountID}
		contacts = append(contacts, contact)
	}
	res, err := sfdc.MarshalCSVSlice(contacts)
	require.NoError(t, err)
	lines := strings.Split(res, "\n")
	assert.Equal(t, len(contacts), len(lines)-2)
}
