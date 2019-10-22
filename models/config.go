package models

import (
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

// InitConfig initializes a new configuration object denoting app configurations
func InitConfig() *Configurations {
	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	viper.SetConfigType("yml")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	// rewrite config keys for heroku's use
	viper.BindEnv("server.port", "PORT")
	viper.BindEnv("database.url", "DATABASE_URL")

	// read config
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	var configuration *Configurations
	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	return configuration
}

// GetServerAddress parses configurations and returns https server address in
// <address:port> format
func (c *Configurations) GetServerAddress() string {
	return fmt.Sprintf("%s:%d", c.Server.Host, c.Server.Port)
}

// GetDbURI parses configurations and returns database server address in
// postgreSQL connection string format
func (c *Configurations) GetDbURI() string {
	if c.Database.URL != "" {
		return c.Database.URL
	}
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		c.Database.Host, c.Database.Port, c.Database.Username, c.Database.Name)
}
