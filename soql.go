package sfdc

import (
	"fmt"
	"strings"
	"time"

	"github.com/stellaraf/go-sfdc/util"
)

type Where struct {
	Key      string
	Operator string
	Value    string
}

type soql struct {
	_select  []string
	_from    string
	_where   []Where
	_groupBy string
	_count   string
	_limit   int
}

func SOQL() *soql {
	return &soql{
		_select:  []string{},
		_from:    "",
		_where:   []Where{},
		_groupBy: "",
		_count:   "",
		_limit:   -1,
	}
}

func (s *soql) Select(keys ...string) *soql {
	s._select = keys
	return s
}

func (s *soql) From(from string) *soql {
	s._from = from
	return s
}

func (s *soql) Where(key string, operator string, value interface{}) *soql {
	var _where Where
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
		if op == "notin" {
			_where = Where{Key: k, Operator: "NOT IN", Value: valuesStr}
		} else {
			_where = Where{Key: k, Operator: "IN", Value: valuesStr}
		}
	} else {
		switch operator {
		case "contains":
			val := fmt.Sprintf("'%%%v%%'", value)
			_where = Where{Key: key, Operator: "LIKE", Value: val}
		case "startswith":
			val := fmt.Sprintf("'%v%%'", value)
			_where = Where{Key: key, Operator: "LIKE", Value: val}
		case "endswith":
			val := fmt.Sprintf("'%%%v'", value)
			_where = Where{Key: key, Operator: "LIKE", Value: val}
		default:
			if util.IsTime(value) {
				val := value.(time.Time)
				_where = Where{Key: key, Operator: operator, Value: val.Format(time.RFC3339)}
			} else if util.IsString(value) {
				_where = Where{Key: key, Operator: operator, Value: fmt.Sprintf("'%v'", value)}
			} else {
				_where = Where{Key: key, Operator: operator, Value: fmt.Sprintf("%v", value)}
			}
		}
	}
	s._where = append(s._where, _where)
	return s
}

func (s *soql) GroupBy(groupBy string) *soql {
	s._groupBy = groupBy
	return s
}

func (s *soql) Count(field string) *soql {
	s._count = field
	return s
}

func (s *soql) Limit(limit int) *soql {
	s._limit = limit
	return s
}

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
			statement := fmt.Sprintf("%s %s %s", w.Key, w.Operator, w.Value)
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
