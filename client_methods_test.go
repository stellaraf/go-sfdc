package sfdc_test

import (
	"fmt"
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
	newCase, err := Client.CreateCase(caseData)
	require.NoError(t, err)
	require.NotNil(t, newCase)
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

func Test_GetObject(t *testing.T) {
	t.Parallel()
	path := fmt.Sprintf("/services/data/%s/sobjects/Account/%s", sfdc.API_VERSION, Env.TestData.AccountID)
	account, err := Client.GetObject(path)
	require.NoError(t, err)
	id := account.GetString("Id")
	assert.Equal(t, Env.TestData.AccountID, id, "mismatched id")
	name := account.GetString("Name")
	assert.Equal(t, Env.TestData.AccountName, name, "mismatched name")
	assert.Equal(t, "", account.GetString("AccountSource"))
}

func Test_PostObject(t *testing.T) {
	t.Parallel()
	t.Run("create contact", func(t *testing.T) {
		t.Parallel()
		// Create (POST)
		path := fmt.Sprintf("/services/data/%s/sobjects/Contact", sfdc.API_VERSION)
		res, err := Client.PostObject(path, map[string]any{"LastName": t.Name()})
		require.NoError(t, err, "failed to create")
		require.True(t, res.Success, "success=%v", res.Success)

		// Update (PATCH)
		objPath := fmt.Sprintf("%s/%s", path, res.ID)
		err = Client.PatchObject(objPath, map[string]any{"FirstName": t.Name()})
		require.NoError(t, err, "failed to update")

		// Verify (GET)
		obj, err := Client.GetObject(objPath)
		require.NoError(t, err, "failed to get after update")
		assert.Equal(t, t.Name(), obj.GetString("FirstName"))

		// Delete (DELETE)
		err = Client.DeleteObject(objPath)
		require.NoError(t, err, "failed to delete")
	})
}
