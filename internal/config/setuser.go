package config

import (
	"encoding/json"
	"os"
)

func (cfg *Config) SetUser(name string) error {
	cfg.CurrentUserName = name //set the config user name

	jsonData, err := json.Marshal(cfg) //decode from go to json
	if err != nil {
		return err
	}

	jsonPath, err := getConfigFilePath() //get the config path
	if err != nil {
		return err
	}

	if err := os.WriteFile(jsonPath, jsonData, 0600); err != nil { //write the decoded config user name to gatorconfig
		return err
	}

	return nil
}
