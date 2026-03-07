package config

import (
	"os"
	"path/filepath"
)

const configFile = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	jsonpath := filepath.Join(homedir, configFile)
	return jsonpath, nil
}
