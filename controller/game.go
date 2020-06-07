package controller

import (
	"log"
	"net/http"

	"github.com/souvikmaji/leaderboard/models"
)

func (e *env) createGame(w http.ResponseWriter, r *http.Request) {

	game := new(models.Game)
	r.ParseForm()
	err := e.decoder.Decode(game, r.PostForm)
	if err != nil {
		log.Println("decode error", err)
	}

	// decoder := json.NewDecoder(req.Body)
	err = e.db.SaveGame(game)
	if err != nil {
		sendError(w, err)
		return
	}

	if err := sendResponse(w, game); err != nil {
		sendError(w, err)
		return
	}

}
