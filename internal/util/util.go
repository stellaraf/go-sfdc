package util

import (
	"encoding/json"
	"sort"
	"strings"
)

type CustomFields map[string]any

func FormatPrivateKey(rawKey string) (key string) {
	lines := []string{}
	for _, line := range strings.Split(rawKey, "\n") {
		newLine := strings.TrimSpace(line)
		lines = append(lines, newLine)
	}
	key = strings.Join(lines, "\n")
	return
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

func MergeCustomFields(obj any, fields []map[string]any) (merged map[string]any, err error) {
	size := 0
	for _, m := range fields {
		size += len(m)
	}
	allFields := make(map[string]any, size)
	for _, m := range fields {
		for k, v := range m {
			allFields[k] = v
		}
	}
	merged, err = MergeStructToMap(obj, allFields)
	return
}

func SortMap[T any](m map[string]T) map[string]T {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	om := make(map[string]T, len(keys))
	for _, k := range keys {
		om[k] = m[k]
	}
	return om
}
