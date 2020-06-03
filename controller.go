package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/souvikmaji/leaderboard/db"
	"github.com/souvikmaji/leaderboard/models"
)

// all routes are implemented as method to this struct,
// so that all routes can share the connection pool and global variables
type env struct {
	db      db.Datastore
	decoder *schema.Decoder
}

// health check endpoint
func (e *env) healthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func sendError(w http.ResponseWriter, errMsg error) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("Error: ", errMsg)
	response := models.NewErrorResponse(errMsg)
	e, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusInternalServerError)
	w.Write(e)
}

func sendResponse(w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	e, err := json.Marshal(response)
	if err != nil {
		return err
	}

	w.Write(e)

	return nil
}

func (e *env) createTeam(w http.ResponseWriter, r *http.Request) {

	team := new(models.Team)
	r.ParseForm()
	err := e.decoder.Decode(team, r.PostForm)
	if err != nil {
		fmt.Println("decode error", err)
	}

	fmt.Println("team:", team)

	// decoder := json.NewDecoder(req.Body)
	err = e.db.SaveTeam(team)
	if err != nil {
		sendError(w, err)
		return
	}

	if err := sendResponse(w, team); err != nil {
		sendError(w, err)
		return
	}

}

func (e *env) getTeamLeaderboard(w http.ResponseWriter, r *http.Request) {

	leaderboardQuery := new(models.LeaderboardQuery)
	e.decoder.Decode(leaderboardQuery, r.URL.Query())

	teams, totalCount, totalFiltered, err := e.db.AllTeams(leaderboardQuery.Length, leaderboardQuery.Start)
	if err != nil {
		sendError(w, err)
		return
	}

	response := models.NewLeaderboardSuccessResponse(leaderboardQuery.Draw, teams, totalCount, totalFiltered)
	if err := sendResponse(w, response); err != nil {
		sendError(w, err)
		return
	}

}
