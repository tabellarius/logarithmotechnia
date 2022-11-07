package dataframe

import (
	"errors"
	"logarithmotechnia/vector"
	"reflect"
	"time"
)

const optionStructHeaderMap = "structHeaderMap"
const optionStructDataframeOptions = "structDataframeOptions"

type confStruct struct {
	headerMap map[string]string
	dfOptions []vector.Option
}

func FromStructs(stArr any, options ...Option) (*Dataframe, error) {
	conf := createStructConf(options...)
	stArrType := reflect.TypeOf(stArr)

	if stArrType.Kind() != reflect.Slice && stArrType.Kind() != reflect.Array &&
		stArrType.Elem().Kind() != reflect.Struct {
		return nil, errors.New("data is not slice or array of structs")
	}

	stArrVal := reflect.ValueOf(stArr)
	if stArrVal.Len() == 0 {
		return New([]vector.Vector{}), nil
	}

	data, types, order := getDataAndTypes(stArrVal, stArrVal.Len())
	if len(order) == 0 {
		return New([]vector.Vector{}, conf.dfOptions...), nil
	}

	df := New(dataToVectors(data, types, order))

	return df, nil
}

func createStructConf(options ...Option) confStruct {
	conf := confStruct{
		headerMap: map[string]string{},
		dfOptions: []vector.Option{},
	}

	for _, option := range options {
		switch option.name {
		case optionStructHeaderMap:
			conf.headerMap = option.val.(map[string]string)
		case optionStructDataframeOptions:
			conf.dfOptions = option.val.([]vector.Option)
		}
	}

	return conf
}

func getDataAndTypes(stArrVal reflect.Value, length int) (map[string][]any, map[string]string, []string) {
	data, types := map[string][]any{}, map[string]string{}

	stVal := stArrVal.Index(0)
	stType := stVal.Type()
	order := make([]string, stType.NumField())
	for i := 0; i < stType.NumField(); i++ {
		name := stType.Field(i).Name
		data[name] = make([]any, length)
		types[name] = getFieldType(stVal.Field(i))
		order[i] = name
	}

	for i := 0; i < stArrVal.Len(); i++ {
		stVal = stArrVal.Index(i)
		for j := 0; j < stType.NumField(); j++ {
			name := stVal.Type().Field(j).Name
			data[name][i] = stVal.Field(j).Interface()
		}
	}

	return data, types, order
}

func dataToVectors(data map[string][]any, types map[string]string, order []string) []Column {
	columns := make([]Column, len(order))

	for i, field := range order {
		switch types[field] {
		case "integer":
			columns[i] = Column{field, vector.Integer(anyArrToTyped[int](data[field]))}
		case "float":
			columns[i] = Column{field, vector.Float(anyArrToTyped[float64](data[field]))}
		case "complex":
			columns[i] = Column{field, vector.Complex(anyArrToTyped[complex128](data[field]))}
		case "string":
			columns[i] = Column{field, vector.String(anyArrToTyped[string](data[field]))}
		case "boolean":
			columns[i] = Column{field, vector.Boolean(anyArrToTyped[bool](data[field]))}
		case "time":
			columns[i] = Column{field, vector.Time(anyArrToTyped[time.Time](data[field]))}
		case "any":
			columns[i] = Column{field, vector.Any(data[field])}
		}
	}

	return columns
}

func getFieldType(fVal reflect.Value) string {
	t := "any"

	switch fVal.Kind() {
	case reflect.Int:
		t = "integer"
	case reflect.Float64:
		t = "float"
	case reflect.Complex128:
		t = "complex"
	case reflect.String:
		t = "string"
	case reflect.Bool:
		t = "boolean"
	case reflect.Struct:
		if _, ok := fVal.Interface().(time.Time); ok {
			t = "time"
		} else {
			t = "any"
		}
	}

	return t
}

func StructOptionHeaderMap(headerMap map[string]string) Option {
	return Option{
		name: optionStructHeaderMap,
		val:  headerMap,
	}
}

func StructOptionDataFrameOptions(options ...vector.Option) Option {
	return Option{
		name: optionStructDataframeOptions,
		val:  options,
	}
}
