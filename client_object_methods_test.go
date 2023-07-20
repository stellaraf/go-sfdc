package sfdc_test

import (
	"fmt"
	"testing"

	"github.com/stellaraf/go-sfdc/types"
	"github.com/stellaraf/go-utils"
	"github.com/stretchr/testify/assert"
)

func Test_Account(t *testing.T) {
	t.Run("get account", func(t *testing.T) {
		account, err := Client.Account(Env.TestData.AccountID)
		assert.NoError(t, err)
		assert.Equal(t, Env.TestData.AccountID, account.ID)
	})
	t.Run("get account custom field", func(t *testing.T) {
		account, err := Client.Account(Env.TestData.AccountID)
		assert.NoError(t, err)
		cfvalue := account.CustomFields[Env.TestData.AccountCustomFieldKey]
		assert.NotEmpty(t, cfvalue)
	})
}

func Test_User(t *testing.T) {
	t.Run("get user", func(t *testing.T) {
		user, err := Client.User(Env.TestData.UserID)
		assert.NoError(t, err)
		assert.Equal(t, Env.TestData.UserID, user.ID)
	})
}

func Test_Group(t *testing.T) {
	t.Run("get group", func(t *testing.T) {
		group, err := Client.Group(Env.TestData.GroupID)
		assert.NoError(t, err)
		assert.Equal(t, Env.TestData.GroupID, group.ID)
	})
}

func Test_Case(t *testing.T) {
	t.Run("get case", func(t *testing.T) {
		_case, err := Client.Case(Env.TestData.CaseID)
		assert.NoError(t, err)
		assert.Equal(t, Env.TestData.CaseID, _case.ID)
	})
}

func Test_Contact(t *testing.T) {
	t.Run("get contact", func(t *testing.T) {
		contact, err := Client.Contact(Env.TestData.ContactID)
		assert.NoError(t, err)
		assert.Equal(t, Env.TestData.ContactID, contact.ID)
	})
}

func Test_CreateCase(t *testing.T) {
	t.Run("create and cancel case", func(t *testing.T) {
		subject := createCaseSubject(t)
		caseData := &types.CaseCreate{
			AccountID:   Env.TestData.AccountID,
			Comments:    "go-sfdc unit test case",
			ContactID:   Env.TestData.ContactID,
			Description: subject,
			Origin:      "Web",
			Status:      "New",
			Subject:     subject,
		}
		result, err := Client.CreateCase(caseData)
		assert.NoError(t, err)
		assert.True(t, result.Success)
		err = Client.UpdateCase(result.ID, &types.CaseUpdate{Status: "Canceled"})
		assert.NoError(t, err)
	})

	t.Run("create a case with custom fields", func(t *testing.T) {
		subject := createCaseSubject(t)
		caseData := &types.CaseCreate{
			AccountID:   Env.TestData.AccountID,
			Comments:    "go-sfdc unit test case",
			ContactID:   Env.TestData.ContactID,
			Description: subject,
			Origin:      "Web",
			Status:      "New",
			Subject:     subject,
		}
		customFieldValue, err := utils.RandomString(16)
		assert.NoError(t, err)
		customFields := types.CustomFields{Env.TestData.CaseCustomFieldKey: customFieldValue}
		result, err := Client.CreateCase(caseData, customFields)
		assert.NoError(t, err)
		assert.True(t, result.Success)
		_case, err := Client.Case(result.ID)
		assert.NoError(t, err)
		assert.Equal(t, customFieldValue, _case.CustomFields[Env.TestData.CaseCustomFieldKey])
		err = Client.UpdateCase(result.ID, &types.CaseUpdate{Status: "Canceled"})
		assert.NoError(t, err)
	})

	t.Run("create case, change owner, verify owner", func(t *testing.T) {
		subject := createCaseSubject(t)
		caseData := &types.CaseCreate{
			AccountID:   Env.TestData.AccountID,
			Comments:    t.Name(),
			ContactID:   Env.TestData.ContactID,
			Description: subject,
			Origin:      "Web",
			Status:      "New",
			Subject:     subject,
		}
		newCaseResult, err := Client.CreateCase(caseData)
		assert.NoError(t, err)
		assert.True(t, newCaseResult.Success)
		err = Client.UpdateCase(newCaseResult.ID, &types.CaseUpdate{
			Comments: fmt.Sprintf("changing owner to %s", Env.TestData.UserID),
			OwnerID:  Env.TestData.UserID,
		})
		assert.NoError(t, err)
		err = Client.CloseCase(newCaseResult.ID)
		assert.NoError(t, err)
		closedCase, err := Client.Case(newCaseResult.ID)
		assert.NoError(t, err)
		assert.Equal(t, Env.TestData.UserID, closedCase.OwnerID)
	})
}
