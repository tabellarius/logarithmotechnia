package vector

const KeyOptionPrecision = "precision"
const KeyOptionFormat = "format"
const KeyOptionInterfacePrinterFunc = "interface_printer_func"
const KeyOptionInterfaceConvertors = "interface_convertors"
const KeyOptionBeforeColumn = "before_column"
const KeyOptionAfterColumn = "after_column"
const KeyOptionColumnNames = "column_names"
const KeyOptionArrangeReverse = "arrange_reverse"
const KeyOptionArrangeReverseColumns = "arrange_reverse_columns"

//deprecated
type Config struct {
	FloatPrinter     *FloatPrinter
	ComplexPrinter   *ComplexPrinter
	TimePrinter      *TimePrinter
	InterfacePrinter func(payload interface{}) string
	Convertors       *InterfaceConvertors
}

type Option interface {
	Key() string
	Value() interface{}
}

type ConfOption struct {
	key   string
	value interface{}
}

func (o ConfOption) Key() string {
	return o.key
}

func (o ConfOption) Value() interface{} {
	return o.value
}

type Configuration struct {
	options map[string]interface{}
}

func (conf Configuration) HasOption(name string) bool {
	_, ok := conf.options[name]

	return ok
}

func (conf Configuration) Value(name string) interface{} {
	return conf.options[name]
}

func MergeOptions(options []Option) Configuration {
	conf := Configuration{
		options: map[string]interface{}{},
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

func OptionInterfacePrinterFunc(fn InterfacePrinterFunc) Option {
	return ConfOption{KeyOptionInterfacePrinterFunc, fn}
}

func OptionInterfaceConvertors(convertors *InterfaceConvertors) Option {
	return ConfOption{KeyOptionInterfaceConvertors, convertors}
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

func OptionArrangeReverseColumns(columns []string) Option {
	return ConfOption{KeyOptionArrangeReverseColumns, columns}
}
