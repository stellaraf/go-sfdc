package sfdc

import (
	"fmt"
	"reflect"

	"github.com/go-resty/resty/v2"
	"github.com/perimeterx/marshmallow"
)

func (client *Client) handleObjectError(res *resty.Response) (err error) {
	if res.IsError() {
		e := res.Error().(*SalesforceErrorResponse)
		return getSFDCError(e)
	}
	return
}

func (client *Client) handleResponse(res *resty.Response, obj any) (err error) {
	if res.IsError() {
		e := res.Error().(*SalesforceErrorResponse)
		return getSFDCError(e)
	}
	p := reflect.ValueOf(obj)
	if p.Kind() != reflect.Pointer {
		err = fmt.Errorf("expected pointer type")
		return
	}
	s := p.Elem()
	body := res.Body()
	extra, err := marshmallow.Unmarshal(body, obj)
	f := s.FieldByName("CustomFields")
	e := reflect.ValueOf(extra)
	if f.CanSet() {
		f.Set(e)
	}
	return
}
