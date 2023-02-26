package vector

const keyOptionPrecision = "precision"
const keyOptionFormat = "format"
const keyOptionTimeFormat = "time_format"
const keyOptionStringToBooleanConverter = "string_boolean_converter"
const keyOptionAnyPrinterFunc = "any_printer_func"
const keyOptionAnyConvertors = "any_convertors"
const keyOptionAnyCallbacks = "any_callbacks"
const keyOptionGroupIndex = "group_index"
const keyOptionVectorName = "vector_name"
const keyOptionMaxPrintElements = "max_print_elements"

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
	return ConfOption{keyOptionPrecision, precision}
}

func OptionFormat(format string) Option {
	return ConfOption{keyOptionFormat, format}
}

func OptionTimeFormat(format string) Option {
	return ConfOption{keyOptionTimeFormat, format}
}

func OptionAnyPrinterFunc(fn AnyPrinterFunc) Option {
	return ConfOption{keyOptionAnyPrinterFunc, fn}
}

func OptionAnyConvertors(convertors AnyConvertors) Option {
	return ConfOption{keyOptionAnyConvertors, convertors}
}

func OptionAnyCallbacks(callbacks AnyCallbacks) Option {
	return ConfOption{keyOptionAnyCallbacks, callbacks}
}

func OptionGroupIndex(index GroupIndex) Option {
	return ConfOption{keyOptionGroupIndex, index}
}

func OptionVectorName(name string) Option {
	return ConfOption{keyOptionVectorName, name}
}

func OptionMaxPrintElements(maxPrintEleements int) Option {
	return ConfOption{keyOptionMaxPrintElements, maxPrintEleements}
}

func OptionStringToBooleanConverter(converter StringToBooleanConverter) Option {
	return ConfOption{keyOptionStringToBooleanConverter, converter}
}
