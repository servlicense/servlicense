package config

import (
	"dario.cat/mergo"
	"github.com/BurntSushi/toml"
)

// Configuration struct to represent the entire TOML file structure
type Configuration struct {
	Port         int
	Host         string
	Smtp         SmtpConfiguration
	Notification NotificationConfiguration
}

type SmtpConfiguration struct {
	Enabled      bool
	SmtpHost     string
	SmtpPort     int
	SmtpUsername string
	SmtpPassword string
	SmtpFrom     string
}

type NotificationConfiguration struct {
	Recipients                  []string
	NotificationOnLicenseCreate bool
	NotificationOnLicenseRevoke bool
	NotificationOnApiKeyCreate  bool
	NotificationOnApiKeyRevoke  bool
}

// Global configuration instance
var config *Configuration

func LoadConfig(filepath string) error {
	defaultConfig := &Configuration{
		Port: 3000,
		Host: "localhost",
		Smtp: SmtpConfiguration{
			Enabled:      false,
			SmtpHost:     "",
			SmtpPort:     587,
			SmtpUsername: "",
			SmtpPassword: "",
			SmtpFrom:     "",
		},
		Notification: NotificationConfiguration{
			Recipients:                  []string{},
			NotificationOnLicenseCreate: true,
			NotificationOnLicenseRevoke: true,
			NotificationOnApiKeyCreate:  true,
			NotificationOnApiKeyRevoke:  true,
		},
	}

	var fileConfig Configuration
	if _, err := toml.DecodeFile(filepath, &fileConfig); err != nil {
		return err
	}

	if err := mergo.Merge(defaultConfig, &fileConfig, mergo.WithOverride); err != nil {
		return err
	}

	config = defaultConfig
	return nil
}

func GetConfig() *Configuration {
	return config
}
