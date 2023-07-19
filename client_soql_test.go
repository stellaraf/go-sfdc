package sfdc_test

import (
	"fmt"
	"testing"

	"github.com/stellaraf/go-sfdc"
	"github.com/stellaraf/go-sfdc/types"
	"github.com/stellaraf/go-sfdc/util"
	"github.com/stretchr/testify/assert"
)

func Test_SOQL(t *testing.T) {
	t.Run("open cases", func(t *testing.T) {
		expected := "SELECT Id FROM Case WHERE IsClosed = false LIMIT 10"
		s, err := sfdc.SOQL().Select("Id").From("Case").Where("IsClosed", "=", false).Limit(10).String()
		assert.NoError(t, err)
		assert.Equal(t, expected, s)
	})
	t.Run("multiple where", func(t *testing.T) {
		env, err := util.LoadEnv()
		assert.NoError(t, err)
		expected := fmt.Sprintf("SELECT Id FROM Case WHERE IsClosed = false AND AccountId = '%s'", env.TestData.AccountID)
		s, err := sfdc.SOQL().Select("Id").From("Case").Where("IsClosed", "=", false).Where("AccountId", "=", env.TestData.AccountID).String()
		assert.NoError(t, err)
		assert.Equal(t, expected, s)
	})
	t.Run("where contains", func(t *testing.T) {
		expected := `SELECT Id FROM Case WHERE Description LIKE '%a%'`
		s, err := sfdc.SOQL().Select("Id").From("Case").Where("Description", "contains", "a").String()
		assert.NoError(t, err)
		assert.Equal(t, expected, s)
	})

	t.Run("where contains query", func(t *testing.T) {
		client, _, err := initClient()
		assert.NoError(t, err)
		q := sfdc.SOQL().Select("Id").From("Case").Where("Subject", "contains", " ").Limit(1)
		sc := sfdc.NewSOQLClient[types.OpenCase](client)
		results, err := sc.Query(q)
		assert.NoError(t, err)
		assert.Equal(t, 1, results.TotalSize)
	})
	t.Run("where starts with query", func(t *testing.T) {
		client, _, err := initClient()
		assert.NoError(t, err)
		q := sfdc.SOQL().Select("Id").From("Case").Where("Subject", "startswith", "A").Limit(1)
		sc := sfdc.NewSOQLClient[types.OpenCase](client)
		results, err := sc.Query(q)
		assert.NoError(t, err)
		assert.Equal(t, 1, results.TotalSize)
	})
	t.Run("where ends with query", func(t *testing.T) {
		client, _, err := initClient()
		assert.NoError(t, err)
		q := sfdc.SOQL().Select("Id").From("Case").Where("Subject", "endswith", "e").Limit(1)
		sc := sfdc.NewSOQLClient[types.OpenCase](client)
		results, err := sc.Query(q)
		assert.NoError(t, err)
		assert.Equal(t, 1, results.TotalSize)
	})
}
