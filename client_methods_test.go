package sfdc_test

import (
	"testing"

	"github.com/stellaraf/go-sfdc/types"
	"github.com/stretchr/testify/assert"
)

func Test_AccountContact(t *testing.T) {
	t.Run("account contact", func(t *testing.T) {
		contact, err := Client.AccountContact(Env.TestData.AccountID)
		assert.NoError(t, err)
		assert.Equal(t, Env.TestData.ContactID, contact.ID)
	})
}

func Test_PostToCase(t *testing.T) {
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
	newCase, _ := Client.CreateCase(caseData)
	t.Run("post plain text update", func(t *testing.T) {
		postResult, err := Client.PostToCase(newCase.ID, "go-sfdc test plain text comment", nil)
		assert.NoError(t, err)
		feedItem, err := Client.FeedItem(postResult.ID)
		assert.NoError(t, err)
		assert.True(t, postResult.Success)
		assert.Equal(t, postResult.ID, feedItem.ID)
		assert.Greater(t, len(postResult.ID), 1)
	})
	t.Run("post html update", func(t *testing.T) {
		body := "<b>go-sfdc test HTML comment</b>"
		postResult, err := Client.PostToCase(newCase.ID, body, &types.FeedItemOptions{
			IsRichText: true,
		})
		assert.NoError(t, err)
		feedItem, err := Client.FeedItem(postResult.ID)
		assert.NoError(t, err)
		assert.True(t, postResult.Success)
		assert.Equal(t, postResult.ID, feedItem.ID)
		assert.Greater(t, len(postResult.ID), 1)
		assert.True(t, feedItem.IsRichText)
		assert.Equal(t, body, feedItem.Body)
	})
	t.Cleanup(func() {
		Client.UpdateCase(newCase.ID, &types.CaseUpdate{Status: "Canceled"})
	})
}

func Test_CloseCase(t *testing.T) {
	t.Run("create a case and close it", func(t *testing.T) {
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
		newCase, err := Client.CreateCase(caseData)
		assert.NoError(t, err)
		err = Client.CloseCase(newCase.ID)
		assert.NoError(t, err)
		closedCase, err := Client.Case(newCase.ID)
		assert.NoError(t, err)
		assert.Equal(t, "Closed", closedCase.Status)
	})
}

func Test_CaseByNumber(t *testing.T) {
	t.Run("get a case by its case number", func(t *testing.T) {
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
		res, err := Client.CreateCase(caseData)
		assert.NoError(t, err)
		assert.True(t, res.Success)
		newCase, err := Client.Case(res.ID)
		assert.NoError(t, err)
		caseByNumber, err := Client.CaseByNumber(newCase.CaseNumber)
		assert.NoError(t, err)
		assert.Equal(t, newCase.CaseNumber, caseByNumber.CaseNumber)
		err = Client.UpdateCase(res.ID, &types.CaseUpdate{Status: "Canceled"})
		assert.NoError(t, err)
	})
}
