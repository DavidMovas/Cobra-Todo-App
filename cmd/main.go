package main

import (
	"cobratodoapp/internal/app"
	"log"
)

func main() {
	app := &app.APP{}

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
