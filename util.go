package sfdc

import (
	"fmt"
)

func getSFDCError(data any) error {
	e, ok := data.(SalesforceErrorResponse)
	if ok {
		return e.GetError()
	}
	ep, ok := data.(*SalesforceErrorResponse)
	if ok && ep != nil {
		return ep.GetError()
	}
	g, ok := data.(GenericResponse)
	if ok {
		return fmt.Errorf("error: %s", g.Message)
	}
	gp, ok := data.(*GenericResponse)
	if ok && gp != nil {
		return fmt.Errorf("error: %s", gp.Message)
	}
	a, ok := data.(AuthErrorResponse)
	if ok {
		return a.GetError()
	}
	ap, ok := data.(*AuthErrorResponse)
	if ok && ap != nil {
		return ap.GetError()
	}
	return fmt.Errorf("unknown error: %v", data)
}

func getSFDCErrorString(data any) string {
	return getSFDCError(data).Error()
}
