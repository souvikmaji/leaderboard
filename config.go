package main

import (
	"fmt"
	"log"

	"github.com/souvikmaji/leaderboard/models"
	"github.com/spf13/viper"
)

// initConfig initializes application configurations from environments
func initConfig() *models.Configurations {

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

	var configuration *models.Configurations
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatal("Unable to decode into struct: ", err)
	}

	return configuration
}

// getServerAddress parses configurations and returns https server address in
// <address:port> format
func getServerAddress(c *models.Configurations) string {
	return fmt.Sprintf("%s:%d", c.Server.Host, c.Server.Port)
}
