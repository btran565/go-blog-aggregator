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
	homeDir, err := os.UserHomeDir() //returns home directory to create path to gatorconfig.json
	if err != nil {
		return Config{}, err
	}
	path := filepath.Join(homeDir, ".gatorconfig.json") //gatorconfig path

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
	homeDir, err := os.UserHomeDir() //returns home directory to create path to gatorconfig.json
	if err != nil {
		return err
	}
	path := filepath.Join(homeDir, ".gatorconfig.json")

	err = os.WriteFile(path, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
