package dataframe

import "logarithmotechnia/vector"

const KeyOptionBeforeColumn = "before_column"
const KeyOptionAfterColumn = "after_column"
const KeyOptionColumnNames = "column_names"
const KeyOptionArrangeReverse = "arrange_reverse"
const KeyOptionArrangeReverseColumns = "arrange_reverse_columns"
const KeyOptionJoinBy = "join_by_columns"
const KeyOptionVectorOptions = "vector_options"

// Option interface
type Option interface {
	Key() string
	Value() any
}

// ConfOption struct
type ConfOption struct {
	key   string
	value any
}

// Key gets a key of an option.
func (o ConfOption) Key() string {
	return o.key
}

// Value gets a value of an option.
func (o ConfOption) Value() any {
	return o.value
}

// Configuration holds a map of options.
type Configuration struct {
	options map[string]any
}

// HasOption returns true if the configuration has an option with the key provided.
func (conf Configuration) HasOption(key string) bool {
	_, ok := conf.options[key]

	return ok
}

// Value returns a value of the option with the key provided.
func (conf Configuration) Value(key string) any {
	return conf.options[key]
}

// MergeOptions merges options into a configuration
func MergeOptions(options []Option) Configuration {
	conf := Configuration{
		options: map[string]any{},
	}

	for _, option := range options {
		conf.options[option.Key()] = option.Value()
	}

	return conf
}

func OptionBeforeColumn(name string) Option {
	return ConfOption{KeyOptionBeforeColumn, name}
}

func OptionAfterColumn(name string) Option {
	return ConfOption{KeyOptionAfterColumn, name}
}

func OptionColumnNames(names []string) Option {
	return ConfOption{KeyOptionColumnNames, names}
}

func OptionArrangeReverse(reverse bool) Option {
	return ConfOption{KeyOptionArrangeReverse, reverse}
}

func OptionArrangeReverseColumns(columns ...string) Option {
	return ConfOption{KeyOptionArrangeReverseColumns, columns}
}

func OptionJoinBy(by ...string) Option {
	return ConfOption{KeyOptionJoinBy, by}
}

func OptionVectorOptions(options []vector.Option) Option {
	return ConfOption{KeyOptionVectorOptions, options}
}
