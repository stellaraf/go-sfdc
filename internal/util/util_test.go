package util_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.stellar.af/go-sfdc/internal/util"
)

const private_key_no_spaces string = `-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----`

const private_key_with_spaces string = `-----BEGIN RSA PRIVATE KEY-----
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 
-----END RSA PRIVATE KEY-----`

func Test_formatPrivateKey(t *testing.T) {
	t.Run("format key without spaces", func(t *testing.T) {
		t.Parallel()
		result := util.FormatPrivateKey(private_key_no_spaces)
		assert.Equal(t, private_key_no_spaces, result)
	})
	t.Run("format key with spaces", func(t *testing.T) {
		t.Parallel()
		result := util.FormatPrivateKey(private_key_with_spaces)
		assert.Equal(t, private_key_no_spaces, result)
	})
}

func Test_EscapeString(t *testing.T) {
	t.Run("replaces string with single quote", func(t *testing.T) {
		t.Parallel()
		in := `John's Bike`
		expected := `John\'s Bike`
		result := util.EscapeString(in)
		assert.Equal(t, expected, result)
	})
	t.Run("replaces string with double quote", func(t *testing.T) {
		t.Parallel()
		in := `I said "that".`
		expected := `I said \"that\".`
		result := util.EscapeString(in)
		assert.Equal(t, expected, result)
	})
	t.Run("replaces string with backslash", func(t *testing.T) {
		t.Parallel()
		in := `This\That.`
		expected := `This\\That.`
		result := util.EscapeString(in)
		assert.Equal(t, expected, result)
	})
}

func Test_MergeStructToMap(t *testing.T) {
	t.Run("merge struct with map", func(t *testing.T) {
		t.Parallel()
		keys := []string{"a", "b", "c", "d"}
		type S struct {
			A string `json:"a"`
			B string `json:"b"`
		}
		s := &S{
			A: "a",
			B: "b",
		}
		m := map[string]any{"c": "c", "d": "d"}
		result, err := util.MergeStructToMap(s, m)
		assert.NoError(t, err)
		assert.IsType(t, map[string]any{}, result)
		for _, k := range keys {
			v := result[k]
			assert.Equal(t, k, v)
		}
	})
}

func Test_SortMap(t *testing.T) {
	t.Parallel()
	m1 := map[string]string{"m": "m", "a": "a", "z": "z"}
	m2 := util.SortMap(m1)
	e := map[string]string{"a": "a", "m": "m", "z": "z"}
	assert.Equal(t, e, m2)
}
