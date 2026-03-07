package config

import (
	"encoding/json"
	"os"
)

func Read() (Config, error) {
	jsonpath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	jsonData, err := os.ReadFile(jsonpath)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := json.Unmarshal(jsonData, &config); err != nil {
		return Config{}, err
	}
	return config, nil
}
