package sfdc

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type ObjectResponse struct {
	Data map[string]any
}

func (obj *ObjectResponse) GetString(key string) string {
	gval, ok := obj.Data[key]
	if !ok {
		return ""
	}
	val := fmt.Sprint(gval)
	if val == "<nil>" {
		return ""
	}
	return val
}

func (obj *ObjectResponse) GetInt(key string) int {
	gval, ok := obj.Data[key]
	if !ok {
		return 0
	}
	val, err := strconv.Atoi(fmt.Sprint(gval))
	if err != nil {
		return 0
	}
	return val
}

func (obj *ObjectResponse) GetFloat32(key string) float32 {
	gval, ok := obj.Data[key]
	if !ok {
		return float32(0.0)
	}
	val, err := strconv.ParseFloat(fmt.Sprint(gval), 32)
	if err != nil {
		return float32(0.0)
	}
	return float32(val)
}

func (obj *ObjectResponse) GetFloat64(key string) float64 {
	gval, ok := obj.Data[key]
	if !ok {
		return float64(0)
	}
	val, err := strconv.ParseFloat(fmt.Sprint(gval), 64)
	if err != nil {
		return float64(0)
	}
	return val
}

func (obj *ObjectResponse) GetBool(key string) bool {
	gval, ok := obj.Data[key]
	if !ok {
		return false
	}
	val, ok := gval.(bool)
	if !ok {
		return false
	}
	return val
}

func (obj *ObjectResponse) GetSlice(key string) []any {
	gval, ok := obj.Data[key]
	if !ok {
		return []any{}
	}
	val, ok := gval.([]any)
	if !ok {
		return []any{}
	}
	return val
}

func (obj *ObjectResponse) GetMap(key string) map[string]any {
	gval, ok := obj.Data[key]
	if !ok {
		return map[string]any{}
	}
	val, ok := gval.(map[string]any)
	if !ok {
		return map[string]any{}
	}
	return val
}

func NewObjectResponse(data []byte) (*ObjectResponse, error) {
	dataMap := make(map[string]any)
	err := json.Unmarshal(data, &dataMap)
	if err != nil {
		return nil, err
	}
	obj := &ObjectResponse{
		Data: dataMap,
	}
	return obj, nil
}
