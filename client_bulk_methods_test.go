package sfdc_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.stellar.af/go-sfdc"
)

func GetObject[R any](maxTime time.Duration, getter func() (*R, error)) (*R, error) {
	bo := backoff.NewExponentialBackOff(backoff.WithMaxElapsedTime(maxTime))
	backoffFn := func() (*R, error) {
		res, err := getter()
		if err != nil {
			return nil, err
		}
		return res, nil
	}
	return backoff.RetryWithData(backoffFn, bo)
}

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

	q := sfdc.SOQL().
		Select("Id").
		From("FeedItem").
		Where("ParentId", sfdc.EQUALS, newCase.ID).
		Where("Title", sfdc.EQUALS, title)
	soql := sfdc.NewSOQL[sfdc.ObjectID](Client)
	res, err := GetObject(60*time.Second, func() (*sfdc.RecordResponse[sfdc.ObjectID], error) {
		res, err := soql.Query(q)
		if err != nil {
			return nil, err
		}
		if res.TotalSize == 0 {
			return nil, fmt.Errorf("feed item not created yet")
		}
		return &res, nil
	})

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
	case2, err := GetObject(60*time.Second, func() (*sfdc.Case, error) {
		res, err := Client.Case(case1.ID)
		if err != nil {
			return nil, err
		}
		if res.Description != "2" {
			return nil, fmt.Errorf("not updated yet")
		}
		return res, nil

	})
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

	case2, err := GetObject(60*time.Second, func() (*sfdc.Case, error) {
		res, err := Client.Case(case1.ID)
		if err != nil {
			return nil, err
		}
		if res.Status != "Closed" {
			return nil, fmt.Errorf("not closed yet")
		}
		return res, nil
	})
	require.NoError(t, err)
	assert.Equal(t, "Closed", case2.Status)
	t.Cleanup(func() {
		if case2.Status != "Closed" {
			err := Client.UpdateCase(case1.ID, &sfdc.CaseUpdate{Status: "Canceled"})
			require.NoError(t, err)
		}
	})
}
