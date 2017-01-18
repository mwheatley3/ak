package pg

import (
	"github.com/satori/go.uuid"
)

// NewParams returns a new *Params object
func NewParams(values ...interface{}) *Params {
	if values == nil {
		values = []interface{}{}
	}

	return &Params{
		values: values,
	}
}

// Params represents a list of db query params
type Params struct {
	values []interface{}
}

// Values returns the underlying []interface{}
func (p *Params) Values() []interface{} {
	if p == nil {
		return []interface{}{}
	}

	return p.values
}

// Add appends values onto the underlying list
func (p *Params) Add(values ...interface{}) *Params {
	p.values = append(p.values, values...)
	return p
}

// AddInt is a typed (int) version of Add
func (p *Params) AddInt(values ...int) *Params {
	l := len(p.values)
	p.values = append(p.values, make([]interface{}, len(values))...)

	for i := 0; i < len(values); i++ {
		p.values[i+l] = values[i]
	}

	return p
}

// AddInt64 is a typed (int) version of Add
func (p *Params) AddInt64(values ...int64) *Params {
	l := len(p.values)
	p.values = append(p.values, make([]interface{}, len(values))...)

	for i := 0; i < len(values); i++ {
		p.values[i+l] = values[i]
	}

	return p
}

// AddString is a typed (string) version of Add
func (p *Params) AddString(values ...string) *Params {
	l := len(p.values)
	p.values = append(p.values, make([]interface{}, len(values))...)

	for i := 0; i < len(values); i++ {
		p.values[l+i] = values[i]
	}

	return p
}

// AddUUID is a typed (uuid.UUID) version of Add
func (p *Params) AddUUID(values ...uuid.UUID) *Params {
	l := len(p.values)
	p.values = append(p.values, make([]interface{}, len(values))...)

	for i := 0; i < len(values); i++ {
		p.values[l+i] = values[i]
	}

	return p
}
