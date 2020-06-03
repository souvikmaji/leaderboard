package controller

import (
	"fmt"
	"github.com/souvikmaji/leaderboard/models"
	"net/http"
)

func (e *env) createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("here")
	user := new(models.User)

	err := r.ParseForm()
	if err != nil {
		sendError(w, err)
	}

	err = e.decoder.Decode(user, r.PostForm)
	if err != nil {
		fmt.Println("decode error", err)
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
	fmt.Println("here")
	userQuery := new(models.User)
	e.decoder.Decode(userQuery, r.URL.Query())

	fmt.Println("r: ", r.URL.Query())
	fmt.Println("userQuery", userQuery)

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
