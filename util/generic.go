package util

import (
	"encoding/json"
	"reflect"
	"strings"
	"time"
)

func IsArray(in any) bool {
	return reflect.TypeOf(in).Kind() == reflect.Slice
}

func IsString(in any) bool {
	return reflect.TypeOf(in).Kind() == reflect.String
}

func IsTime(in any) bool {
	_, isTime := in.(time.Time)
	return isTime
}

func EscapeString(in string) (out string) {
	replacer := strings.NewReplacer(`'`, `\'`, `"`, `\"`, `\`, `\\`)
	out = replacer.Replace(in)
	return
}

func MergeStructToMap(data any, extra map[string]any) (result map[string]any, err error) {
	b, err := json.Marshal(&data)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return
	}
	for k, v := range extra {
		result[k] = v
	}
	return
}
