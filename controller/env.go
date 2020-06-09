package controller

import (
	"github.com/gorilla/schema"
	"github.com/souvikmaji/leaderboard/db"
	"github.com/urfave/negroni"
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
