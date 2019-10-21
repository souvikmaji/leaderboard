package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/souvikmaji/leaderboard/models"
)

type Env struct {
	db models.Datastore
}

// health check endpoint
func healthCheck(w http.ResponseWriter, r *http.Request) {
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

func sendResponse(w http.ResponseWriter, draw int64, teams []*models.Team, totalCount int64) error {
	response := models.NewSuccessResponse(draw, teams, totalCount)
	e, err := json.Marshal(response)
	if err != nil {
		return err
	}

	w.Write(e)

	return nil
}

func (env *Env) Teams(w http.ResponseWriter, r *http.Request) {
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

	teams, totalCount, err := env.db.AllTeams(length, offset)
	if err != nil {
		sendError(w, err)
		return
	}

	if err := sendResponse(w, draw, teams, totalCount); err != nil {
		sendError(w, err)
		return
	}

}
