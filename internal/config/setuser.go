package config

import (
	"encoding/json"
	"os"
)

func (cfg *Config) SetUser(name string) error {
	cfg.CurrentUserName = name

	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	jsonPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	if err := os.WriteFile(jsonPath, jsonData, 0600); err != nil {
		return err
	}

	return nil
}
