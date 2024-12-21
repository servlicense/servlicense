package config

import (
	"github.com/BurntSushi/toml"
)

// Configuration struct to represent the entire TOML file structure
type Configuration struct {
	Port int
	Host string
}

// Global configuration instance
var config *Configuration

// LoadConfig loads the configuration file and stores it globally
func LoadConfig(filepath string) error {
	var cfg Configuration
	if _, err := toml.DecodeFile(filepath, &cfg); err != nil {
		return err
	}
	config = &cfg
	return nil
}

func GetConfig() *Configuration {
	return config
}
