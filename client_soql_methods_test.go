package sfdc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_OpenCases(t *testing.T) {
	t.Run("soql open cases", func(t *testing.T) {
		client, _, err := initClient()
		assert.NoError(t, err)
		cases, err := client.OpenCases()
		assert.NoError(t, err)
		assert.True(t, len(cases.Records) > 0)
	})
}

func Test_UserName(t *testing.T) {
	t.Run("user name by ID", func(t *testing.T) {
		client, env, err := initClient()
		assert.NoError(t, err)
		user, err := client.User(env.TestData.UserID)
		assert.NoError(t, err)
		name, err := client.UserName(env.TestData.UserID)
		assert.NoError(t, err)
		assert.Equal(t, user.Name, name)
	})
}

func Test_GroupName(t *testing.T) {
	t.Run("group name by ID", func(t *testing.T) {
		client, env, err := initClient()
		assert.NoError(t, err)
		group, err := client.Group(env.TestData.GroupID)
		assert.NoError(t, err)
		name, err := client.GroupName(env.TestData.GroupID)
		assert.NoError(t, err)
		assert.Equal(t, group.Name, name)
	})
}

func Test_AccountIDFromName(t *testing.T) {
	t.Run("account ID from account name", func(t *testing.T) {
		client, env, err := initClient()
		assert.NoError(t, err)
		account, err := client.Account(env.TestData.AccountID)
		assert.NoError(t, err)
		id, err := client.AccountIDFromName(env.TestData.AccountName)
		assert.NoError(t, err)
		assert.Equal(t, account.ID, id)
	})
}
