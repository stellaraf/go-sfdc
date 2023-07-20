package util_test

import (
	"fmt"
	"testing"

	"github.com/stellaraf/go-sfdc/types"
	"github.com/stellaraf/go-sfdc/util"
	"github.com/stretchr/testify/assert"
)

func Test_GetSFDCError(t *testing.T) {
	t.Run("server error", func(t *testing.T) {
		e := types.SalesforceErrorResponse{
			&types.SalesforceError{
				ErrorCode: "ERROR_CODE",
				Message:   "Error Message",
				Fields:    []string{},
			},
		}
		expected := "Error: [ERROR_CODE] Error Message"
		result := util.GetSFDCErrorString(e)
		assert.Equal(t, expected, result)
	})
	t.Run("generic error", func(t *testing.T) {
		e := types.GenericResponse{
			Message: "Message",
			Data:    map[string]any{},
		}
		expected := "Error: Message"
		result := util.GetSFDCErrorString(e)
		assert.Equal(t, expected, result)
	})
	t.Run("auth error", func(t *testing.T) {
		e := types.AuthErrorResponse{
			Error:            "AUTH_ERROR",
			ErrorDescription: "Authentication Error",
		}
		expected := "Error: [AUTH_ERROR] Authentication Error"
		result := util.GetSFDCErrorString(e)
		assert.Equal(t, expected, result)
	})
	t.Run("default", func(t *testing.T) {
		e := fmt.Errorf("some other error")
		expected := "Error: some other error"
		result := util.GetSFDCErrorString(e)
		assert.Equal(t, expected, result)
	})
}
