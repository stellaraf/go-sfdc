package sfdc

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/go-resty/resty/v2"
	"github.com/perimeterx/marshmallow"
)

func (client *Client) handleObjectError(res *resty.Response) error {
	if res.IsError() {
		return getSFDCError(res.Error())
	}
	return nil
}

func (client *Client) handleResponse(res *resty.Response, obj any) error {
	if res.IsError() {
		return getSFDCError(res.Error())
	}
	p := reflect.ValueOf(obj)
	if p.Kind() != reflect.Pointer {
		return fmt.Errorf("expected pointer type")
	}
	s := p.Elem()
	body := res.Body()
	extra, err := marshmallow.Unmarshal(body, obj)
	if err != nil {
		err = errors.Join(err, fmt.Errorf("failed to parse response body '%s'", string(body)))
		return err
	}
	f := s.FieldByName("CustomFields")
	e := reflect.ValueOf(extra)
	if f.CanSet() {
		f.Set(e)
	}
	return nil
}
