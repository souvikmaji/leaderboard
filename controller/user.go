package controller

import (
	"log"
	"net/http"

	"github.com/souvikmaji/leaderboard/models"
)

func (e *env) createUser(w http.ResponseWriter, r *http.Request) {
	user := new(models.User)

	err := r.ParseForm()
	if err != nil {
		sendError(w, err)
	}

	err = e.decoder.Decode(user, r.PostForm)
	if err != nil {
		log.Println("decode error", err)
	}

	// decoder := json.NewDecoder(req.Body)
	err = e.db.SaveUser(user)
	if err != nil {
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
