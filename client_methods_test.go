package sfdc_test

import (
	"testing"

	"github.com/stellaraf/go-sfdc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_PostToCase(t *testing.T) {
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
	t.Run("post plain text update", func(t *testing.T) {
		t.Parallel()
		postResult, err := Client.PostToCase(newCase.ID, "go-sfdc test plain text comment", nil)
		require.NoError(t, err)
		feedItem, err := Client.FeedItem(postResult.ID)
		require.NoError(t, err)
		assert.True(t, postResult.Success)
		assert.Equal(t, postResult.ID, feedItem.ID)
		assert.Greater(t, len(postResult.ID), 1)
	})
	t.Run("post html update", func(t *testing.T) {
		t.Parallel()
		body := "<b>go-sfdc test HTML comment</b>"
		postResult, err := Client.PostToCase(newCase.ID, body, &sfdc.FeedItemOptions{
			IsRichText: true,
		})
		require.NoError(t, err)
		feedItem, err := Client.FeedItem(postResult.ID)
		require.NoError(t, err)
		assert.True(t, postResult.Success)
		assert.Equal(t, postResult.ID, feedItem.ID)
		assert.Greater(t, len(postResult.ID), 1)
		assert.True(t, feedItem.IsRichText)
		assert.Equal(t, body, feedItem.Body)
	})
	t.Cleanup(func() {
		Client.UpdateCase(newCase.ID, &sfdc.CaseUpdate{Status: "Canceled"})
	})
}

func Test_CloseCase(t *testing.T) {
	t.Run("create a case and close it", func(t *testing.T) {
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
		newCase, err := Client.CreateCase(caseData)
		require.NoError(t, err)
		err = Client.CloseCase(newCase.ID)
		require.NoError(t, err)
		closedCase, err := Client.Case(newCase.ID)
		require.NoError(t, err)
		assert.Equal(t, "Closed", closedCase.Status)
	})
}
