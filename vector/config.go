package vector

type Config struct {
	Names    []string
	NamesMap map[string]int
	NA       []bool
	NAMap    map[int]bool
}

func NA(na []bool) Config {
	return Config{NA: na}
}

func Names(names []string) Config {
	return Config{Names: names}
}

func NamesMap(namesMap map[string]int) Config {
	return Config{NamesMap: namesMap}
}

func mergeConfigs(configs []Config) Config {
	config := Config{}

	for _, c := range configs {
		if c.Names != nil {
			config.Names = c.Names
		}
		if c.NamesMap != nil {
			config.NamesMap = c.NamesMap
		}
		if c.NA != nil {
			config.NA = c.NA
		}
		if c.NAMap != nil {
			config.NAMap = c.NAMap
		}
	}

	return config
}
