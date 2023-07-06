package models

import (
	"log"
	"sync"

	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/pop/v6"
)

// db is a connection to the database to be used
// throughout your application.
var db *pop.Connection
var once sync.Once

func Init() *pop.Connection {
	var err error
	env := envy.Get("GO_ENV", "dev")
	once.Do(func() {
		db, err = pop.Connect(env)
		if err != nil {
			log.Fatal(err)
		}
	})
	pop.Debug = env == "dev"
	return db
}
