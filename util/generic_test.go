package util_test

import (
	"testing"
	"time"

	"github.com/stellaraf/go-sfdc/util"
	"github.com/stretchr/testify/assert"
)

func Test_EscapeString(t *testing.T) {
	t.Run("replaces string with single quote", func(t *testing.T) {
		in := `John's Bike`
		expected := `John\'s Bike`
		result := util.EscapeString(in)
		assert.Equal(t, expected, result)
	})
	t.Run("replaces string with double quote", func(t *testing.T) {
		in := `I said "that".`
		expected := `I said \"that\".`
		result := util.EscapeString(in)
		assert.Equal(t, expected, result)
	})
	t.Run("replaces string with backslash", func(t *testing.T) {
		in := `This\That.`
		expected := `This\\That.`
		result := util.EscapeString(in)
		assert.Equal(t, expected, result)
	})
}

func Test_IsArray(t *testing.T) {
	t.Run("true when array", func(t *testing.T) {
		in := []string{}
		result := util.IsArray(in)
		assert.True(t, result)
	})
	t.Run("true when not array", func(t *testing.T) {
		in := ""
		result := util.IsArray(in)
		assert.False(t, result)
	})
}

func Test_IsString(t *testing.T) {
	t.Run("true when string", func(t *testing.T) {
		in := ""
		result := util.IsString(in)
		assert.True(t, result)
	})
	t.Run("true when not string", func(t *testing.T) {
		in := []byte{}
		result := util.IsString(in)
		assert.False(t, result)
	})
}

func Test_IsTime(t *testing.T) {
	t.Run("true when time", func(t *testing.T) {
		in := time.Now()
		result := util.IsTime(in)
		assert.True(t, result)
	})
	t.Run("true when not time", func(t *testing.T) {
		in := ""
		result := util.IsTime(in)
		assert.False(t, result)
	})
}
