package sfdc_test

import (
	"fmt"
	"testing"

	"github.com/stellaraf/go-sfdc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_SOQL(t *testing.T) {

	t.Run("where equals", func(t *testing.T) {
		t.Parallel()
		expected := "SELECT Id FROM Case WHERE IsClosed = false LIMIT 10"
		s, err := sfdc.SOQL().Select("Id").From("Case").Where("IsClosed", sfdc.EQUALS, false).Limit(10).String()
		require.NoError(t, err)
		assert.Equal(t, expected, s)
	})
	t.Run("where multiple", func(t *testing.T) {
		t.Parallel()
		expected := fmt.Sprintf("SELECT Id FROM Case WHERE IsClosed = false AND AccountId = '%s'", Env.TestData.AccountID)
		s, err := sfdc.SOQL().Select("Id").From("Case").Where("IsClosed", sfdc.EQUALS, false).Where("AccountId", "=", Env.TestData.AccountID).String()
		require.NoError(t, err)
		assert.Equal(t, expected, s)
	})
	t.Run("where contains", func(t *testing.T) {
		t.Parallel()
		expected := `SELECT Id FROM Case WHERE Description LIKE '%a%'`
		s, err := sfdc.SOQL().Select("Id").From("Case").Where("Description", sfdc.CONTAINS, "a").String()
		require.NoError(t, err)
		assert.Equal(t, expected, s)
	})
	t.Run("where contains query", func(t *testing.T) {
		t.Parallel()
		q := sfdc.SOQL().Select("Id").From("Case").Where("Subject", sfdc.CONTAINS, " ").Limit(1)
		sc := sfdc.NewSOQL[sfdc.OpenCase](Client)
		results, err := sc.Query(q)
		require.NoError(t, err)
		assert.Equal(t, 1, results.TotalSize)
	})
	t.Run("where starts with query", func(t *testing.T) {
		t.Parallel()
		q := sfdc.SOQL().Select("Id").From("Case").Where("Subject", sfdc.STARTS_WITH, "A").Limit(1)
		sc := sfdc.NewSOQL[sfdc.OpenCase](Client)
		results, err := sc.Query(q)
		require.NoError(t, err)
		assert.Equal(t, 1, results.TotalSize)
	})
	t.Run("where ends with query", func(t *testing.T) {
		t.Parallel()
		q := sfdc.SOQL().Select("Id").From("Case").Where("Subject", sfdc.ENDS_WITH, "e").Limit(1)
		sc := sfdc.NewSOQL[sfdc.OpenCase](Client)
		results, err := sc.Query(q)
		require.NoError(t, err)
		assert.Equal(t, 1, results.TotalSize)
	})
	t.Run("where in", func(t *testing.T) {
		t.Parallel()
		q := sfdc.SOQL().Select("Id").From("Case").Where("Subject", sfdc.IN, []string{"One", "Two"})
		expected := `SELECT Id FROM Case WHERE Subject IN ('One','Two')`
		result, err := q.String()
		require.NoError(t, err)
		assert.Equal(t, expected, result)
	})
	t.Run("where not in", func(t *testing.T) {
		t.Parallel()
		q := sfdc.SOQL().Select("Id").From("Case").Where("Subject", sfdc.NOT_IN, []string{"One", "Two"})
		expected := `SELECT Id FROM Case WHERE Subject NOT IN ('One','Two')`
		result, err := q.String()
		require.NoError(t, err)
		assert.Equal(t, expected, result)
	})
	t.Run("where not equals", func(t *testing.T) {
		t.Parallel()
		q := sfdc.SOQL().Select("Id").From("Case").Where("Subject", sfdc.NOT_EQUALS, "One")
		expected := `SELECT Id FROM Case WHERE Subject != 'One'`
		result, err := q.String()
		require.NoError(t, err)
		assert.Equal(t, expected, result)
	})
	t.Run("where greater than", func(t *testing.T) {
		t.Parallel()
		q := sfdc.SOQL().Select("Id").From("Case").Where("Subject", sfdc.GREATER_THAN, 1)
		expected := `SELECT Id FROM Case WHERE Subject > 1`
		result, err := q.String()
		require.NoError(t, err)
		assert.Equal(t, expected, result)
	})
	t.Run("where less than", func(t *testing.T) {
		t.Parallel()
		q := sfdc.SOQL().Select("Id").From("Case").Where("Subject", sfdc.LESS_THAN, 1)
		expected := `SELECT Id FROM Case WHERE Subject < 1`
		result, err := q.String()
		require.NoError(t, err)
		assert.Equal(t, expected, result)
	})
	t.Run("where gequal", func(t *testing.T) {
		t.Parallel()
		q := sfdc.SOQL().Select("Id").From("Case").Where("Subject", sfdc.GEQUAL, 1)
		expected := `SELECT Id FROM Case WHERE Subject >= 1`
		result, err := q.String()
		require.NoError(t, err)
		assert.Equal(t, expected, result)
	})
	t.Run("where lequal", func(t *testing.T) {
		t.Parallel()
		q := sfdc.SOQL().Select("Id").From("Case").Where("Subject", sfdc.LEQUAL, 1)
		expected := `SELECT Id FROM Case WHERE Subject <= 1`
		result, err := q.String()
		require.NoError(t, err)
		assert.Equal(t, expected, result)
	})
	t.Run("where includes", func(t *testing.T) {
		t.Parallel()
		q := sfdc.SOQL().Select("Id").From("Case").Where("Subject", sfdc.INCLUDES, "One")
		expected := `SELECT Id FROM Case WHERE Subject INCLUDES 'One'`
		result, err := q.String()
		require.NoError(t, err)
		assert.Equal(t, expected, result)
	})
	t.Run("where excludes", func(t *testing.T) {
		t.Parallel()
		q := sfdc.SOQL().Select("Id").From("Case").Where("Subject", sfdc.EXCLUDES, "One")
		expected := `SELECT Id FROM Case WHERE Subject EXCLUDES 'One'`
		result, err := q.String()
		require.NoError(t, err)
		assert.Equal(t, expected, result)
	})
	t.Run("sort", func(t *testing.T) {
		t.Parallel()
		q := sfdc.SOQL().Select("Id").From("Case").Where("Subject", sfdc.INCLUDES, "One").Sort("Id")
		expected := `SELECT Id FROM Case WHERE Subject INCLUDES 'One' ORDER BY Id ASC NULLS FIRST`
		result, err := q.String()
		require.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("sort ascending", func(t *testing.T) {
		t.Parallel()
		q := sfdc.SOQL().Select("Id").From("Case").Where("Subject", sfdc.INCLUDES, "One").Sort("Id").Ascending()
		expected := `SELECT Id FROM Case WHERE Subject INCLUDES 'One' ORDER BY Id ASC NULLS FIRST`
		result, err := q.String()
		require.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("sort descending", func(t *testing.T) {
		t.Parallel()
		q := sfdc.SOQL().Select("Id").From("Case").Where("Subject", sfdc.INCLUDES, "One").Sort("Id").Descending()
		expected := `SELECT Id FROM Case WHERE Subject INCLUDES 'One' ORDER BY Id DESC NULLS FIRST`
		result, err := q.String()
		require.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("sort ascending nulls last", func(t *testing.T) {
		t.Parallel()
		q := sfdc.SOQL().Select("Id").From("Case").Where("Subject", sfdc.INCLUDES, "One").Sort("Id").NullsLast()
		expected := `SELECT Id FROM Case WHERE Subject INCLUDES 'One' ORDER BY Id ASC NULLS LAST`
		result, err := q.String()
		require.NoError(t, err)
		assert.Equal(t, expected, result)
	})
}
