package models

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

// Configurations exported
type Configurations struct {
	Server   ServerConfigurations
	Database DatabaseConfigurations
}

// ServerConfigurations represents http server configurations
type ServerConfigurations struct {
	Host string
	Port int
}

// DatabaseConfigurations represents postgre db credentials
type DatabaseConfigurations struct {
	URL      string
	Host     string
	Port     int
	Name     string
	Username string
	Password string
	Logmode  bool
}

// InitConfig initializes application configurations from environments
func InitConfig() (c *Configurations, err error) {

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	// rewrite config keys for heroku's use
	viper.BindEnv("server.port", "PORT")
	viper.BindEnv("database.url", "DATABASE_URL")

	// read config.yml file
	viper.SetConfigName("config")

	// Path for config file is the project root
	viper.AddConfigPath(".")

	viper.SetConfigType("yml")

	// read config
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("Error reading config.yml file: %v", err)
	}

	// Load .env file if present
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	viper.SetConfigType("env")

	if err := viper.MergeInConfig(); err != nil {
		// If config file not found; ignore error; otherwise throw
		if !errors.As(err, &viper.ConfigFileNotFoundError{}) {
			return nil, fmt.Errorf("Error reading .env file: %v", err)
		}
	}

	if err := viper.Unmarshal(&c); err != nil {
		return nil, fmt.Errorf("Unable to decode into struct: %v", err)
	}

	return c, nil
}
