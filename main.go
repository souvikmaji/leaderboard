package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

const (
	staticDir = "static"
	indexFile = "index.html"
)

// health check endpoint
func healthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func setupRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/teams", env.Teams)

	router.PathPrefix("/").Handler(negroni.Classic())

	return router
}

func main() {
	configuration := models.InitConfig()

	listenAddr := configuration.GetServerAddress()
	srv := &http.Server{
		Handler:      setupRouter(),
		Addr:         listenAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	db, err := gorm.Open("postgres", configuration.GetDbURI())
	if err != nil {
		panic("failed to connect database" + err.Error())
	}
	defer db.Close()

	log.Println("Server is ready to handle requests at", listenAddr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", listenAddr, err)
	}

}
