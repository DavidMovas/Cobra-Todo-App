package models

import "errors"

var (
	ErrTaskAlreadyExists = errors.New("[ERROR] Task already exists")
	ErrTaskNotFound      = errors.New("[ERROR] Task not found")
)

type TaskList []*Task
