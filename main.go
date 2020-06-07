package main

import (
	"log"
	"net/http"
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/souvikmaji/leaderboard/controller"
	"github.com/souvikmaji/leaderboard/db"
	"github.com/souvikmaji/leaderboard/models"
)

func main() {
	configuration := models.InitConfig()

	db, err := db.NewDB(configuration.GetDbURI())
	if err != nil {
		log.Fatalf("failed to connect database: %s", err.Error())
	}
	defer db.Close()

	db.DB.AutoMigrate(&models.Game{})
	db.DB.AutoMigrate(&models.User{})

	listenAddr := configuration.GetServerAddress()
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
