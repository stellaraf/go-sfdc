package sfdc

import (
	"fmt"
	"testing"
	"time"

	"github.com/stellaraf/go-sfdc/types"
	"github.com/stretchr/testify/assert"
)

func Test_Account(t *testing.T) {
	t.Run("get account", func(t *testing.T) {
		client, env, err := initClient()
		assert.NoError(t, err)
		account, err := client.Account(env.TestData.AccountID)
		assert.NoError(t, err)
		assert.Equal(t, env.TestData.AccountID, account.ID)
	})
}

func Test_User(t *testing.T) {
	t.Run("get user", func(t *testing.T) {
		client, env, err := initClient()
		assert.NoError(t, err)
		user, err := client.User(env.TestData.UserID)
		assert.NoError(t, err)
		assert.Equal(t, env.TestData.UserID, user.ID)
	})
}

func Test_Group(t *testing.T) {
	t.Run("get group", func(t *testing.T) {
		client, env, err := initClient()
		assert.NoError(t, err)
		group, err := client.Group(env.TestData.GroupID)
		assert.NoError(t, err)
		assert.Equal(t, env.TestData.GroupID, group.ID)
	})
}

func Test_Case(t *testing.T) {
	t.Run("get case", func(t *testing.T) {
		client, env, err := initClient()
		assert.NoError(t, err)
		_case, err := client.Case(env.TestData.CaseID)
		assert.NoError(t, err)
		assert.Equal(t, env.TestData.CaseID, _case.ID)
	})
}

func Test_Contact(t *testing.T) {
	t.Run("get contact", func(t *testing.T) {
		client, env, err := initClient()
		assert.NoError(t, err)
		contact, err := client.Contact(env.TestData.ContactID)
		assert.NoError(t, err)
		assert.Equal(t, env.TestData.ContactID, contact.ID)
	})
}

func Test_CreateCase(t *testing.T) {
	t.Run("create and cancel case", func(t *testing.T) {
		client, env, err := initClient()
		assert.NoError(t, err)
		now := time.Now()
		subject := fmt.Sprintf("go-sfdc Test_CreateCase case %s", now.Format(time.RFC3339))
		caseData := &types.CaseCreate{
			AccountID:   env.TestData.AccountID,
			Comments:    "go-sfdc unit test case",
			ContactID:   env.TestData.ContactID,
			Description: subject,
			Origin:      "Web",
			Status:      "New",
			Subject:     subject,
		}
		result, err := client.CreateCase(caseData)
		assert.NoError(t, err)
		assert.True(t, result.Success)
		err = client.UpdateCase(result.ID, &types.CaseUpdate{Status: "Canceled"})
		assert.NoError(t, err)
	})
}
