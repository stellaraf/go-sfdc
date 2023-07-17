package sfdc

import (
	"fmt"
	"testing"
	"time"

	"github.com/stellaraf/go-sfdc/types"
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

func Test_PostToCase(t *testing.T) {
	client, env, _ := initClient()
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
	newCase, _ := client.CreateCase(caseData)
	t.Run("post plain text update", func(t *testing.T) {
		postResult, err := client.PostToCase(newCase.ID, "go-sfdc test plain text comment", nil)
		assert.NoError(t, err)
		feedItem, err := client.FeedItem(postResult.ID)
		assert.NoError(t, err)
		assert.True(t, postResult.Success)
		assert.Equal(t, postResult.ID, feedItem.ID)
		assert.Greater(t, len(postResult.ID), 1)
	})
	t.Run("post html update", func(t *testing.T) {
		body := "<b>go-sfdc test HTML comment</b>"
		postResult, err := client.PostToCase(newCase.ID, body, &types.FeedItemOptions{
			IsRichText: true,
		})
		assert.NoError(t, err)
		feedItem, err := client.FeedItem(postResult.ID)
		assert.NoError(t, err)
		assert.True(t, postResult.Success)
		assert.Equal(t, postResult.ID, feedItem.ID)
		assert.Greater(t, len(postResult.ID), 1)
		assert.True(t, feedItem.IsRichText)
		assert.Equal(t, body, feedItem.Body)
	})
	t.Cleanup(func() {
		client.UpdateCase(newCase.ID, &types.CaseUpdate{Status: "Canceled"})
	})
}

func Test_CloseCase(t *testing.T) {
	t.Run("create a case and close it", func(t *testing.T) {
		client, env, err := initClient()
		assert.NoError(t, err)
		now := time.Now()
		subject := fmt.Sprintf("go-sfdc Test_CloseCase case %s", now.Format(time.RFC3339))
		caseData := &types.CaseCreate{
			AccountID:   env.TestData.AccountID,
			Comments:    "go-sfdc unit test case",
			ContactID:   env.TestData.ContactID,
			Description: subject,
			Origin:      "Web",
			Status:      "New",
			Subject:     subject,
		}
		newCase, err := client.CreateCase(caseData)
		assert.NoError(t, err)
		err = client.CloseCase(newCase.ID)
		assert.NoError(t, err)
		closedCase, err := client.Case(newCase.ID)
		assert.NoError(t, err)
		assert.Equal(t, "Closed", closedCase.Status)
	})
}

func Test_CaseByNumber(t *testing.T) {
	t.Run("get a case by its case number", func(t *testing.T) {
		client, env, err := initClient()
		assert.NoError(t, err)
		now := time.Now()
		subject := fmt.Sprintf("go-sfdc Test_CaseByNumber case %s", now.Format(time.RFC3339))
		caseData := &types.CaseCreate{
			AccountID:   env.TestData.AccountID,
			Comments:    "go-sfdc unit test case",
			ContactID:   env.TestData.ContactID,
			Description: subject,
			Origin:      "Web",
			Status:      "New",
			Subject:     subject,
		}
		res, err := client.CreateCase(caseData)
		assert.NoError(t, err)
		assert.True(t, res.Success)
		newCase, err := client.Case(res.ID)
		assert.NoError(t, err)
		caseByNumber, err := client.CaseByNumber(newCase.CaseNumber)
		assert.NoError(t, err)
		assert.Equal(t, newCase.CaseNumber, caseByNumber.CaseNumber)
		err = client.UpdateCase(res.ID, &types.CaseUpdate{Status: "Canceled"})
		assert.NoError(t, err)
	})
}
