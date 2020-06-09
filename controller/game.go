package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/souvikmaji/leaderboard/models"
)

func (e *env) createGame(w http.ResponseWriter, r *http.Request) {
	game, err := e.parsePostRequest(r)
	if err != nil {
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

func (e *env) parsePostRequest(r *http.Request) (game *models.Game, err error) {

	switch contentType := r.Header.Get("Content-type"); contentType {
	case contentApplicationJSON:
		return e.parseJSONRequest(r)
	case contentXFormURLEncoded:
		return e.parseFormRequest(r)
	default:
		return nil, errors.New("Unknown content type")
	}

}

func (e *env) parseJSONRequest(r *http.Request) (game *models.Game, err error) {
	game = new(models.Game)
	if err := json.NewDecoder(r.Body).Decode(game); err != nil {
		return nil, err
	}

	return game, nil
}

func (e *env) parseFormRequest(r *http.Request) (game *models.Game, err error) {
	game = new(models.Game)
	if err := r.ParseForm(); err != nil {
		return nil, err
	}

	if err := e.decoder.Decode(game, r.PostForm); err != nil {
		return nil, err
	}
	return game, nil
}
