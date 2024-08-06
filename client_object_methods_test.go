package sfdc_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.stellar.af/go-sfdc"
	"go.stellar.af/go-utils/random"
)

func Test_Account(t *testing.T) {
	t.Run("get account", func(t *testing.T) {
		t.Parallel()
		account, err := Client.Account(Env.TestData.AccountID)
		require.NoError(t, err)
		assert.Equal(t, Env.TestData.AccountID, account.ID)
	})
	t.Run("get account custom field", func(t *testing.T) {
		t.Parallel()
		account, err := Client.Account(Env.TestData.AccountID)
		require.NoError(t, err)
		cfvalue := account.CustomFields[Env.TestData.AccountCustomFieldKey]
		assert.NotEmpty(t, cfvalue)
	})
}

func Test_User(t *testing.T) {
	t.Run("get user", func(t *testing.T) {
		t.Parallel()
		user, err := Client.User(Env.TestData.UserID)
		require.NoError(t, err)
		assert.Equal(t, Env.TestData.UserID, user.ID)
	})
}

func Test_Group(t *testing.T) {
	t.Run("get group", func(t *testing.T) {
		t.Parallel()
		group, err := Client.Group(Env.TestData.GroupID)
		require.NoError(t, err)
		assert.Equal(t, Env.TestData.GroupID, group.ID)
	})
}

func Test_Case(t *testing.T) {
	t.Run("get case", func(t *testing.T) {
		t.Parallel()
		_case, err := Client.Case(Env.TestData.CaseID)
		require.NoError(t, err)
		assert.Equal(t, Env.TestData.CaseID, _case.ID)
	})
}

func Test_ServiceContract(t *testing.T) {
	t.Run("get service contract", func(t *testing.T) {
		t.Parallel()
		contract, err := Client.ServiceContract(Env.TestData.ServiceContractID)
		require.NoError(t, err)
		assert.Equal(t, Env.TestData.ServiceContractID, contract.ID)
	})
}

func Test_Contact(t *testing.T) {
	t.Run("get contact", func(t *testing.T) {
		t.Parallel()
		contact, err := Client.Contact(Env.TestData.ContactID)
		require.NoError(t, err)
		assert.Equal(t, Env.TestData.ContactID, contact.ID)
	})
}

func Test_CreateCase(t *testing.T) {
	t.Run("create and cancel case", func(t *testing.T) {
		t.Parallel()
		subject := createCaseSubject(t)
		caseData := &sfdc.CaseCreate{
			AccountID:   Env.TestData.AccountID,
			Comments:    "go-sfdc unit test case",
			ContactID:   Env.TestData.ContactID,
			Description: subject,
			Origin:      "Web",
			Status:      "New",
			Subject:     subject,
		}
		result, err := Client.CreateCase(caseData)
		require.NoError(t, err)
		assert.True(t, result.Success)
		err = Client.UpdateCase(result.ID, &sfdc.CaseUpdate{Status: "Canceled"})
		require.NoError(t, err)
	})

	t.Run("create a case with custom fields", func(t *testing.T) {
		t.Parallel()
		subject := createCaseSubject(t)
		caseData := &sfdc.CaseCreate{
			AccountID:   Env.TestData.AccountID,
			Comments:    "go-sfdc unit test case",
			ContactID:   Env.TestData.ContactID,
			Description: subject,
			Origin:      "Web",
			Status:      "New",
			Subject:     subject,
		}
		customFieldValue, err := random.String(16)
		require.NoError(t, err)
		customFields := sfdc.CustomFields{Env.TestData.CaseCustomFieldKey: customFieldValue}
		result, err := Client.CreateCase(caseData, customFields)
		require.NoError(t, err)
		assert.True(t, result.Success)
		_case, err := Client.Case(result.ID)
		require.NoError(t, err)
		assert.Equal(t, customFieldValue, _case.CustomFields[Env.TestData.CaseCustomFieldKey])
		err = Client.UpdateCase(result.ID, &sfdc.CaseUpdate{Status: "Canceled"})
		require.NoError(t, err)
	})

}

func Test_CreateFeedItem(t *testing.T) {
	subject := createCaseSubject(t)
	caseData := &sfdc.CaseCreate{
		AccountID:   Env.TestData.AccountID,
		Comments:    "go-sfdc unit test case",
		ContactID:   Env.TestData.ContactID,
		Description: subject,
		Origin:      "Web",
		Status:      "New",
		Subject:     subject,
	}
	newCase, _ := Client.CreateCase(caseData)
	t.Run("base", func(t *testing.T) {
		t.Parallel()
		req := &sfdc.FeedItemOptions{
			Body:     fmt.Sprintf("test at %s", time.Now().Format(time.RFC3339Nano)),
			ParentID: newCase.ID,
			Type:     "TextPost",
			Title:    fmt.Sprintf("Title: %s", subject),
		}
		result, err := Client.CreateFeedItem(req)
		require.NoError(t, err)
		assert.True(t, result.Success)
		feedItem, err := Client.FeedItem(result.ID)
		require.NoError(t, err)
		assert.Equal(t, req.Body, feedItem.Body)
		assert.Equal(t, req.Title, req.Title)
	})
	t.Cleanup(func() {
		Client.UpdateCase(newCase.ID, &sfdc.CaseUpdate{Status: "Canceled"})
	})
}
