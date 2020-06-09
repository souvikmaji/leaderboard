package controller

import (
	"log"
	"net/http"

	"github.com/souvikmaji/leaderboard/models"
)

func (e *env) createUser(w http.ResponseWriter, r *http.Request) {
	user := new(models.User)

	if err := r.ParseForm(); err != nil {
		sendError(w, err)
	}

	if err := e.decoder.Decode(user, r.PostForm); err != nil {
		log.Println("decode error", err)
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
	e.decoder.Decode(userQuery, r.URL.Query())

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
