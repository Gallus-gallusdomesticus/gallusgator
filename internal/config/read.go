package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func Read() (Config, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}
	jsonpath := filepath.Join(homedir, ".gatorconfig.json")

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
