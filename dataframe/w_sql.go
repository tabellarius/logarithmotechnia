package dataframe

import (
	"database/sql"
	"errors"
	"fmt"
	"logarithmotechnia/vector"
	"math"
	"reflect"
	"time"
)

const optionSQLDataframeOptions = "sqlDataframeOptions"
const optionSQLDataframeTransformers = "sqlDataframeTransformers"
const SQLTypeDateTime = "DATETIME"
const SQLTypeDate = "DATE"

type transformerFunc = func(vector.Vector) vector.Vector

type confSQL struct {
	dfOptions    []vector.Option
	transformers map[string]transformerFunc
}

func combineSQLConfig(options ...ConfOption) confSQL {
	conf := confSQL{
		dfOptions:    []vector.Option{},
		transformers: DefaultTransformers(),
	}

	for _, option := range options {
		switch option.Key() {
		case optionSQLDataframeOptions:
			conf.dfOptions = option.Value().([]vector.Option)
		}
	}

	return conf
}

func FromSQL(tx *sql.Tx, query string, args []interface{}, options ...ConfOption) (*Dataframe, error) {
	conf := combineSQLConfig(options...)

	stmt, err := tx.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}

	return ReadSQLRows(rows, conf)
}

func ReadSQLRows(rows *sql.Rows, conf confSQL) (*Dataframe, error) {
	columnNames, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}

	columns := make([]any, len(columnNames))
	for i := 0; i < len(columnNames); i++ {
		columns[i] = &SQLColumn{}
	}

	for rows.Next() {
		err = rows.Scan(columns...)
		if err != nil {
			return nil, err
		}
	}

	vectors := make([]vector.Vector, len(columns))
	for i, column := range columns {
		col := column.(*SQLColumn)
		switch col.kind {
		case SQLBoolean:
			vec := vector.BooleanWithNA(col.data.booleans, col.data.na)
			if transformer, ok := conf.transformers["boolean"]; ok {
				vec = transformer(vec)
			}
			vectors[i] = vec
		case SQLTime:
			vec := vector.TimeWithNA(col.data.times, col.data.na)
			if transformer, ok := conf.transformers["time"]; ok {
				vec = transformer(vec)
			}
			vectors[i] = vec
		case SQLNone:
			fallthrough
		case SQLString:
			switch columnTypes[i].DatabaseTypeName() {
			case SQLTypeDateTime:
				vec := vector.StringWithNA(col.data.strings, col.data.na)
				if transformer, ok := conf.transformers[SQLTypeDateTime]; ok {
					vec = transformer(vec)
				}
				vectors[i] = vec
			case SQLTypeDate:
				vec := vector.StringWithNA(col.data.strings, col.data.na)
				if transformer, ok := conf.transformers[SQLTypeDate]; ok {
					vec = transformer(vec)
				}
				vectors[i] = vec
			default:
				vec := vector.StringWithNA(col.data.strings, col.data.na)
				if transformer, ok := conf.transformers["string"]; ok {
					vec = transformer(vec)
				}
				vectors[i] = vec
			}
		case SQLInteger:
			vectors[i] = vector.IntegerWithNA(col.data.integers, col.data.na)
		case SQLFloat:
			vectors[i] = vector.FloatWithNA(col.data.floats, col.data.na)
		}
	}

	options := append(conf.dfOptions, OptionColumnNames(columnNames))
	df := New(vectors, options...)

	return df, nil
}

func DefaultTransformers() map[string]transformerFunc {
	return map[string]func(vector.Vector) vector.Vector{
		"DATETIME": func(vec vector.Vector) vector.Vector {
			vec.SetOption(vector.OptionTimeFormat("2006-01-02 15:04:05"))
			return vec.AsTime()
		},
		"DATE": func(vec vector.Vector) vector.Vector {
			vec.SetOption(vector.OptionTimeFormat("2006-01-02"))
			return vec.AsTime()
		},
	}
}

type SQLColumnType int

const (
	SQLNone SQLColumnType = iota
	SQLBoolean
	SQLFloat
	SQLInteger
	SQLString
	SQLTime
)

type SQLColumn struct {
	kind    SQLColumnType
	kindSet bool
	nulls   int
	data    struct {
		booleans []bool
		floats   []float64
		integers []int
		strings  []string
		times    []time.Time
		na       []bool
	}
}

func (c *SQLColumn) Boolean(v bool) {
	if !c.kindSet {
		c.data.booleans = make([]bool, 0, 100)
		c.data.na = make([]bool, 0, 100)
		c.kindSet = true
		c.kind = SQLBoolean

		if c.nulls > 0 {
			for i := 0; i < c.nulls; i++ {
				c.data.booleans = append(c.data.booleans, false)
				c.data.na = append(c.data.na, true)
			}
		}
	}
	c.data.booleans = append(c.data.booleans, v)
	c.data.na = append(c.data.na, false)
}

func (c *SQLColumn) Float(v float64) {
	if !c.kindSet {
		c.data.floats = make([]float64, 0, 100)
		c.data.na = make([]bool, 0, 100)
		c.kindSet = true
		c.kind = SQLFloat

		if c.nulls > 0 {
			for i := 0; i < c.nulls; i++ {
				c.data.floats = append(c.data.floats, math.NaN())
				c.data.na = append(c.data.na, true)
			}
		}
	}
	c.data.floats = append(c.data.floats, v)
	c.data.na = append(c.data.na, false)
}

func (c *SQLColumn) Integer(v int) {
	if !c.kindSet {
		c.data.integers = make([]int, 0, 100)
		c.data.na = make([]bool, 0, 100)
		c.kindSet = true
		c.kind = SQLInteger

		if c.nulls > 0 {
			for i := 0; i < c.nulls; i++ {
				c.data.integers = append(c.data.integers, 0)
				c.data.na = append(c.data.na, true)
			}
		}
	}
	c.data.integers = append(c.data.integers, v)
	c.data.na = append(c.data.na, false)
}

func (c *SQLColumn) String(v string) {
	if !c.kindSet {
		c.data.strings = make([]string, 0, 100)
		c.data.na = make([]bool, 0, 100)
		c.kindSet = true
		c.kind = SQLString

		if c.nulls > 0 {
			for i := 0; i < c.nulls; i++ {
				c.data.strings = append(c.data.strings, "")
				c.data.na = append(c.data.na, true)
			}
		}
	}
	c.data.strings = append(c.data.strings, v)
	c.data.na = append(c.data.na, false)
}

func (c *SQLColumn) Time(v time.Time) {
	if !c.kindSet {
		c.data.times = make([]time.Time, 0, 100)
		c.data.na = make([]bool, 0, 100)
		c.kindSet = true
		c.kind = SQLTime

		if c.nulls > 0 {
			for i := 0; i < c.nulls; i++ {
				c.data.times = append(c.data.times, time.Time{})
				c.data.na = append(c.data.na, true)
			}
		}
	}
	c.data.times = append(c.data.times, v)
	c.data.na = append(c.data.na, false)
}

func (c *SQLColumn) Null() error {
	if !c.kindSet {
		c.nulls++
		return nil
	}

	switch c.kind {
	case SQLBoolean:
		c.data.booleans = append(c.data.booleans, false)
		c.data.na = append(c.data.na, true)
	case SQLFloat:
		c.data.floats = append(c.data.floats, math.NaN())
		c.data.na = append(c.data.na, true)
	case SQLInteger:
		c.data.floats = append(c.data.floats, 0)
		c.data.na = append(c.data.na, true)
	case SQLString:
		c.data.strings = append(c.data.strings, "")
		c.data.na = append(c.data.na, true)
	case SQLTime:
		c.data.times = append(c.data.times, time.Time{})
		c.data.na = append(c.data.na, true)
	default:
		return errors.New(fmt.Sprintf("non-nullable type %v", c.kind))
	}

	return nil
}

func (c *SQLColumn) Scan(val interface{}) error {
	switch v := val.(type) {
	case bool:
		c.Boolean(v)
	case float64:
		c.Float(v)
	case int64:
		c.Integer(int(v))
	case string:
		c.String(v)
	case []byte:
		c.String(string(v))
	case time.Time:
		c.Time(v)
	case nil:
		err := c.Null()
		if err != nil {
			return err
		}
	default:
		return errors.New(fmt.Sprintf("unsupported scan type: %s for %v", reflect.ValueOf(v).Kind(), v))
	}

	return nil
}

func SQLOptionDataframeOptions(options ...vector.Option) ConfOption {
	return ConfOption{optionSQLDataframeOptions, options}
}

func SQLOptionTransformers(transformers map[string]transformerFunc) ConfOption {
	return ConfOption{optionSQLDataframeTransformers, transformers}
}
