package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Env string
	DB  struct {
		Host     string
		Login    string
		Password string
		Port     string
	}
	DBSQLite struct {
		FileName string
	}
	PortApi       int
	PortConcert   int
	APIConcertKey string
}

func (c Config) String() string {
	return "not authorized to display"
}

func New() *Config {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	return &config
}
