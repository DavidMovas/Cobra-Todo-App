package app

import "cobratodoapp/internal/db"

type APP struct {
	DB db.DB
}

func (app *APP) Run() error {

	return nil
}
