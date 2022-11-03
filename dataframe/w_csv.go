package dataframe

import (
	"encoding/csv"
	"golang.org/x/exp/slices"
	"io"
	"logarithmotechnia/vector"
	"math"
	"os"
	"strconv"
)

const optionCSVSkipFirstLine = "skipFirstLine"
const optionCSVSeparator = "separator"

type confCSV struct {
	colTypes      []string
	colNames      []string
	skipFirstLine bool
	separator     rune
}

type OptionCSV struct {
	name string
	val  any
}

func FromCSVFile(filename string, options ...OptionCSV) (*Dataframe, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	df, err := FromCSV(file, options...)

	return df, err
}

func FromCSV(reader io.Reader, options ...OptionCSV) (*Dataframe, error) {
	conf := combineCSVConfig(options...)

	r := csv.NewReader(reader)
	r.Comma = conf.separator

	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	rowNum, colNum := len(records), 0
	if rowNum == 0 {
		return New([]vector.Vector{}), nil
	}
	colNum = len(records[0])
	if colNum == 0 {
		return New([]vector.Vector{}), nil
	}

	conf.colNames = make([]string, colNum)
	for i := 0; i < colNum; i++ {
		conf.colNames[i] = strconv.Itoa(i)
	}

	if conf.skipFirstLine {
		conf.colNames = records[0]
		records = records[1:]
		rowNum = rowNum - 1
	}

	vecs := make([]vector.Vector, colNum)
	for i := 0; i < colNum; i++ {
		arr := make([]string, rowNum)
		for j := 0; j < rowNum; j++ {
			arr[j] = records[j][i]
		}
		vecs[i] = vector.String(arr)
	}

	types := defaultTypes(colNum)
	if len(records) > 0 {
		types = detectTypes(records[0], vector.DefaultStringToBoolConverter())
	}
	vecs = convertVectors(vecs, types)

	df := New(vecs, vector.OptionColumnNames(conf.colNames))

	return df, nil
}

func combineCSVConfig(options ...OptionCSV) confCSV {
	conf := confCSV{
		colTypes:      []string{},
		colNames:      []string{},
		skipFirstLine: true,
		separator:     ',',
	}

	for _, option := range options {
		switch option.name {
		case optionCSVSkipFirstLine:
			conf.skipFirstLine = option.val.(bool)
		case optionCSVSeparator:
			conf.separator = option.val.(rune)
		}
	}

	return conf
}

func detectTypes(templateRow []string, boolConv vector.StringToBooleanConverter) []string {
	types := make([]string, len(templateRow))
	for i := 0; i < len(templateRow); i++ {
		types[i] = "string"
		fVal, err := strconv.ParseFloat(templateRow[i], 64)
		if err == nil {
			_, frac := math.Modf(fVal)
			if frac == 0 {
				types[i] = "integer"
			} else {
				types[i] = "float"
			}
			continue
		}

		if slices.Contains(boolConv.TrueValues(), templateRow[i]) ||
			slices.Contains(boolConv.TrueValues(), templateRow[i]) {
			types[i] = "boolean"
			continue
		}
	}

	return types
}

func defaultTypes(length int) []string {
	types := make([]string, length)
	for i := 0; i < length; i++ {
		types[i] = "string"
	}

	return types
}

func convertVectors(vecs []vector.Vector, types []string) []vector.Vector {
	for i, vec := range vecs {
		switch types[i] {
		case "integer":
			vecs[i] = vec.AsInteger()
		case "float":
			vecs[i] = vec.AsFloat()
		case "boolean":
			vecs[i] = vec.AsBoolean()
		}
	}

	return vecs
}

func CSVOptionSkipFirstLine(skip bool) OptionCSV {
	return OptionCSV{
		name: optionCSVSkipFirstLine,
		val:  skip,
	}
}

func CSVOptionSeparator(separator rune) OptionCSV {
	return OptionCSV{
		name: optionCSVSeparator,
		val:  separator,
	}
}
