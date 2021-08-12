package dataframe

type Config struct {
	columnNames []string
}

func mergeConfigs(configs []Config) Config {
	config := Config{}

	for _, c := range configs {
		if c.columnNames != nil {
			config.columnNames = c.columnNames
		}
	}

	return config
}

func OptionColumnNames(columnNames []string) Config {
	return Config{
		columnNames: columnNames,
	}
}
