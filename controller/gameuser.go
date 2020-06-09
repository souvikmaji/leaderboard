package controller

import (
	"net/http"

	"github.com/souvikmaji/leaderboard/models"
)

func (e *env) getGameLeaderboard(w http.ResponseWriter, r *http.Request) {

	leaderboardQuery := new(models.LeaderboardQuery)
	if err := e.decoder.Decode(leaderboardQuery, r.URL.Query()); err != nil {
		sendError(w, err)
		return
	}

	games, totalCount, totalFiltered, err := e.db.GetAllSortedGameUser(leaderboardQuery.Length, leaderboardQuery.Start)
	if err != nil {
		sendError(w, err)
		return
	}

	response := models.NewLeaderboardSuccessResponse(leaderboardQuery.Draw, games, totalCount, totalFiltered)
	if err := sendResponse(w, response); err != nil {
		sendError(w, err)
		return
	}

}
