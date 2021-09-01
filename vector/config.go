package vector

const OPTION_PRECISION = "precision"
const OPTION_FORMAT = "format"
const OPTION_INTERFACE_PRINTER_FUNC = "interface_printer_func"
const OPTION_INTERFACE_CONVERTORS = "interface_convertors"

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

type confOption struct {
	key   string
	value interface{}
}

func (o confOption) Key() string {
	return o.key
}

func (o confOption) Value() interface{} {
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

func mergeOptions(options []Option) Configuration {
	conf := Configuration{
		options: map[string]interface{}{},
	}

	for _, option := range options {
		conf.options[option.Key()] = option.Value()
	}

	return conf
}

func OptionPrecision(precision int) Option {
	return confOption{OPTION_PRECISION, precision}
}

func OptionFormat(format string) Option {
	return confOption{OPTION_FORMAT, format}
}

func OptionInterfacePrinterFunc(fn InterfacePrinterFunc) Option {
	return confOption{OPTION_INTERFACE_PRINTER_FUNC, fn}
}

func OptionInterfaceConvertors(convertors *InterfaceConvertors) Option {
	return confOption{OPTION_INTERFACE_CONVERTORS, convertors}
}
