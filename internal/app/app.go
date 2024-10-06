package app

import (
	"cobratodoapp/internal/db"
	"cobratodoapp/internal/models"
	"flag"
	"log"

	"github.com/boltdb/bolt"
)

type APP struct {
	db db.DB
}

var (
	action    string
	taskTitle string
	taskID    uint64
)

func (app *APP) Run() error {
	boltdb, err := bolt.Open("../tasks.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer boltdb.Close()

	app.db = db.NewBoltDB(boltdb)

	flag.StringVar(&action, "action", "list", "Command to Add, Remove, List, or Complete Task")
	flag.StringVar(&taskTitle, "title", "", "Task Title")
	flag.Uint64Var(&taskID, "id", 0, "Task ID")

	flag.Parse()

	switch action {
	case "add":
		task := models.NewTask(taskTitle)

		err = app.db.Add(task)
		if err != nil {
			return err
		}

		log.Println("Task was added")
	case "list":
		tasks := app.db.Tasks()
		if len(*tasks) == 0 {
			log.Println("No tasks found")
		}
		for _, task := range *tasks {
			log.Printf("ID: %d | Title: %s | Completed: %t | Created At: %s | Completed At: %s",
				task.ID,
				task.Title,
				task.Done,
				task.CreatedAt.Format("2006-01-02 15:04:05"),
				task.GetCompletedAt(),
			)
		}
	case "complete":
		err = app.db.Complete(taskID)
		if err != nil {
			return err
		}

		log.Println("Congrats! Task was completed")
	case "remove":
		err = app.db.Remove(taskID)
		if err != nil {
			return err
		}

		log.Println("Task was removed")
	default:
		log.Fatalf("Wrong action: %s", action)
	}

	return nil
}
