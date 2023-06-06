package types

import "reflect"

func IsServerError(data any) bool {
	reflection := reflect.ValueOf(data)
	field := reflection.FieldByName("Error")
	return field.IsValid()
}

func IsQueryError(data any) bool {
	reflection := reflect.ValueOf(data)
	field := reflection.FieldByName("ErrorCode")
	return field.IsValid()
}
