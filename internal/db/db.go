package db

import "cobratodoapp/internal/models"

type DB interface {
	Tasks() *models.TaskList
	Add(task *models.Task) error
	Complete(id uint64) error
	Remove(id uint64) error
}
