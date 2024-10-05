package main

import (
	"cobratodoapp/internal/app"
	"cobratodoapp/internal/db"
	"cobratodoapp/internal/models"
	"log"

	"github.com/boltdb/bolt"
)

func main() {
	boltdb, err := bolt.Open("tasks.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer boltdb.Close()

	app := &app.APP{
		DB: &db.BoltDB{
			DB:    boltdb,
			Tasks: &models.TaskList{},
		},
	}

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
