package vector

import "logarithmotechnia/option"

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

func MergeOptions(options []option.Option) Configuration {
	conf := Configuration{
		options: map[string]any{},
	}

	for _, opt := range options {
		conf.options[opt.Key()] = opt.Value()
	}

	return conf
}

func OptionPrecision(precision int) option.Option {
	return ConfOption{keyOptionPrecision, precision}
}

func OptionFormat(format string) option.Option {
	return ConfOption{keyOptionFormat, format}
}

func OptionTimeFormat(format string) option.Option {
	return ConfOption{keyOptionTimeFormat, format}
}

func OptionAnyPrinterFunc(fn AnyPrinterFunc) option.Option {
	return ConfOption{keyOptionAnyPrinterFunc, fn}
}

func OptionAnyConvertors(convertors AnyConvertors) option.Option {
	return ConfOption{keyOptionAnyConvertors, convertors}
}

func OptionAnyCallbacks(callbacks AnyCallbacks) option.Option {
	return ConfOption{keyOptionAnyCallbacks, callbacks}
}

func OptionGroupIndex(index GroupIndex) option.Option {
	return ConfOption{keyOptionGroupIndex, index}
}

func OptionVectorName(name string) option.Option {
	return ConfOption{keyOptionVectorName, name}
}

func OptionMaxPrintElements(maxPrintElements int) option.Option {
	return ConfOption{keyOptionMaxPrintElements, maxPrintElements}
}

func OptionStringToBooleanConverter(converter StringToBooleanConverter) option.Option {
	return ConfOption{keyOptionStringToBooleanConverter, converter}
}
