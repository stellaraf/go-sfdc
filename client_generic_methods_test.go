package sfdc_test

import (
	"fmt"
	"testing"

	"github.com/stellaraf/go-sfdc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_NewObjectResponse(t *testing.T) {
	t.Run("base", func(t *testing.T) {
		t.Parallel()
		data := []byte(`{"key":"value"}`)
		obj, err := sfdc.NewObjectResponse(data)
		require.NoError(t, err)
		assert.NotNil(t, obj)
	})
	t.Run("invalid data", func(t *testing.T) {
		t.Parallel()
		data := []byte(`this%is*not(json$)`)
		obj, err := sfdc.NewObjectResponse(data)
		require.Error(t, err)
		assert.Nil(t, obj)
	})
}

func TestObjectResponse_GetString(t *testing.T) {
	t.Parallel()
	data := []byte(`{"key":"value"}`)
	obj, err := sfdc.NewObjectResponse(data)
	require.NoError(t, err)
	require.NotNil(t, obj)
	val := obj.GetString("key")
	assert.Equal(t, "value", val, "value mismatch")
}

func TestObjectResponse_GetInt(t *testing.T) {
	t.Parallel()
	t.Run("base", func(t *testing.T) {
		t.Parallel()
		data := []byte(`{"key":1234}`)
		obj, err := sfdc.NewObjectResponse(data)
		require.NoError(t, err)
		require.NotNil(t, obj)
		val := obj.GetInt("key")
		assert.Equal(t, 1234, val, "value mismatch")
	})
	t.Run("invalid data", func(t *testing.T) {
		t.Parallel()
		data := []byte(`{"key":"not-an-int"}`)
		obj, err := sfdc.NewObjectResponse(data)
		require.NoError(t, err)
		require.NotNil(t, obj)
		val := obj.GetInt("key")
		assert.Equal(t, 0, val, "value mismatch")
	})
}

func TestObjectResponse_GetFloat32(t *testing.T) {
	t.Parallel()
	t.Run("base", func(t *testing.T) {
		t.Parallel()
		data := []byte(`{"key":1.2}`)
		obj, err := sfdc.NewObjectResponse(data)
		require.NoError(t, err)
		require.NotNil(t, obj)
		val := obj.GetFloat32("key")
		assert.Equal(t, float32(1.2), val, "value mismatch")
	})
	t.Run("invalid data", func(t *testing.T) {
		t.Parallel()
		data := []byte(`{"key":"not-a-float"}`)
		obj, err := sfdc.NewObjectResponse(data)
		require.NoError(t, err)
		require.NotNil(t, obj)
		val := obj.GetFloat32("key")
		assert.Equal(t, float32(0), val, "value mismatch")
	})
}

func TestObjectResponse_GetFloat64(t *testing.T) {
	t.Parallel()
	t.Run("base", func(t *testing.T) {
		t.Parallel()
		data := []byte(`{"key":1.2}`)
		obj, err := sfdc.NewObjectResponse(data)
		require.NoError(t, err)
		require.NotNil(t, obj)
		val := obj.GetFloat64("key")
		assert.Equal(t, float64(1.2), val, "value mismatch")
	})
	t.Run("invalid data", func(t *testing.T) {
		t.Parallel()
		data := []byte(`{"key":"not-a-float"}`)
		obj, err := sfdc.NewObjectResponse(data)
		require.NoError(t, err)
		require.NotNil(t, obj)
		val := obj.GetFloat64("key")
		assert.Equal(t, float64(0), val, "value mismatch")
	})
}

func TestObjectResponse_GetBool(t *testing.T) {
	t.Parallel()
	t.Run("base", func(t *testing.T) {
		t.Parallel()
		data := []byte(`{"key":true}`)
		obj, err := sfdc.NewObjectResponse(data)
		require.NoError(t, err)
		require.NotNil(t, obj)
		val := obj.GetBool("key")
		assert.True(t, val, "value mismatch")
	})
	t.Run("invalid data", func(t *testing.T) {
		t.Parallel()
		data := []byte(`{"key":1234}`)
		obj, err := sfdc.NewObjectResponse(data)
		require.NoError(t, err)
		require.NotNil(t, obj)
		val := obj.GetBool("key")
		assert.False(t, val, "value mismatch")
	})
}

func TestObjectResponse_GetSlice(t *testing.T) {
	t.Parallel()
	t.Run("base", func(t *testing.T) {
		t.Parallel()
		data := []byte(`{"key":["one","two","three"]}`)
		obj, err := sfdc.NewObjectResponse(data)
		require.NoError(t, err)
		require.NotNil(t, obj)
		val := obj.GetSlice("key")
		sval := make([]string, 0, len(val))
		for _, el := range val {
			sval = append(sval, fmt.Sprint(el))
		}
		assert.ElementsMatch(t, []string{"one", "two", "three"}, sval)
	})
	t.Run("invalid data", func(t *testing.T) {
		t.Parallel()
		data := []byte(`{"key":"not-a-slice"}`)
		obj, err := sfdc.NewObjectResponse(data)
		require.NoError(t, err)
		require.NotNil(t, obj)
		val := obj.GetSlice("key")
		assert.Equal(t, []any{}, val)
	})
}

func TestObjectResponse_GetMap(t *testing.T) {
	t.Parallel()
	t.Run("base", func(t *testing.T) {
		t.Parallel()
		data := []byte(`{"key":{"one":"two"}}`)
		obj, err := sfdc.NewObjectResponse(data)
		require.NoError(t, err)
		require.NotNil(t, obj)
		val := obj.GetMap("key")
		exp := map[string]any{"one": "two"}
		assert.Equal(t, exp, val, "value mismatch")
		assert.Equal(t, exp["one"], val["one"], "nested value mismatch")
	})
	t.Run("invalid data", func(t *testing.T) {
		t.Parallel()
		data := []byte(`{"key":"not-a-map"}`)
		obj, err := sfdc.NewObjectResponse(data)
		require.NoError(t, err)
		require.NotNil(t, obj)
		val := obj.GetMap("key")
		assert.Equal(t, map[string]any{}, val)
	})
}
