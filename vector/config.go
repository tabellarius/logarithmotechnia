package vector

const KeyOptionPrecision = "precision"
const KeyOptionFormat = "format"
const KeyOptionTimeFormat = "time_format"
const KeyOptionStringToBooleanConverter = "string_boolean_converter"
const KeyOptionAnyPrinterFunc = "any_printer_func"
const KeyOptionAnyConvertors = "any_convertors"
const KeyOptionAnyCallbacks = "any_callbacks"
const KeyOptionGroupIndex = "group_index"
const KeyOptionVectorName = "vector_name"
const KeyOptionMaxPrintElements = "max_print_elements"

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

func OptionGroupIndex(index GroupIndex) Option {
	return ConfOption{KeyOptionGroupIndex, index}
}

func OptionVectorName(name string) Option {
	return ConfOption{KeyOptionVectorName, name}
}

func OptionMaxPrintElements(maxPrintEleements int) Option {
	return ConfOption{KeyOptionMaxPrintElements, maxPrintEleements}
}
