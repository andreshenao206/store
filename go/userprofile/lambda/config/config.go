package config

import (
	"os"
)

const ( //
	User = "user"
)

// Config stores all configuration of the application.
type Config struct {
	NetworkHost string
	NetworkDB   string
	CreateUser  string
}

// LoadConfig reads configuration from environments variables.
func LoadConfig() (config Config) {
	config = Config{
		NetworkHost: os.Getenv("NETWORK_HOST"),
		NetworkDB:   os.Getenv("NETWORK_HOST"),
		CreateUser:  os.Getenv("CREATE_USER"),
	}
	return
}
