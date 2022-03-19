package config

import (
	"github.com/spf13/viper"
	"log"
)

var AllConfig *Configuration

// Configuration struct
type Configuration struct {
	Database DatabaseConfiguration
}

// DatabaseConfiguration struct
type DatabaseConfiguration struct {
	UserName string `mapstructure:"user"`
	PassWord string `mapstructure:"pass"`
	NameDB   string `mapstructure:"name"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Type     string `mapstructure:"type"`
}

// Init func initialize configuration
func Init(configPath string) {
	var configuration *Configuration

	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	AllConfig = configuration
}

func GetAllConfig() *Configuration {
	return AllConfig
}