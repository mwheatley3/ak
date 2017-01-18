package pg

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// https://gist.github.com/adharris/4163702
var (
	// unquoted array values must not contain: (" , \ { } whitespace NULL)
	// and must be at least one char
	unquotedChar  = `[^",\\{}\s(NULL)]`
	unquotedValue = fmt.Sprintf("(%s)+", unquotedChar)

	// quoted array values are surrounded by double quotes, can be any
	// character except " or \, which must be backslash escaped:
	quotedChar  = `[^"\\]|\\"|\\\\`
	quotedValue = fmt.Sprintf("\"(%s)*\"", quotedChar)

	// an array value may be either quoted or unquoted:
	arrayValue = fmt.Sprintf("(?P<value>(%s|%s))", unquotedValue, quotedValue)

	// Array values are separated with a comma IF there is more than one value:
	arrayExp = regexp.MustCompile(fmt.Sprintf("((%s)(,)?)", arrayValue))

	valueIndex int
)

// StringSlice is a postgres serializable representation of an array of strings
type StringSlice []string

// Scan implements sql.Scanner for the String slice type
// Scanners take the database value (in this case as a byte slice)
// and sets the value of the type.  Here we cast to a string and
// do a regexp based parse
func (s *StringSlice) Scan(src interface{}) error {
	asBytes, ok := src.([]byte)
	if !ok {
		return error(errors.New("Scan source was not []bytes"))
	}

	asString := string(asBytes)
	parsed := parseArray(asString)
	(*s) = StringSlice(parsed)

	return nil
}

func parseArray(array string) []string {
	results := []string{}
	matches := arrayExp.FindAllStringSubmatch(array, -1)
	for _, match := range matches {
		s := match[valueIndex]
		// the string _might_ be wrapped in quotes, so trim them:
		s = strings.Trim(s, "\"")
		results = append(results, s)
	}
	return results
}

// IntSlice is a postgres serializable representation of an array of ints
type IntSlice []int

// Scan implements the sql interfaces
func (i *IntSlice) Scan(src interface{}) error {
	asBytes, ok := src.([]byte)
	if !ok {
		return error(errors.New("Scan source was not []bytes"))
	}

	asString := string(asBytes)
	parsed, err := parseIntArray(asString)
	if err != nil {
		return err
	}

	(*i) = IntSlice(parsed)

	return nil
}

func parseIntArray(array string) ([]int, error) {
	results := []int{}
	matches := arrayExp.FindAllStringSubmatch(array, -1)
	for _, match := range matches {
		s := match[valueIndex]
		// the string _might_ be wrapped in quotes, so trim them:
		s = strings.Trim(s, "\"")
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("Error converting %s to int: %s", s, err)
		}
		results = append(results, i)
	}
	return results, nil
}

// Value implements the db/driver.Value interface
func (i IntSlice) Value() (driver.Value, error) {
	strValues := []string{}
	for _, intValue := range i {
		s := strconv.Itoa(intValue)
		strValues = append(strValues, s)
	}
	v := fmt.Sprintf("{%s}", strings.Join(strValues, ","))
	return v, nil
}

// Find the index of the 'value' named expression
func init() {
	for i, subexp := range arrayExp.SubexpNames() {
		if subexp == "value" {
			valueIndex = i
			break
		}
	}
}
