package main

import (
	"encoding/json"
	"github.com/gorilla/schema"
	"github.com/souvikmaji/leaderboard/models"
	"log"
	"net/http"
	"strconv"
)

// all routes are implemented as method to this struct,
// so that all routes can share the connection pool
type env struct {
	db models.Datastore
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

func sendResponse(w http.ResponseWriter, draw int64, teams []*models.Team, totalCount, totalFiltered int64) error {
	response := models.NewSuccessResponse(draw, teams, totalCount, totalFiltered)
	e, err := json.Marshal(response)
	if err != nil {
		return err
	}

	w.Write(e)

	return nil
}

func (e *env) createTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()

	drawStr := params.Get("draw")
	draw, err := strconv.ParseInt(drawStr, 10, 64)
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

	if err := sendResponse(w, draw, teams, totalCount, totalFiltered); err != nil {
		sendError(w, err)
		return
	}

}

func (e *env) getTeamLeaderboard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()

	leaderboardQuery := new(models.LeaderboardQuery)
	decoder := schema.NewDecoder()
	decoder.Decode(leaderboardQuery, params)

	teams, totalCount, totalFiltered, err := e.db.AllTeams(leaderboardQuery.Length, leaderboardQuery.Start)
	if err != nil {
		sendError(w, err)
		return
	}

	if err := sendResponse(w, leaderboardQuery.Draw, teams, totalCount, totalFiltered); err != nil {
		sendError(w, err)
		return
	}

}
