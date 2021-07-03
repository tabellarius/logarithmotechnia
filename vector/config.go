package vector

type Config struct {
	NamesMap         map[string]int
	FloatPrinter     *FloatPrinter
	TimePrinter      *TimePrinter
	InterfacePrinter func(payload interface{}) string
	Convertors       *InterfaceConvertors
}

func OptionNamesMap(namesMap map[string]int) Config {
	return Config{NamesMap: namesMap}
}

func OptionFloatPrinter(printer FloatPrinter) Config {
	return Config{FloatPrinter: &printer}
}

func OptionTimePrinter(printer TimePrinter) Config {
	return Config{TimePrinter: &printer}
}

func OptionInterfacePrinter(printer func(payload interface{}) string) Config {
	return Config{InterfacePrinter: printer}
}

func OptionConvertors(convertors InterfaceConvertors) Config {
	return Config{Convertors: &convertors}
}

func mergeConfigs(configs []Config) Config {
	config := Config{}

	for _, c := range configs {
		if c.NamesMap != nil {
			config.NamesMap = c.NamesMap
		}

		if c.FloatPrinter != nil {
			config.FloatPrinter = c.FloatPrinter
		}

		if c.TimePrinter != nil {
			config.TimePrinter = c.TimePrinter
		}

		if c.InterfacePrinter != nil {
			config.InterfacePrinter = c.InterfacePrinter
		}

		if c.Convertors != nil {
			config.Convertors = c.Convertors
		}
	}

	return config
}
