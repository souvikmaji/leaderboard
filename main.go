package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/souvikmaji/leaderboard/controller"
	"github.com/souvikmaji/leaderboard/db"
	"github.com/souvikmaji/leaderboard/models"
)

func main() {
	configuration, err := models.InitConfig()
	if err != nil {
		log.Fatalln("Error initializing configurations", err)
	}

	db, err := db.NewDB(configuration)
	if err != nil {
		log.Fatalln("Failed to connect database: ", err)
	}
	defer db.Close()

	listenAddr := getServerAddress(configuration)
	srv := &http.Server{
		Handler:      controller.Handlers(db),
		Addr:         listenAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	db.LogMode(configuration.Database.Logmode)

	log.Println("Server is ready to handle requests at", listenAddr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", listenAddr, err)
	}

}

// getServerAddress parses configurations and returns https server address in
// <address:port> format
func getServerAddress(c *models.Configurations) string {
	return fmt.Sprintf("%s:%d", c.Server.Host, c.Server.Port)
}
