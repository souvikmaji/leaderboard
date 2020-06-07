package models

import (
	"fmt"
	"log"

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
		log.Fatal("Error reading config file1: ", err)
	}

	// Load .env file if present
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	viper.SetConfigType("env")

	if err := viper.MergeInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error
		} else {
			log.Fatal("Error reading .env file: ", err)
		}
	}

	var configuration *Configurations
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatal("Unable to decode into struct: ", err)
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

	connectionString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		c.Database.Host, c.Database.Port, c.Database.Username, c.Database.Name)

	if c.Database.Password != "" {
		connectionString = fmt.Sprintf("%s password=%s", connectionString, c.Database.Password)
	}

	return connectionString
}
