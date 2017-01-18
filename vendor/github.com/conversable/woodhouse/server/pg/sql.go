package pg

import (
	"strconv"
	"strings"
)

func fmtRange(count, start int, sep string, fn func(int) string) string {
	p := make([]string, count)

	if start < 0 {
		start = 1
	}

	for i := 0; i < count; i++ {
		p[i] = fn(i + start)
	}

	return strings.Join(p, sep)
}

// EscapeLike escapes the input string so that it is suitable to be used
// inside of a LIKE clause
func EscapeLike(str string) string {
	return strings.Replace(strings.Replace(str, "_", `\_`, -1), "%", `\%`, -1)
}

// SQLIn returns an IN clause with count postgres paramaters beginning with $1
func SQLIn(count int) string {
	return SQLInStart(count, 1)
}

// SQLInStart returns an IN clause with count postgres paramaters beginning with $start
func SQLInStart(count, start int) string {
	return "(" + fmtRange(count, start, ", ", SQLParam) + ")"
}

// SQLValuesFn returns a VALUES clause with count entries, beginning with index start
// and with each entry generated by fn
func SQLValuesFn(count int, start int, fn func(int) []string) string {
	return fmtRange(count, start, ", ", func(i int) string {
		return "(" + strings.Join(fn(i), ", ") + ")"
	})
}

// SQLValues returns a VALUES clause with count postgres paramater entries, beginning with $start
func SQLValues(count int, start int) string {
	return SQLValuesFn(count, start, func(i int) []string {
		return []string{SQLParam(i)}
	})
}

// SQLParam returns a postgres input parameter representation
func SQLParam(i int) string {
	return "$" + strconv.Itoa(i)
}

// ParamsSQLValues is a string representing the positional parameters
// used in a VALUES($1, $2, ...) type of expression
func ParamsSQLValues(p *Params) string {
	return ParamsSQLValuesSlice(p, len(p.values), 1)
}

// ParamsSQLValuesSlice is a string representing the positional parameters
// used in a VALUES($1, $2, ...) type of expression
func ParamsSQLValuesSlice(p *Params, count int, start int) string {
	// Invalid parameters
	if start < 1 || count > len(p.values) {
		return ""
	}
	return SQLValuesFn(count, start, func(idx int) []string {
		// SQL Param index is 1 indexed so we convert it to 0 index
		param := p.values[idx-1]

		// make params type specific
		// e.g. $1::int, $2::bigint
		return []string{SQLParam(idx) + SQLParamsCast(param)}
	})
}

// SQLParamsCast determines the SQL type cast used for a given go Type
// e.g. int64 would be "::bigint"
func SQLParamsCast(v interface{}) string {
	switch v.(type) {
	default:
		// no specific SQL type (treated as a string)
		return ""
	case int64:
		return "::bigint"
	case int:
		return "::int"
	}
}