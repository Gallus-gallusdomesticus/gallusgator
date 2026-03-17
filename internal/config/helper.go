package config

import (
	"os"
	"path/filepath"
)

const configFile = ".gatorconfig.json" //config file name

func getConfigFilePath() (string, error) {
	homedir, err := os.UserHomeDir() //get the directory of the config file
	if err != nil {
		return "", err
	}
	jsonpath := filepath.Join(homedir, configFile) //combine the directory with the config file name
	return jsonpath, nil
}
