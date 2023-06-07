package sfdc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AccountContact(t *testing.T) {
	t.Run("account contact", func(t *testing.T) {
		client, env, err := initClient()
		assert.NoError(t, err)
		contact, err := client.AccountContact(env.TestData.AccountID)
		assert.NoError(t, err)
		assert.Equal(t, env.TestData.ContactID, contact.ID)
	})
}
