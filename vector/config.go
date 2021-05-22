package vector

type Config struct {
	NamesMap map[string]int
}

func OptionNamesMap(namesMap map[string]int) Config {
	return Config{NamesMap: namesMap}
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
