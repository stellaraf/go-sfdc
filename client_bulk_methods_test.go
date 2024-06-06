package sfdc_test

import (
	"testing"
	"time"

	"github.com/stellaraf/go-sfdc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_SendFeedItem(t *testing.T) {
	subj := createCaseSubject(t)
	newCase, err := Client.CreateCase(&sfdc.CaseCreate{
		AccountID: Env.TestData.AccountID,
		ContactID: Env.TestData.ContactID,
		Subject:   subj,
		Status:    "New",
	})
	require.NoError(t, err, "failed to create case")
	title := subj + "--feed-item"
	_, err = Client.SendFeedItem(&sfdc.FeedItemOptions{
		CreatedByID: Env.TestData.UserID,
		Title:       title,
		Body:        "bulk api test",
		ParentID:    newCase.ID,
	})
	require.NoError(t, err, "failed to send feed item")
	time.Sleep(time.Second * 15)
	q := sfdc.SOQL().
		Select("Id").
		From("FeedItem").
		Where("ParentId", sfdc.EQUALS, newCase.ID).
		Where("Title", sfdc.EQUALS, title)
	soql := sfdc.NewSOQL[sfdc.ObjectID](Client)
	res, err := soql.Query(q)
	require.NoError(t, err, "failed to retrieve feed item")
	assert.Len(t, res.Records, 1)
	t.Cleanup(func() {
		err := Client.UpdateCase(newCase.ID, &sfdc.CaseUpdate{Status: "Canceled"})
		require.NoError(t, err)
	})
}

func TestClient_SendCaseUpdate(t *testing.T) {
	subj := createCaseSubject(t)
	case1, err := Client.CreateCase(&sfdc.CaseCreate{
		AccountID:   Env.TestData.AccountID,
		ContactID:   Env.TestData.ContactID,
		Subject:     subj,
		Status:      "New",
		Description: "1",
	})
	require.NoError(t, err, "failed to create case")

	_, err = Client.SendCaseUpdate(case1.ID, &sfdc.Case{
		Description: "2",
	})
	require.NoError(t, err, "failed to update case")
	time.Sleep(time.Second * 15)
	case2, err := Client.Case(case1.ID)
	require.NoError(t, err)
	assert.Equal(t, "2", case2.Description)
	t.Cleanup(func() {
		err := Client.UpdateCase(case1.ID, &sfdc.CaseUpdate{Status: "Canceled"})
		require.NoError(t, err)
	})
}

func TestClient_SendCloseCase(t *testing.T) {
	subj := createCaseSubject(t)
	case1, err := Client.CreateCase(&sfdc.CaseCreate{
		AccountID: Env.TestData.AccountID,
		ContactID: Env.TestData.ContactID,
		Subject:   subj,
		Status:    "New",
	})
	require.NoError(t, err, "failed to create case")

	err = Client.SendCloseCase(case1.ID)
	require.NoError(t, err, "failed to close case")

	time.Sleep(time.Second * 15)

	case2, err := Client.Case(case1.ID)
	require.NoError(t, err)
	assert.Equal(t, "Closed", case2.Status)
	t.Cleanup(func() {
		if case2.Status != "Closed" {
			err := Client.UpdateCase(case1.ID, &sfdc.CaseUpdate{Status: "Canceled"})
			require.NoError(t, err)
		}
	})
}
