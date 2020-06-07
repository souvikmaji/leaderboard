package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func (e *env) setupRouter() *negroni.Negroni {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/health", e.healthCheck)

	userRouter := router.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/", e.createUser).Methods("POST")
	userRouter.HandleFunc("/", e.getUser).Methods("GET")

	gameRouter := router.PathPrefix("/game").Subrouter()
	gameRouter.HandleFunc("/", e.createGame).Methods("POST")
	// gameRouter.HandleFunc("/", e.createGame).Methods("POST")
	gameRouter.HandleFunc("/leaderboard", e.getGameLeaderboard).Methods("GET")

	n := negroni.Classic()
	n.UseHandler(router)

	return n
}

// health check endpoint
func (e *env) healthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
