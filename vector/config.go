package vector

const KeyOptionPrecision = "precision"
const KeyOptionFormat = "format"
const KeyOptionTimeFormat = "time_format"
const KeyOptionAnyPrinterFunc = "any_printer_func"
const KeyOptionAnyConvertors = "any_convertors"
const KeyOptionAnyCallbacks = "any_callbacks"
const KeyOptionBeforeColumn = "before_column"
const KeyOptionAfterColumn = "after_column"
const KeyOptionColumnNames = "column_names"
const KeyOptionArrangeReverse = "arrange_reverse"
const KeyOptionArrangeReverseColumns = "arrange_reverse_columns"
const KeyOptionJoinBy = "join_by_columns"
const KeyOptionGroupIndex = "group_index"
const KeyOptionVectorName = "vector_name"

// deprecated
type Config struct {
	FloatPrinter     *FloatPrinter
	ComplexPrinter   *ComplexPrinter
	TimePrinter      *TimePrinter
	InterfacePrinter func(payload any) string
	Convertors       *AnyConvertors
}

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

type Configuration struct {
	options map[string]any
}

func (conf Configuration) HasOption(name string) bool {
	_, ok := conf.options[name]

	return ok
}

func (conf Configuration) Value(name string) any {
	return conf.options[name]
}

func (conf Configuration) SetOptions(payload Payload) {
	for str, val := range conf.options {
		payload.SetOption(str, val)
	}
}

func MergeOptions(options []Option) Configuration {
	conf := Configuration{
		options: map[string]any{},
	}

	for _, option := range options {
		conf.options[option.Key()] = option.Value()
	}

	return conf
}

func OptionPrecision(precision int) Option {
	return ConfOption{KeyOptionPrecision, precision}
}

func OptionFormat(format string) Option {
	return ConfOption{KeyOptionFormat, format}
}

func OptionTimeFormat(format string) Option {
	return ConfOption{KeyOptionTimeFormat, format}
}

func OptionAnyPrinterFunc(fn AnyPrinterFunc) Option {
	return ConfOption{KeyOptionAnyPrinterFunc, fn}
}

func OptionAnyConvertors(convertors AnyConvertors) Option {
	return ConfOption{KeyOptionAnyConvertors, convertors}
}

func OptionAnyCallbacks(callbacks AnyCallbacks) Option {
	return ConfOption{KeyOptionAnyCallbacks, callbacks}
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

func OptionGroupIndex(index GroupIndex) Option {
	return ConfOption{KeyOptionGroupIndex, index}
}

func OptionVectorName(name string) Option {
	return ConfOption{KeyOptionVectorName, name}
}
