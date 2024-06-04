package sfdc_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stellaraf/go-sfdc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBulkClient_Insert(t *testing.T) {
	t.Parallel()
	client := sfdc.NewBulkClient(Client)
	contact := sfdc.Contact{LastName: t.Name(), AccountID: Env.TestData.AccountID}
	job := client.NewJob("Contact")
	res, err := client.Insert(job, &contact)
	require.NoError(t, err)
	assert.NotEmpty(t, res.ID)

	t.Cleanup(func() {
		time.Sleep(time.Second * 2)
		soql := sfdc.NewSOQL[sfdc.ObjectID](Client)
		q := sfdc.SOQL().Select("Id").From("Contact").Where("LastName", sfdc.EQUALS, t.Name())
		res, err := soql.Query(q)
		require.NoError(t, err)
		for _, record := range res.Records {
			path := fmt.Sprintf(sfdc.PATH_CONTACT, sfdc.API_VERSION) + fmt.Sprintf("/%s", record.ID)
			err := Client.DeleteObject(path)
			require.NoError(t, err)
		}
	})
}

func Test_MarshalCSV(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		t.Parallel()
		contact := sfdc.Contact{Name: t.Name()}
		csv, err := sfdc.MarshalCSV(&contact)
		expected := fmt.Sprintf("Name\n%s\n", t.Name())
		require.NoError(t, err)
		assert.Equal(t, expected, csv)
	})
	t.Run("custom fields", func(t *testing.T) {
		t.Parallel()
		contact := sfdc.Contact{Name: "John", CustomFields: map[string]any{}}
		contact.CustomFields["Field__c"] = "Value"
		csv, err := sfdc.MarshalCSV(&contact, contact.CustomFields)
		require.NoError(t, err)
		assert.Equal(t, "Field__c,Name\nValue,John\n", csv)
	})
}
