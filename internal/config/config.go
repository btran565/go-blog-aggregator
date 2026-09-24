package config

import (
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
	path := filepath.Join(homeDir, ".gatorconfig.json")
}
