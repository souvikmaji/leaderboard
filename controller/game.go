package controller

import (
	"net/http"

	"github.com/souvikmaji/leaderboard/models"
)

func (e *env) createGame(w http.ResponseWriter, r *http.Request) {

	game := new(models.Game)

	if err := r.ParseForm(); err != nil {
		sendError(w, err)
		return
	}

	if err := e.decoder.Decode(game, r.PostForm); err != nil {
		sendError(w, err)
		return
	}

	if err := e.db.SaveGame(game); err != nil {
		sendError(w, err)
		return
	}

	if err := sendResponse(w, game); err != nil {
		sendError(w, err)
		return
	}

}
