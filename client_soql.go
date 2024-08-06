package sfdc

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"go.stellar.af/go-sfdc/internal/util"
	"go.stellar.af/go-utils/is"
)

type where struct {
	Key      string
	Operator SOQLOperator
	Value    string
}

type soql struct {
	_select  []string
	_from    string
	_where   []where
	_groupBy string
	_count   string
	_limit   int
	_errors  []error
	_sort    string
	_order   int
	_nulls   int
}

type SOQLOperator string

const IN SOQLOperator = "IN"
const NOT_IN SOQLOperator = "NOT IN"
const CONTAINS SOQLOperator = "CONTAINS"
const STARTS_WITH SOQLOperator = "STARTS WITH"
const ENDS_WITH SOQLOperator = "ENDS WITH"
const LIKE SOQLOperator = "LIKE"
const EQUALS SOQLOperator = "="
const NOT_EQUALS SOQLOperator = "!="
const GREATER_THAN SOQLOperator = ">"
const LESS_THAN SOQLOperator = "<"
const GEQUAL SOQLOperator = ">="
const LEQUAL SOQLOperator = "<="
const INCLUDES SOQLOperator = "INCLUDES"
const EXCLUDES SOQLOperator = "EXCLUDES"

const (
	ascending = iota
	descending
)

const (
	nullsFirst = iota
	nullsLast
)

// Create a (limited, for now) SOQL queries using a Go API.
func SOQL() *soql {
	return &soql{
		_select:  []string{},
		_from:    "",
		_where:   []where{},
		_groupBy: "",
		_count:   "",
		_limit:   -1,
		_sort:    "",
		_order:   ascending,
		_nulls:   nullsFirst,
	}
}

// SOQL 'SELECT' function.
func (s *soql) Select(keys ...string) *soql {
	s._select = keys
	return s
}

// SOQL 'FROM' function.
func (s *soql) From(from string) *soql {
	s._from = from
	return s
}

func sliceType[T comparable](value any) ([]T, bool) {
	v, ok := value.([]T)
	return v, ok
}

// SOQL 'WHERE' function.
func (s *soql) Where(key string, operator SOQLOperator, value interface{}) *soql {
	var _where where
	if (operator == IN || operator == NOT_IN) && is.Slice(value) {
		values := []string{}
		s_str, is_str := sliceType[string](value)
		s_int, is_int := sliceType[int](value)
		s_int8, is_int8 := sliceType[int8](value)
		s_int16, is_int16 := sliceType[int16](value)
		s_int32, is_int32 := sliceType[int32](value)
		s_int64, is_int64 := sliceType[int64](value)
		s_uint, is_uint := sliceType[uint](value)
		s_uint8, is_uint8 := sliceType[uint8](value)
		s_uint16, is_uint16 := sliceType[uint16](value)
		s_uint32, is_uint32 := sliceType[uint32](value)
		s_uint64, is_uint64 := sliceType[uint64](value)
		s_float32, is_float32 := sliceType[float32](value)
		s_float64, is_float64 := sliceType[float64](value)
		switch {
		case is_str:
			for _, s := range s_str {
				values = append(values, fmt.Sprintf("'%s'", util.EscapeString(s)))
			}
		case is_int:
			for _, i := range s_int {
				values = append(values, fmt.Sprintf("'%d',", i))
			}
		case is_int8:
			for _, i := range s_int8 {
				values = append(values, fmt.Sprintf("'%d'", i))
			}
		case is_int16:
			for _, i := range s_int16 {
				values = append(values, fmt.Sprintf("'%d'", i))
			}
		case is_int32:
			for _, i := range s_int32 {
				values = append(values, fmt.Sprintf("'%d'", i))
			}
		case is_int64:
			for _, i := range s_int64 {
				values = append(values, fmt.Sprintf("'%d'", i))
			}
		case is_uint:
			for _, i := range s_uint {
				values = append(values, fmt.Sprintf("'%d'", i))
			}
		case is_uint8:
			for _, i := range s_uint8 {
				values = append(values, fmt.Sprintf("'%d'", i))
			}
		case is_uint16:
			for _, i := range s_uint16 {
				values = append(values, fmt.Sprintf("'%d'", i))
			}
		case is_uint32:
			for _, i := range s_uint32 {
				values = append(values, fmt.Sprintf("'%d'", i))
			}
		case is_uint64:
			for _, i := range s_uint64 {
				values = append(values, fmt.Sprintf("'%d'", i))
			}
		case is_float32:
			for _, i := range s_float32 {
				values = append(values, fmt.Sprintf("'%.9f'", i))
			}
		case is_float64:
			for _, i := range s_float64 {
				values = append(values, fmt.Sprintf("'%.9f'", i))
			}
		default:
			t := reflect.TypeOf(value).Kind().String()
			s._errors = append(s._errors, fmt.Errorf("failed to determine type for value '%v' (type %s)", value, t))
		}
		valuesStr := fmt.Sprintf("(%s)", strings.Join(values, ","))
		if operator == NOT_IN {
			_where = where{Key: key, Operator: NOT_IN, Value: valuesStr}
		} else {
			_where = where{Key: key, Operator: IN, Value: valuesStr}
		}
	} else {
		switch operator {
		case CONTAINS:
			val := fmt.Sprintf("'%%%v%%'", value)
			_where = where{Key: key, Operator: LIKE, Value: val}
		case STARTS_WITH:
			val := fmt.Sprintf("'%v%%'", value)
			_where = where{Key: key, Operator: LIKE, Value: val}
		case ENDS_WITH:
			val := fmt.Sprintf("'%%%v'", value)
			_where = where{Key: key, Operator: LIKE, Value: val}
		default:
			if is.Time(value) {
				val := value.(time.Time)
				_where = where{Key: key, Operator: operator, Value: val.Format(time.RFC3339)}
			} else if is.String(value) {
				val := util.EscapeString(value.(string))
				_where = where{Key: key, Operator: operator, Value: fmt.Sprintf("'%s'", val)}
			} else {
				_where = where{Key: key, Operator: operator, Value: fmt.Sprint(value)}
			}
		}
	}
	s._where = append(s._where, _where)
	return s
}

// SOQL 'GROUP_BY' function.
func (s *soql) GroupBy(groupBy string) *soql {
	s._groupBy = groupBy
	return s
}

// SOQL 'COUNT' function.
func (s *soql) Count(field string) *soql {
	s._count = field
	return s
}

// SOQL 'LIMIT' function.
func (s *soql) Limit(limit int) *soql {
	s._limit = limit
	return s
}

// SOQL 'SORT' function.
func (s *soql) Sort(field string) *soql {
	s._sort = field
	return s
}

// In conjunction with Sort(), order results A-Z.
func (s *soql) Ascending() *soql {
	s._order = ascending
	return s
}

// In conjunction with Sort(), order results Z-A.
func (s *soql) Descending() *soql {
	s._order = descending
	return s
}

// In conjunction with Sort(), place null results first.
func (s *soql) NullsFirst() *soql {
	s._nulls = nullsFirst
	return s
}

// In conjunction with Sort(), place null results last.
func (s *soql) NullsLast() *soql {
	s._nulls = nullsLast
	return s
}

// Convert the SOQL query object to an SOQL query string.
func (s *soql) String() (result string, err error) {
	if s._from == "" {
		err = fmt.Errorf("'FROM' is required for SOQL queries")
		return
	}
	if len(s._select) == 0 {
		err = fmt.Errorf("'SELECT' is required for SOQL queries")
		return
	}
	for _, e := range s._errors {
		err = e
		return
	}
	_select := strings.Join(s._select, ",")
	parts := []string{fmt.Sprintf("SELECT %s", _select), fmt.Sprintf("FROM %s", s._from)}

	if len(s._where) > 0 {
		var filters []string
		for _, w := range s._where {
			statement := strings.Join([]string{w.Key, string(w.Operator), w.Value}, " ")
			filters = append(filters, statement)
		}
		_where := fmt.Sprintf("WHERE %s", strings.Join(filters, " AND "))
		parts = append(parts, _where)
	}

	if s._groupBy != "" {
		_groupBy := fmt.Sprintf("GROUP BY %s", s._groupBy)
		parts = append(parts, _groupBy)
	}

	if s._count != "" {
		_select := parts[0]
		rest := parts[1:]
		_count := fmt.Sprintf("COUNT(%s)", s._count)
		parts = []string{_select, _count}
		parts = append(parts, rest...)
	}

	if s._sort != "" {
		_sort := fmt.Sprintf("ORDER BY %s", s._sort)
		parts = append(parts, _sort)
		if s._order == ascending {
			parts = append(parts, "ASC")
		} else if s._order == descending {
			parts = append(parts, "DESC")
		}
		if s._nulls == nullsFirst {
			parts = append(parts, "NULLS FIRST")
		} else if s._nulls == nullsLast {
			parts = append(parts, "NULLS LAST")
		}
	}

	if s._limit != -1 {
		_limit := fmt.Sprintf("LIMIT %d", s._limit)
		parts = append(parts, _limit)
	}

	result = strings.Join(parts, " ")
	return
}
