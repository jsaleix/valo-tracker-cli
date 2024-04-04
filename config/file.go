package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

const FILE_NAME = "./config.json"

func getConfigPath() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", errors.New("Error while retrieving executable path")
	}
	configPath := filepath.Join(filepath.Dir(exePath), FILE_NAME)

	return configPath, nil
}

type ConfigFile struct {
	Tag      string
	Username string
}

func ApplyConfig(tag, username string) error {
	configFile := ConfigFile{Tag: tag, Username: username}
	jsonData, _ := json.MarshalIndent(configFile, "", "	")
	filePath, er := getConfigPath()
	if er != nil {
		return errors.New(er.Error())
	}
	err := os.WriteFile(filePath, jsonData, 0644)
	return err
}

func ReadConfig() (ConfigFile, error) {
	var config ConfigFile

	filePath, er := getConfigPath()
	if er != nil {
		return config, errors.New(er.Error())
	}

	// Read file contents
	jsonData, err := os.ReadFile(filePath)
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
