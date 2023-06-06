package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findProjectRoot(t *testing.T) {
	t.Run("find project root", func(t *testing.T) {
		root, err := findProjectRoot()
		assert.NoError(t, err)
		assert.NotEqual(t, "", root)
	})
}
