package sfdc_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.stellar.af/go-sfdc"
)

func Test_PicklistValues(t *testing.T) {
	t.Run("get", func(t *testing.T) {
		t.Parallel()
		res, err := Client.PicklistValues(Env.TestData.PicklistObject, sfdc.DEFAULT_RECORD_TYPE_ID, Env.TestData.PicklistField)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(res.Values), 1)
	})
}

func Test_StandardValueSet(t *testing.T) {
	t.Run("get", func(t *testing.T) {
		t.Parallel()
		res, err := Client.StandardValueSet("Industry")
		require.NoError(t, err)
		assert.Len(t, res.Records, 1)
	})
	t.Run("items method", func(t *testing.T) {
		t.Parallel()
		res, err := Client.StandardValueSet("Industry")
		require.NoError(t, err)
		items := res.Items()[0]
		assert.False(t, len(items) == 0, "empty map")
		tech, ok := items["Technology"]
		assert.True(t, ok, "map missing value")
		assert.Equal(t, "Technology", tech)
	})
}
