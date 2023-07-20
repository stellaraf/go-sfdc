package types

import (
	"errors"
	"fmt"
	"strings"
)

type SalesforceError struct {
	Fields    []string `json:"fields"`
	ErrorCode string   `json:"errorCode"`
	Message   string   `json:"message"`
}

type SalesforceErrorResponse []*SalesforceError

func (er *SalesforceErrorResponse) GetErrorString() string {
	parts := []string{}
	for _, e := range *er {
		f := strings.Join(e.Fields, ", ")
		fs := "[%s] %s"
		if f != "" {
			fs += fmt.Sprintf(" (%s)", f)
		}
		s := fmt.Sprintf(fs, e.ErrorCode, e.Message)
		parts = append(parts, s)
	}
	result := fmt.Sprintf("Error: %s", strings.Join(parts, "; "))
	return result
}

func (er *SalesforceErrorResponse) GetError() error {
	return errors.New(er.GetErrorString())
}

type AuthErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func (a *AuthErrorResponse) GetErrorString() string {
	return fmt.Sprintf("Error: [%s] %s", a.Error, a.ErrorDescription)
}
func (a *AuthErrorResponse) GetError() error {
	return errors.New(a.GetErrorString())
}

type QueryErrorResponse struct {
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
}
