package dataframe

import "logarithmotechnia/vector"

type Config struct {
	columnNames       []string
	columnNamesVector vector.Vector
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
