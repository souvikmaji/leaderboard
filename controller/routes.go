package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
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

// health check endpoint
func (e *env) healthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
