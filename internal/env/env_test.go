package env_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.stellar.af/go-sfdc/internal/env"
)

func Test_FindProjectRoot(t *testing.T) {
	t.Run("find project root", func(t *testing.T) {
		t.Parallel()
		root, err := env.FindProjectRoot()
		assert.NoError(t, err)
		assert.NotEqual(t, "", root)
	})
}
