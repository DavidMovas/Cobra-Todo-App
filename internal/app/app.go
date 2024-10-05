package app

import (
	"cobratodoapp/internal/db"
	"log"

	"github.com/boltdb/bolt"
)

type APP struct {
	db db.DB
}

func (app *APP) Run() error {
	boltdb, err := bolt.Open("tasks.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer boltdb.Close()

	app.db = db.NewBoltDB(boltdb)

	return nil
}
