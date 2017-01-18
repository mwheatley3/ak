package pg

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/jmoiron/sqlx/reflectx"
)

// Column represents a sql table column
type Column struct {
	schema      string
	table       string
	name        string
	aliasSchema string
	aliasTable  string
	aliasName   string
}

func (c *Column) clone() *Column {
	cl := *c
	return &cl
}

func (c *Column) String() string {
	var (
		tbl = c.tableString()
		al  = c.aliasString()
	)

	if al == "" {
		return tbl
	}

	return tbl + " AS " + al
}

func (c *Column) tableString() string {
	if c.table == "" {
		return fmt.Sprintf("\"%s\"", c.name)
	}

	tbl := `"` + c.table + `"`

	if c.schema != "" {
		tbl = `"` + c.schema + `".` + tbl
	}

	return tbl + `."` + c.name + `"`
}

func (c *Column) aliasString() string {
	if c.aliasName == "" {
		return ""
	}

	if c.aliasTable == "" {
		return fmt.Sprintf("\"%s\"", c.aliasName)
	}

	parts := []string{c.aliasTable, c.aliasName}

	if c.aliasSchema != "" {
		parts = append([]string{c.aliasSchema}, parts...)
	}

	return `"` + strings.Join(parts, ".") + `"`
}

// Cols is a sortable alias for a group of columns
type Cols []*Column

func (c Cols) Len() int {
	return len(c)
}

func (c Cols) Less(i, j int) bool {
	return c[i].name < c[j].name
}

func (c Cols) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c Cols) clone() Cols {
	c2 := make([]*Column, len(c))

	for i, Col := range c {
		c2[i] = Col.clone()
	}

	return c2
}

// Collapse returns a new set of columns such that
// "Collections"."created_at" AS "Collections.created_at"
// becomes
// "Collections.created_at"
func (c Cols) Collapse() Cols {
	c2 := c.clone()

	for _, col := range c2 {
		colAlias := col.aliasString()

		if colAlias != "" {
			// remove quotes
			col.name = colAlias[1 : len(colAlias)-1]
		}

		col.aliasName = ""
		col.table = ""
		col.aliasTable = ""
		col.schema = ""
		col.aliasSchema = ""
	}

	return c2
}

// WithSchema returns a group of columns with the schema property replaced
func (c Cols) WithSchema(schema string) Cols {
	c2 := c.clone()

	for _, col := range c2 {
		col.schema = schema
	}

	return c2
}

// WithAliasSchema returns a group of columns with the alias schema property replaced
func (c Cols) WithAliasSchema(schema string) Cols {
	c2 := c.clone()

	for _, col := range c2 {
		col.aliasSchema = schema
	}

	return c2
}

// WithTable returns a group of columns with the table property replaced
func (c Cols) WithTable(tbl string) Cols {
	c2 := c.clone()

	for _, Col := range c2 {
		Col.table = tbl
	}

	return c2
}

// WithAliasTable returns a group of columns with the alias table property replaced
func (c Cols) WithAliasTable(tbl string) Cols {
	c2 := c.clone()

	for _, Col := range c2 {
		Col.aliasTable = tbl
	}

	return c2
}

func (c Cols) String() string {
	strs := make([]string, len(c))

	for i, Col := range c {
		strs[i] = Col.String()
	}

	return strings.Join(strs, ", ")
}

// Columns reflects over the passed in value and generates a group of
// columns using the `db` struct tag
func Columns(val interface{}) Cols {
	if mp, ok := val.(MapProxyer); ok {
		val = mp.MapProxy()
	}

	t := reflect.TypeOf(val)

	if t.Kind() == reflect.Map {
		return mapCols(reflect.ValueOf(val))
	}

	return structCols(t)
}

func mapCols(v reflect.Value) Cols {
	cols := make(Cols, v.Len())
	i := 0

	for _, k := range v.MapKeys() {
		cols[i] = colFromString(k.String())
		i++
	}

	sort.Sort(cols)

	return cols
}

func structCols(typ reflect.Type) Cols {
	m := reflectx.NewMapper("db").TypeMap(typ)
	cols := Cols{}

	for c, fi := range m.Names {
		if fi.Field.Tag.Get("db") != "" {
			cols = append(cols, colFromString(c))
		}
	}

	sort.Sort(cols)

	return cols
}

func colFromString(str string) *Column {
	col := &Column{name: str}

	if i := strings.Index(str, "."); i >= 0 {
		parts := strings.SplitN(str, ".", 3)

		if len(parts) == 3 {
			col.schema = parts[0]
			col.aliasSchema = parts[0]
			parts = parts[1:]
		}

		col.table = parts[0]
		col.name = parts[1]
		col.aliasTable = parts[0]
		col.aliasName = parts[1]
	}

	return col
}
