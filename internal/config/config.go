package config

import "os"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() {
	homeDir, err := os.UserHomeDir() //returns home directory to create path to gatorconfig.json

}
