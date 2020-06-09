package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/souvikmaji/leaderboard/db"
	"github.com/urfave/negroni"
)

const (
	contentApplicationJSON = "application/json"
	contentXFormURLEncoded = "application/x-www-form-urlencoded"
)

// env represets a controller Environment
// all routes are implemented as method to this struct,
// so that all routes can share the connection pool and global variables
type env struct {
	db      db.Datastore
	decoder *schema.Decoder
}

// Handlers setsup routers and initializes controller common environments
func Handlers(datastore *db.DB) *negroni.Negroni {

	e := &env{
		db:      datastore,
		decoder: setupDecoder(),
	}

	return e.setupRouter()
}

func setupDecoder() *schema.Decoder {
	d := schema.NewDecoder()
	d.IgnoreUnknownKeys(true)
	return d
}

// dst must be a pointer to a struct
func (e *env) parsePostRequest(r *http.Request, dst interface{}) error {

	switch contentType := r.Header.Get("Content-type"); contentType {
	case contentApplicationJSON:
		return e.parseJSONRequest(r, dst)
	case contentXFormURLEncoded:
		return e.parseFormRequest(r, dst)
	default:
		return errors.New("Unknown content type")
	}

}

func (e *env) parseJSONRequest(r *http.Request, dst interface{}) error {

	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		return err
	}

	return nil
}

func (e *env) parseFormRequest(r *http.Request, dst interface{}) error {

	if err := r.ParseForm(); err != nil {
		return err
	}

	if err := e.decoder.Decode(dst, r.PostForm); err != nil {
		return err
	}

	return nil
}
