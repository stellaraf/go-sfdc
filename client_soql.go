package sfdc

import (
	"fmt"
	"strings"
	"time"

	"github.com/stellaraf/go-sfdc/util"
)

type where struct {
	Key      string
	Operator string
	Value    string
}

type soql struct {
	_select  []string
	_from    string
	_where   []where
	_groupBy string
	_count   string
	_limit   int
}

// Create a (limited, for now) SOQL queries using a Go API.
func SOQL() *soql {
	return &soql{
		_select:  []string{},
		_from:    "",
		_where:   []where{},
		_groupBy: "",
		_count:   "",
		_limit:   -1,
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

// SOQL 'WHERE' function.
func (s *soql) Where(key string, operator string, value interface{}) *soql {
	var _where where
	if util.IsArray(value) {
		va := value.([]interface{})
		k := va[0].(string)
		op := va[1]
		val := va[2]
		var values []string
		if util.IsArray(val) {
			for _, v := range val.([]interface{}) {
				values = append(values, fmt.Sprintf("'%v'", v))
			}
		} else {
			values = append(values, fmt.Sprintf("'%v'", val))
		}
		valuesStr := fmt.Sprintf("(%s)", strings.Join(values, ", "))
		valuesStr = util.EscapeString(valuesStr)
		if op == "notin" {
			_where = where{Key: k, Operator: "NOT IN", Value: valuesStr}
		} else {
			_where = where{Key: k, Operator: "IN", Value: valuesStr}
		}
	} else {
		switch operator {
		case "contains":
			val := fmt.Sprintf("'%%%v%%'", value)
			_where = where{Key: key, Operator: "LIKE", Value: val}
		case "startswith":
			val := fmt.Sprintf("'%v%%'", value)
			_where = where{Key: key, Operator: "LIKE", Value: val}
		case "endswith":
			val := fmt.Sprintf("'%%%v'", value)
			_where = where{Key: key, Operator: "LIKE", Value: val}
		default:
			if util.IsTime(value) {
				val := value.(time.Time)
				_where = where{Key: key, Operator: operator, Value: val.Format(time.RFC3339)}
			} else if util.IsString(value) {
				val := util.EscapeString(value.(string))
				_where = where{Key: key, Operator: operator, Value: fmt.Sprintf("'%v'", val)}
			} else {
				_where = where{Key: key, Operator: operator, Value: fmt.Sprintf("%v", value)}
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
	_select := strings.Join(s._select, ",")
	parts := []string{fmt.Sprintf("SELECT %s", _select), fmt.Sprintf("FROM %s", s._from)}

	if len(s._where) > 0 {
		var filters []string
		for _, w := range s._where {
			statement := strings.Join([]string{w.Key, w.Operator, w.Value}, " ")
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

	if s._limit != -1 {
		_limit := fmt.Sprintf("LIMIT %d", s._limit)
		parts = append(parts, _limit)
	}
	result = strings.Join(parts, " ")
	return
}
