package vector

type Config struct {
	NamesMap     map[string]int
	FloatPrinter *FloatPrinter
	TimePrinter  *TimePrinter
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

func mergeConfigs(configs []Config) Config {
	config := Config{}

	for _, c := range configs {
		if c.NamesMap != nil {
			config.NamesMap = c.NamesMap
		}
	}

	return config
}
