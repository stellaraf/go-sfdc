package sfdc

import (
	"fmt"
	"testing"

	"github.com/stellaraf/go-sfdc/util"
	"github.com/stretchr/testify/assert"
)

func Test_SOQL(t *testing.T) {
	t.Run("open cases", func(t *testing.T) {
		expected := "SELECT Id FROM Case WHERE IsClosed = false LIMIT 10"
		s, err := SOQL().Select("Id").From("Case").Where("IsClosed", "=", false).Limit(10).String()
		assert.NoError(t, err)
		assert.Equal(t, expected, s)
	})
	t.Run("multiple where", func(t *testing.T) {
		env, err := util.LoadEnv()
		assert.NoError(t, err)
		expected := fmt.Sprintf("SELECT Id FROM Case WHERE IsClosed = false AND AccountId = '%s'", env.TestData.AccountID)
		s, err := SOQL().Select("Id").From("Case").Where("IsClosed", "=", false).Where("AccountId", "=", env.TestData.AccountID).String()
		assert.NoError(t, err)
		assert.Equal(t, expected, s)
	})
	t.Run("where contains", func(t *testing.T) {
		expected := `SELECT Id FROM Case WHERE Description LIKE \'%a%\'`
		s, err := SOQL().Select("Id").From("Case").Where("Description", "contains", "a").String()
		assert.NoError(t, err)
		assert.Equal(t, expected, s)
	})
}
