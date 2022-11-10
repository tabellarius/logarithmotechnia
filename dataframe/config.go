package dataframe

const KeyOptionBeforeColumn = "before_column"
const KeyOptionAfterColumn = "after_column"
const KeyOptionColumnNames = "column_names"
const KeyOptionArrangeReverse = "arrange_reverse"
const KeyOptionArrangeReverseColumns = "arrange_reverse_columns"
const KeyOptionJoinBy = "join_by_columns"

type Option interface {
	Key() string
	Value() any
}

type ConfOption struct {
	key   string
	value any
}

func (o ConfOption) Key() string {
	return o.key
}

func (o ConfOption) Value() any {
	return o.value
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
