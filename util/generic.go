package util

import (
	"reflect"
)

func IsArray(in any) bool {
	return reflect.TypeOf(in).Kind() == reflect.Slice
}

func IsString(in any) bool {
	return reflect.TypeOf(in).Kind() == reflect.String
}
