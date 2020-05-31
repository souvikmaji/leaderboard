package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/schema"
	"github.com/souvikmaji/leaderboard/models"
	"log"
	"net/http"
)

// all routes are implemented as method to this struct,
// so that all routes can share the connection pool
type env struct {
	db      models.Datastore
	decoder *schema.Decoder
}

// health check endpoint
func (e *env) healthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func sendError(w http.ResponseWriter, errMsg error) {
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
	e, err := json.Marshal(response)
	if err != nil {
		return err
	}

	w.Write(e)

	return nil
}

func (e *env) createTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	team := new(models.Team)
	err := e.decoder.Decode(team, r.PostForm)
	if err != nil {
		sendError(w, err)
		return
	}

	lengthStr := params.Get("length")
	length, err := strconv.ParseInt(lengthStr, 10, 64)
	if err != nil {
		sendError(w, err)
		return
	}

	offsetStr := params.Get("start")
	offset, err := strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		sendError(w, err)
		return
	}

	teams, totalCount, totalFiltered, err := e.db.AllTeams(length, offset)
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
	w.Header().Set("Content-Type", "application/json")

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
