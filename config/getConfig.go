package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

func GetConfig(args ...string) (*Config, error) {
	//init config struct
	config := &Config{}

	// set default file path
	filePath := "../settings/default.yaml"
	// collect the filepath from varidac arguments if provided
	if len(args) > 0 {
		filePath = args[0]
	}

	// read the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// unmarshal the yaml file into conifg struct
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
