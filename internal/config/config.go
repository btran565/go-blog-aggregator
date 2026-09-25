package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var c Config
	err = json.Unmarshal(data, &c)
	if err != nil {
		return Config{}, err
	}
	return c, nil
}

func (c *Config) SetUser(name string) error {
	c.CurrentUserName = name

	jsonData, err := json.Marshal(c)
	if err != nil {
		return err
	}

	path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	err = os.WriteFile(path, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir() //returns home directory to create path to gatorconfig.json
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, ".gatorconfig.json"), nil
}
