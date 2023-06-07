package sfdc

import (
	"testing"

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
