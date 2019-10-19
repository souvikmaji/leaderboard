package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const (
	staticDir = "static"
	indexFile = "index.html"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	// health check endpoint
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func setupRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/health", healthCheck)

	spa := handler{staticPath: staticDir, indexPath: indexFile}
	router.PathPrefix("/").Handler(spa)

	return router
}

func main() {
	port := flag.Int("port", 8000, "port number for the http server")
	flag.Parse()

	listenAddr := fmt.Sprintf("127.0.0.1:%d", *port)

	srv := &http.Server{
		Handler:      setupRouter(),
		Addr:         listenAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server is ready to handle requests at", listenAddr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", listenAddr, err)
	}

}
