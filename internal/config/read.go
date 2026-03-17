package config

import (
	"encoding/json"
	"os"
)

//function to rread the json file in ~/.gatorconfig.json

func Read() (Config, error) {
	jsonpath, err := getConfigFilePath() //get the file path
	if err != nil {
		return Config{}, err
	}

	jsonData, err := os.ReadFile(jsonpath) //read the file on the path
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := json.Unmarshal(jsonData, &config); err != nil { //decode the json data to config struct
		return Config{}, err
	}
	return config, nil
}
