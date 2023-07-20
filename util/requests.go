package util

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/stellaraf/go-sfdc/types"
)

func CheckForError(response *resty.Response) (err error) {
	var possibleError any
	body := response.Body()
	err = json.Unmarshal(body, &possibleError)
	if err != nil {
		return
	}
	var errorDetail any = "unknown"

	if IsString(possibleError) {
		errorDetail = possibleError.(string)
		err = fmt.Errorf("request failed with error '%s'", errorDetail)
		return
	}
	if IsArray(possibleError) {
		arr := possibleError.([]interface{})
		var errs []types.QueryErrorResponse
		for _, item := range arr {
			m := item.(map[string]interface{})
			msg := m["message"]
			errorCode := m["errorCode"]
			if msg != nil && errorCode != nil {
				e := types.QueryErrorResponse{
					Message:   m["message"].(string),
					ErrorCode: m["errorCode"].(string),
				}
				errs = append(errs, e)
			}
		}
		messages := []string{}
		for _, e := range errs {
			messages = append(messages, e.Message)
		}
		errorDetail = strings.Join(messages, ", ")
		err = fmt.Errorf("request failed with error '%s'", errorDetail)
		return
	}
	if !IsArray(possibleError) {
		data := possibleError.(map[string]any)

	loop:
		for key := range data {
			switch key {
			case "error_code":
				errorDetail = data[key]
				break loop
			case "error_description":
				errorDetail = data[key]
				break loop
			case "error":
				errorDetail = data[key]
				break loop
			}
		}
	}
	if errorDetail == "unknown" {
		return nil
	}
	err = fmt.Errorf("request failed with %d error '%s'", response.StatusCode(), errorDetail)
	return
}

func GetSFDCError(data any) error {
	e, ok := data.(types.SalesforceErrorResponse)
	if ok {
		return e.GetError()
	}
	ep, ok := data.(*types.SalesforceErrorResponse)
	if ok && ep != nil {
		return ep.GetError()
	}
	g, ok := data.(types.GenericResponse)
	if ok {
		return fmt.Errorf("Error: %s", g.Message)
	}
	gp, ok := data.(*types.GenericResponse)
	if ok && gp != nil {
		return fmt.Errorf("Error: %s", gp.Message)
	}
	a, ok := data.(types.AuthErrorResponse)
	if ok {
		return a.GetError()
	}
	ap, ok := data.(*types.AuthErrorResponse)
	if ok && ap != nil {
		return ap.GetError()
	}
	return fmt.Errorf("Unknown Error: %v", data)
}

func GetSFDCErrorString(data any) string {
	return GetSFDCError(data).Error()
}
