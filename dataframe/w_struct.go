package dataframe

import (
	"errors"
	"fmt"
	"logarithmotechnia/vector"
	"reflect"
	"strings"
	"time"
)

func FromStructs(stArr any) (*Dataframe, error) {
	stArrType := reflect.TypeOf(stArr)

	if stArrType.Kind() != reflect.Slice && stArrType.Kind() != reflect.Array &&
		stArrType.Elem().Kind() != reflect.Struct {
		return nil, errors.New("data is not slice or array of structs")
	}

	stArrVal := reflect.ValueOf(stArr)
	if stArrVal.Len() == 0 {
		return New([]vector.Vector{}), nil
	}

	data, types := getEmptyDataAndTypes(stArrVal.Index(0), stArrVal.Len())

	fmt.Println(data)
	fmt.Println(types)

	//	stArrVal := reflect.ValueOf(stArr)
	//	examiner(stArrType, 0)

	return nil, nil
}

func getEmptyDataAndTypes(stVal reflect.Value, length int) (map[string][]any, map[string]string) {
	data, types := map[string][]any{}, map[string]string{}

	stType := stVal.Type()
	examiner(stType, 0)
	for i := 0; i < stType.NumField(); i++ {
		fType := stType.Field(i)
		fVal := stVal.Field(i)
		name := fType.Name
		data[name] = make([]any, length)

		switch fVal.Kind() {
		case reflect.Int:
			types[name] = "integer"
		case reflect.Float64:
			types[name] = "float"
		case reflect.Complex128:
			types[name] = "complex"
		case reflect.String:
			types[name] = "string"
		case reflect.Bool:
			types[name] = "boolean"
		case reflect.Struct:
			if _, ok := fVal.Interface().(time.Time); ok {
				types[name] = "time"
			} else {
				types[name] = "any"
			}
		}
	}

	return data, types
}

func examiner(t reflect.Type, depth int) {
	fmt.Println(strings.Repeat("\t", depth), "Type is", t.Name(), "and kind is", t.Kind())
	switch t.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		fmt.Println(strings.Repeat("\t", depth+1), "Contained type:")
		examiner(t.Elem(), depth+1)
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			fmt.Println(strings.Repeat("\t", depth+1), "Field", i+1, "name is", f.Name, "type is", f.Type.Name(), "and kind is", f.Type.Kind())
			if f.Tag != "" {
				fmt.Println(strings.Repeat("\t", depth+2), "Tag is", f.Tag)
				fmt.Println(strings.Repeat("\t", depth+2), "tag1 is", f.Tag.Get("tag1"), "tag2 is", f.Tag.Get("tag2"))
			}
		}
	}
}
