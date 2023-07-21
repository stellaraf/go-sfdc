package sfdc

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getSFDCError(t *testing.T) {
	t.Run("server error", func(t *testing.T) {
		t.Parallel()
		e := SalesforceErrorResponse{
			&SalesforceError{
				ErrorCode: "ERROR_CODE",
				Message:   "Error Message",
				Fields:    []string{},
			},
		}
		expected := "error: [ERROR_CODE] Error Message"
		result := getSFDCErrorString(e)
		assert.Equal(t, expected, result)
	})
	t.Run("generic error", func(t *testing.T) {
		t.Parallel()
		e := GenericResponse{
			Message: "Message",
			Data:    map[string]any{},
		}
		expected := "error: Message"
		result := getSFDCErrorString(e)
		assert.Equal(t, expected, result)
	})
	t.Run("auth error", func(t *testing.T) {
		t.Parallel()
		e := AuthErrorResponse{
			Error:            "AUTH_ERROR",
			ErrorDescription: "Authentication Error",
		}
		expected := "error: [AUTH_ERROR] Authentication Error"
		result := getSFDCErrorString(e)
		assert.Equal(t, expected, result)
	})
	t.Run("default", func(t *testing.T) {
		t.Parallel()
		e := fmt.Errorf("some other error")
		expected := "unknown error: some other error"
		result := getSFDCErrorString(e)
		assert.Equal(t, expected, result)
	})
}
