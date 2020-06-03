package controller

import (
	"fmt"
	"github.com/souvikmaji/leaderboard/models"
	"net/http"
)

func (e *env) createTeam(w http.ResponseWriter, r *http.Request) {

	team := new(models.Team)
	r.ParseForm()
	err := e.decoder.Decode(team, r.PostForm)
	if err != nil {
		fmt.Println("decode error", err)
	}

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
