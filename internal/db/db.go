package db

import "cobratodoapp/internal/models"

type DB interface {
	Tasks() *models.TaskList
	Add(task *models.Task)
	Complete(id uint64) error
	Remove(id uint64) error
}
