package pg

import (
	"database/sql/driver"
	"encoding/json"
)

// JSON is a simple value that can be encoded/decoded
// as json from the db
type JSON json.RawMessage

// MarshalJSON satisfies the json.Marshaler interface
func (j *JSON) MarshalJSON() ([]byte, error) {
	return (*json.RawMessage)(j).MarshalJSON()
}

// UnmarshalJSON satisfies the json.Unmarshaler interface
func (j *JSON) UnmarshalJSON(b []byte) error {
	return (*json.RawMessage)(j).UnmarshalJSON(b)
}

// Value implements the driver.Valuer interface
func (j JSON) Value() (driver.Value, error) {
	return JSONValue(j)
}

// Scan implements the database.Scanner interface
func (j *JSON) Scan(v interface{}) error {
	return JSONScan(j, v)
}

// JSONValue is a helper that can be used to help satisfy the
// database/sql/driver.Value interface for json columns in pg
func JSONValue(val interface{}) (driver.Value, error) {
	b, err := json.Marshal(val)

	if err != nil {
		return nil, err
	}

	return string(b), nil
}

// JSONScan is a helper that can be used to help satisfy the
// database/sql.Scanner interface for json columns in pg
func JSONScan(dest interface{}, v interface{}) error {
	if v == nil {
		return nil
	}

	return json.Unmarshal([]byte(v.(string)), dest)
}

// JSONWrap returns wraps a value with json
// Scan and Value methods
func JSONWrap(v interface{}) interface{} {
	return &jsonWrap{v: v}
}

type jsonWrap struct {
	v interface{}
}

// Value implements the driver.Valuer interface
func (j jsonWrap) Value() (driver.Value, error) {
	return JSONValue(j.v)
}

// Scan implements the database.Scanner interface
func (j *jsonWrap) Scan(b interface{}) error {
	return JSONScan(j.v, b)
}
