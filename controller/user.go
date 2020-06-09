package controller

import (
	"net/http"

	"github.com/souvikmaji/leaderboard/models"
)

func (e *env) createUser(w http.ResponseWriter, r *http.Request) {
	user := new(models.User)

	if err := e.parsePostRequest(r, user); err != nil {
		sendError(w, err)
		return
	}

	if err := e.db.SaveUser(user); err != nil {
		sendError(w, err)
		return
	}

	if err := sendResponse(w, user); err != nil {
		sendError(w, err)
		return
	}

}

func (e *env) getUser(w http.ResponseWriter, r *http.Request) {
	userQuery := new(models.User)

	if err := e.decoder.Decode(userQuery, r.URL.Query()); err != nil {
		sendError(w, err)
		return
	}

	user, err := e.db.GetUser(userQuery)
	if err != nil {
		sendError(w, err)
		return
	}

	if err := sendResponse(w, user); err != nil {
		sendError(w, err)
		return
	}

}
