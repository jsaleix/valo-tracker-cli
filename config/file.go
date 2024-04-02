package config

import (
	"encoding/json"
	"os"
)

const FILE_PATH = "./config.json"

type ConfigFile struct {
	Tag      string
	Username string
}

func ApplyConfig(tag, username string) error {
	configFile := ConfigFile{Tag: tag, Username: username}
	jsonData, _ := json.MarshalIndent(configFile, "", "	")
	err := os.WriteFile(FILE_PATH, jsonData, 0644)
	return err
}

func ReadConfig() (ConfigFile, error) {
	var config ConfigFile

	// Read file contents
	jsonData, err := os.ReadFile(FILE_PATH)
	if err != nil {
		return config, err
	}

	// Unmarshal JSON data into Config struct
	err = json.Unmarshal(jsonData, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
