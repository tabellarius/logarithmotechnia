package dataframe

import (
	"errors"
	"golang.org/x/exp/slices"
	"logarithmotechnia/vector"
	"reflect"
	"strings"
	"time"
)

const optionStructHeaderMap = "structHeaderMap"
const optionStructDataframeOptions = "structDataframeOptions"
const optionStructSkipFields = "structSkipFields"

type confStruct struct {
	headerMap   map[string]string
	dfOptions   []vector.Option
	skipColumns []string
}

func FromStructs(stArr any, options ...ConfOption) (*Dataframe, error) {
	conf := createStructConf(options...)
	stArrType := reflect.TypeOf(stArr)

	if stArrType.Kind() != reflect.Slice && stArrType.Kind() != reflect.Array &&
		stArrType.Elem().Kind() != reflect.Struct {
		return nil, errors.New("data is not slice or array of structs")
	}

	stArrVal := reflect.ValueOf(stArr)
	if stArrVal.Len() == 0 {
		return New([]vector.Vector{}, conf.dfOptions...), nil
	}

	data, types, order := getDataAndTypes(stArrVal, stArrVal.Len(), conf)
	if len(order) == 0 {
		return New([]vector.Vector{}, conf.dfOptions...), nil
	}

	df := New(dataToVectors(data, types, order), conf.dfOptions...)

	return df, nil
}

func createStructConf(options ...ConfOption) confStruct {
	conf := confStruct{
		headerMap:   map[string]string{},
		dfOptions:   []vector.Option{},
		skipColumns: []string{},
	}

	for _, option := range options {
		switch option.Key() {
		case optionStructHeaderMap:
			conf.headerMap = option.Value().(map[string]string)
		case optionStructDataframeOptions:
			conf.dfOptions = option.Value().([]vector.Option)
		case optionStructSkipFields:
			conf.skipColumns = option.Value().([]string)
		}
	}

	return conf
}

func getDataAndTypes(stArrVal reflect.Value, length int, conf confStruct) (map[string][]any, map[string]string, []string) {
	data, types := map[string][]any{}, map[string]string{}

	stVal := stArrVal.Index(0)
	stType := stVal.Type()
	order := make([]string, 0, stType.NumField())
	nameMap := map[string]string{}
	for i := 0; i < stType.NumField(); i++ {
		name := stType.Field(i).Name
		tagName := stType.Field(i).Tag.Get("lth")
		if tagName == "" {
			tagName = name
		}
		if _, ok := conf.headerMap[name]; ok {
			tagName = conf.headerMap[name]
		}

		fOpt := stType.Field(i).Tag.Get("lto")
		if slices.Contains(strings.Split(fOpt, ","), "skip") || slices.Contains(conf.skipColumns, name) {
			continue
		}

		nameMap[name] = tagName
		data[nameMap[name]] = make([]any, length)
		types[nameMap[name]] = getFieldType(stVal.Field(i))
		order = append(order, nameMap[name])
	}

	for i := 0; i < stArrVal.Len(); i++ {
		stVal = stArrVal.Index(i)
		for j := 0; j < stType.NumField(); j++ {
			name := stVal.Type().Field(j).Name
			if _, ok := nameMap[name]; !ok {
				continue
			}
			data[nameMap[name]][i] = stVal.Field(j).Interface()
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

func StructOptionHeaderMap(headerMap map[string]string) ConfOption {
	return ConfOption{optionStructHeaderMap, headerMap}
}

func StructOptionDataFrameOptions(options ...vector.Option) ConfOption {
	return ConfOption{optionStructDataframeOptions, options}
}

func StructOptionSkipFields(fields ...string) ConfOption {
	return ConfOption{optionStructSkipFields, fields}
}
