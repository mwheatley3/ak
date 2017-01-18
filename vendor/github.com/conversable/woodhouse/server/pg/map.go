package pg

import (
	"github.com/jmoiron/sqlx"
)

// A MapProxy is a value that maps column names to pointer values
// The MapProxy is used to map database values to go values on
// another struct
type MapProxy map[string]interface{}

// A MapProxyer provides a MapProxy for database deserialization
// operations on a single row
type MapProxyer interface {
	MapProxy() MapProxy
}

// A SliceMapProxyer represents a MapProxy for a collection
// of rows
type SliceMapProxyer interface {
	Append() MapProxyer
}

// mapScan is a lot like sqlx.MapScan except that the map values of dest
// are used in row.Scan
func mapScan(r sqlx.ColScanner, dest map[string]interface{}) error {
	columns, err := r.Columns()

	if err != nil {
		return err
	}

	values := make([]interface{}, len(columns))
	for i := range values {
		values[i] = dest[columns[i]]
	}

	if err = r.Scan(values...); err != nil {
		return err
	}

	return nil
}
