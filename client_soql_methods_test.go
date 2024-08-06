package sfdc_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.stellar.af/go-sfdc"
)

func Test_OpenCases(t *testing.T) {
	t.Run("soql open cases", func(t *testing.T) {
		t.Parallel()
		cases, err := Client.OpenCases()
		require.NoError(t, err)
		assert.True(t, len(cases) > 0)
	})
}

func Test_UserName(t *testing.T) {
	t.Run("user name by ID", func(t *testing.T) {
		t.Parallel()
		user, err := Client.User(Env.TestData.UserID)
		require.NoError(t, err)
		name, err := Client.UserName(Env.TestData.UserID)
		require.NoError(t, err)
		assert.Equal(t, user.Name, name)
	})
}

func Test_GroupName(t *testing.T) {
	t.Run("group name by ID", func(t *testing.T) {
		t.Parallel()
		group, err := Client.Group(Env.TestData.GroupID)
		require.NoError(t, err)
		name, err := Client.GroupName(Env.TestData.GroupID)
		require.NoError(t, err)
		assert.Equal(t, group.Name, name)
	})
}

func Test_AccountIDFromName(t *testing.T) {
	t.Run("account ID from account name", func(t *testing.T) {
		t.Parallel()
		account, err := Client.Account(Env.TestData.AccountID)
		require.NoError(t, err)
		id, err := Client.AccountIDFromName(Env.TestData.AccountName)
		require.NoError(t, err)
		assert.Equal(t, account.ID, id)
	})
}

func Test_Customers(t *testing.T) {
	t.Run("accounts with type customer", func(t *testing.T) {
		t.Parallel()
		customers, err := Client.Customers()
		require.NoError(t, err)
		assert.True(t, len(customers) > 0)
	})
}

func Test_CaseByNumber(t *testing.T) {
	t.Run("get a case by its case number", func(t *testing.T) {
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
		res, err := Client.CreateCase(caseData)
		require.NoError(t, err)
		assert.True(t, res.Success)
		newCase, err := Client.Case(res.ID)
		require.NoError(t, err)
		caseByNumber, err := Client.CaseByNumber(newCase.CaseNumber)
		require.NoError(t, err)
		assert.Equal(t, newCase.CaseNumber, caseByNumber.CaseNumber)
		err = Client.UpdateCase(res.ID, &sfdc.CaseUpdate{Status: "Canceled"})
		require.NoError(t, err)
	})
}

func Test_UserByEmail(t *testing.T) {
	t.Run("get a user by its email address", func(t *testing.T) {
		t.Parallel()
		result, err := Client.UserByEmail(Env.TestData.UserEmail)
		require.NoError(t, err)
		assert.Equal(t, Env.TestData.UserEmail, result.Email)
	})
}
