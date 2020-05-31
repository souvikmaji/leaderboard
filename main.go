package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/souvikmaji/leaderboard/models"
	"github.com/urfave/negroni"
)

func (e *env) setupRouter() *negroni.Negroni {
	router := mux.NewRouter()
	router.HandleFunc("/health", e.healthCheck)

	teamRouter := router.PathPrefix("/team").Subrouter()
	teamRouter.HandleFunc("/", e.createTeam).Methods("POST")
	teamRouter.HandleFunc("/leaderboard", e.getTeamLeaderboard)

	n := negroni.Classic()
	n.UseHandler(router)

	return n
}

func main() {
	configuration := models.InitConfig()

	db, err := models.NewDB(configuration.GetDbURI())
	if err != nil {
		log.Fatalf("failed to connect database: %s", err.Error())
	}
	defer db.Close()

	e := &env{db}

	listenAddr := configuration.GetServerAddress()
	srv := &http.Server{
		Handler:      e.setupRouter(),
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
