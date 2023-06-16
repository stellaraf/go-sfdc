package util

import (
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
