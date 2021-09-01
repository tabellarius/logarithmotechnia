package vector

//deprecated
type Config struct {
	NamesMap         map[string]int
	FloatPrinter     *FloatPrinter
	ComplexPrinter   *ComplexPrinter
	TimePrinter      *TimePrinter
	InterfacePrinter func(payload interface{}) string
	Convertors       *InterfaceConvertors
}
